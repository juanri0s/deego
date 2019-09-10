package date

import (
	"time"
)

func sampleTime() time.Time {
	s := "2019-01-22"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)
	return t
}

// now returns the current datetime
func now() time.Time {
	t := time.Now()
	return t
	// 2019-09-08 00:00:00.619429 -0400 EDT
}

// tomorrow retuns the datetime of the next day from the datetime given
func tomorrow(t time.Time) time.Time {
	// add 1 int day
	// years, months, days
	tomorrow := t.AddDate(0, 0, 1)
	return tomorrow
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// tomorrow -> 2019-09-09 13:23:52.73977 -0400 EDT
}

// nextWeek returns the datetime of one week from the datetime given
func nextWeek(t time.Time) time.Time {
	// years, months, days
	// 7 days * the number of weeks we want to add
	nextWeek := t.AddDate(0, 0, 1)
	return nextWeek
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// next week -> 2019-09-15 13:23:52.73977 -0400 EDT
}

// isWeekend returns true if a given datetime lands on a weekend, false when it doesn't.
func isWeekend(t time.Time) bool {
	if (t.Weekday().String() == "Saturday") || (t.Weekday().String() == "Sunday") {
		return true
	}

	return false
}

// isWeekday returns true if a given datetime lands on a weekday, false when it doesn't.
func isWeekday(t time.Time) bool {
	if (t.Weekday().String() != "Saturday") || (t.Weekday().String() != "Sunday") {
		return true
	}

	return false
}

// getMonthLong returns the full month of a given datetime
func getMonthLong(t time.Time) string {
	// We want the abbreviated form so lets take the first 3 letters of the string
	month := t.Month().String()

	return month
}

// getMonthShort returns the abbreviated month of a given datetime
func getMonthShort(t time.Time) string {
	// We want the abbreviated form so lets take the first 3 letters of the string
	month := t.Month().String()[:3]

	return month
}

// last returns the datetime from the last day that was given, for example last Thursday
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

// next returns the dt of the next given day, for example next Thursday
func next(wd time.Weekday) time.Time {
	t := sampleTime()
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
