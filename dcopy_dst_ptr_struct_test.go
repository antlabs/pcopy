// Copyright [2020-2023] [guonaihong]
package pcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestDstStructPtr_Dst struct {
	TestDstPtr *struct {
		X int
		Y int
	}
}

type TestDstPtr2 struct {
	X int
	Y int
}

type TestDstStructPtr_Src struct {
	TestDstPtr TestDstPtr2
}

func Test_DstPtr_Struct(t *testing.T) {
	var dst TestDstStructPtr_Dst
	src := TestDstStructPtr_Src{
		TestDstPtr: TestDstPtr2{X: 3, Y: 4},
	}

	Copy(&dst, &src)
	assert.NotNil(t, dst.TestDstPtr)
	assert.Equal(t, dst.TestDstPtr.X, 3)
	assert.Equal(t, dst.TestDstPtr.Y, 4)
}
