package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var res []string

	for i := 0; i < len(rhyme); i++ {
		if i == len(rhyme)-1 {
			res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			break
		}
		res = append(res, fmt.Sprintf(
			"For want of a %s the %s was lost.",
			rhyme[i],
			rhyme[i+1],
		))
	}
	return res

}
