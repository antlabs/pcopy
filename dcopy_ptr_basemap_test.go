// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// 测试基础类型的指针
type test_BaseMapType_ptr_Dst struct {
	A *map[int]int
	B *map[int8]int8
	C *map[int16]int16
	D *map[int32]int32
	E *map[int64]int64
	F *map[uint]uint
	G *map[uint8]uint8
	H *map[uint16]uint16
	I *map[uint32]uint32
	J *map[uint64]uint64
	K *map[uintptr]uintptr
	L *map[float32]float32
	M *map[float64]float64
	N *map[complex64]complex64
	O *map[complex128]complex128
	P *map[bool]bool
	Q *map[string]string
	// R *interface{} TODO
}

type test_BaseMapType_ptr_Src struct {
	A map[int]int
	B map[int8]int8
	C map[int16]int16
	D map[int32]int32
	E map[int64]int64
	F map[uint]uint
	G map[uint8]uint8
	H map[uint16]uint16
	I map[uint32]uint32
	J map[uint64]uint64
	K map[uintptr]uintptr
	L map[float32]float32
	M map[float64]float64
	N map[complex64]complex64
	O map[complex128]complex128
	P map[bool]bool
	Q map[string]string
}

var local_test_BaseMapType_ptr_Dst = test_BaseMapType_ptr_Dst{
	A: newDef(map[int]int{1: 1, 2: 2}),
	B: newDef(map[int8]int8{3: 3, 4: 4}),
	C: newDef(map[int16]int16{5: 5, 6: 6}),
	D: newDef(map[int32]int32{7: 7, 8: 8}),
	E: newDef(map[int64]int64{9: 9, 10: 10}),
	F: newDef(map[uint]uint{11: 11, 12: 12}),
	G: newDef(map[uint8]uint8{13: 13, 14: 14}),
	H: newDef(map[uint16]uint16{15: 15, 16: 16}),
	I: newDef(map[uint32]uint32{17: 17, 18: 18}),
	J: newDef(map[uint64]uint64{19: 19, 20: 20}),
	K: newDef(map[uintptr]uintptr{21: 21, 22: 22}),
	L: newDef(map[float32]float32{23: 23, 24: 24}),
	M: newDef(map[float64]float64{25: 25, 26: 26}),
	N: newDef(map[complex64]complex64{27: 27, 28: 28}),
	O: newDef(map[complex128]complex128{29: 29, 30: 30}),
	P: newDef(map[bool]bool{true: true, false: false}),
	Q: newDef(map[string]string{"a": "a", "b": "b"}),
}

var local_test_BaseMapType_ptr_Src = test_BaseMapType_ptr_Src{
	A: map[int]int{1: 1, 2: 2},
	B: map[int8]int8{3: 3, 4: 4},
	C: map[int16]int16{5: 5, 6: 6},
	D: map[int32]int32{7: 7, 8: 8},
	E: map[int64]int64{9: 9, 10: 10},
	F: map[uint]uint{11: 11, 12: 12},
	G: map[uint8]uint8{13: 13, 14: 14},
	H: map[uint16]uint16{15: 15, 16: 16},
	I: map[uint32]uint32{17: 17, 18: 18},
	J: map[uint64]uint64{19: 19, 20: 20},
	K: map[uintptr]uintptr{21: 21, 22: 22},
	L: map[float32]float32{23: 23, 24: 24},
	M: map[float64]float64{25: 25, 26: 26},
	N: map[complex64]complex64{27: 27, 28: 28},
	O: map[complex128]complex128{29: 29, 30: 30},
	P: map[bool]bool{true: true, false: false},
	Q: map[string]string{"a": "a", "b": "b"},
}

func Test_Ptr_BaseMapType1_1(t *testing.T) {
	err := Preheat(&test_BaseMapType_ptr_Dst{}, &local_test_BaseMapType_ptr_Dst)
	assert.NoError(t, err, "Preheat(&test_BaseMapType_ptr_Dst{}, &local_test_BaseMapType_ptr_Dst)")

	var opts []Option
	for _, b := range []bool{true, false} {
		var dst test_BaseMapType_ptr_Dst
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		err = Copy(&dst, &local_test_BaseMapType_ptr_Dst, opts...)
		assert.NoError(t, err, "Copy(&dst, &local_test_BaseMapType_ptr_Dst, WithUsePreheat()")
		assert.Equal(t, local_test_BaseMapType_ptr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseMapType_ptr_Dst.A), unsafe.Pointer(dst.A), fmt.Sprintf("local_test_BaseMapType_ptr_Dst.A: %p, dst.A: %p", local_test_BaseMapType_ptr_Dst.A, dst.A))
		assert.Equal(t, *local_test_BaseMapType_ptr_Dst.A, *dst.A)
	}
}

func Test_Ptr_BaseMapType1_2(t *testing.T) {
	err := Preheat(&test_BaseMapType_ptr_Dst{}, &test_BaseMapType_ptr_Dst{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		var dst test_BaseMapType_ptr_Dst
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		err = Copy(&dst, &local_test_BaseMapType_ptr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseMapType_ptr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseMapType_ptr_Dst.A), unsafe.Pointer(dst.A))
		assert.Equal(t, *local_test_BaseMapType_ptr_Dst.A, *dst.A)
	}
}

func Test_Ptr_BaseMapType2(t *testing.T) {
	err := Preheat(&test_BaseMapType_ptr_Src{}, &test_BaseMapType_ptr_Dst{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		var dst test_BaseMapType_ptr_Src
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		err = Copy(&dst, &local_test_BaseMapType_ptr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseMapType_ptr_Src, dst)
	}
}

func Test_Ptr_BaseMapType3(t *testing.T) {
	err := Preheat(&test_BaseMapType_ptr_Dst{}, &test_BaseMapType_ptr_Src{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseMapType_ptr_Dst
		err = Copy(&dst, &local_test_BaseMapType_ptr_Src, opts...)
		assert.NoError(t, err)
		fmt.Printf("dst: %+v\n", dst)
		fmt.Printf(".A %p\n", dst.A)
		assert.Equal(t, local_test_BaseMapType_ptr_Dst, dst)
	}
}
