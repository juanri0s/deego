package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_randomDate(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{"random date", time.Time{}},
		{"random date", time.Time{}},
		{"random date", time.Time{}},
		{"random date", time.Time{}},
		{"random date", time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RandomDate()
			assert.IsType(t, time.Time{}, actual, "Random datetime not of type time")
		})
	}
}
