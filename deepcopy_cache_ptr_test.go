package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FastCopyDst struct {
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

type FastCopySrc FastCopyDst

var testSrc = FastCopySrc{
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
func TestFastCopyBase(t *testing.T) {
	var dst FastCopyDst

	err := Preheat(&dst, &testSrc)
	assert.NoError(t, err)

	printCacheAllFunc()
	// fmt.Printf("%v\n", cacheAllFunc)
	err = CopyEx(&dst, &testSrc, WithUsePreheat())

	var dst2 FastCopyDst
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
type FastCopyDst_BasePtr struct {
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
type FastCopyDst_BaseSlice struct {
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
	// SliceComplex64  []complex64
	// SliceComplex128 []complex128
}

type FastCopySrc_BaseSlice FastCopyDst_BaseSlice

var testSrc_BaseSlice = FastCopySrc_BaseSlice{
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
	SliceFloat32: []float32{23.0, 24.0},
	SliceFloat64: []float64{25.0, 26.0},
	// SliceComplex64:  []complex64{27.0, 28.0},
	// SliceComplex128: []complex128{29.0, 30.0},
}

func TestFastCopy_BaseWithSlice(t *testing.T) {
	var dst FastCopyDst_BaseSlice

	err := Preheat(&dst, &testSrc_BaseSlice)
	assert.NoError(t, err)

	printCacheAllFunc()
	// fmt.Printf("%v\n", cacheAllFunc)
	err = CopyEx(&dst, &testSrc_BaseSlice, WithUsePreheat())
	assert.NoError(t, err)

	var dst2 FastCopyDst_BaseSlice = FastCopyDst_BaseSlice(testSrc_BaseSlice)
	assert.Equal(t, dst, dst2)
}
