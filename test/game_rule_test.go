package test

import (
	"fmt"
	"testing"

	"github.com/bo-er/poker/rules"
	"github.com/stretchr/testify/require"
)

var cardsGroup [][]int

func init() {
	cardsGroup = [][]int{
		{1},                                 //0
		{2},                                 //1
		{52, 53},                            //2
		{3, 20},                             //3
		{0, 13, 26, 39},                     //4
		{1, 27, 14, 40},                     //5
		{1, 14, 27, 5},                      //6
		{1, 14, 27, 5, 18},                  //7
		{1, 2, 3, 4, 5, 6},                  //8
		{1, 15, 29, 4, 18, 6},               //9
		{1, 14, 2, 28, 7, 20},               //10
		{1, 14, 2, 28, 7, 20, 3, 16, 5, 18}, //11
		{1, 14, 2, 28, 7, 20, 40, 4},        //12
		{1, 14, 2, 28},                      //13
		{1, 14, 27, 40, 7},                  //14
		{1, 2, 3, 4, 8, 6},                  //15
	}

	for i := 0; i < len(cardsGroup); i++ {
		fmt.Println(i, generateCardsNames(cardsGroup[i]))
	}
}

func TestCheckIfGameRuleOk(t *testing.T) {
	previousType := rules.CheckCardsRules(cardsGroup[4])
	fmt.Println(cardsGroup[4])
	laterType := rules.CheckCardsRules(cardsGroup[5])
	fmt.Println(cardsGroup[5])
	require.Equal(t, previousType, laterType)
}
