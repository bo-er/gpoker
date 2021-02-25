package rules

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bo-er/poker/utils"
)

var cards []int

// Init 是初始化扑克牌的函数
func Init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("初始化扑克牌")
	cards = make([]int, 54)
	for i := 0; i < 54; i++ {
		cards[i] = i
	}
	// 洗牌算法
	for i := 53; i > 0; i-- {
		index := rand.Intn(i)
		utils.SwitchArrayElement(cards, index, 53-i)
	}
}

// GetRandomCard 从已经洗好牌的数组中获取特定数量的牌，结果以字符串返回，逗号隔开。
func GetRandomCard(numbers int) (string, error) {
	if len(cards) < numbers {
		return "", fmt.Errorf("Cards are less than %d", numbers)
	}
	pickedCards := cards[(len(cards)-numbers):]
	cards = cards[:len(cards)-numbers]
	result := utils.IntArrayToString(pickedCards, ",")
	return result, nil
}
