package codes

// MessageType 定义消息类型
type MessageType = int32

const (
	// GameStart is message for game start
	GameStart MessageType = 1
	// InitialCardsDistribute is message for system's initial distribution of cards
	InitialCardsDistribute MessageType = 2
	// MasterCardsDistribute happens after master is picked and three cards are being distributed
	MasterCardsDistribute MessageType = 3
	// CompetitionMaster is a message server sends to clients for signaling master competition
	CompetitionMaster MessageType = 3
	// PlayerMasterCall is message client sends to server and competes for master
	PlayerMasterCall MessageType = 4
	// MasterBroadCast happens after a master is picked randomly
	MasterBroadCast MessageType = 5
	// PlayCards is the message when players are playing cards
	PlayCards MessageType = 6
	// PlayerDeclining is the mesaage when player gives up his turn
	PlayerDeclining MessageType = 7
	// GameAbruptlyOver for example when a client abruptly disconnects
	GameAbruptlyOver MessageType = 8
	// GameOver is for a normal game over with winners and losers
	GameOver MessageType = 9
	// Heartbeat is sent for both client and server to confirm connectivity
	Heartbeat MessageType = 3
	// SystemMessage is distributed when server needs to told client something
	SystemMessage MessageType = 3
)
