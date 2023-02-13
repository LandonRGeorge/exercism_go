package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	var sum int
	for i, j := len(id)-1, 1; i >= 0; i, j = i-1, j+1 {
		v, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		if j%2 == 0 {
			v = v * 2
			if v > 9 {
				v = v - 9
			}
		}
		sum += v
	}
	return sum%10 == 0
}
