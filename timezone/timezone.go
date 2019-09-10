package timezone

import (
	"time"
)

// timezone returns the timezone from a given dt.
// https://play.golang.org/p/vDwLPVND9Sg
func timezone(t time.Time) string {
	// z - timezone returned as the zone
	z, _ := t.Zone()
	return z
	// timezone: EST
}
