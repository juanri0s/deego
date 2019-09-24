package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_equal(t *testing.T) {

	type args struct {
		t1 time.Time
		t2 time.Time
	}

	d1 := RandomDate()
	d2 := RandomDate()
	d3 := RandomDate()

	t1 := args{d1, d1}
	t2 := args{d2, d2}
	t3 := args{d3, time.Now()}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"compare times #1", t1, true},
		{"compare times #2", t2, true},
		{"compare times #3", t3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := equal(tt.args.t1, tt.args.t2)
			assert.Equal(t, want, actual, "dates do not match")
		})
	}
}
