package date

import "time"

// isWeekday accepts a dt and returns true if it lands on a weekday, false if it doesn't.
// https://play.golang.org/p/UToSOoUrI0H
func isWeekday(t time.Time) bool {
	if (t.Weekday().String() == "Saturday") || (t.Weekday().String() == "Sunday") {
		return false
	}

	return true
}
