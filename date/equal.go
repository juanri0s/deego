package date

import "time"

// equal accepts two date strings and checks if they are equal.
// https://play.golang.org/p/aErKIb3Eg53
func equal(t1 time.Time, t2 time.Time) bool {
	if t1.Equal(t2) {
		return true
	}

	return false
}
