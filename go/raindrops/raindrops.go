package raindrops

import (
	"strconv"
	"strings"
)

const testVersion = 2

func Convert(n int) string {

	// slice to append output, which is Joined later
	var output []string

	if n%3 == 0 {
		output = append(output, "Pling")
	}
	if n%5 == 0 {
		output = append(output, "Plang")
	}
	if n%7 == 0 {
		output = append(output, "Plong")
	}
	if len(output) == 0 {
		output = append(output, strconv.Itoa(n))
	}

	// strings.Join is faster than +=
	return strings.Join(output, "")
}
