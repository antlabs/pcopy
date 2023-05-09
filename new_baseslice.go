package dcopy

import (
	"reflect"
	"unsafe"
)

func newIntSlice(n int) unsafe.Pointer {
	rv := make([]int, n)
	return unsafe.Pointer(&rv)
}

func newInt8Slice(n int) unsafe.Pointer {
	rv := make([]int8, n)
	return unsafe.Pointer(&rv)
}

func newInt16Slice(n int) unsafe.Pointer {
	rv := make([]int16, n)
	return unsafe.Pointer(&rv)
}

func newInt32Slice(n int) unsafe.Pointer {
	rv := make([]int32, n)
	return unsafe.Pointer(&rv)
}

func newInt64Slice(n int) unsafe.Pointer {
	rv := make([]int64, n)
	return unsafe.Pointer(&rv)
}

func newUintSlice(n int) unsafe.Pointer {
	rv := make([]uint, n)
	return unsafe.Pointer(&rv)
}

func newUint8Slice(n int) unsafe.Pointer {
	rv := make([]uint8, n)
	return unsafe.Pointer(&rv)
}

func newUint16Slice(n int) unsafe.Pointer {
	rv := make([]uint16, n)
	return unsafe.Pointer(&rv)
}

func newUint32Slice(n int) unsafe.Pointer {
	rv := make([]uint32, n)
	return unsafe.Pointer(&rv)
}

func newUint64Slice(n int) unsafe.Pointer {
	rv := make([]uint64, n)
	return unsafe.Pointer(&rv)
}

func newStringSlice(n int) unsafe.Pointer {
	rv := make([]string, n)
	return unsafe.Pointer(&rv)
}

func newFloat32Slice(n int) unsafe.Pointer {
	rv := make([]float32, n)
	return unsafe.Pointer(&rv)
}

func newFloat64Slice(n int) unsafe.Pointer {
	rv := make([]float64, n)
	return unsafe.Pointer(&rv)
}

func newBoolSlice(n int) unsafe.Pointer {
	rv := make([]bool, n)
	return unsafe.Pointer(&rv)
}

func newComplex64Slice(n int) unsafe.Pointer {
	rv := make([]complex64, n)
	return unsafe.Pointer(&rv)
}

func newComplex128Slice(n int) unsafe.Pointer {
	rv := make([]complex128, n)
	return unsafe.Pointer(&rv)
}

func getNewBaseSliceType(k reflect.Kind, len int) unsafe.Pointer {
	newBaseSliceType, ok := newBaseSliceTypeTable[k]
	if ok {
		return newBaseSliceType(len)
	}
	return nil
}

var newBaseSliceTypeTable = map[reflect.Kind]func(len int) unsafe.Pointer{
	reflect.Int:        newIntSlice,
	reflect.Int8:       newInt8Slice,
	reflect.Int16:      newInt16Slice,
	reflect.Int32:      newInt32Slice,
	reflect.Int64:      newInt64Slice,
	reflect.Uint:       newUintSlice,
	reflect.Uint8:      newUint8Slice,
	reflect.Uint16:     newUint16Slice,
	reflect.Uint32:     newUint32Slice,
	reflect.Uint64:     newUint64Slice,
	reflect.String:     newStringSlice,
	reflect.Float32:    newFloat32Slice,
	reflect.Float64:    newFloat64Slice,
	reflect.Bool:       newBoolSlice,
	reflect.Complex64:  newComplex64Slice,
	reflect.Complex128: newComplex128Slice,
}
