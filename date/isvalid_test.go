package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	type args struct {
		t time.Time
	}

	d1 := args{RandomDate()}
	d2 := args{RandomDate()}
	d3 := args{RandomDate()}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"random date validate", d1, true},
		{"random date validate", d2, true},
		{"random date validate", d3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := isValid(tt.args.t)
			assert.Equal(t, want, actual, "unable to determine if given time was valid")
		})
	}
}
