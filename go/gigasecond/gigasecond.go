package gigasecond

import "time"

const GigaSecond = 1_000_000_000

func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Second * GigaSecond)
}
