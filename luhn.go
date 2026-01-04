package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {

	new_string := strings.ReplaceAll(id, " ", "")

	for _, r := range new_string {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	if len(new_string) <= 1 {
		return false
	}

	for i := len(new_string) - 2; i >= 0; i -= 2 {
		char := new_string[i]
		new_replacement_int := int(char - '0')
		new_replacement_int *= 2
		if new_replacement_int > 9 {
			new_replacement_int -= 9
		}
		replacementstr := strconv.Itoa(new_replacement_int)
		new_string = new_string[:i] + replacementstr + new_string[i+1:]
	}
	sum := 0

	for j := 0; j < len(new_string); j++ {
		char := new_string[j]
		new_replacement_int := int(char - '0')
		sum += new_replacement_int
	}

	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}
