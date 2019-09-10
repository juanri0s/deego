package date

import "time"

// utc accepts a date string and a timezone and converts the time to utc
// utc is not possible without an initial tz.
// https://play.golang.org/p/2ITa-xA4YpA
func utc(s string, tz string) time.Time {
	layout := "2006-01-02"

	// normally we would handle the errors here instead of escaping them
	z, _ := time.LoadLocation(tz)
	t, _ := time.ParseInLocation(layout, s, z)

	return t.UTC()
}
