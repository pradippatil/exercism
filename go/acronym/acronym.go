package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 1

func abbreviate(s string) string {
	var abbr []rune

	// Separate words by spaces and hyphens
	f := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})

	for _, v := range f {
		r := []rune(v)
		for i := range r {
			//First letter of each word and UpperCase letter from word if any
			if i == 0 || !unicode.IsUpper(r[i-1]) && unicode.IsUpper(r[i]) {
				abbr = append(abbr, r[i])
			}
		}

	}
	return strings.ToUpper(string(abbr))
}
