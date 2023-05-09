package dcopy

import (
	"reflect"
	"unsafe"
)

func newInt() unsafe.Pointer {
	return unsafe.Pointer(new(int))
}

func newInt8() unsafe.Pointer {
	return unsafe.Pointer(new(int8))
}

func newInt16() unsafe.Pointer {
	return unsafe.Pointer(new(int16))
}

func newInt32() unsafe.Pointer {
	return unsafe.Pointer(new(int32))
}

func newInt64() unsafe.Pointer {
	return unsafe.Pointer(new(int64))
}

func newUint() unsafe.Pointer {
	return unsafe.Pointer(new(uint))
}

func newUint8() unsafe.Pointer {
	return unsafe.Pointer(new(uint8))
}

func newUint16() unsafe.Pointer {
	return unsafe.Pointer(new(uint16))
}

func newUint32() unsafe.Pointer {
	return unsafe.Pointer(new(uint32))
}

func newUint64() unsafe.Pointer {
	return unsafe.Pointer(new(uint64))
}

func newString() unsafe.Pointer {
	return unsafe.Pointer(new(string))
}

func newFloat32() unsafe.Pointer {
	return unsafe.Pointer(new(float32))
}

func newFloat64() unsafe.Pointer {
	return unsafe.Pointer(new(float64))
}

func newBool() unsafe.Pointer {
	return unsafe.Pointer(new(bool))
}

func newComplex64() unsafe.Pointer {
	return unsafe.Pointer(new(complex64))
}

func newComplex128() unsafe.Pointer {
	return unsafe.Pointer(new(complex128))
}

func getNewBaseType(k reflect.Kind) unsafe.Pointer {
	newBaseType, ok := newBaseTypeTable[k]
	if ok {
		return newBaseType()
	}
	return nil
}

var newBaseTypeTable = map[reflect.Kind]func() unsafe.Pointer{
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
}
