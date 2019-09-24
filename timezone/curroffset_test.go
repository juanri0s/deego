package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_currOffset(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"current timezone test", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := currOffset()
			assert.Equal(t, want, actual, "unable to determine offset from now time")
		})
	}
}
