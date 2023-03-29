package grains

import "fmt"

func Square(number int) (uint64, error) {
	if number <= 0 || number >= 65 {
		return 0, fmt.Errorf("%d not be between 1 and 64", number)
	}
	var x uint64 = 1
	for i := 1; i < number; i++ {
		x *= 2
	}
	return x, nil
}

func Total() uint64 {
	var x uint64
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		x += s
	}
	return x
}
