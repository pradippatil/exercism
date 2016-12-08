package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	// Calculate total minutes
	tmin := hour*60 + minute
	// Discard time more than 24 hours
	rmin := tmin % (24 * 60)

	// If total minutes is negative, rewind the clock (subtract from 24 hours)
	if rmin < 0 {
		rmin += (24 * 60)
	}

	// Convert minutes to hours
	h := (rmin) / 60
	m := (rmin) % 60

	return Clock{h, m}
}

func (c Clock) String() string {
	//Padding 2 zeros
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
