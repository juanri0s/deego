package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_MonthShort(t *testing.T) {

	type args struct {
		t time.Time
	}

	jan := args{SampleTime()}
	feb := args{SampleTime().AddDate(0, 1, 0)}
	mar := args{SampleTime().AddDate(0, 2, 0)}
	apr := args{SampleTime().AddDate(0, 3, 0)}
	may := args{SampleTime().AddDate(0, 4, 0)}
	jun := args{SampleTime().AddDate(0, 5, 0)}
	jul := args{SampleTime().AddDate(0, 6, 0)}
	aug := args{SampleTime().AddDate(0, 7, 0)}
	sep := args{SampleTime().AddDate(0, 8, 0)}
	oct := args{SampleTime().AddDate(0, 9, 0)}
	nov := args{SampleTime().AddDate(0, 10, 0)}
	dec := args{SampleTime().AddDate(0, 11, 0)}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"January", jan, "Jan"},
		{"February", feb, "Feb"},
		{"March", mar, "Mar"},
		{"April", apr, "Apr"},
		{"May", may, "May"},
		{"June", jun, "Jun"},
		{"July", jul, "Jul"},
		{"August", aug, "Aug"},
		{"September", sep, "Sep"},
		{"October", oct, "Oct"},
		{"November", nov, "Nov"},
		{"December", dec, "Dec"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := monthShort(tt.args.t)
			assert.Equal(t, want, actual, "months did not match")
		})
	}
}
