package series

func All(n int, s string) []string {
	var strs []string
	for i := 0; i <= len(s)-n; i++ {
		strs = append(strs, s[i:i+n])
	}
	return strs
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return s, false
	}
	return s[:n], true
}
