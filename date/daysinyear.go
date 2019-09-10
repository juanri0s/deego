package date

// daysInYear returns the number of days in a year including leap years
// https://play.golang.org/p/WBacxQ5q7zC
func daysInYear(s string) float64 {
	// gives us the first day of the year
	t1 := startOfYear(s)
	// gives us the last day of the year
	// we need to add 1 day to account for the last day
	t2 := endOfYear(s).AddDate(0, 0, 1)
	// get number of days between the first day and last day of the year
	days := diff(t2, t1)

	return days
}
