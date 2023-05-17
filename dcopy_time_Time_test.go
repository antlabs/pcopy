// Copyright [2020-2023] [guonaihong]
package pcopy

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

	err := Copy(&t1, &t2)
	assert.NoError(t, err)
	assert.Equal(t, t1.T1, t2.T1)
	assert.Equal(t, t1.I, t2.I)
}

func Test_TimeCopyEx(t *testing.T) {
	var t1 testCaseWithTime1
	t2 := testCaseWithTime2{T1: time.Now(), I: 3}

	err := Copy(&t1, &t2)
	assert.NoError(t, err)
	assert.Equal(t, t1.T1, t2.T1)
	assert.Equal(t, t1.I, t2.I)
}
