package rules

import (
	"fmt"

	"github.com/bo-er/poker/utils"
)

// CheckGameRuleFollowed 用于检查出牌是否符合游戏规则
func CheckGameRuleFollowed(previous []int, later []int) bool {
	laterType := CheckCardsRules(later)
	previousType := CheckCardsRules(previous)

	if previousType == "bomb" {

		if len(later) != 4 && len(later) != 2 {
			return false
		}
		if laterType == "twoJokers" {
			return true
		}
		if later[0]%13 > previous[0]%13 {
			return true
		}
		return false
	}

	if previousType == "twoJokers" {
		return false
	}

	if previousType != laterType {
		return false
	}

	switch previousType {
	case "singleCard":
		return later[0]%13 > previous[0]%13
	case "doubleCards":
		return later[0]%13 > previous[0]%13
	case "threeWithOne":
		previousPivot := getPivotOfThree(previous)
		laterPivot := getPivotOfThree(later)
		return laterPivot > previousPivot
	case "threeWithTwo":
		previousPivot := getPivotOfThree(previous)
		laterPivot := getPivotOfThree(later)
		return laterPivot > previousPivot
	case "straight":
		if len(previous) != len(later) {
			return false
		}
		previousSorted := utils.Qsort(stripCards(previous))
		laterSorted := utils.Qsort(stripCards(later))
		if previousSorted[0]%13 < laterSorted[0]%13 {
			return true
		}
		return false
	case "pairs":
		if len(previous) != len(later) {
			return false
		}
		previousSorted := utils.Qsort(stripCards(previous))
		laterSorted := utils.Qsort(stripCards(later))
		if previousSorted[0]%13 < laterSorted[0]%13 {
			return true
		}
		return false
	default:
		fmt.Println("游戏规则匹配失效，默认拒绝")
		return false
	}

}
