package cipher

import (
	"strings"
)

func clean(input string) string {
	input = strings.ToLower(input)
	cleanFunc := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r
		}
		return -1
	}
	return strings.Map(cleanFunc, input)
}

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return caesar{}
}

func (c caesar) Encode(input string) string {
	input = clean(input)
	rot3 := func(r rune) rune {
		return 'a' + (r-'a'+3)%26
	}
	return strings.Map(rot3, input)
}

func (c caesar) Decode(input string) string {
	input = clean(input)
	rot3 := func(r rune) rune {
		return 'z' - ('z'-r+3)%26

	}
	return strings.Map(rot3, input)
}

func NewShift(distance int) Cipher {
	if !(distance >= 1 && distance <= 25) && !(distance <= -1 && distance >= -25) {
		return nil
	}
	return shift{distance}
}

func (c shift) Encode(input string) string {
	input = clean(input)
	rot := func(r rune) rune {
		return rune('a' + (((int(r-'a')+c.distance)%26)+26)%26)
	}
	return strings.Map(rot, input)
}

func (c shift) Decode(input string) string {
	input = clean(input)
	rot := func(r rune) rune {
		return rune('z' - (((int('z'-r)+c.distance)%26)+26)%26)
	}
	return strings.Map(rot, input)
}

func NewVigenere(key string) Cipher {
	for _, r := range key {
		if !(r >= 'a' && r <= 'z') {
			return nil
		}
	}
	if strings.Count(key, "a") == len(key) {
		return nil
	}
	return vigenere{key}
}

func expandKey(key, input string) string {
	keyLen := len(key)
	expandedKey := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		idx := i % keyLen
		v := key[idx]
		expandedKey[i] = string(v)
	}
	return strings.Join(expandedKey, "")
}

func (v vigenere) Encode(input string) string {
	input = clean(input)
	expandedKey := expandKey(v.key, input)
	var res []rune
	for i, r := range input {
		diff := int(expandedKey[i] - 'a')
		v := rune('a' + (((int(r-'a')+diff)%26)+26)%26)
		res = append(res, v)
	}
	return string(res)
}

func (v vigenere) Decode(input string) string {
	input = clean(input)
	expandedKey := expandKey(v.key, input)
	var res []rune
	for i, r := range input {
		diff := int(expandedKey[i] - 'a')
		v := rune('z' - (((int('z'-r)+diff)%26)+26)%26)
		res = append(res, v)
	}
	return string(res)
}
