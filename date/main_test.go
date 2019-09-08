package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_last(t *testing.T) {

	type Weekday struct {
		s time.Weekday
	}

	sun := Weekday{time.Sunday}
	mon := Weekday{time.Monday}
	tues := Weekday{time.Tuesday}
	wed := Weekday{time.Wednesday}
	thurs := Weekday{time.Thursday}
	fri := Weekday{time.Friday}
	sat := Weekday{time.Saturday}

	tests := []struct {
		name string
		args Weekday
		want time.Time
	}{
		{"Sunday", sun, time.Now()},
		{"Monday", mon, time.Now()},
		{"Tuesday", tues, time.Now()},
		{"Wednesday", wed, time.Now()},
		{"Thursday", thurs, time.Now()},
		{"Friday", fri, time.Now()},
		{"Saturday", sat, time.Now()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEqual(t, last(tt.args.s), tt.want, " marshallindent")
		})
	}
}

func Test_next(t *testing.T) {

	type Weekday struct {
		s time.Weekday
	}

	sun := Weekday{time.Sunday}
	mon := Weekday{time.Monday}
	tues := Weekday{time.Tuesday}
	wed := Weekday{time.Wednesday}
	thurs := Weekday{time.Thursday}
	fri := Weekday{time.Friday}
	sat := Weekday{time.Saturday}

	tests := []struct {
		name string
		args Weekday
		want time.Time
	}{
		{"Sunday", sun, time.Now()},
		{"Monday", mon, time.Now()},
		{"Tuesday", tues, time.Now()},
		{"Wednesday", wed, time.Now()},
		{"Thursday", thurs, time.Now()},
		{"Friday", fri, time.Now()},
		{"Saturday", sat, time.Now()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEqual(t, next(tt.args.s), tt.want, " marshallindent")
		})
	}
}
