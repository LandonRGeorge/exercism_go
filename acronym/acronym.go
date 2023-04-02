package acronym

import (
	"unicode"
)

func Abbreviate(s string) string {
	var rs []rune
	var isSpace bool
	for i, r := range s {
		if !unicode.IsLetter(r) {
			if unicode.IsSpace(r) || r == '-' {
				isSpace = true
			}
			continue
		}
		if isSpace || i == 0 {
			rs = append(rs, unicode.ToUpper(r))
			isSpace = false
		}
	}
	return string(rs)
}
