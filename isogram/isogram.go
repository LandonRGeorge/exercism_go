package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	letters := make(map[string]struct{})
	wordLower := strings.ToLower(word)
	type empty struct{}
	var charString string
	for _, char := range wordLower {
		charString = string(char)
		if charString == "-" || charString == " " {
			continue
		}
		_, exists := letters[charString]
		if exists {
			return false
		} else {
			letters[charString] = empty{}
		}
	}
	return true
}
