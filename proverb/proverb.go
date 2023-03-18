package proverb

import (
	"fmt"
)

func Proverb(rhyme []string) []string {
	lines := make([]string, len(rhyme))
	if len(rhyme) == 0 {
		return lines
	}
	for i := 0; i < len(rhyme)-1; i++ {
		lines[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	ending := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	lines[len(rhyme)-1] = ending
	return lines
}
