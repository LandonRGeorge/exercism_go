package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freqMap := make(Frequency)
	fields := strings.FieldsFunc(strings.ToLower(phrase), func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})
	for _, f := range fields {
		f = strings.TrimFunc(f, func(r rune) bool {
			return !(unicode.IsLetter(r) || unicode.IsDigit(r))
		})
		freqMap[f]++
	}
	return freqMap
}
