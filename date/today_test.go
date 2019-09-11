package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_today(t *testing.T) {
	actual := today()
	want := roundDate(time.Now())

	assert.Equal(t, want, actual, "times do not match")
}
