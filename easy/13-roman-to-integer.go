package easy

import (
	"strings"
)

func romanToInt(s string) int {
	var result int = 0
	var twoSum int = 0
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
	newSt := strings.Split(s, "")

	for i := 0; i < len(newSt); i++ {
		if i < len(newSt)-1 {
			if romanNumerals[newSt[i]] < romanNumerals[newSt[i+1]] {
				twoSum = romanNumerals[newSt[i+1]] - romanNumerals[newSt[i]]
				result = result + twoSum
				if i <= len(newSt) {
					i = i + 1
				}
			} else {
				result = result + romanNumerals[newSt[i]]
			}
		} else {
			result = result + romanNumerals[newSt[i]]
			break
		}
	}
	return result
}
