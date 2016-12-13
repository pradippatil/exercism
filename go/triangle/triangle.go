package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

type Kind struct {
	name string
}

var NaT = Kind{"NaT"}
var Iso = Kind{"Iso"}
var Sca = Kind{"Sca"}
var Equ = Kind{"Equ"}

// KindFromSides returns type of triangle for given sides
func KindFromSides(a, b, c float64) Kind {

	// Check for Not-a-Number and Infinity values
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	// Triangle inequality theorem (sum of two sides should be greater than third)
	// Better to sort the sides by length, so that we can just compare sum of smallest pair
	// Exception here is a 'line' is allowed (sum of two sides equal to third)
	var sides sort.Float64Slice = []float64{a, b, c}
	sides.Sort()
	if sides[0]+sides[1] < sides[2] {
		return NaT
	}

	// Check for -ve values
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// If all sides are equal
	if a == b && b == c {
		return Equ
	}

	// If any two sides are equal
	if a == b || b == c || c == a {
		return Iso
	}

	// If none of the sides are equal
	if a != b && b != c && c != a {
		return Sca
	}

	return NaT

}
