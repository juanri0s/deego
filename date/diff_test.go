package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_diff(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}

	d1 := RandomDate()
	d2 := RandomDate()
	d3 := RandomDate()

	t1 := args{d1, d1.AddDate(0, 0, 1)}
	t2 := args{d2, d2.AddDate(0, 0, 30)}
	t3 := args{d3, d3}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1 day diff", t1, 1},
		{"2 day diff", t2, 30},
		{"0 days diff", t3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := diff(tt.args.t1, tt.args.t2)
			assert.Equal(t, want, actual, "diff could not be calculated")
		})
	}
}
