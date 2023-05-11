package easy

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, x := range nums {
		if set[x] {
			return true
		}
		set[x] = true
	}

	return false
}
