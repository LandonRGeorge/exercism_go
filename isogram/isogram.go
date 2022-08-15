package isogram

import (
    "strings"
)

func IsIsogram(word string) bool {
	letters := make(map[string]int)
	var charString string
	for _, char := range strings.ToLower(word) {
		charString = string(char)
		if charString == "-" || charString == " " {
			continue
		}
		letters[charString] += 1
		if letters[charString] > 1 {
			return false
		}
	}
	return true
}
