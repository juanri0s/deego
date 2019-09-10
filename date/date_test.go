package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_last(t *testing.T) {
	st := sampleTime()

	ttSun := st.AddDate(0, 0, -2)
	ttMon := st.AddDate(0, 0, -1)
	ttTues := st.AddDate(0, 0, -7)
	ttWed := st.AddDate(0, 0, -6)
	ttThurs := st.AddDate(0, 0, -5)
	ttFri := st.AddDate(0, 0, -4)
	ttSat := st.AddDate(0, 0, -3)

	type Weekday struct {
		weekday time.Weekday
	}

	sun := Weekday{time.Sunday}
	mon := Weekday{time.Monday}
	tues := Weekday{time.Tuesday}
	wed := Weekday{time.Wednesday}
	thurs := Weekday{time.Thursday}
	fri := Weekday{time.Friday}
	sat := Weekday{time.Saturday}

	tests := []struct {
		name    string
		inputs  Weekday
		outputs time.Time
	}{
		{"Sunday", sun, ttSun},
		{"Monday", mon, ttMon},
		{"Tuesday", tues, ttTues},
		{"Wednesday", wed, ttWed},
		{"Thursday", thurs, ttThurs},
		{"Friday", fri, ttFri},
		{"Saturday", sat, ttSat},
	}
	for _, tt := range tests {
		actual := last(tt.inputs.weekday)
		want := tt.outputs

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, actual, want, "times do not match")
		})
	}
}

func Test_next(t *testing.T) {
	st := sampleTime()

	ttSun := st.AddDate(0, 0, 5)
	ttMon := st.AddDate(0, 0, 6)
	ttTues := st.AddDate(0, 0, 7)
	ttWed := st.AddDate(0, 0, 1)
	ttThurs := st.AddDate(0, 0, 2)
	ttFri := st.AddDate(0, 0, 3)
	ttSat := st.AddDate(0, 0, 4)

	type Weekday struct {
		weekday time.Weekday
	}

	sun := Weekday{time.Sunday}
	mon := Weekday{time.Monday}
	tues := Weekday{time.Tuesday}
	wed := Weekday{time.Wednesday}
	thurs := Weekday{time.Thursday}
	fri := Weekday{time.Friday}
	sat := Weekday{time.Saturday}

	tests := []struct {
		name    string
		inputs  Weekday
		outputs time.Time
	}{
		{"Sunday", sun, ttSun},
		{"Monday", mon, ttMon},
		{"Tuesday", tues, ttTues},
		{"Wednesday", wed, ttWed},
		{"Thursday", thurs, ttThurs},
		{"Friday", fri, ttFri},
		{"Saturday", sat, ttSat},
	}
	for _, tt := range tests {
		actual := next(tt.inputs.weekday)
		want := tt.outputs

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, actual, want, "times do not match")
		})
	}
}

func Test_MonthShort(t *testing.T) {

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
			assert.Equal(t, actual, want, "months did not match")
		})
	}
}

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
			assert.Equal(t, actual, want, "months did not match")
		})
	}
}

func Test_isWeekday(t *testing.T) {
	type args struct {
		t time.Time
	}

	st := sampleTime()

	ttSun := args{st.AddDate(0, 0, 5)}
	ttMon := args{st.AddDate(0, 0, 6)}
	ttTues := args{st.AddDate(0, 0, 7)}
	ttWed := args{st.AddDate(0, 0, 1)}
	ttThurs := args{st.AddDate(0, 0, 2)}
	ttFri := args{st.AddDate(0, 0, 3)}
	ttSat := args{st.AddDate(0, 0, 4)}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Sunday", ttSun, false},
		{"Monday", ttMon, true},
		{"Tuesday", ttTues, true},
		{"Wednesday", ttWed, true},
		{"Thursday", ttThurs, true},
		{"Friday", ttFri, true},
		{"Saturday", ttSat, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := isWeekday(tt.args.t)
			assert.Equal(t, actual, want, "expected day failed weekday check")
		})
	}
}

func Test_isWeekend(t *testing.T) {
	type args struct {
		t time.Time
	}

	st := sampleTime()

	ttSun := args{st.AddDate(0, 0, 5)}
	ttMon := args{st.AddDate(0, 0, 6)}
	ttTues := args{st.AddDate(0, 0, 7)}
	ttWed := args{st.AddDate(0, 0, 1)}
	ttThurs := args{st.AddDate(0, 0, 2)}
	ttFri := args{st.AddDate(0, 0, 3)}
	ttSat := args{st.AddDate(0, 0, 4)}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Sunday", ttSun, true},
		{"Monday", ttMon, false},
		{"Tuesday", ttTues, false},
		{"Wednesday", ttWed, false},
		{"Thursday", ttThurs, false},
		{"Friday", ttFri, false},
		{"Saturday", ttSat, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := isWeekend(tt.args.t)
			assert.Equal(t, actual, want, "expected day not failed weekend check")
		})
	}
}

//func Test_nextWeek(t *testing.T) {
//	type args struct {
//		t time.Time
//	}
//	tests := []struct {
//		name string
//		args args
//		want time.Time
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := nextWeek(tt.args.t); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("nextWeek() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_tomorrow(t *testing.T) {
	st := sampleTime()
	got := tomorrow(st)

	s := "2019-01-23"
	layout := "2006-01-02"
	want, _ := time.Parse(layout, s)

	assert.Equal(t, got, want, "invalid date for tomorrow")
}

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
			actual := randomDate()
			assert.IsType(t, time.Time{}, actual, "Random datetime not of type time")
		})
	}
}
