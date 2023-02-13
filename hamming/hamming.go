package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("%q not same length as %q", a, b)
	}
	var count int
	for i, _ := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
