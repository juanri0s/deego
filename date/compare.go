package date

import "time"

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
