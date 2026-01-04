package raindrops

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	result := ""
	counter := 0
	if number%3 == 0 {
		counter++
		if number%5 == 0 {
			if number%7 == 0 {
				result = "PlingPlangPlong"
			} else {
				result = "PlingPlang"
			}
		} else if number%7 == 0 {
			result = "PlingPlong"
		} else {
			result = "Pling"
		}
	} else if number%5 == 0 {
		counter++
		if number%7 == 0 {
			result = "PlangPlong"
		} else {
			result = "Plang"
		}
	} else if number%7 == 0 {
		counter++
		result = "Plong"
	}
	if counter != 0 {
		return result
	}
	result = strconv.Itoa(number)
	return fmt.Sprintf("%s", result)
}
