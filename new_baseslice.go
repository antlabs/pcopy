package dcopy

import (
	"reflect"
	"unsafe"
)

func newIntSlice(sh *reflect.SliceHeader) interface{} {
	rv := make([]int, sh.Len)
	copy(rv, *(*[]int)(unsafe.Pointer(sh)))
	return rv
}

func newInt8Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]int8, sh.Len)
	copy(rv, *(*[]int8)(unsafe.Pointer(sh)))
	return rv
}

func newInt16Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]int16, sh.Len)
	copy(rv, *(*[]int16)(unsafe.Pointer(sh)))
	return rv
}

func newInt32Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]int32, sh.Len)
	copy(rv, *(*[]int32)(unsafe.Pointer(sh)))
	return rv
}

func newInt64Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]int64, sh.Len)
	copy(rv, *(*[]int64)(unsafe.Pointer(sh)))
	return rv
}

func newUintSlice(sh *reflect.SliceHeader) interface{} {
	rv := make([]uint, sh.Len)
	copy(rv, *(*[]uint)(unsafe.Pointer(sh)))
	return rv
}

func newUint8Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]uint8, sh.Len)
	copy(rv, *(*[]uint8)(unsafe.Pointer(sh)))
	return rv
}

func newUint16Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]uint16, sh.Len)
	copy(rv, *(*[]uint16)(unsafe.Pointer(sh)))
	return rv
}

func newUint32Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]uint32, sh.Len)
	copy(rv, *(*[]uint32)(unsafe.Pointer(sh)))
	return rv
}

func newUint64Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]uint64, sh.Len)
	copy(rv, *(*[]uint64)(unsafe.Pointer(sh)))
	return rv
}

func newStringSlice(sh *reflect.SliceHeader) interface{} {
	rv := make([]string, sh.Len)
	copy(rv, *(*[]string)(unsafe.Pointer(sh)))
	return rv
}

func newFloat32Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]float32, sh.Len)
	copy(rv, *(*[]float32)(unsafe.Pointer(sh)))
	return rv
}

func newFloat64Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]float64, sh.Len)
	copy(rv, *(*[]float64)(unsafe.Pointer(sh)))
	return rv
}

func newComplex64Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]complex64, sh.Len)
	copy(rv, *(*[]complex64)(unsafe.Pointer(sh)))
	return rv
}

func newComplex128Slice(sh *reflect.SliceHeader) interface{} {
	rv := make([]complex128, sh.Len)
	copy(rv, *(*[]complex128)(unsafe.Pointer(sh)))
	return rv
}

func newBoolSlice(sh *reflect.SliceHeader) interface{} {
	rv := make([]bool, sh.Len)
	copy(rv, *(*[]bool)(unsafe.Pointer(sh)))
	return rv
}

func newUintptrSlice(sh *reflect.SliceHeader) interface{} {
	rv := make([]uintptr, sh.Len)
	copy(rv, *(*[]uintptr)(unsafe.Pointer(sh)))
	return rv
}

func getNewBaseSliceType(k reflect.Kind, sh *reflect.SliceHeader) interface{} {
	newBaseSliceType, ok := newBaseSliceTypeTable[k]
	if ok {
		return newBaseSliceType(sh)
	}
	return nil
}

var newBaseSliceTypeTable = map[reflect.Kind]func(sh *reflect.SliceHeader) interface{}{
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
	reflect.Uintptr:    newUintptrSlice,
}
