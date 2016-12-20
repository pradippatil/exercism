package twelve

import (
	"fmt"
)

const testVersion = 1

var day = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gift = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Song() string {
	var song string
	for i := 0; i < len(day); i++ {
		song += fmt.Sprintf("%s\n", Verse(i+1))
	}
	return song
}

func Verse(i int) string {
	var gifts string
	if i == 1 {
		gifts = ", " + gift[i-1]

	} else if i > 1 {
		for j := i; j > 0; j-- {
			if j == 1 {
				gifts += fmt.Sprintf(", and %s", gift[j-1])
			} else {
				gifts += fmt.Sprintf(", %s", gift[j-1])
			}
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me%s.", day[i-1], gifts)
}
