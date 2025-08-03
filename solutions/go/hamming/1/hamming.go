package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hammerTime := 0

	if a == b { // if identical hamming distance is zero
		return hammerTime, nil
	} else if len(a) != len(b) {
		return hammerTime, errors.New("bad dna")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammerTime++
			}
		}
	}

	return hammerTime, nil
}
