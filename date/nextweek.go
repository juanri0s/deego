package date

import "time"

// nextWeek accepts a dt and returns the dt one week later.
// https://play.golang.org/p/LlZIlpaiHJ3
func nextWeek(t time.Time) time.Time {
	// years, months, days
	// 7 days * the number of weeks we want to add
	nextWeek := t.AddDate(0, 0, 7)
	return nextWeek
	// today -> 2019-09-08 13:23:52.73977 -0400 EDT
	// next week -> 2019-09-15 13:23:52.73977 -0400 EDT
}
