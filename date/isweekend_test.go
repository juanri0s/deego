package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_isWeekend(t *testing.T) {
	type args struct {
		t time.Time
	}

	st := SampleTime()

	ttSun := args{st.AddDate(0, 0, 5)}
	ttMon := args{st.AddDate(0, 0, 6)}
	ttTues := args{st.AddDate(0, 0, 7)}
	ttWed := args{st.AddDate(0, 0, 1)}
	ttThurs := args{st.AddDate(0, 0, 2)}
	ttFri := args{st.AddDate(0, 0, 3)}
	ttSat := args{st.AddDate(0, 0, 4)}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Sunday", ttSun, true},
		{"Monday", ttMon, false},
		{"Tuesday", ttTues, false},
		{"Wednesday", ttWed, false},
		{"Thursday", ttThurs, false},
		{"Friday", ttFri, false},
		{"Saturday", ttSat, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := isWeekend(tt.args.t)
			assert.Equal(t, want, actual, "expected day not failed weekend check")
		})
	}
}
