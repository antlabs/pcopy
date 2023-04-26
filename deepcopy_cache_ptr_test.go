package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FastCopyDst struct {
	// 所有基础类型
	Bool       bool
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	String     string
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
}

type FastCopySrc struct {
	Bool       bool
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	String     string
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
}

var testSrc = FastCopySrc{
	Bool:       true,
	Int:        1,
	Int8:       2,
	Int16:      3,
	Int32:      4,
	Int64:      5,
	Uint:       6,
	Uint8:      7,
	Uint16:     8,
	Uint32:     9,
	Uint64:     10,
	String:     "11",
	Float32:    12.0,
	Float64:    13.0,
	Complex64:  14.0,
	Complex128: 15.0,
}

func TestFastCopy(t *testing.T) {
	var dst FastCopyDst

	err := Preheat(&dst, &testSrc)
	assert.NoError(t, err)

	printCacheAllFunc()
	// fmt.Printf("%v\n", cacheAllFunc)
	err = CopyEx(&dst, &testSrc, WithPreheat())

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
	dst2.Complex64 = 14.0
	dst2.Complex128 = 15.0
	assert.NoError(t, err)
	assert.Equal(t, dst, dst2)
}
