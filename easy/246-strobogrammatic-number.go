package easy

import (
	"strings"
)

func isStrobogrammatic(num string) bool {
	stopogrammaticNumbers := map[string]string{"0": "0", "1": "1", "6": "9", "8": "8", "9": "6"}
	sSlice := strings.Split(num, "")
	var rotatedString string
	for i := len(sSlice) - 1; i >= 0; i-- {
		if _, ok := stopogrammaticNumbers[sSlice[i]]; ok {
			rotatedString += stopogrammaticNumbers[sSlice[i]]
		}
	}

	if rotatedString == num {
		return true
	}
	return false
}
