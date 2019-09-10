package date

import "time"

// monthLong accepts a dt and returns the full month name.
// https://play.golang.org/p/JDtYIqCipbu
func monthLong(t time.Time) string {
	month := t.Month().String()
	return month
}
