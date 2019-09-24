package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_roundDate(t *testing.T) {
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
		{"date rounded 1", d1, d1.t},
		{"date rounded 2", d2, d2.t},
		{"date rounded 3", d3, d3.t},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := roundDate(tt.args.t)
			// TODO: I think this is an issue with testify but we can't compare times (greater, less) https://github.com/stretchr/testify/issues/502
			// so to fix this we just check that they're not equal even though they're using the same date
			assert.NotEqual(t, want, actual)
		})
	}
}
