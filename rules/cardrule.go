package rules

import (
	"github.com/bo-er/poker/utils"
)

// func handlePlayCard(playerCards *[]int, toPlayCards []int) *[]int {

// }

// CheckCardsRules 根据牌检查是否合法，如果合法则返回所出牌的类型
func CheckCardsRules(cards []int) string {
	switch len(cards) {
	case 1:
		return "singleCard"
	case 2:
		if TwoJokers(cards) {
			return "twoJokers"
		}
		return "doubleCards"
	case 4:
		if IsBomb(cards) {
			return "bomb"
		}
		if IsThreeWithOne(cards) {
			return "threeWithOne"
		}
		return "invalid"

	case 5:
		if IsThreeWithTwo(cards) {
			return "threeWithTwo"
		}
		if IsStraight(cards) {
			return "straight"
		}
		return "invalid"
	default:
		if IsStraight(cards) {
			return "straight"
		}
		if IsPairs(cards) {
			return "pairs"
		}
		return "invalid"
	}
}

// IsBomb 检查是否是炸弹
func IsBomb(cards []int) bool {
	var pivot int
	for i, card := range cards {
		if i == 0 {
			pivot = card % 13
		}
		if cards[i] == 52 || cards[i] == 53 {
			return false
		}

		if pivot != card%13 {
			return false
		}
	}
	return true
}

// TwoJokers 检查是否是王炸
func TwoJokers(cards []int) bool {
	if len(cards) != 2 {
		return false
	}
	if cards[0] == 52 && cards[1] == 53 {
		return true
	}
	if cards[0] == 53 && cards[1] == 52 {
		return true
	}
	return false
}

// IsThreeWithOne 是否是三带一
func IsThreeWithOne(cards []int) bool {
	length := len(cards)
	if length != 4 {
		return false
	}
	cardsMap := make(map[int]int, 2)
	sortedCards := utils.QuickSort(cards)
	count := 0
	for i, card := range sortedCards {
		if sortedCards[i] >= 52 {
			return false
		}
		c, ok := cardsMap[card%13]
		if ok {
			count = c + 1
			cardsMap[card%13]++

		} else {
			cardsMap[card%13] = 1
		}
		if len(cardsMap) > 2 {
			return false
		}
	}
	if count == 3 {
		return true
	}
	return false

}

// IsThreeWithTwo 检查是否是三带二
func IsThreeWithTwo(cards []int) bool {
	length := len(cards)
	if length != 5 {
		return false
	}
	cardsMap := make(map[int]int, 2)
	sortedCards := utils.QuickSort(cards)
	count := 0
	for i, card := range sortedCards {
		if sortedCards[i] >= 52 {
			return false
		}
		c, ok := cardsMap[card%13]
		if ok {
			count = c + 1
			cardsMap[card%13]++

		} else {
			cardsMap[card%13] = 1
		}
		if len(cardsMap) > 2 {
			return false
		}
	}
	if count == 3 {
		return true
	}
	return false

}

// IsPairs 检查牌是否是连对
func IsPairs(cards []int) bool {
	if len(cards) < 6 {
		return false
	}
	if len(cards)%2 != 0 {
		return false
	}
	pairCount := len(cards) / 2
	countsMap := make(map[int]int, pairCount)
	for i, card := range cards {
		if cards[i] >= 52 {
			return false
		}
		c, ok := countsMap[card%13]
		if ok {
			if c == 2 {
				return false
			}
			countsMap[card%13]++
		} else {
			countsMap[card%13] = 1
		}
	}
	if len(countsMap) != pairCount {
		return false
	}
	return true

}

// IsStraight 检查牌是否是顺子
func IsStraight(cards []int) bool {
	length := len(cards)
	if length < 5 {
		return false
	}
	for i := 0; i < length; i++ {
		cards[i] = cards[i] % 13
	}

	sortedCards := utils.QuickSort(cards)

	for i, card := range sortedCards {
		if sortedCards[i] >= 52 {
			return false
		}
		if card%13 == 11 || card%13 == 12 {
			return false
		}
		if i < length-1 {
			if (card%13)+1 != sortedCards[i+1]%13 {
				return false
			}
		}
	}
	return true
}

func getPivotOfThree(cards []int) int {
	pivot1 := 0
	count1 := 0
	pivot2 := 0
	count2 := 0

	for _, card := range cards {
		if pivot1 == 0 {
			pivot1 = card % 13
			count1++
			continue
		}
		if pivot2 == 0 {
			pivot2 = card % 13
			count2++
			continue
		}
		if card%13 == pivot1 {
			count1++
		} else {
			count2++
		}
	}
	if count1 > count2 {
		return pivot1
	}
	return pivot2
}

func stripCards(cards []int)[]int {
	for i, card := range cards {
		cards[i] = card % 13
	}
	return cards
}
