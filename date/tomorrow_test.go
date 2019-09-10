package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_tomorrow(t *testing.T) {
	st := sampleTime()
	got := tomorrow(st)

	s := "2019-01-23"
	layout := "2006-01-02"
	want, _ := time.Parse(layout, s)

	assert.Equal(t, got, want, "invalid date for tomorrow")
}
