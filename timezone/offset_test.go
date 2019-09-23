package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	d "../date"
)

func Test_offset(t *testing.T) {
	type args struct {
		t time.Time
	}

	t1 := args{d.RandomDate()}
	t2 := args{d.RandomDate()}
	t3 := args{d.RandomDate()}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Timezone from time test #1", t1, -14400},
		{"Timezone from time test #2", t2, -14400},
		{"Timezone from time test #3", t3, -14400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := offset(tt.args.t)
			assert.Equal(t, want, actual, "unable to determine offset from time")
		})
	}
}
