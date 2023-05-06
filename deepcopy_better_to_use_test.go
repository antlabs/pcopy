// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testBetterToUse struct {
	Str string
	ID  int
}

// src是指针类型的结构体
// dst是普通结构体
func Test_srcPtr_DstBaseType(t *testing.T) {
	t1 := testBetterToUse{Str: "hello", ID: 1}
	t2 := testBetterToUse{}
	err := Copy(&t2, t1)
	assert.NoError(t, err)
	assert.Equal(t, t1, t2)
}

func Test_srcPtr_DstBaseType_NotPanics(t *testing.T) {
	t1 := testBetterToUse{Str: "hello", ID: 1}
	assert.NotPanics(t, func() {
		Copy((*testBetterToUse)(nil), t1)
	})
}
