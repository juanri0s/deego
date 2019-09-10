package date

import "time"

// monthShort accepts a dt returns the abbreviated month name.
// https://play.golang.org/p/oVorpqtpb7D
func monthShort(t time.Time) string {
	// We want the abbreviated form so lets take the first 3 letters of the string
	month := t.Month().String()[:3]
	return month
}
