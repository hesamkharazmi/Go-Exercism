package dna

import "fmt"

type (
	DNA       string
	Histogram map[rune]int
)

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, dna := range d {
		switch dna {
		case 'A':
			h['A']++
		case 'C':
			h['C']++
		case 'G':
			h['G']++
		case 'T':
			h['T']++
		default:
			return h, fmt.Errorf("Synatx in invalid!")
		}
	}

	return h, nil
}
