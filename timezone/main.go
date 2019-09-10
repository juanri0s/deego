package timezone

import (
	"time"
)

func now() time.Time {
	return time.Now()
}

func sampleTime() time.Time {
	s := "2019-01-22"
	t, _ := time.Parse("2019-01-22", s)
	return t
}

// currTimezone returns the timezone from the current datetime.
func currTimezone() string {
	z, _ := now().Zone()
	return z
	// Current timezone: EST
}

// getTimezone returns the timezone from a given datetime
func getTimezone() string {
	z, _ := sampleTime().Zone()
	return z
	// timezone: EST
}
