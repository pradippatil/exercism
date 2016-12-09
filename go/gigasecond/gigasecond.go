package gigasecond

import "time"

const testVersion = 4

//AddGigasecond returns time 10^9 seconds after given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
