package date

import "time"

// roundDate accepts a datetime and returns the same date with time rounded down.
// https://play.golang.org/p/pr2c_g8nOFH
func roundDate(t time.Time) time.Time {
	// only assign the day, month, year to achieve the time being 00:00:00
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return rounded
	// t -> 2019-09-10 23:00:00 +0000 UTC
	// rounded -> 2019-09-10 00:00:00 +0000 UTC
}
