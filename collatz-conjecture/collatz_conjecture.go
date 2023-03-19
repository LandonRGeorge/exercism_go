package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	var count int
	if n < 1 {
		return 0, fmt.Errorf("n must be positive")
	}
	for n != 1 {
		switch {
		case n%2 == 0:
			n = n / 2
		default:
			n = n*3 + 1

		}
		count++
	}
	return count, nil
}
