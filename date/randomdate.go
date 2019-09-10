package date

import (
	"math/rand"
	"time"
)

// randomDate generates a random date
// https://play.golang.org/p/ZD1WdZgffaA
func randomDate() time.Time {
	min := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
