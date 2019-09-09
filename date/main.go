package date

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome from deego/date")
}

var Now = time.Now

// now returns the current system datetime, mainly used for the other datetime functions.
func now() time.Time {
	return time.Now()
}


// tomorrow retuns the datetime of the next day from the datetime given
func tomorrow(t time.Time) time.Time {
	// add 1 int day
	// years, months, days
	to := t.AddDate(0, 0, 1)
	return to
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// tomorrow -> 2019-09-09 13:23:52.73977 -0400 EDT
}

// nextWeek returns the datetime of one week from the datetime given
func nextWeek(t time.Time) time.Time {
	// years, months, days
	// 7 days * the number of weeks we want to add
	nw := t.AddDate(0, 0, 1)
	return nw
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// next week -> 2019-09-15 13:23:52.73977 -0400 EDT
}

// isWeekend returns true if a given datetime lands on a weekend, false when it doesn't.
func isWeekend(s string) bool {
	dt, _ := time.Parse("2019-09-10 15:04:05", s)

	if (dt.Weekday().String() == "Saturday") || (dt.Weekday().String() == "Sunday") {
		return true
	}

	return false
}

// isWeekday returns true if a given datetime lands on a weekday, false when it doesn't.
func isWeekday(s string) bool {
	dt, _ := time.Parse("2019-09-10 15:04:05", s)

	if (dt.Weekday().String() != "Saturday") || (dt.Weekday().String() != "Sunday") {
		return true
	}

	return false
}

// getMonthLong returns the full month of a given datetime
func getMonthLong(s string) string {
	dt, _ := time.Parse("2019-09-10 15:04:05", s)
	// We want the abbreviated form so lets take the first 3 letters of the string
	m := dt.Month().String()

	return m
}

// getMonthShort returns the abbreviated month of a given datetime
func getMonthShort(s string) string {
	dt, _ := time.Parse("2019-09-10 15:04:05", s)
	// We want the abbreviated form so lets take the first 3 letters of the string
	m := dt.Month().String()[:3]

	return m
}

// last returns the datetime from the last day that was given, for example last Thursday
func last(s time.Weekday) time.Time {
	t := now()

	// Converts day of week to its int form
	// Sunday = 0, Saturday = 6
	day := int(t.Weekday())
	ds := int(s)
	diff := ds - day

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
func next(s time.Weekday) time.Time {
	t := now()

	// Converts day of week to its int form
	// Sunday = 0, Saturday = 6
	day := int(t.Weekday())
	ds := int(s)
	diff := ds - day

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
