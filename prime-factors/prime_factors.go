package prime

func Factors(n int64) (fs []int64) {
	var divisor int64 = 2
	remaining := n

	for remaining > 1 {
		mod := remaining % divisor
		if mod != 0 {
			divisor++
			continue
		}
		fs = append(fs, divisor)
		remaining = remaining / divisor
	}
	return fs
}
