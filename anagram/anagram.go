package anagram

import (
	"bytes"
	"sort"
	"strings"
	"unicode/utf8"
)

func runesToUTF8(rs []rune) []byte {
	bs := make([]byte, len(rs)*utf8.UTFMax)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs[:count]
}

func sorted(text string) []byte {
	rs := []rune(text)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})
	return runesToUTF8(rs)
}

func Detect(subject string, candidates []string) []string {
	sLower := strings.ToLower(subject)
	sSorted := sorted(sLower)
	var matching []string
	for _, c := range candidates {
		cLower := strings.ToLower(c)
		if sLower == cLower {
			continue
		}
		cSorted := sorted(cLower)
		if !bytes.Equal(sSorted, cSorted) {
			continue
		}
		matching = append(matching, c)
	}
	return matching
}
