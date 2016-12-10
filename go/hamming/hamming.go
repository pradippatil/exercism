package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {

	var c int
	lengthErr := errors.New("Strings length mismatch")

	if len(a) != len(b) {
		return 0, lengthErr
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			c++
		}
	}
	return c, nil
}
