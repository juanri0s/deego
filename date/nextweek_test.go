package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_nextWeek(t *testing.T) {
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
		{"nextWeek 1", d1, d1.t.AddDate(0, 0, 7)},
		{"nextWeek 2", d2, d2.t.AddDate(0, 0, 7)},
		{"nextWeek 3", d3, d3.t.AddDate(0, 0, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := nextWeek(tt.args.t)
			assert.Equal(t, want, actual, "dates do not match")
		})
	}
}
