package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_compare(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}

	date1 := SampleTime()
	date2 := today()

	equal := args{date1, date1}
	less := args{date1, date2}
	greater := args{date2, date1}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"compare -> equal", equal, 0},
		{"compare -> less than", less, -1},
		{"compare -> greater than", greater, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := compare(tt.args.t1, tt.args.t2)
			assert.Equal(t, want, actual, "unable to compare times")
		})
	}
}
