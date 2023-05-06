package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DCopyDst_BaseSlice struct {
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

type DCopySrc_BaseSlice DCopyDst_BaseSlice

var testSrc_BaseSlice = DCopySrc_BaseSlice{
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

// 基础类型slice测试
func TestDCopy_BaseWithSlice(t *testing.T) {
	var dst DCopyDst_BaseSlice

	err := Preheat(&dst, &testSrc_BaseSlice)
	dst = DCopyDst_BaseSlice{}
	assert.NoError(t, err)

	// fmt.Printf("%v\n", cacheAllFunc)
	// fmt.Printf("%x\n", getSliceHeaderPtr(unsafe.Pointer(&dst.SliceBool)).Data)
	// fmt.Printf("%d\n", getSliceHeaderPtr(unsafe.Pointer(&dst.SliceBool)).Len)
	// fmt.Printf("%d\n", getSliceHeaderPtr(unsafe.Pointer(&dst.SliceBool)).Cap)

	err = Copy(&dst, &testSrc_BaseSlice, WithUsePreheat())
	assert.NoError(t, err)

	var dst2 DCopyDst_BaseSlice = DCopyDst_BaseSlice(testSrc_BaseSlice)
	assert.Equal(t, dst, dst2)
	// fmt.Printf("%x\n", getSliceHeaderPtr(unsafe.Pointer(&dst.SliceBool)).Data)
	// fmt.Printf("%d\n", getSliceHeaderPtr(unsafe.Pointer(&dst.SliceBool)).Len)
	// fmt.Printf("%d\n", getSliceHeaderPtr(unsafe.Pointer(&dst.SliceBool)).Cap)
	// fmt.Printf("%x\n", getSliceHeaderPtr(unsafe.Pointer(&testSrc_BaseSlice.SliceBool)).Data)
	// fmt.Printf("%d\n", getSliceHeaderPtr(unsafe.Pointer(&testSrc_BaseSlice.SliceBool)).Len)
	// fmt.Printf("%d\n", getSliceHeaderPtr(unsafe.Pointer(&testSrc_BaseSlice.SliceBool)).Cap)
}
