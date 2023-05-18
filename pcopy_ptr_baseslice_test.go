// Copyright [2020-2023] [guonaihong]
package pcopy

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// 测试基础类型的指针
type test_BaseSliceType_ptr_Dst struct {
	A *[]int
	B *[]int8
	C *[]int16
	D *[]int32
	E *[]int64
	F *[]uint
	G *[]uint8
	H *[]uint16
	I *[]uint32
	J *[]uint64
	K *[]uintptr
	L *[]float32
	M *[]float64
	N *[]complex64
	O *[]complex128
	P *[]bool
	Q *[]string
	// R *interface{} TODO
}

type test_BaseSliceType_ptr_Src struct {
	A []int
	B []int8
	C []int16
	D []int32
	E []int64
	F []uint
	G []uint8
	H []uint16
	I []uint32
	J []uint64
	K []uintptr
	L []float32
	M []float64
	N []complex64
	O []complex128
	P []bool
	Q []string
	// R interface{}
}

var local_test_BaseSliceType_ptr_Dst = test_BaseSliceType_ptr_Dst{
	A: newDef([]int{1, 2}),
	B: newDef([]int8{3, 4}),
	C: newDef([]int16{5, 6}),
	D: newDef([]int32{7, 8}),
	E: newDef([]int64{9, 10}),
	F: newDef([]uint{11, 12}),
	G: newDef([]uint8{13, 14}),
	H: newDef([]uint16{15, 16}),
	I: newDef([]uint32{17, 18}),
	J: newDef([]uint64{19, 20}),
	K: newDef([]uintptr{21, 22}),
	L: newDef([]float32{23.23, 24.24}),
	M: newDef([]float64{25.25, 26.26}),
	N: newDef([]complex64{27.27, 28.28}),
	O: newDef([]complex128{29.29, 30.30}),
	P: newDef([]bool{true, false}),
	Q: newDef([]string{"hello", "world"}),
}

var local_test_BaseSliceType_ptr_Src = test_BaseSliceType_ptr_Src{
	A: []int{1, 2},
	B: []int8{3, 4},
	C: []int16{5, 6},
	D: []int32{7, 8},
	E: []int64{9, 10},
	F: []uint{11, 12},
	G: []uint8{13, 14},
	H: []uint16{15, 16},
	I: []uint32{17, 18},
	J: []uint64{19, 20},
	K: []uintptr{21, 22},
	L: []float32{23.23, 24.24},
	M: []float64{25.25, 26.26},
	N: []complex64{27.27, 28.28},
	O: []complex128{29.29, 30.30},
	P: []bool{true, false},
	Q: []string{"hello", "world"},
}

func Test_Ptr_BaseSliceType1_1(t *testing.T) {
	err := Preheat(&test_BaseSliceType_ptr_Dst{}, &local_test_BaseSliceType_ptr_Dst)
	assert.NoError(t, err, "Preheat(&test_BaseSliceType_ptr_Dst{}, &local_test_BaseSliceType_ptr_Dst)")

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseSliceType_ptr_Dst
		err = Copy(&dst, &local_test_BaseSliceType_ptr_Dst, opts...)
		assert.NoError(t, err, "Copy(&dst, &local_test_BaseSliceType_ptr_Dst, WithUsePreheat()")
		assert.Equal(t, local_test_BaseSliceType_ptr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseSliceType_ptr_Dst.A), unsafe.Pointer(dst.A), fmt.Sprintf("local_test_BaseSliceType_ptr_Dst.A: %p, dst.A: %p", local_test_BaseSliceType_ptr_Dst.A, dst.A))
		assert.Equal(t, *local_test_BaseSliceType_ptr_Dst.A, *dst.A)
	}
}

func Test_Ptr_BaseSliceType1_2(t *testing.T) {
	err := Preheat(&test_BaseSliceType_ptr_Dst{}, &test_BaseSliceType_ptr_Dst{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseSliceType_ptr_Dst
		err = Copy(&dst, &local_test_BaseSliceType_ptr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseSliceType_ptr_Dst, dst)
		assert.NotEqual(t, unsafe.Pointer(local_test_BaseSliceType_ptr_Dst.A), unsafe.Pointer(dst.A))
		assert.Equal(t, *local_test_BaseSliceType_ptr_Dst.A, *dst.A)
	}
}

func Test_Ptr_BaseSliceType2(t *testing.T) {
	err := Preheat(&test_BaseSliceType_ptr_Src{}, &test_BaseSliceType_ptr_Dst{})
	assert.NoError(t, err)
	// Preheat(&test_BaseSliceType_ptr_Dst{}, &test_BaseSliceType_ptr_Src{})

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst test_BaseSliceType_ptr_Src
		err = Copy(&dst, &local_test_BaseSliceType_ptr_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_test_BaseSliceType_ptr_Src, dst)
	}
}

func Test_Ptr_BaseSliceType3(t *testing.T) {
	err := Preheat(&test_BaseSliceType_ptr_Dst{}, &test_BaseSliceType_ptr_Src{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}

		var dst test_BaseSliceType_ptr_Dst
		err = Copy(&dst, &local_test_BaseSliceType_ptr_Src, opts...)
		assert.NoError(t, err)
		fmt.Printf("dst: %+v\n", dst)
		fmt.Printf(".A %p\n", dst.A)
		assert.Equal(t, local_test_BaseSliceType_ptr_Dst, dst)
	}
}
