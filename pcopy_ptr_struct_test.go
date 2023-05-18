// Copyright [2020-2023] [guonaihong]
package pcopy

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type test_BaseStructType_ptr_Dst_A struct {
	A int
	B int8
	C int16
	D int32
}

type test_BaseStructType_ptr_Dst_B struct {
	A int
	B int8
	C int16
	D int32
}

// 测试基础类型的指针
type test_BaseStructType_ptr_Dst struct {
	A *test_BaseStructType_ptr_Dst_A
	B *test_BaseStructType_ptr_Dst_B
}

type test_BaseStructType_ptr_Src struct {
	A test_BaseStructType_ptr_Dst_A
	B test_BaseStructType_ptr_Dst_B
}

var local_test_BaseStructType_ptr_Dst = test_BaseStructType_ptr_Dst{
	A: &test_BaseStructType_ptr_Dst_A{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
	},
	B: &test_BaseStructType_ptr_Dst_B{
		A: 5,
		B: 6,
		C: 7,
		D: 8,
	},
}

var local_test_BaseStructType_ptr_Src = test_BaseStructType_ptr_Src{
	A: test_BaseStructType_ptr_Dst_A{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
	},
	B: test_BaseStructType_ptr_Dst_B{
		A: 5,
		B: 6,
		C: 7,
		D: 8,
	},
}

func Test_Ptr_BaseStructType1_1(t *testing.T) {
	err := Preheat(&test_BaseStructType_ptr_Dst{}, &local_test_BaseStructType_ptr_Dst)
	assert.NoError(t, err, "Preheat(&test_BaseStructType_ptr_Dst{}, &local_test_BaseStructType_ptr_Dst)")

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseStructType_ptr_Dst
		err = Copy(&dst, &local_test_BaseStructType_ptr_Dst, opts...)
		assert.NoError(t, err, "Copy(&dst, &local_test_BaseStructType_ptr_Dst, WithUsePreheat()")
		assert.Equal(t, local_test_BaseStructType_ptr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseStructType_ptr_Dst.A), unsafe.Pointer(dst.A), fmt.Sprintf("local_test_BaseStructType_ptr_Dst.A: %p, dst.A: %p", local_test_BaseStructType_ptr_Dst.A, dst.A))
		assert.Equal(t, *local_test_BaseStructType_ptr_Dst.A, *dst.A)
	}
}

func Test_Ptr_BaseStructType1_2(t *testing.T) {
	err := Preheat(&test_BaseStructType_ptr_Dst{}, &test_BaseStructType_ptr_Dst{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseStructType_ptr_Dst
		err = Copy(&dst, &local_test_BaseStructType_ptr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseStructType_ptr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseStructType_ptr_Dst.A), unsafe.Pointer(dst.A))
		assert.Equal(t, *local_test_BaseStructType_ptr_Dst.A, *dst.A)
	}
}

func Test_Ptr_BaseStructType2(t *testing.T) {
	err := Preheat(&test_BaseStructType_ptr_Src{}, &test_BaseStructType_ptr_Dst{})
	assert.NoError(t, err)
	// Preheat(&test_BaseStructType_ptr_Dst{}, &test_BaseStructType_ptr_Src{})

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseStructType_ptr_Src
		err = Copy(&dst, &local_test_BaseStructType_ptr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseStructType_ptr_Src, dst)
	}
}

func Test_Ptr_BaseStructType3(t *testing.T) {
	err := Preheat(&test_BaseStructType_ptr_Dst{}, &test_BaseStructType_ptr_Src{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseStructType_ptr_Dst
		err = Copy(&dst, &local_test_BaseStructType_ptr_Src, opts...)
		assert.NoError(t, err)
		fmt.Printf("dst: %+v\n", dst)
		fmt.Printf(".A %p\n", dst.A)
		assert.Equal(t, local_test_BaseStructType_ptr_Dst, dst)
	}
}
