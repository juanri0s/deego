package date

import "time"

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
