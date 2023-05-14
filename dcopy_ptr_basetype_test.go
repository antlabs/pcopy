// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试基础类型的指针
type test_BaseType_ptr_Dst struct {
	A *int
	B *int8
	C *int16
	D *int32
	E *int64
	F *uint
	G *uint8
	H *uint16
	I *uint32
	J *uint64
	K *uintptr
	L *float32
	M *float64
	N *complex64
	O *complex128
	P *bool
	Q *string
	// R *interface{} TODO
}

type test_BaseType_ptr_Src struct {
	A int
	B int8
	C int16
	D int32
	E int64
	F uint
	G uint8
	H uint16
	I uint32
	J uint64
	K uintptr
	L float32
	M float64
	N complex64
	O complex128
	P bool
	Q string
	// R interface{}
}

func newDef[T any](t T) (rv *T) {
	return &t
}

var local_test_BaseType_ptr_Dst = test_BaseType_ptr_Dst{
	A: newDef(1),
	B: newDef(int8(2)),
	C: newDef(int16(3)),
	D: newDef(int32(4)),
	E: newDef(int64(5)),
	F: newDef(uint(6)),
	G: newDef(uint8(7)),
	H: newDef(uint16(8)),
	I: newDef(uint32(9)),
	J: newDef(uint64(10)),
	K: newDef(uintptr(11)),
	L: newDef(float32(12.12)),
	M: newDef(float64(13.13)),
	N: newDef(complex64(14.14)),
	O: newDef(complex128(15.15)),
	P: newDef(true),
	Q: newDef("hello"),
	// R: newDef("world"),
}

func Test_Ptr_BaseType1_1(t *testing.T) {
	err := Preheat(&test_BaseType_ptr_Dst{}, &local_test_BaseType_ptr_Dst)
	assert.NoError(t, err, "Preheat(&test_BaseType_ptr_Dst{}, &local_test_BaseType_ptr_Dst)")

	var dst test_BaseType_ptr_Dst
	err = Copy(&dst, &local_test_BaseType_ptr_Dst, WithUsePreheat())
	assert.NoError(t, err, "Copy(&dst, &local_test_BaseType_ptr_Dst, WithUsePreheat()")
	assert.Equal(t, local_test_BaseType_ptr_Dst, dst)
}

func Test_Ptr_BaseType1_2(t *testing.T) {
	err := Preheat(&test_BaseType_ptr_Dst{}, &test_BaseType_ptr_Dst{})
	assert.NoError(t, err)
	// Preheat(&test_BaseType_ptr_Src{}, &test_BaseType_ptr_Dst{})
	// Preheat(&test_BaseType_ptr_Dst{}, &test_BaseType_ptr_Src{})

	var dst test_BaseType_ptr_Dst
	err = Copy(&dst, &local_test_BaseType_ptr_Dst, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, local_test_BaseType_ptr_Dst, dst)
}

// TODO
func Test_Ptr_BaseType2(t *testing.T) {
	err := Preheat(&test_BaseType_ptr_Dst{}, &test_BaseType_ptr_Dst{})
	// Preheat(&test_BaseType_ptr_Src{}, &test_BaseType_ptr_Dst{})
	assert.NoError(t, err)
	// Preheat(&test_BaseType_ptr_Dst{}, &test_BaseType_ptr_Src{})

	var dst test_BaseType_ptr_Dst
	err = Copy(&dst, &local_test_BaseType_ptr_Dst, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, local_test_BaseType_ptr_Dst, dst)
}

// TODO
func Test_Ptr_BaseType3(t *testing.T) {
	err := Preheat(&test_BaseType_ptr_Dst{}, &test_BaseType_ptr_Dst{})
	// Preheat(&test_BaseType_ptr_Src{}, &test_BaseType_ptr_Dst{})
	assert.NoError(t, err)
	// Preheat(&test_BaseType_ptr_Dst{}, &test_BaseType_ptr_Src{})

	var dst test_BaseType_ptr_Dst
	err = Copy(&dst, &local_test_BaseType_ptr_Dst, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, local_test_BaseType_ptr_Dst, dst)
}

// 测试指针
func Test_Ptr_OK(t *testing.T) {
	type interfaceTest struct {
		Iptr *int
		Fptr *float64
	}

	for _, tc := range []testCase{
		func() testCase {
			d := interfaceTest{}
			src := interfaceTest{
				Iptr: new(int),
				Fptr: new(float64),
			}

			*src.Iptr = 3
			*src.Fptr = 3.3
			Copy(&d, &src)
			return testCase{got: d, need: src}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

func Test_Ptr_OKCopyEx(t *testing.T) {
	type interfaceTest struct {
		Iptr *int
		Fptr *float64
	}

	for _, tc := range []testCase{
		func() testCase {
			d := interfaceTest{}
			src := interfaceTest{
				Iptr: new(int),
				Fptr: new(float64),
			}

			*src.Iptr = 3
			*src.Fptr = 3.3
			err := Copy(&d, &src)
			assert.NoError(t, err)
			return testCase{got: d, need: src}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 测试指针特殊情况
// 只要不崩溃就是对的
func Test_Ptr_Special(t *testing.T) {
	for _, tc := range []testCase{
		// dst 是空指针
		func() testCase {
			Copy((*int)(nil), new(int))
			return testCase{}
		}(),

		// dst, src是不同类型
		func() testCase {
			Copy("hello", new(int))
			return testCase{}
		}(),
		// dst 是双指针
		func() testCase {
			n := 3
			dst := 0
			dstPtr := &dst
			dstPtrPtr := &dstPtr
			err := Copy(dstPtrPtr, &n)
			assert.NoError(t, err)
			return testCase{}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
