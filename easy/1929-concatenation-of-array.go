package easy

func getConcatenation(nums []int) []int {
	count := 2
	ans := make([]int, 0)
	result := fill(nums, count, ans)
	return result
}

func fill(nums []int, count int, ans []int) []int {
	if count == 0 {
		return ans
	}
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[i])
	}
	count -= 1
	return fill(nums, count, ans)
}
