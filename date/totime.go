package date

import "time"

// toTime converts a string to a datetime based on a given layout
// layout here is ISO8601 YYYY-MM-DD but the layout can be changed to fit your needs.
// https://play.golang.org/p/P8dLlTWnihY
func toTime(s string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)

	return t
}
