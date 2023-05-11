package easy

import (
	"sort"
)

func frequencySort(nums []int) []int {
	seen := make(map[int]int)

	for _, value := range nums {
		seen[value]++
	}

	sort.Slice(nums, func(i, j int) bool {

		if seen[nums[i]] == seen[nums[j]] {
			return nums[j] < nums[i]
		}

		return seen[nums[i]] < seen[nums[j]]
	})

	return nums
}
