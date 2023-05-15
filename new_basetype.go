package dcopy

import (
	"reflect"
	"unsafe"
)

func newUint(dst interface{}) interface{} {
	d := *(*uint)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newUint8(dst interface{}) interface{} {
	d := *(*uint8)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newUint16(dst interface{}) interface{} {
	d := *(*uint16)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newUint32(dst interface{}) interface{} {
	d := *(*uint32)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newUint64(dst interface{}) interface{} {
	d := *(*uint64)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newInt(dst interface{}) interface{} {
	d := *(*int)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newInt8(dst interface{}) interface{} {
	d := *(*int8)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newInt16(dst interface{}) interface{} {
	d := *(*int16)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newInt32(dst interface{}) interface{} {
	d := *(*int32)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newInt64(dst interface{}) interface{} {
	d := *(*int64)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

// func newString(dst interface{}) interface{} {
// 	d := *(*string)((*emptyInterface)(unsafe.Pointer(&dst)).word)
// 	return d
// }

func newString(dst interface{}) interface{} {
	d := *(*string)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newFloat32(dst interface{}) interface{} {
	d := *(*float32)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newFloat64(dst interface{}) interface{} {
	d := *(*float64)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newBool(dst interface{}) interface{} {
	d := *(*bool)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newComplex64(dst interface{}) interface{} {
	d := *(*complex64)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newComplex128(dst interface{}) interface{} {
	d := *(*complex128)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func newUintptr(dst interface{}) interface{} {
	d := *(*uintptr)((*emptyInterface)(unsafe.Pointer(&dst)).word)
	return d
}

func getNewBaseType(k reflect.Kind, dst interface{}) interface{} {
	newBaseType, ok := newBaseTypeTable[k]
	if ok {
		return newBaseType(dst)
	}
	return nil
}

var newBaseTypeTable = map[reflect.Kind]newBaseTypeFunc{
	reflect.Int:        newInt,
	reflect.Int8:       newInt8,
	reflect.Int16:      newInt16,
	reflect.Int32:      newInt32,
	reflect.Int64:      newInt64,
	reflect.Uint:       newUint,
	reflect.Uint8:      newUint8,
	reflect.Uint16:     newUint16,
	reflect.Uint32:     newUint32,
	reflect.Uint64:     newUint64,
	reflect.String:     newString,
	reflect.Float32:    newFloat32,
	reflect.Float64:    newFloat64,
	reflect.Bool:       newBool,
	reflect.Complex64:  newComplex64,
	reflect.Complex128: newComplex128,
	reflect.Uintptr:    newUintptr,
}
