package timezone

import (
	"github.com/juanri0s/deego/date"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_timezone(t *testing.T) {
	type args struct {
		t time.Time
	}

	t1 := args{date.SampleTime()}
	t2 := args{date.SampleTime()}
	t3 := args{date.SampleTime()}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Timezone from time test #1", t1, "UTC"},
		{"Timezone from time test #2", t2, "UTC"},
		{"Timezone from time test #3", t3, "UTC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := timezone(tt.args.t)
			assert.Equal(t, want, actual, "unable to determine timezone from time")
		})
	}
}
