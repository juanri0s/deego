package date

import "time"

// last accepts a weekday (Saturday) and returns the dt of the last occurrence.
// https://play.golang.org/p/GABxlzkbWXb
func last(wd time.Weekday) time.Time {
	t := sampleTime()
	// Converts day of week to its int form
	// Sunday = 0, Saturday = 7
	cd := int(t.Weekday())
	ld := int(wd)
	diff := ld - cd

	// subtract a week since the requested day is equal to current day
	if diff == 0 {
		return t.AddDate(0, 0, -7)
	}

	// just subtract the days
	if diff < 0 {
		return t.AddDate(0, 0, diff)
	}

	return t.AddDate(0, 0, -(7 - diff))
}
