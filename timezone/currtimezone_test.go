package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_currTimezone(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"current timezone test", "UTC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := currTimezone()
			assert.Equal(t, want, actual, "unable to determine current timezone from now time")
		})
	}
}
