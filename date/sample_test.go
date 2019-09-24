package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_sampleTime(t *testing.T) {
	actual := SampleTime()
	assert.IsType(t, time.Time{}, actual, "output is not a datetime")
}
