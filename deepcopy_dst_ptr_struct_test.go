package deepcopy

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

type TestDstPtr struct {
	X int
	Y int
}

type TestDstStructPtr_Src struct {
	TestDstPtr TestDstPtr
}

func Test_DstPtr_Struct(t *testing.T) {
	var dst TestDstStructPtr_Dst
	src := TestDstStructPtr_Src{
		TestDstPtr: TestDstPtr{X: 3, Y: 4},
	}

	Copy(&dst, &src).Do()
	assert.NotNil(t, src.TestDstPtr)
	assert.Equal(t, src.TestDstPtr.X, 3)
	assert.Equal(t, src.TestDstPtr.Y, 4)
}
