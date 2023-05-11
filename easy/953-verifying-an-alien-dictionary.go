package easy

import (
	"strings"
)

func isAlienSorted(words []string, order string) bool {
	orderMap := map[string]int{}
	var word []string
	var nextWord []string
	for key, value := range strings.Split(order, "") {
		orderMap[value] = key
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			word = strings.Split(words[i], "")
			nextWord = strings.Split(words[i+1], "")

			if j >= len(words[i+1]) {
				return false
			}

			if orderMap[word[j]] != orderMap[nextWord[j]] {
				if orderMap[word[j]] > orderMap[nextWord[j]] {
					return false
				}
				break
			}
		}
	}
	return true
}
