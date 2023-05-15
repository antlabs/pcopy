package dcopy

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// 测试基础类型的指针
type test_BaseType_DoublePtr_Dst struct {
	A **int
	B **int8
	C **int16
	D **int32
	E **int64
	F **uint
	G **uint8
	H **uint16
	I **uint32
	J **uint64
	K **uintptr
	L **float32
	M **float64
	N **complex64
	O **complex128
	P **bool
	Q **string
	// R *interface{} TODO
}

type test_BaseType_DoublePtr_Src struct {
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

func newDoubleDef[T any](t T) (rv **T) {
	p := &t
	return &p
}

var local_test_BaseType_DoublePtr_Dst = test_BaseType_DoublePtr_Dst{
	A: newDoubleDef(1),
	B: newDoubleDef(int8(2)),
	C: newDoubleDef(int16(3)),
	D: newDoubleDef(int32(4)),
	E: newDoubleDef(int64(5)),
	F: newDoubleDef(uint(6)),
	G: newDoubleDef(uint8(7)),
	H: newDoubleDef(uint16(8)),
	I: newDoubleDef(uint32(9)),
	J: newDoubleDef(uint64(10)),
	K: newDoubleDef(uintptr(11)),
	L: newDoubleDef(float32(12.12)),
	M: newDoubleDef(float64(13.13)),
	N: newDoubleDef(complex64(14.14)),
	O: newDoubleDef(complex128(15.15)),
	P: newDoubleDef(true),
	Q: newDoubleDef("hello"),
	// R: newDef("world"),
}

var local_test_BaseType_DoublePtr_Src = test_BaseType_DoublePtr_Src{
	A: 1,
	B: int8(2),
	C: int16(3),
	D: int32(4),
	E: int64(5),
	F: uint(6),
	G: uint8(7),
	H: uint16(8),
	I: uint32(9),
	J: uint64(10),
	K: uintptr(11),
	L: float32(12.12),
	M: float64(13.13),
	N: complex64(14.14),
	O: complex128(15.15),
	P: true,
	Q: "hello",
	// R: "world",
}

func Test_DoublePtr_BaseType1_1(t *testing.T) {
	err := Preheat(&test_BaseType_DoublePtr_Dst{}, &local_test_BaseType_DoublePtr_Dst)
	assert.NoError(t, err, "Preheat(&test_BaseType_DoublePtr_Dst{}, &local_test_BaseType_DoublePtr_Dst)")

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseType_DoublePtr_Dst
		err = Copy(&dst, &local_test_BaseType_DoublePtr_Dst, opts...)
		assert.NoError(t, err, "Copy(&dst, &local_test_BaseType_DoublePtr_Dst, WithUsePreheat()")
		assert.Equal(t, local_test_BaseType_DoublePtr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseType_DoublePtr_Dst.A), unsafe.Pointer(dst.A), fmt.Sprintf("local_test_BaseType_DoublePtr_Dst.A: %p, dst.A: %p", local_test_BaseType_DoublePtr_Dst.A, dst.A))
		assert.Equal(t, *local_test_BaseType_DoublePtr_Dst.A, *dst.A)
	}
}

func Test_DoublePtr_BaseType1_2(t *testing.T) {
	err := Preheat(&test_BaseType_DoublePtr_Dst{}, &test_BaseType_DoublePtr_Dst{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseType_DoublePtr_Dst
		err = Copy(&dst, &local_test_BaseType_DoublePtr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseType_DoublePtr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseType_DoublePtr_Dst.A), unsafe.Pointer(dst.A))
		assert.Equal(t, *local_test_BaseType_DoublePtr_Dst.A, *dst.A)
	}
}

func Test_DoublePtr_BaseType2(t *testing.T) {
	err := Preheat(&test_BaseType_DoublePtr_Src{}, &test_BaseType_DoublePtr_Dst{})
	assert.NoError(t, err)
	// Preheat(&test_BaseType_DoublePtr_Dst{}, &test_BaseType_DoublePtr_Src{})

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseType_DoublePtr_Src
		err = Copy(&dst, &local_test_BaseType_DoublePtr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseType_DoublePtr_Src, dst)
	}
}

func Test_DoublePtr_BaseType3(t *testing.T) {
	err := Preheat(&test_BaseType_DoublePtr_Dst{}, &test_BaseType_DoublePtr_Src{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseType_DoublePtr_Dst
		err = Copy(&dst, &local_test_BaseType_DoublePtr_Src, opts...)
		assert.NoError(t, err)
		fmt.Printf("dst: %+v\n", dst)
		fmt.Printf(".A %p\n", dst.A)
		assert.Equal(t, local_test_BaseType_DoublePtr_Dst, dst)
	}
}

type Msg struct {
	FromUid int32  `protobuf:"varint,1,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`
	ToUid   int32  `protobuf:"varint,2,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"`
	Msg     string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

// 测试src是双指针的情况
// 为了提升dcopy库的兼容性, src是双指针, dst是单指针, 也可以正常拷贝
func Test_Src_DoublePtr_Test(t *testing.T) {
	m1 := Msg{FromUid: 1, ToUid: 2, Msg: "1to2"}

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		// 单指针
		m1Ptr := &m1

		var m2 Msg
		// dst(m2) 是单指针
		// src(m1Ptr) 是双指针
		err := Copy(&m2, &m1Ptr, opts...)

		assert.NoError(t, err)
		assert.Equal(t, m1, m2)
	}
}

func Test_Src_DoublePtr_Test_CopyEx(t *testing.T) {
	m1 := Msg{FromUid: 1, ToUid: 2, Msg: "1to2"}

	// 单指针
	m1Ptr := &m1

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var m2 Msg
		// dst(m2) 是单指针
		// src(m1Ptr) 是双指针
		err := Copy(&m2, &m1Ptr, opts...)

		assert.NoError(t, err)
		assert.Equal(t, m1, m2)
	}
}
