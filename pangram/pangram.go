package pangram

import (
	"unicode"
)

const (
	lettersInAlphabet = 26
)

func IsPangram(input string) bool {
	m := make(map[rune]struct{})
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		m[unicode.ToLower(r)] = struct{}{}
	}
	return len(m) == lettersInAlphabet
}
