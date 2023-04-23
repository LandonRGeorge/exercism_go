package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`\b[\w']+\b`)

func WordCount(phrase string) Frequency {
	freqMap := make(Frequency)
	strs := re.FindAllString(strings.ToLower(phrase), -1)
	for _, str := range strs {
		freqMap[str]++
	}
	return freqMap
}
