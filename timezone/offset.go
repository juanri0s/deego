package timezone

import "time"

// offset returns the timezone offset (in seconds)
// from the given dt.
// https://play.golang.org/p/xXDggA2xr9o
func offset(t time.Time) int {
	// o - offset returned in seconds
	_, o := t.Zone()
	// to convert from seconds to hour divide by 3600
	// _, o := now().Zone() / 3600
	return o
	// offset UTC: 0
}
