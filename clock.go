package clock

import "time"

type nowFuncT func() time.Time

var NowFunc nowFuncT

func init() {
	resetClockImplementation()
}

func resetClockImplementation() {
	NowFunc = func() time.Time {
		return time.Now()
	}
}

// function to return current time stamp
func Now() time.Time {
	return NowFunc()
}
