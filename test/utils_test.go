package test

import (
	"fmt"
	"testing"

	"github.com/bo-er/poker/utils"
)

func TestMapToString(t *testing.T) {
	m := map[string]int{
		"steve": 1,
		"eve":   2,
		"jack":  3,
	}
	result := utils.MapToString(m)
	fmt.Println(result)
}
