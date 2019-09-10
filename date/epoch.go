package date

import "time"

// epochSec accepts a time and returns the number of elapsed seconds from unix epoch.
// https://play.golang.org/p/cq1y6r7qiYd
func epochSec(t time.Time) int64 {
	secs := t.Unix()
	return secs
}

// epochNano accepts a time and returns the number of elapsed nanoseconds from unix epoch
// https://play.golang.org/p/H4t7b3KeW8t
func epochNano(t time.Time) int64 {
	nanos := t.UnixNano()
	return nanos
}

// epochMilli accepts a time and returns the number of elapsed milliseconds from unix epoch
// https://play.golang.org/p/6g_v4v4XvTu
func epochMilli(t time.Time) int64 {
	nanos := t.UnixNano()
	// go does not have UnixMillis, so we have to divide manually from nanos
	// 1000000 nanoseconds in 1 millisecond
	millis := nanos / 1000000
	return millis
}
