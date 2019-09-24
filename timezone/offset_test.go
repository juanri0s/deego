package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/juanri0s/deego/date"
)

func Test_offset(t *testing.T) {
	type args struct {
		t time.Time
	}

	t1 := args{date.SampleTime()}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Timezone from time test #1", t1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := offset(tt.args.t)
			assert.Equal(t, want, actual, "unable to determine offset from time")
		})
	}
}
