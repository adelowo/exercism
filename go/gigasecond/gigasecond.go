package gigasecond

import "time"

const (
	testVersion = 4
)

func AddGigasecond(startTime time.Time) time.Time {
	return startTime.Add(time.Second * 1e9)
}
