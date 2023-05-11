package easy

import (
	"strings"
)

func firstUniqChar(s string) int {
	sSlice := strings.Split(s, "")
	freq := map[string]int{}
	var result int

	for _, value := range sSlice {
		freq[value]++
	}

	for _, value := range sSlice {
		if freq[value] == 1 {
			result = strings.Index(s, value)
			break
		}
		result = -1
	}
	return result
}
