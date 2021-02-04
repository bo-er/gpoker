package rules

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 54; i++ {
		remainNumber[i] = i
	}
}

var takenNumber []int = make([]int, 54)
var remainNumber []int = make([]int, 54)

func GetRandomCard() int {
	index := rand.Intn(len(remainNumber))
	chosen := remainNumber[index]
	remainNumber = append(remainNumber[:index], remainNumber[index+1:]...)
	takenNumber = append(takenNumber, chosen)
	return chosen
}

func PickAMaster() string {
	players := []string{"PlayerA", "PlayerB", "PlayerC"}
	index := rand.Intn(len(players))
	return players[index]
}
