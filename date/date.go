package date

import (
	"math/rand"
	"strconv"
	"time"
)

// NOTES:
// - For functions that accept strings, if your string doesn't follow the year-month-day
// ISO8601 layout, replace the layout string and it should still work.
// - For functions that accept go time, if you want to use a string, then you will first need to convert
// using time.Parse.
// - If you prefer accepting time over strings, then the conversion in some methods is not needed and only
// the logic needs to be used.

// sampleTime generates a time that is used for testing
func sampleTime() time.Time {
	s := "2019-01-22"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)
	return t
}

// toTime converts a string to a datetime based on a given layout
// layout here is ISO8601 YYYY-MM-DD but the layout can be changed to fit your needs.
// https://play.golang.org/p/P8dLlTWnihY
func toTime(s string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)

	return t
}

// roundDate accepts a datetime and returns the same date with time rounded down.
// https://play.golang.org/p/pr2c_g8nOFH
func roundDate(t time.Time) time.Time {
	// only assign the day, month, year to achieve the time being 00:00:00
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return rounded
	// t -> 2019-09-10 23:00:00 +0000 UTC
	// rounded -> 2019-09-10 00:00:00 +0000 UTC
}

// randomDate generates a random date
// https://play.golang.org/p/ZD1WdZgffaA
func randomDate() time.Time {
	min := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// today returns the current day rounded down to the start of the day.
// https://play.golang.org/p/hwWlvlX-Joq
func today() time.Time {
	t := time.Now()
	// only assign the day, month, year to achieve the time being 00:00:00
	rounded := roundDate(t)
	return rounded
	// t -> 2019-09-10 23:00:00 +0000 UTC
	// rounded -> 2019-09-10 00:00:00 +0000 UTC
}

// tomorrow accepts a dt and returns the next day.
// https://play.golang.org/p/waBNYtVGEjZ
func tomorrow(t time.Time) time.Time {
	// add 1 int day
	// years, months, days
	tomorrow := t.AddDate(0, 0, 1)
	return tomorrow
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// tomorrow -> 2019-09-09 13:23:52.73977 -0400 EDT
}

// yesterday accepts a dt and returns the previous day.
// https://play.golang.org/p/pkt5oNApY3O
func yesterday(t time.Time) time.Time {
	// add 1 int day
	// years, months, days
	yesterday := t.AddDate(0, 0, -1)
	return yesterday
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// tomorrow -> 2019-09-09 13:23:52.73977 -0400 EDT
}

// nextWeek accepts a dt and returns the dt one week later.
// https://play.golang.org/p/LlZIlpaiHJ3
func nextWeek(t time.Time) time.Time {
	// years, months, days
	// 7 days * the number of weeks we want to add
	nextWeek := t.AddDate(0, 0, 7)
	return nextWeek
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// next week -> 2019-09-15 13:23:52.73977 -0400 EDT
}

// isWeekend accepts a dt and returns true if it lands on a weekend, false if it doesn't.
// https://play.golang.org/p/mNX65LSJey4
func isWeekend(t time.Time) bool {
	if (t.Weekday().String() == "Saturday") || (t.Weekday().String() == "Sunday") {
		return true
	}

	return false
}

// isWeekday accepts a dt and returns true if it lands on a weekday, false if it doesn't.
// https://play.golang.org/p/UToSOoUrI0H
func isWeekday(t time.Time) bool {
	if (t.Weekday().String() == "Saturday") || (t.Weekday().String() == "Sunday") {
		return false
	}

	return true
}

// monthLong accepts a dt and returns the full month name.
// https://play.golang.org/p/JDtYIqCipbu
func monthLong(t time.Time) string {
	month := t.Month().String()
	return month
}

// monthShort accepts a dt returns the abbreviated month name.
// https://play.golang.org/p/oVorpqtpb7D
func monthShort(t time.Time) string {
	// We want the abbreviated form so lets take the first 3 letters of the string
	month := t.Month().String()[:3]
	return month
}

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

// next accepts a weekday (Saturday) and returns the dt of the next occurrence.
// https://play.golang.org/p/4BMDqpG-QRH
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
// utc is not possible without an initial tz.
// https://play.golang.org/p/2ITa-xA4YpA
func utc(s string, tz string) time.Time {
	layout := "2006-01-02"

	// normally we would handle the errors here instead of escaping them
	z, _ := time.LoadLocation(tz)
	t, _ := time.ParseInLocation(layout, s, z)

	return t.UTC()
}

// equal accepts two date strings and checks if they are equal.
// https://play.golang.org/p/aErKIb3Eg53
func equal(t1 time.Time, t2 time.Time) bool {
	if t1.Equal(t2) {
		return true
	}

	return false
}

// compare accepts two strings and returns the comparison based on the first parameter
// 1 = greater, -1 = less than, 0 = equal
// https://play.golang.org/p/kHBExOiL1tW
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
// If you want a more robust solution, switch my diff() for Sam's daysBetween()
// which covers which handles the conversion much better.
// https://play.golang.org/p/83--9vlZQt1
func diff(t1 time.Time, t2 time.Time) float64 {
	// We'll get a negative number if subtract a time that's after
	if t2.After(t1) {
		t1, t2 = t2, t1
	}

	diff := t1.Sub(t2)
	// Sub returns hours so we have to divide by 24 if it's more than 24
	if diff.Hours() >= 24 {
		return diff.Hours() / 24
	}

	return diff.Hours()
}

// epochSec accepts a time and returns the number of elapsed seconds from unix epoch.
// https://play.golang.org/p/cq1y6r7qiYd
func epochSec(t time.Time) int64 {
	secs := t.Unix()
	return secs
}

// epochNano accepts a time and returns the number of elapsed nanoseconds from unix epoch
// https://play.golang.org/p/H4t7b3KeW8t
func epochNano(t time.Time) int64 {
	nanos := t.UnixNano()
	return nanos
}

// epochMilli accepts a time and returns the number of elapsed milliseconds from unix epoch
// https://play.golang.org/p/6g_v4v4XvTu
func epochMilli(t time.Time) int64 {
	nanos := t.UnixNano()
	// go does not have UnixMillis, so we have to divide manually from nanos
	// 1000000 nanoseconds in 1 millisecond
	millis := nanos / 1000000
	return millis
}

// daysInYear returns the number of days in a year including leap years
// https://play.golang.org/p/WBacxQ5q7zC
func daysInYear(s string) float64 {
	// gives us the first day of the year
	t1 := startOfYear(s)
	// gives us the last day of the year
	// we need to add 1 day to account for the last day
	t2 := endOfYear(s).AddDate(0, 0, 1)
	// get number of days between the first day and last day of the year
	days := diff(t2, t1)

	return days
}

// startOfYear accepts a year string and returns the start of that given year
// https://play.golang.org/p/ZqlMo8nXSas
func startOfYear(s string) time.Time {
	// convert year string to int
	y, _ := strconv.Atoi(s)
	start := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)
	return start
}

// endOfYear accepts a year string and returns the end of that given year
// https://play.golang.org/p/VOK_B3uezGt
func endOfYear(s string) time.Time {
	// convert year string to int
	y, _ := strconv.Atoi(s)
	end := time.Date(y, time.December, 31, 0, 0, 0, 0, time.UTC)
	return end
}

// diff() can cause issues since it also accounts for time, daysBetween() solves this
// Courtesy of Sam Rose's post on dev.to
// https://dev.to/samwho/get-the-number-of-days-between-two-dates-in-go-5bf3
// daysBetween returns the number of days in between two dt
// https://play.golang.org/p/1bIGdQk25mR
func daysBetween(t1, t2 time.Time) int {
	if t1.After(t2) {
		t1, t2 = t2, t1
	}

	days := -t1.YearDay()
	for year := t1.Year(); year < t2.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += t2.YearDay()

	return days
}
