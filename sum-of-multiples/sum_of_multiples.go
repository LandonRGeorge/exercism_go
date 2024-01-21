package summultiples

func getMultiples(limit int, divisor int) (arr []int) {
	for i := divisor; i < limit && i > 0; i++ {
		if i%divisor == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func SumMultiples(limit int, divisors ...int) (sum int) {
	pm := make(map[int]struct{})
	for _, d := range divisors {
		for _, e := range getMultiples(limit, d) {
			if _, ok := pm[e]; !ok {
				pm[e] = struct{}{}
				sum += e
			}
		}
	}
	return sum
}
