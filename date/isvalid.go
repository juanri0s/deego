package date

import "time"

// isValid checks whether a time is valid against Go's zero date
// 0001-01-01 00:00:00 +0000 UTC
// https://play.golang.org/p/QtMzBDVYHpB
func isValid(t time.Time) bool {
	if t.IsZero() {
		return false
	}

	return true
}
