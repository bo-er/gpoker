package utils

import (
	"fmt"
	"strings"
)

// IntArrayToString 将[]int{1,2,3,4,5}转化为"1,2,3,4,5",分隔符delim可以自定
func IntArrayToString(intArray []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(intArray), " ", delim, -1), "[]")
}

// SwitchArrayElement 将切片两个不同位置的元素交换位置
func SwitchArrayElement(elements interface{}, positionA int, positionB int) {
	switch elements.(type) {
	case []int:
		e := elements.([]int)
		eleA := e[positionA]
		eleB := e[positionB]
		e[positionA] = eleB
		e[positionB] = eleA
	case []string:
		e := elements.([]string)
		eleA := e[positionA]
		eleB := e[positionB]
		e[positionA] = eleB
		e[positionB] = eleA
	}

}
