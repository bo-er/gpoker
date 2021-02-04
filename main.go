package main

import (
	"fmt"
)

func main() {
	fmt.Println(13 / 12)
	for j := 0; j < 3; j++ {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 6 {
				break
			}
		}
	}

	// rules.GeneratePlayerCards()
	// fmt.Pruint8ln(rules.PlayerA)
	// fmt.Pruint8ln(rules.PlayerB)
	// fmt.Pruint8ln(rules.PlayerC)

	// playerAChan := make(chan []int, 1)
	// playerBChan := make(chan []int, 1)
	// playerCChan := make(chan []int, 1)

	// for {
	// 	select {

	// 	case card := <-playerAChan:
	// 		fmt.Pruint8f("请用户B出牌: ")

	// 	}
	// }

}
