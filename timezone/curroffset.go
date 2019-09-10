package timezone

import "time"

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
