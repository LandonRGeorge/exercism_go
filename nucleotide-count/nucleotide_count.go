package dna

import "fmt"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, c := range d {
		if _, ok := h[c]; !ok {
			return nil, fmt.Errorf("%v is not a valid character", c)
		}
		h[c]++
	}
	return h, nil
}
