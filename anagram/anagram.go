package anagram

import (
	"sort"
	"strings"
)

func sorted(text string) []rune {
	rs := []rune(text)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})
	return rs
}

func isEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
		if !isEqual(sSorted, cSorted) {
			continue
		}
		matching = append(matching, c)
	}
	return matching
}
