package isogram

import (
	"fmt"
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	charsReplace := []string{" ", "-"}
	for _, charReplace := range charsReplace {
		word = strings.ReplaceAll(word, charReplace, "")
	}

	letters := make(map[rune]struct{})
	type empty struct{}
	for _, char := range word {
		lowerChar := unicode.ToLower(char)
		fmt.Println(lowerChar)
		_, exists := letters[lowerChar]
		if exists {
			return false
		} else {
			letters[lowerChar] = empty{}
		}
	}
	return true
}
