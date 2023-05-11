package easy

func countNegatives(grid [][]int) int {
	count := 0
	for _, value := range grid {
		for i := 0; i < len(value); i++ {
			if value[i] < 0 {
				count += 1
			}
		}
	}
	return count
}
