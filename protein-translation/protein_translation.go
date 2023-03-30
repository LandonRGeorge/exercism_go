package protein

import "errors"

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid base")
)

var cMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var cs []string
	for i := 3; i <= len(rna); i += 3 {
		c, err := FromCodon(rna[i-3 : i])
		if err != nil {
			if errors.Is(err, ErrStop) {
				return cs, nil
			}
			return []string{}, err
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func FromCodon(codon string) (string, error) {
	p, ok := cMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if p == "STOP" {
		return "", ErrStop
	}
	return p, nil
}
