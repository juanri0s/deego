package date

import "time"

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
