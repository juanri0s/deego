package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_daysInYear(t *testing.T) {
	type args struct {
		s string
	}

	y2019 := args{"2019"}
	y2020 := args{"2020"} // leap year
	y1999 := args{"1999"}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"daysInYear 2019", y2019, 365},
		{"daysInYear 2020", y2020, 366},
		{"daysInYear 1999", y1999, 365},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := daysInYear(tt.args.s)
			assert.Equal(t, want, actual, "unable to calculate number of days in the year")
		})
	}
}
