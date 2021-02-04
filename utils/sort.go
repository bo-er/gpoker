package utils

import (
	"math/rand"
)

var count int

func FastSort(numbers []int) []int {
	pivot := rand.Intn(len(numbers))
	// id := count
	frontPart := numbers[:pivot]
	// fmt.Println("前面的数组是:", frontPart)
	backPart := numbers[pivot:]
	// fmt.Println("后面的数组是:", backPart)
	if len(frontPart) > 2 {
		frontPart = FastSort(frontPart)
	}
	if len(backPart) > 2 {
		backPart = FastSort(backPart)
	}
	if len(frontPart) == 2 && (frontPart[0] > frontPart[1]) {
		frontPart[0], frontPart[1] = frontPart[1], frontPart[0]

	}
	if len(backPart) == 2 && (backPart[0] > backPart[1]) {
		backPart[0], backPart[1] = backPart[1], backPart[0]

	}
	// fmt.Println(id, "前面的数组是:", frontPart)

	// fmt.Println(id, "后面的数组是:", backPart)

	if len(frontPart) > len(backPart) {
		result := MergeTwoSliceV2(frontPart, backPart)
		// fmt.Println(id, "合并结果是:", result)
		return result
	}
	result := MergeTwoSliceV2(backPart, frontPart)

	// fmt.Println(id, "合并结果是:", result)
	return result

}

// BrutalForceSort 暴力排序
func BrutalForceSort(elements []int) []int {

	length := len(elements)
	result := make([]int, length)
	for i := 0; i < length; i++ {
		index := 0
		pivot := elements[0]
		for j := 0; j < len(elements); j++ {
			if pivot > elements[j] {
				pivot = elements[j]
				index = j
			}
			count++
		}
		result[i] = pivot
		elements = append(elements[:index], elements[index+1:]...)
	}

	return result
}

func MergeTwoSliceV2(bigger []int, lesser []int) []int {
	index := 0
	for i := 0; i < len(lesser); i++ {
		for j := index; j < len(bigger); j++ {
			index++
			count++
			if bigger[j] < lesser[i] {
				count++
				if j == len(bigger)-1 {
					bigger = append(bigger, lesser[i:]...)
					return bigger
				}
			} else {
				temp := make([]int, 0)
				temp = append(temp, bigger[:j]...)
				temp = append(temp, lesser[i])
				bigger = append(temp, bigger[j:]...)
				break
			}

		}
	}
	return bigger
}

// Qsort是另外一个快速排序的例子
func Qsort(ori []int) []int {

	copy := append([]int{}, ori...)

	var inner func(ori []int)
	inner = func(ori []int) {

		if len(ori) == 0 || len(ori) == 1 {
			return
		}
		//找一个参考点最左边 元素个数>=2之后进行比较
		ref := ori[0]
		var i, j int
		loopj := true
		for i, j = 0, len(ori)-1; i != j; {
			if loopj {
				if ori[j] < ref {
					ori[i] = ori[j]
					i++ //该位置已经替换需要更新
					loopj = false
				} else {
					j--
				}
			} else {
				if ori[i] > ref {
					ori[j] = ori[i]
					j--
					loopj = true
				} else {
					i++
				}
			}
		}
		ori[i] = ref

		inner(ori[0:i])
		//if i < len(ori) {
		inner(ori[i+1:]) //equal ori[i+1:len(ori)]
		//	}
	}
	inner(copy)
	return copy

}

//___________________________________________________________________________________

// QuickSort 菜鸟教程的例子
func QuickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
