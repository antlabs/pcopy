// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHaveValue struct {
	Int     int
	Float32 float32
}

func Test_have_value(t *testing.T) {
	t1 := testHaveValue{Int: 3, Float32: 3.14}
	t2 := testHaveValue{}
	Copy(&t1, &t2).Do()
	assert.Equal(t, t1, testHaveValue{3, 3.14})
}
