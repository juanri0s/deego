package date

import "time"

// sampleTime generates a time that is used for testing
func sampleTime() time.Time {
	s := "2019-01-22"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)
	return t
}
