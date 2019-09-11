package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_MonthLong(t *testing.T) {
	type args struct {
		t time.Time
	}

	jan := args{sampleTime()}
	feb := args{sampleTime().AddDate(0, 1, 0)}
	mar := args{sampleTime().AddDate(0, 2, 0)}
	apr := args{sampleTime().AddDate(0, 3, 0)}
	may := args{sampleTime().AddDate(0, 4, 0)}
	jun := args{sampleTime().AddDate(0, 5, 0)}
	jul := args{sampleTime().AddDate(0, 6, 0)}
	aug := args{sampleTime().AddDate(0, 7, 0)}
	sep := args{sampleTime().AddDate(0, 8, 0)}
	oct := args{sampleTime().AddDate(0, 9, 0)}
	nov := args{sampleTime().AddDate(0, 10, 0)}
	dec := args{sampleTime().AddDate(0, 11, 0)}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"January", jan, "January"},
		{"February", feb, "February"},
		{"March", mar, "March"},
		{"April", apr, "April"},
		{"May", may, "May"},
		{"June", jun, "June"},
		{"July", jul, "July"},
		{"August", aug, "August"},
		{"September", sep, "September"},
		{"October", oct, "October"},
		{"November", nov, "November"},
		{"December", dec, "December"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := monthLong(tt.args.t)
			assert.Equal(t, want, actual, "months did not match")
		})
	}
}
