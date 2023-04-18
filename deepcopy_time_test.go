package deepcopy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testCaseWithTime1 struct {
	T1 time.Time
	I  int
}

type testCaseWithTime2 struct {
	T1 time.Time
	I  int
}

func Test_Time(t *testing.T) {
	var t1 testCaseWithTime1
	t2 := testCaseWithTime2{T1: time.Now(), I: 3}

	err := Copy(&t1, &t2).Do()
	assert.NoError(t, err)
	assert.Equal(t, t1.T1, t2.T1)
	assert.Equal(t, t1.I, t2.I)
}

func Test_TimeCopyEx(t *testing.T) {
	var t1 testCaseWithTime1
	t2 := testCaseWithTime2{T1: time.Now(), I: 3}

	err := CopyEx(&t1, &t2)
	assert.NoError(t, err)
	assert.Equal(t, t1.T1, t2.T1)
	assert.Equal(t, t1.I, t2.I)
}
