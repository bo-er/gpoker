package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"github.com/bo-er/poker/pb"
	"github.com/bo-er/poker/rules"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

// var gameStore *GameStore

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	connectionMap = make(map[string]*Connection)

}

// Connection 是存储连接信息的结构体
type Connection struct {
	stream pb.BroadCast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

// Server 是游戏服务结构体
type Server struct {
	Connections       []*Connection
	InMemoryGameStore *GameStore
}

// GameStore 是存储游戏数据的结构体
type GameStore struct {
	PlayerStatus   map[string]bool  // 存储玩家准备状态的map,key为playerID,value为布尔类型，ture表示玩家准备就绪，false表示未准备
	Master         string           //存储地主的playerID
	MasterNominees []string         //存储叫地主的playerID
	CurrentPlayer  string           //存储当前出牌的玩家
	PlayerCards    map[string][]int //存储各个玩家当前的牌
}

var connectionMap map[string]*Connection

// CreateStream 是创建连接的方法
func (s *Server) CreateStream(pconn *pb.Connect, stream pb.BroadCast_CreateStreamServer) error {
	if len(s.Connections) == 3 {
		grpcLog.Error("已经有三名玩家正在进行游戏，新玩家尝试加入失败")
		return errors.New("已经有三名玩家正在进行游戏，您无法加入")
	}
	conn := &Connection{
		stream: stream,
		id:     pconn.Player.Id,
		active: pconn.Active,
		error:  make(chan error),
	}
	s.Connections = append(s.Connections, conn)
	connectionMap[conn.id] = conn
	return <-conn.error
}

// DeliverMessage 是服务器给特定客户端发送消息的方法
func (s *Server) DeliverMessage(ctx context.Context, msg *pb.Message) (*pb.Close, error) {
	for _, conn := range s.Connections {
		if conn.id == msg.PlayerID {
			err := conn.stream.Send(msg)
			grpcLog.Infof("给玩家%s发送数据:%v", conn.id, conn.stream)
			if err != nil {
				grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
				conn.active = false
				conn.error <- err
			}
		}
	}
	return &pb.Close{}, nil
}

// BroadcastMessage 是服务器广播消息的方法
func (s *Server) BroadcastMessage(ctx context.Context, msg *pb.Message) (*pb.Close, error) {
	_, ok := connectionMap[msg.PlayerID]
	if !ok && msg.PlayerID != "system001" {
		return &pb.Close{}, errors.New("Illegal player id, service has been denied")
	}
	if len(s.Connections) != 3 {
		return &pb.Close{}, errors.New("wait until all 3 players are connected")
	}
	if msg.Content == "ready" {
		s.InMemoryGameStore.PlayerStatus[msg.PlayerID] = true
	}
	processed := s.processPlayerMessage(msg)
	if processed {
		return &pb.Close{}, nil
	}
	fmt.Println("gameStore player status is:", s.InMemoryGameStore.PlayerStatus)
	wait := sync.WaitGroup{}
	done := make(chan int)
	for _, conn := range s.Connections {
		wait.Add(1)
		go func(msg *pb.Message, conn *Connection) {
			defer wait.Done()
			if conn.active {
				// Send方法在stream接口里
				err := conn.stream.Send(msg)
				grpcLog.Infof("给客户端%s发送数据:%v\n", conn.id, msg.Content)
				if err != nil {
					grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()
	<-done
	return &pb.Close{}, nil

}

func main() {
	var connections []*Connection
	server := &Server{
		Connections: connections,
		InMemoryGameStore: &GameStore{
			PlayerStatus: make(map[string]bool),
			PlayerCards:  make(map[string][]int),
		},
	}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server: %v", err)
	}
	grpcLog.Info("Starting server at port : 8080")
	// 由于server实现了BroadCast服务的两个方法，因此可以作为BroadCast服务注册
	pb.RegisterBroadCastServer(grpcServer, server)
	go func() {
	CHECK:
		for {
			time.Sleep(time.Second)
			ctx := context.Background()
			fmt.Printf("connections are:%v,and current time is:%v\n", server.Connections, time.Now())
			if len(server.Connections) == 3 {
				// for _, value := range server.InMemoryGameStore.PlayerStatus {
				// 	if !value {
				// 		goto CHECK
				// 	}
				// }
				rules.Init()
				err := server.systemBroadCastMessage(ctx, "游戏开始", 0)
				if err != nil {
					fmt.Printf("Error sending message:%s", err)

				}

				for _, conn := range server.Connections {
					cards := getRandomCards(17)
					err = server.deliverMessage(ctx, cards, conn.id, 1)

				}
				for _, conn := range server.Connections {
					err = server.deliverMessage(ctx,
						"开始抢地主,发送`叫地主`来抢地主,系统将随机挑选一名叫地主玩家作为地主,如果20秒没有人叫地主游戏将重新开始", conn.id, 3)
				}
				timer := time.NewTimer(20 * time.Second)
				<-timer.C
				if len(server.InMemoryGameStore.MasterNominees) != 0 {
					master := server.pickerARandomMaster()
					server.systemBroadCastMessage(ctx, master, 5)
					remainCards, err := rules.GetRandomCard(3)
					if err != nil {
						fmt.Printf("获取地主的三张牌时出错:%v\n", err)
					}
					err = server.deliverMessage(ctx, remainCards, master, 2)
					if err != nil {
						fmt.Printf("给抢到地主的玩家:%s发送地主的三张牌时产生错误:%v", master, err)
					}
				} else {
					fmt.Println("由于没有玩家叫地主，游戏重新开始")
					goto CHECK
				}

				return
			}
		}

	}()
	grpcServer.Serve(listener)
}

func getRandomCards(total int) string {
	result, err := rules.GetRandomCard(total)
	if err != nil {
		fmt.Printf("error when trying to get random cards:%v", err)
		return ""
	}
	return result
}

func (s *Server) systemBroadCastMessage(ctx context.Context, message string, messageType int32) error {
	timestamp := time.Now()
	id := sha256.Sum256([]byte(timestamp.String()))
	_, err := s.BroadcastMessage(ctx, &pb.Message{
		Id:          hex.EncodeToString(id[:]),
		PlayerID:    "system001",
		Content:     message,
		Timestamp:   timestamp.String(),
		MessageType: messageType,
	})
	if err != nil {
		return fmt.Errorf("Error sending message:%s", err)
	}
	return nil
}

// deliverMessage 跟广播不同的是 deliverMessage是向特定客户端发送消息
func (s *Server) deliverMessage(ctx context.Context, message, playerID string, messageType int32) error {
	timestamp := time.Now()
	id := sha256.Sum256([]byte(timestamp.String() + playerID))
	_, err := s.DeliverMessage(ctx, &pb.Message{
		Id:          hex.EncodeToString(id[:]),
		PlayerID:    playerID,
		Content:     message,
		Timestamp:   timestamp.String(),
		MessageType: messageType,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) processPlayerMessage(msg *pb.Message) bool {
	if msg.MessageType == 4 {
		s.InMemoryGameStore.MasterNominees = append(s.InMemoryGameStore.MasterNominees, msg.PlayerID)
		return true
	}
	return false
}

func (s *Server) pickerARandomMaster() string {
	index := rand.Intn(len(s.InMemoryGameStore.MasterNominees))
	return s.InMemoryGameStore.MasterNominees[index]
}
