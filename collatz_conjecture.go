package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var count int = 0
	if n < 1 {
		return 0, errors.New("Values less than 1 not valid Collatz Conjecture input")
	} else {
		for n > 1 {
			if n%2 == 0 {
				n /= 2
				count++
			} else {
				n *= 3
				n++
				count++
			}

		}
		return count, nil
	}
}
