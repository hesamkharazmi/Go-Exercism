package darts

import "math"

func Score(x, y float64) int {
	res := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	if res <= 1 {
		return 10
	} else if res > 1 && res <= 5 {
		return 5
	} else if res > 5 && res <= 10 {
		return 1
	}
	return 0
}
