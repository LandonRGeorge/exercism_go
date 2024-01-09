package armstrong

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IsNumber(n int) bool {
	nStr := strings.Split(fmt.Sprintf("%d", n), "")
	nStrLen := len(nStr)
	var cumVal float64
	for i := 0; i < len(nStr); i++ {
		nbr, _ := strconv.Atoi(nStr[i])
		cumVal += math.Pow(float64(nbr), float64(nStrLen))
	}
	return n == int(cumVal)
}
