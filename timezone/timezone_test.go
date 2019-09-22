package timezone

import (
	"testing"
	"time"
)

func Test_timezone(t *testing.T) {
	type args struct {
		t time.Time
	}

	// t1 := args{}
	// t2 := args{d.randomDate()}
	// t3 := args{d.randomDate()}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timezone(tt.args.t); got != tt.want {
				t.Errorf("timezone() = %v, want %v", got, tt.want)
			}
		})
	}
}
