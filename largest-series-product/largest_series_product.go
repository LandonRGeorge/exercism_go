package lsproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) {
		return 0, errors.New("span longer than string length")
	}
	if span < 0 {
		return 0, errors.New("span less than zero")
	}
	var max int64
	for i := 0; i <= len(digits)-span; i++ {
		chars := digits[i : i+span]
		var product int64 = 1
		for _, r := range chars {
			if !unicode.IsNumber(r) {
				return 0, errors.New("not a number")
			}
			nbr := int64(r - '0')
			product *= nbr
		}
		if product <= max {
			continue
		}
		max = product
	}
	return max, nil
}
