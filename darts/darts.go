package darts

import "math"

func Score(x, y float64) int {
	hypot := math.Hypot(x, y)
	switch {
	case hypot <= 1:
		return 10
	case hypot <= 5:
		return 5
	case hypot <= 10:
		return 1
	default:
		return 0
	}
	// outer := 10 -> 1 pts
	// middle := 5 -> 5 pts
	// inner := 1 -> 10 pts

}
