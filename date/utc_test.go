package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_utc(t *testing.T) {
	type args struct {
		s  string
		tz string
	}

	s1 := "2019-09-10"
	tz1 := "EST"

	s2 := "2000-09-10"
	tz2 := "HST"

	s3 := "2005-01-20"
	tz3 := "America/New_York"

	t1 := args{s1, tz1}
	t2 := args{s2, tz2}
	t3 := args{s3, tz3}

	tests := []struct {
		name string
		args args
	}{
		{"string to time test1", t1},
		{"string to time test1", t2},
		{"string to time test1", t3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := utc(tt.args.s, tt.args.tz)
			z, _ := actual.Zone()
			assert.NotEqual(t, z, tt.args.tz, "timezone does not match utc time or invalid timezone")
		})
	}
}
