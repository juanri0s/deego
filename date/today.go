package date

import "time"

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
