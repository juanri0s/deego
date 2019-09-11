package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_endOfYear(t *testing.T) {
	type args struct {
		s string
	}

	y2019 := args{"2019"}
	y2020 := args{"2020"} // leap year
	y1999 := args{"1999"}

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"endofyear 2019", y2019, toTime("2019-12-31")},
		{"endofyear 2020", y2020, toTime("2020-12-31")},
		{"endofyear 1999", y1999, toTime("1999-12-31")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := endOfYear(tt.args.s)
			assert.Equal(t, want, actual, "unable to calculate end of year")
		})
	}
}
