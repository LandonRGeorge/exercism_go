package resistorcolor

var m = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors should return the list of all colors.
func Colors() []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return m[color]
}
