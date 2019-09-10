package timezone

import (
	"time"
)

// NOTES:
// - .Zone() returns both timezone and offset depending on the
// function, I ignore one of the two to keep the return simple.
// - time.Now() defaults to UTC time, so the offset is going to
// always be 0 unless the timezone is specified

// sampleTime generates a time that is used for testing.
func sampleTime() time.Time {
	s := "2019-01-22"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)
	return t
}

// currTimezone returns the timezone from the current dt.
// https://play.golang.org/p/-W0J32E3DKf
func currTimezone() string {
	z, _ := time.Now().Zone()
	return z
	// Current timezone: EST
}

// timezone returns the timezone from a given dt.
// https://play.golang.org/p/skYyMq9dne6
func timezone(t time.Time) string {
	// z - timezone returned as the zone
	z, _ := t.Zone()
	return z
	// timezone: EST
}

// currOffset returns the timezone offset (in seconds)
// from the current dt.
// https://play.golang.org/p/d5NPpJmu801
func currOffset() int {
	// now() might not be the best option since it defaults
	// to UTC so offset will always be 0

	// o - offset returned in seconds
	_, o := time.Now().Zone()
	// to convert from seconds to hour divide by 3600
	// _, o := now().Zone() / 3600
	return o
	// Current offset UTC: 0
}

// offset returns the timezone offset (in seconds)
// from the given dt.
// https://play.golang.org/p/iQdonhxZzZH
func offset(t time.Time) int {
	// o - offset returned in seconds
	_, o := t.Zone()
	// to convert from seconds to hour divide by 3600
	// _, o := now().Zone() / 3600
	return o
	// offset UTC: 0
}
