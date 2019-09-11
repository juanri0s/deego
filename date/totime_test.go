package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_toTime(t *testing.T) {
	type args struct {
		s string
	}

	s1 := args{"2019-12-25"}
	s2 := args{"2000-01-25"}
	s3 := args{"2004-5-25"}

	tests := []struct {
		name string
		args args
	}{
		{"string to time test1", s1},
		{"string to time test2", s2},
		{"string to time test3", s3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := randomDate()
			assert.IsType(t, time.Time{}, actual, "unable to convert string to time")
		})
	}
}
