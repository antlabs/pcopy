// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// dcopy指优化原来dcopy的优化版本，正式发布也行会改名为dcopy
type DCopyDst struct {
	// 所有基础类型
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	String  string
	Float32 float32
	Float64 float64
	// Complex64  complex64
	// Complex128 complex128
}

type DCopySrc DCopyDst

var testSrc = DCopySrc{
	Bool:    true,
	Int:     1,
	Int8:    2,
	Int16:   3,
	Int32:   4,
	Int64:   5,
	Uint:    6,
	Uint8:   7,
	Uint16:  8,
	Uint32:  9,
	Uint64:  10,
	String:  "11",
	Float32: 12.0,
	Float64: 13.0,
	// Complex64:  14.0,
	// Complex128: 15.0,
}

// 基础类型
func TestDCopyBase(t *testing.T) {
	var dst DCopyDst

	err := Preheat(&dst, &testSrc)
	assert.NoError(t, err)
	dst = DCopyDst{}

	err = Copy(&dst, &testSrc, WithUsePreheat())

	var dst2 DCopyDst
	dst2.Bool = true
	dst2.Int = 1
	dst2.Int8 = 2
	dst2.Int16 = 3
	dst2.Int32 = 4
	dst2.Int64 = 5
	dst2.Uint = 6
	dst2.Uint8 = 7
	dst2.Uint16 = 8
	dst2.Uint32 = 9
	dst2.Uint64 = 10
	dst2.String = "11"
	dst2.Float32 = 12.0
	dst2.Float64 = 13.0
	// dst2.Complex64 = 14.0
	// dst2.Complex128 = 15.0
	assert.NoError(t, err)
	assert.Equal(t, dst, dst2)
}

// 基础指针类型
// TODO
type DopyDst_BasePtr struct {
	// 所有基础类型
	Bool    *bool
	Int     *int
	Int8    *int8
	Int16   *int16
	Int32   *int32
	Int64   *int64
	Uint    *uint
	Uint8   *uint8
	Uint16  *uint16
	Uint32  *uint32
	Uint64  *uint64
	String  *string
	Float32 *float32
	Float64 *float64
	// Complex64  complex64
	// Complex128 complex128
}

// func TestFastCopyBasePtr(t *testing.T) {
// 	var dst FastCopyDst_BasePtr

// 	err := Preheat(&dst, &testSrc)
// 	assert.NoError(t, err)

// 	printCacheAllFunc()
// 	err = CopyEx(&dst, &testSrc, WithPreheat())

// 	var dst2 FastCopyDst_BasePtr
// 	dst2.Bool = new(bool)
// 	*dst2.Bool = true

// 	dst2.Int = new(int)
// 	*dst2.Int = 1

// 	dst2.Int8 = new(int8)
// 	*dst2.Int8 = 2

// 	dst2.Int16 = new(int16)
// 	*dst2.Int16 = 3

// 	dst2.Int32 = new(int32)
// 	*dst2.Int32 = 4

// 	dst2.Int64 = new(int64)
// 	*dst2.Int64 = 5

// 	dst2.Uint = new(uint)
// 	*dst2.Uint = 6

// 	dst2.Uint8 = new(uint8)
// 	*dst2.Uint8 = 7

// 	dst2.Uint16 = new(uint16)
// 	*dst2.Uint16 = 8

// 	dst2.Uint32 = new(uint32)
// 	*dst2.Uint32 = 9

// 	dst2.Uint64 = new(uint64)
// 	*dst2.Uint64 = 10

// 	// dst2.Complex64 = 14.0
// 	// dst2.Complex128 = 15.0
// 	assert.NoError(t, err)
// 	assert.Equal(t, dst, dst2)
// }

// 基础类型带slice

func getSliceHeaderPtr(s unsafe.Pointer) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(s)
}

type DCopySrc_BaseStruct DCopyDst_BaseStruct

// 基础下struct 套struct的情况
type DCopyDst_BaseStruct struct {
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	String  string
	Float32 float32
	Float64 float64
	// Complex64  complex64
	// Complex128 complex128
	SliceBool    []bool
	SliceInt     []int
	SliceInt8    []int8
	SliceInt16   []int16
	SliceInt32   []int32
	SliceInt64   []int64
	SliceUint    []uint
	SliceUint8   []uint8
	SliceUint16  []uint16
	SliceUint32  []uint32
	SliceUint64  []uint64
	SliceString  []string
	SliceFloat32 []float32
	SliceFloat64 []float64

	Sub2 DCopyDst_BaseStruct_Sub2
}

type DCopyDst_BaseStruct_Sub2 struct {
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	String  string
	Float32 float32
	Float64 float64
	// Complex64  complex64
	// Complex128 complex128
	SliceBool    []bool
	SliceInt     []int
	SliceInt8    []int8
	SliceInt16   []int16
	SliceInt32   []int32
	SliceInt64   []int64
	SliceUint    []uint
	SliceUint8   []uint8
	SliceUint16  []uint16
	SliceUint32  []uint32
	SliceUint64  []uint64
	SliceString  []string
	SliceFloat32 []float32
	SliceFloat64 []float64

	Sub3 DCopyDst_BaseStruct_Sub3
}

type DCopyDst_BaseStruct_Sub3 struct {
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	String  string
	Float32 float32
	Float64 float64
	// Complex64  complex64
	// Complex128 complex128
	SliceBool    []bool
	SliceInt     []int
	SliceInt8    []int8
	SliceInt16   []int16
	SliceInt32   []int32
	SliceInt64   []int64
	SliceUint    []uint
	SliceUint8   []uint8
	SliceUint16  []uint16
	SliceUint32  []uint32
	SliceUint64  []uint64
	SliceString  []string
	SliceFloat32 []float32
	SliceFloat64 []float64
}

var testBaseStructSrc DCopySrc_BaseStruct = DCopySrc_BaseStruct{
	Bool:    true,
	Int:     1,
	Int8:    2,
	Int16:   3,
	Int32:   4,
	Int64:   5,
	Uint:    6,
	Uint8:   7,
	Uint16:  8,
	Uint32:  9,
	Uint64:  10,
	String:  "11",
	Float32: 12,
	Float64: 13,
	// Complex64:  14,
	// Complex128: 15,
	SliceBool:    []bool{true, false},
	SliceInt:     []int{1, 2},
	SliceInt8:    []int8{3, 4},
	SliceInt16:   []int16{5, 6},
	SliceInt32:   []int32{7, 8},
	SliceInt64:   []int64{9, 10},
	SliceUint:    []uint{11, 12},
	SliceUint8:   []uint8{13, 14},
	SliceUint16:  []uint16{15, 16},
	SliceUint32:  []uint32{17, 18},
	SliceUint64:  []uint64{19, 20},
	SliceString:  []string{"21", "22"},
	SliceFloat32: []float32{23, 24},
	SliceFloat64: []float64{25, 26},
	Sub2: DCopyDst_BaseStruct_Sub2{
		Bool:    true,
		Int:     1,
		Int8:    2,
		Int16:   3,
		Int32:   4,
		Int64:   5,
		Uint:    6,
		Uint8:   7,
		Uint16:  8,
		Uint32:  9,
		Uint64:  10,
		String:  "11",
		Float32: 12,
		Float64: 13,
		// Complex64:  14,
		// Complex128: 15,
		SliceBool:    []bool{true, false},
		SliceInt:     []int{1, 2},
		SliceInt8:    []int8{3, 4},
		SliceInt16:   []int16{5, 6},
		SliceInt32:   []int32{7, 8},
		SliceInt64:   []int64{9, 10},
		SliceUint:    []uint{11, 12},
		SliceUint8:   []uint8{13, 14},
		SliceUint16:  []uint16{15, 16},
		SliceUint32:  []uint32{17, 18},
		SliceUint64:  []uint64{19, 20},
		SliceString:  []string{"21", "22"},
		SliceFloat32: []float32{23, 24},
		SliceFloat64: []float64{25, 26},
		Sub3: DCopyDst_BaseStruct_Sub3{
			Bool:    true,
			Int:     1,
			Int8:    2,
			Int16:   3,
			Int32:   4,
			Int64:   5,
			Uint:    6,
			Uint8:   7,
			Uint16:  8,
			Uint32:  9,
			Uint64:  10,
			String:  "11",
			Float32: 12,
			Float64: 13,
			// Complex64:  14,
			// Complex128: 15,
			SliceBool:    []bool{true, false},
			SliceInt:     []int{1, 2},
			SliceInt8:    []int8{3, 4},
			SliceInt16:   []int16{5, 6},
			SliceInt32:   []int32{7, 8},
			SliceInt64:   []int64{9, 10},
			SliceUint:    []uint{11, 12},
			SliceUint8:   []uint8{13, 14},
			SliceUint16:  []uint16{15, 16},
			SliceUint32:  []uint32{17, 18},
			SliceUint64:  []uint64{19, 20},
			SliceString:  []string{"21", "22"},
			SliceFloat32: []float32{23, 24},
			SliceFloat64: []float64{25, 26},
		},
	},
}

func Test_StructWithStruct(t *testing.T) {
	var dst DCopyDst_BaseStruct

	err := Preheat(&dst, &testBaseStructSrc)
	// err := Preheat(&dst, &testSrc)
	assert.NoError(t, err)
	dst = DCopyDst_BaseStruct{}

	// fmt.Printf("%v\n", cacheAllFunc)
	err = Copy(&dst, &testBaseStructSrc, WithUsePreheat())

	// 直接赋值对于slice是浅拷贝，这里只是为了测试
	dst2 := testBaseStructSrc
	assert.NoError(t, err)
	assert.Equal(t, dst, DCopyDst_BaseStruct(dst2))
}
