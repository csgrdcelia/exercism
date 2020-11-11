package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	distance := 0

	if (len(a) != len(b)) {
		return -1, errors.New("a & b should have same length")
	}

	for index := range a {
		if a[index] != b[index] {
			distance++
		}
	}

	return distance, nil
}
