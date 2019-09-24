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

	s := "2019-01-22"
	layout := "0001-01-01 00:00:00 +0000 UTC"
	t3, _ := time.Parse(layout, s)
	
	d1 := args{RandomDate()}
	d2 := args{RandomDate()}
	d3 := args{t3}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"random date validate true #1", d1, true},
		{"random date validate true #2", d2, true},
		{"random date validate false #!", d3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := isValid(tt.args.t)
			assert.Equal(t, want, actual, "unable to determine if given time was valid")
		})
	}
}
