package timezone

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome from deego/timezone")
}

// currTimezone returns the timezone from the current datetime.
func currTimezone() string {
	t := time.Now()
	z, _ := t.Zone()
	return z
	// Current timezone: EST
}

// getTimezone returns the timezone from a given datetime
func getTimezone(t time.Time) string {
	z, _ := t.Zone()
	return z
	// timezone: EST
}
