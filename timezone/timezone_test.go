package timezone

import (
	"testing"
	"time"

	d "../date"
	"github.com/stretchr/testify/assert"
)

func Test_timezone(t *testing.T) {
	type args struct {
		t time.Time
	}

	t1 := args{d.RandomDate()}
	t2 := args{d.RandomDate()}
	t3 := args{d.RandomDate()}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Grabbing timezone", t1, "EDT"},
		{"Grabbing timezone", t2, "EDT"},
		{"Grabbing timezone", t3, "EDT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := timezone(tt.args.t)
			assert.Equal(t, want, actual, "unable to determine timezone from time")
		})
	}
}
