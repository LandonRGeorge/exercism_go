package atbash

import (
	"unicode"
)

func AtbashMap() map[rune]rune {
	m := make(map[rune]rune)
	var a []rune
	for r := 'a'; r <= 'z'; r++ {
		a = append(a, r)
	}
	for i := 0; i < len(a); i++ {
		k := a[i]
		v := a[len(a)-1-i]
		m[k] = v
	}
	return m
}

func Atbash(s string) string {
	m := AtbashMap()
	var rs []rune
	for _, r := range s {
		if unicode.IsNumber(r) {
			rs = append(rs, r)
		}
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		newR := m[r]
		rs = append(rs, newR)
	}
	var rs2 []rune
	for i, r := range rs {
		if i > 0 && i%5 == 0 {
			rs2 = append(rs2, ' ')
		}
		rs2 = append(rs2, r)
	}
	return string(rs2)
}
