package date

import (
	"strconv"
	"time"
)

// startOfYear accepts a year string and returns the start of that given year
// https://play.golang.org/p/ZqlMo8nXSas
func startOfYear(s string) time.Time {
	// convert year string to int
	y, _ := strconv.Atoi(s)
	start := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)
	return start
}
