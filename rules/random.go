package rules

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("初始化扑克牌")
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
	fmt.Println("the card chosen is :", chosen,"remain cards are:",remainNumber)
	return chosen
}

func GetRemainingCards() string{
	
}

