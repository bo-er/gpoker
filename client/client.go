package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bo-er/poker/pb"
	"google.golang.org/grpc"
)

var client pb.BroadCastClient
var wait *sync.WaitGroup

var cards []int

//通过初始化函数来初始化WaitGroup
func init() {
	wait = &sync.WaitGroup{}
}

func connect(player *pb.Player) error {
	var streamError error
	stream, err := client.CreateStream(context.Background(), &pb.Connect{
		Player: player,
		Active: true,
	})
	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}
	fmt.Println("stream created")
	wait.Add(1)
	go func(str pb.BroadCast_CreateStreamClient) {
		defer wait.Done()
		for {
			msg, err := str.Recv()
			if err != nil {
				streamError = fmt.Errorf("Error reading message: %v", err)
				break
			}
			// if msg.Type == 0 && msg.PlayerID == player.Id {
			// 	fmt.Println("接收到系统初次发牌")
			// }
			fmt.Printf("%v : %s\n", msg.Id, msg.Content)
		}
	}(stream)
	return streamError
}

func main() {
	timestamp := time.Now()
	done := make(chan int)

	name := flag.String("name", "player", "The name of poker player")
	flag.Parse()

	id := sha256.Sum256([]byte(timestamp.String() + *name))

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to service: %v", err)
	}

	client = pb.NewBroadCastClient(conn)
	player := &pb.Player{
		// Id:   hex.EncodeToString(id[:]),
		Id:   *name,
		Name: *name,
	}

	err = connect(player)
	if err != nil {
		log.Fatalf("Couldn't connect to service: %v", err)
	}
	wait.Add(1)
	go func() {
		defer wait.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message, messageType := parsePlayerInstruction(scanner.Text())
			msg := &pb.Message{
				Id:          hex.EncodeToString(id[:]),
				PlayerID:    player.Id,
				Content:     message,
				Timestamp:   timestamp.String(),
				MessageType: int32(messageType),
			}
			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				fmt.Printf("Error sending message:%s", err)
				break
			}
		}

	}()

	go func() {
		wait.Wait()
		close(done)
	}()
	<-done
}

func parsePlayerInstruction(instruction string) (parsedMessage string, messageType int) {
	fmt.Println("instruction is:", instruction)
	switch instruction {
	case "call master":
		messageType = 4
		return
	case "pass":
		messageType = 7
		return
	default:
		parsedMessage = instruction
		messageType = 6
		return
	}
}
