package hamming

import "errors"

func Distance(a, b string) (Dis int, err error) {
	Dis = 0
	if len(a) != len(b) {
		return -1, errors.New("Input strings must have same length.")
	}
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			Dis++
		}
	}
	return Dis, err
	//panic("Implement the Distance function")
}
