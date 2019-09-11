package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_last(t *testing.T) {
	st := sampleTime()

	ttSun := st.AddDate(0, 0, -2)
	ttMon := st.AddDate(0, 0, -1)
	ttTues := st.AddDate(0, 0, -7)
	ttWed := st.AddDate(0, 0, -6)
	ttThurs := st.AddDate(0, 0, -5)
	ttFri := st.AddDate(0, 0, -4)
	ttSat := st.AddDate(0, 0, -3)

	type Weekday struct {
		weekday time.Weekday
	}

	sun := Weekday{time.Sunday}
	mon := Weekday{time.Monday}
	tues := Weekday{time.Tuesday}
	wed := Weekday{time.Wednesday}
	thurs := Weekday{time.Thursday}
	fri := Weekday{time.Friday}
	sat := Weekday{time.Saturday}

	tests := []struct {
		name    string
		inputs  Weekday
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
		actual := last(tt.inputs.weekday)
		want := tt.outputs

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, want, actual, "times do not match")
		})
	}
}
