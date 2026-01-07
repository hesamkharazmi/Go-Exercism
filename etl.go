package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)
	for i, str := range in {
		for _, val := range str {
			val = strings.ToLower(val)
			res[val] = i
		}
	}
	return res
}
