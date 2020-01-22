package mock

import (
	"github.com/jonboulle/clockwork"
	"time"
)

func NewFixedClock() clockwork.Clock {
	now := time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC)
	return clockwork.NewFakeClockAt(now)
}
