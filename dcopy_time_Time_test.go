package dcopy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_timeTime_BaseType(t *testing.T) {
	err := Preheat(&time.Time{}, &time.Time{})
	assert.NoError(t, err)

	src := time.Now()
	var dst time.Time
	err = Copy(&dst, &src, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, src, dst)
}

func Test_timeTime_Slice(t *testing.T) {
	err := Preheat(&[]time.Time{}, &[]time.Time{})
	assert.NoError(t, err)

	src := []time.Time{time.Now(), time.Now(), time.Now()}
	var dst []time.Time
	err = Copy(&dst, &src, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, src, dst)
}
