package strand

import "bytes"

var m = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var b bytes.Buffer
	for _, oldChar := range dna {
		newChar := m[oldChar]
		b.WriteRune(newChar)
	}
	return b.String()
}
