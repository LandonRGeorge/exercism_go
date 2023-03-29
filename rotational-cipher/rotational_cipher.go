package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	rs := []rune(plain)
	for i, r := range rs {
		if !unicode.IsLetter(r) {
			continue
		}
		base := 'a'
		if unicode.IsUpper(r) {
			base = 'A'
		}
		var rShifted rune
		rDiffFromBase := r - base
		shiftFromBase := (rDiffFromBase + int32(shiftKey)) % 26
		rShifted = base + shiftFromBase
		rs[i] = rShifted
	}
	return string(rs)
}
