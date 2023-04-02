package prime

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("there is no zeroth prime")
	}
	var counter int
	for i := 2; counter <= n; i++ {
		if !isPrime(i) {
			continue
		}
		counter++
		if counter == n {
			return i, nil
		}
	}
	return 0, nil
}
