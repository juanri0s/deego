package date

import "time"

// isWeekend accepts a dt and returns true if it lands on a weekend, false if it doesn't.
// https://play.golang.org/p/mNX65LSJey4
func isWeekend(t time.Time) bool {
	if (t.Weekday().String() == "Saturday") || (t.Weekday().String() == "Sunday") {
		return true
	}

	return false
}
