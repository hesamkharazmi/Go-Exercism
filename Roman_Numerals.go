package romannumerals

import (
	"errors"
)

var values = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("")
	}

	final_str := ""

	for _, key := range keys {
		if input < key {
			continue
		}
		for input >= key {
			input -= key
			final_str += values[key]
		}
	}
	return final_str, nil
}
