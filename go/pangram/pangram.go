package pangram

import "unicode"

const testVersion = 1

//IsPangram checks if given string is pangram
func IsPangram(s string) bool {

	//empty string
	if len(s) == 0 {
		return false
	}

	//ASCII A-Z are 26 letters in range 65-90
	//Uppercase all letters and check unique count using map
	a := make(map[rune]bool)
	for _, v := range s {
		v := unicode.ToUpper(v)
		if v >= 65 && v <= 90 {
			a[v] = true
		}
	}
	if len(a) == 26 {
		return true
	}
	return false

}
