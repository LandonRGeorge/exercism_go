package anagram

import (
	"sort"
	"strings"
)

func sorted(text string) string {
	r := []rune(text)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func isAnagram(subject, candidate string) bool {
	subject = strings.ToLower(subject)
	candidate = strings.ToLower(candidate)
	if subject == candidate {
		return false
	}
	return sorted(subject) == sorted(candidate)
}

func Detect(subject string, candidates []string) []string {
	var matching []string
	for _, c := range candidates {
		if match := isAnagram(subject, c); match {
			matching = append(matching, c)
		}
	}
	return matching
}
