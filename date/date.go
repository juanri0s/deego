package date

import (
	"math/rand"
	"time"
)

// For functions that accept strings, if your string doesn't follow the year-month-day
// ISO8601 layout, replace the layout string and it should still work.
// For functions that accept go time, if you want to use a string, then you will first need to convert
// using time.Parse.
// If you prefer accepting time over strings, then the conversion in some methods is not needed and only
// the logic needs to be used.

// sampleTime generates a time that is used for testing
func sampleTime() time.Time {
	s := "2019-01-22"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)
	return t
}

// toTime converts a string to a datetime based on a given layout
// layout here is ISO8601 YYYY-MM-DD but the layout can be changed to fit your needs
func toTime(s string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)

	return t
}

// roundT accepts a datetime and returns the same date with time rounded down
func roundT(t time.Time) time.Time {
	// only assign the day, month, year to achieve the time being 00:00:00
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return rounded
	// t -> 2019-09-10 23:00:00 +0000 UTC
	// rounded -> 2019-09-10 00:00:00 +0000 UTC
}

// randomDate generates a random date
func randomDate() time.Time {
	min := time.Date(1990, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2030, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// now returns the exact current datetime
func now() time.Time {
	t := time.Now()
	return t
	// 2019-09-08 00:00:00.619429 -0400 EDT
}

// today returns the current day rounded down to the start of the day
func today() time.Time {
	t := time.Now()
	// only assign the day, month, year to achieve the time being 00:00:00
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return rounded
	// t -> 2019-09-10 23:00:00 +0000 UTC
	// rounded -> 2019-09-10 00:00:00 +0000 UTC
}

// tomorrow accepts a dt and returns the next day
func tomorrow(t time.Time) time.Time {
	// add 1 int day
	// years, months, days
	tomorrow := t.AddDate(0, 0, 1)
	return tomorrow
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// tomorrow -> 2019-09-09 13:23:52.73977 -0400 EDT
}

// nextWeek accepts a dt and returns the dt one week later
func nextWeek(t time.Time) time.Time {
	// years, months, days
	// 7 days * the number of weeks we want to add
	nextWeek := t.AddDate(0, 0, 1)
	return nextWeek
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// next week -> 2019-09-15 13:23:52.73977 -0400 EDT
}

// isWeekend accepts a dt and returns true if it lands on a weekend, false if it doesn't.
func isWeekend(t time.Time) bool {
	if (t.Weekday().String() == "Saturday") || (t.Weekday().String() == "Sunday") {
		return true
	}

	return false
}

// isWeekday accepts a dt and returns true if it lands on a weekday, false if it doesn't.
func isWeekday(t time.Time) bool {
	if (t.Weekday().String() == "Saturday") || (t.Weekday().String() == "Sunday") {
		return false
	}

	return true
}

// monthLong accepts a dt and returns the full month name.
func monthLong(t time.Time) string {
	// We want the abbreviated form so lets take the first 3 letters of the string
	month := t.Month().String()

	return month
}

// getMonthShort accepts a dt returns the abbreviated month name.
func monthShort(t time.Time) string {
	// We want the abbreviated form so lets take the first 3 letters of the string
	month := t.Month().String()[:3]

	return month
}

// last accepts a weekday (Saturday) and returns the dt of the last occurence
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

// next accepts a weekday (Saturday) and returns the dt of the next occurence
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

// utc accepts a date string and a timezone and converts the time to utc
// utc is not possible without an initial tz
func utc(s string, tz string) time.Time {
	layout := "2006-01-02"

	// normally we would handle the errors here instead of escaping them
	z, _ := time.LoadLocation(tz)
	t, _ := time.ParseInLocation(layout, s, z)

	return t.UTC()
}

// equal accepts two date strings and checks if they are equal
func equal(t1 time.Time, t2 time.Time) bool {
	if t1 == t2 {
		return true
	}

	return false
}

// compare accepts two strings and returns the comparison based on the first parameter
// 1 = greater, -1 = less than, 0 = equal
func compare(t1 time.Time, t2 time.Time) int {
	if t1.After(t2) {
		return 1
	}

	if t1.Before(t2) {
		return -1
	}

	// can also be checked with ==
	// but since it's the last condition we just return it
	return 0
}

// diff accepts two strings and returns the difference in days
func diff(t1 time.Time, t2 time.Time) float64 {
	diff := t1.Sub(t2)
	// Sub returns hours so we have to divide by 24 if it's more than 24
	if diff.Hours() >= 24 {
		return diff.Hours() / 24
	}

	return diff.Hours()
}

// epochSec accepts a time and returns the number of elapsed seconds from unix epoch
func epochSec(t time.Time) int64 {
	secs := t.Unix()
	return secs
}

// epochNano accepts a time and returns the number of elapsed nanoseconds from unix epoch
func epochNano(t time.Time) int64 {
	nanos := t.UnixNano()
	return nanos
}

// epochMilli accepts a time and returns the number of elapsed milliseconds from unix epoch
func epochMilli(t time.Time) int64 {
	nanos := t.UnixNano()
	// go does not have UnixMillis, so we have to divide manually from nanos
	// 1000000 nanoseconds in 1 millisecond
	millis := nanos / 1000000
	return millis
}
