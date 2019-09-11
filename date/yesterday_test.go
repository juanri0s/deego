package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_yesterday(t *testing.T) {
	type args struct {
		t time.Time
	}

	d1 := args{randomDate()}
	d2 := args{randomDate()}
	d3 := args{randomDate()}

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"yesterdays date 1", d1, d1.t.AddDate(0, 0, -1)},
		{"yesterdays date 2", d2, d2.t.AddDate(0, 0, -1)},
		{"yesterdays date 3", d3, d3.t.AddDate(0, 0, -1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := yesterday(tt.args.t)
			assert.Equal(t, want, actual, "dates do not match")
		})
	}
}
