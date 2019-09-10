package timezone

import "time"

// currTimezone returns the timezone from the current dt.
// https://play.golang.org/p/-W0J32E3DKf
func currTimezone() string {
	z, _ := time.Now().Zone()
	return z
	// Current timezone: EST
}
