package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_isWeekday(t *testing.T) {
	type args struct {
		t time.Time
	}

	st := sampleTime()

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
		{"Sunday", ttSun, false},
		{"Monday", ttMon, true},
		{"Tuesday", ttTues, true},
		{"Wednesday", ttWed, true},
		{"Thursday", ttThurs, true},
		{"Friday", ttFri, true},
		{"Saturday", ttSat, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := isWeekday(tt.args.t)
			assert.Equal(t, actual, want, "expected day failed weekday check")
		})
	}
}
