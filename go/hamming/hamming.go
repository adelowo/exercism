package hamming

import (
	"errors"
)

const testVersion = 5

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Both strings should be of the same length")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if b[i] != a[i] {
			distance++
		}
	}

	return distance, nil
}
