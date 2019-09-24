package date

import "time"

// diff accepts two strings and returns the difference in days
// If you want a more robust solution, switch my diff() for Sam's daysBetween()
// which covers which handles the conversion much better.
// https://play.golang.org/p/83--9vlZQt1
func diff(t1 time.Time, t2 time.Time) float64 {
	// We'll get a negative number if subtract a time that's after
	if t2.After(t1) {
		t1, t2 = t2, t1
	}

	diff := t1.Sub(t2)
	// Sub returns hours so we have to divide by 24 if it's more than 24
	if diff.Hours() >= 24 {
		return diff.Hours() / 24
	}

	return diff.Hours()
}

// diff() can cause issues since it also accounts for time, daysBetween() solves this
// Courtesy of Sam Rose's post on dev.to
// https://dev.to/samwho/get-the-number-of-days-between-two-dates-in-go-5bf3
// daysBetween returns the number of days in between two dt
// https://play.golang.org/p/1bIGdQk25mR
func daysBetween(t1, t2 time.Time) float64 {
	if t1.After(t2) {
		t1, t2 = t2, t1
	}

	days := -t1.YearDay()
	for year := t1.Year(); year < t2.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += t2.YearDay()

	return float64(days)
}
