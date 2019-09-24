package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_tomorrow(t *testing.T) {
	type args struct {
		t time.Time
	}

	d1 := args{RandomDate()}
	d2 := args{RandomDate()}
	d3 := args{RandomDate()}

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"tomorrows date 1", d1, d1.t.AddDate(0, 0, 1)},
		{"tomorrows date 2", d2, d2.t.AddDate(0, 0, 1)},
		{"tomorrows date 3", d3, d3.t.AddDate(0, 0, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := tomorrow(tt.args.t)
			assert.Equal(t, want, actual, "dates do not match")
		})
	}
}
