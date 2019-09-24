package date

import "time"

// next accepts a weekday (Saturday) and returns the dt of the next occurrence.
// https://play.golang.org/p/4BMDqpG-QRH
func next(wd time.Weekday) time.Time {
	t := SampleTime()
	// Converts day of week to its int form
	// Sunday = 0, Saturday = 7
	cd := int(t.Weekday())
	ld := int(wd)
	diff := ld - cd

	// add a week since the requested day is equal to current day
	if diff == 0 {
		return t.AddDate(0, 0, 7)
	}

	// just add the days
	if diff > 0 {
		return t.AddDate(0, 0, diff)
	}

	return t.AddDate(0, 0, 7+(diff))
}
