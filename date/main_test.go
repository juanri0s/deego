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
		inputs Weekday
		outputs time.Time
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
		actual := last(tt.inputs.s)
		want := tt.outputs

		t.Run(tt.name, func(t *testing.T) {
			assert.NotEqual(t, actual, want, " marshallindent")
		})
	}
}

func Test_next(t *testing.T) {

	// Mock time, represents a Tuesday
	tt := time.Date(2019, 1, 22, 18, 0, 0, 0, time.Now().Location())

	ttSun := tt.AddDate(0, 0, 5)
	ttMon :=  tt.AddDate(0, 0, 6)
	ttTues := tt.AddDate(0, 0, 7)
	ttWed := tt.AddDate(0, 0, 1)
	ttThurs := tt.AddDate(0, 0, 2)
	ttFri := tt.AddDate(0, 0, 3)
	ttSat := tt.AddDate(0, 0, 4)

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
		inputs Weekday
		outputs time.Time
	}{
		{"Sunday", sun, ttSun},
		{"Monday", mon, ttMon},
		{"Tuesday", tues, ttTues},
		{"Wednesday", wed, ttWed},
		{"Thursday", thurs, ttThurs},
		{"Friday", fri, ttFri},
		{"Saturday", sat, ttSat},
	}
	for _, tt := range tests {
		actual := last(tt.inputs.s)
		want := tt.outputs

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, actual, want, " marshallindent")
		})
	}
}
