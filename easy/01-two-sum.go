package easy

import (
	"fmt"
)

var nums = []int{2, 7, 11, 15}

func twoSum(nums []int, target int) []int {
	itemToIndexMap := make(map[int]int)

	// target = cur + x
	for index, current := range nums {
		itemToIndexMap[current] = index
	}

	for index, current := range nums {
		x := target - current
		if _, ok := itemToIndexMap[x]; ok && itemToIndexMap[x] != index {
			fmt.Println(index, itemToIndexMap[x])
			result := []int{index, itemToIndexMap[x]}
			return result
		}
	}
	return nums
}
