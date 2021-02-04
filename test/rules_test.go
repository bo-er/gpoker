package test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bo-er/poker/play"
	"github.com/bo-er/poker/rules"
	"github.com/bo-er/poker/utils"
	"github.com/stretchr/testify/require"
)

func TestRules(t *testing.T) {
	cards := []int{1, 14, 27, 40}
	isSame := rules.IsBomb(cards)
	require.True(t, isSame)
}

func TestThreeWithTwo(t *testing.T) {
	cards := []int{1, 14, 27, 2, 15}
	isThreeWithTwo := rules.IsThreeWithTwo(cards)
	require.True(t, isThreeWithTwo)
}

func TestIsPairs(t *testing.T) {
	cards := []int{1, 14, 2, 15, 3, 16}
	IsPairs := rules.IsPairs(cards)
	require.True(t, IsPairs)
}

func TestIsThreeWithOne(t *testing.T) {
	cards := []int{1, 14, 27, 4}
	fmt.Println("持有的牌是:", generateCardsNames(cards))
	IsThreeWithOne := rules.IsThreeWithOne(cards)
	require.True(t, IsThreeWithOne)
}

func TestIsThreeWithTwo(t *testing.T) {
	cards := []int{1, 14, 27, 4, 17}
	fmt.Println("持有的牌是:", generateCardsNames(cards))
	IsThreeWithTwo := rules.IsThreeWithTwo(cards)
	require.True(t, IsThreeWithTwo)
}

func TestCheckCardsRules(t *testing.T) {
	names := []string{
		"singleCard",
		"twoJokers",
		"doubleCards",
		"bomb_0",
		"bomb_1",
		"threeWithOne",
		"threeWithTwo",
		"straight",
		"straight",
		"pairs",
		"pairs",
		"invalid",
		"invalid",
		"invalid",
		"invalid",
	}

	cardsGroup := [][]int{
		{2},
		{52, 53},
		{3, 20},
		{0, 13, 26, 39},
		{1, 14, 27, 40},
		{1, 14, 27, 5},
		{1, 14, 27, 5, 18},
		{1, 2, 3, 4, 5, 6},
		{1, 15, 29, 4, 18, 6},
		{1, 14, 2, 28, 7, 20},
		{1, 14, 2, 28, 7, 20,3,16,5,18},
		{1, 14, 2, 28, 7, 20, 40, 4},
		{1, 14, 2, 28},
		{1, 14, 27, 40, 7},
		{1, 2, 3, 4, 8, 6},
	}

	for i, cards := range cardsGroup {
		fmt.Println("数字牌是:", cards, "持有的牌是:", generateCardsNames(cards))
		rule := rules.CheckCardsRules(cards)
		fmt.Println("rule is :",rule)
		require.True(t, names[i] == rule)
	}

}

func TestIsStraight(t *testing.T) {
	cards := []int{1, 15, 29, 4, 18, 6}
	fmt.Println("持有的牌是:", generateCardsNames(cards))
	IsStraight := rules.IsStraight(cards)
	require.True(t, IsStraight)
}

func TestFastSort(t *testing.T) {
	// startTime := time.Now()

	// fmt.Println("cards are:", cards)
	// rightResult := utils.Qsort(cards)
	// result := utils.FastSort(cards)
	// endTime := time.Now()
	// totalTime := endTime.Nanosecond() - startTime.Nanosecond()
	// fmt.Println("用时:", totalTime, len(result))
	// fmt.Println("正确答案是:", rightResult)
	// check := utils.CheckSlices(result, rightResult)
	// require.True(t, check)
}

func TestBrutalForceSort(t *testing.T) {
	for count := 0; count < 100; count++ {
		numbers := generateRandomSlice(10, false)
		copy := append([]int{}, numbers...)
		brutalResult := utils.BrutalForceSort(numbers)
		fmt.Println("brutalResult is", brutalResult)
		fmt.Println("numbers are:", numbers)
		qSortResult := utils.QuickSort(copy)
		fmt.Println("qSortResult is", qSortResult)
		isEqual := utils.CheckSlices(brutalResult, qSortResult)
		require.True(t, isEqual)
	}

}

func generateRandomSlice(length int, allowDuplicate bool) []int {
	numbers := make([]int, length)
	if !allowDuplicate {
		record := make(map[int]int, length)

		for i := 0; i < length; i++ {
			for {
				number := rand.Intn(length * 10)
				_, ok := record[number]
				if !ok {
					record[number] = number
					numbers[i] = number
					break
				}
			}
		}

	} else {
		for i := 0; i < length; i++ {
			number := rand.Intn(length * 10)
			numbers[i] = number
		}

	}
	return numbers
}

func generateCardsNames(cards []int) []string {
	cardsNames := make([]string, len(cards))
	for i, card := range cards {
		name := play.NumberToString[card]
		cardsNames[i] = name
	}
	return cardsNames

}
