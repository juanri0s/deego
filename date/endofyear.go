package date

import (
	"strconv"
	"time"
)

// endOfYear accepts a year string and returns the end of that given year
// https://play.golang.org/p/VOK_B3uezGt
func endOfYear(s string) time.Time {
	// convert year string to int
	y, _ := strconv.Atoi(s)
	end := time.Date(y, time.December, 31, 0, 0, 0, 0, time.UTC)
	return end
}
