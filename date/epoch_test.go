package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_epochSec(t *testing.T) {

	type args struct {
		t time.Time
	}

	d1 := args{randomDate()}
	d2 := args{randomDate()}
	d3 := args{randomDate()}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{"time to epoch1", d1, d1.t.Unix()},
		{"time to epoch2", d2, d2.t.Unix()},
		{"time to epoch3", d3, d3.t.Unix()},
	}
	for _, tt := range tests {
		want := tt.want
		actual := epochSec(tt.args.t)
		assert.Equal(t, want, actual, "epochs do not match")
	}
}

func Test_epochNano(t *testing.T) {
	type args struct {
		t time.Time
	}

	d1 := args{randomDate()}
	d2 := args{randomDate()}
	d3 := args{randomDate()}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{"time to epoch nano", d1, d1.t.UnixNano()},
		{"time to epoch nano", d2, d2.t.UnixNano()},
		{"time to epoch nano", d3, d3.t.UnixNano()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := epochNano(tt.args.t)
			assert.Equal(t, want, actual, "epochs do not match")
		})
	}
}

func Test_epochMilli(t *testing.T) {
	type args struct {
		t time.Time
	}

	d1 := args{randomDate()}
	d2 := args{randomDate()}
	d3 := args{randomDate()}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{"time to epoch milliseconds", d1, d1.t.UnixNano() / 1000000},
		{"time to epoch milliseconds", d2, d2.t.UnixNano() / 1000000},
		{"time to epoch milliseconds", d3, d3.t.UnixNano() / 1000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := epochMilli(tt.args.t)
			assert.Equal(t, want, actual, "epochs do not match")
		})
	}
}
