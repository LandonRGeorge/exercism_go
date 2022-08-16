package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	letters := make(map[rune]struct{})
	type empty struct{}
	for _, char := range word {
		if string(char) == " " || string(char) == "-" {
			continue
		}
		lowerChar := unicode.ToLower(char)
		_, exists := letters[lowerChar]
		if exists {
			return false
		} else {
			letters[lowerChar] = empty{}
		}
	}
	return true
}
