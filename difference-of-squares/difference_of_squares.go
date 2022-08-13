package diffsquares

import "math"

func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	a := SumOfSquares(n)
	b := SquareOfSum(n)
	return int(math.Abs(float64(a - b)))
}
