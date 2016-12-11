package raindrops

import "strconv"

const testVersion = 2

func Convert(n int) string {
	var output string

	isPling := n % 3
	isPlang := n % 5
	isPlong := n % 7

	if isPling == 0 {
		output += "Pling"
	}
	if isPlang == 0 {
		output += "Plang"
	}
	if isPlong == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(n)
	}

	return output
}
