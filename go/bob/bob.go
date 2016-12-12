package bob

import (
	"regexp"
	"strings"
)

const testVersion = 2

func Hey(in string) string {

	//compile regex
	question := regexp.MustCompile(`.*\?\s*$`)
	yell := regexp.MustCompile(`[A-Z]`)

	switch {

	//silence
	case strings.TrimSpace(in) == "":
		return "Fine. Be that way!"

		//yell
	case yell.MatchString(in) && strings.ToUpper(in) == in:
		return "Whoa, chill out!"

		//question
	case question.MatchString(in):
		return "Sure."

		//unkown
	default:
		return "Whatever."

	}

}
