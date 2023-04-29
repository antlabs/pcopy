package deepcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

type (
	setFunc    func(dstAddr, srcAddr unsafe.Pointer)
	setFuncTab map[reflect.Kind]setFunc
)

var copyTab = setFuncTab{
	reflect.Bool:       setBool,
	reflect.Int:        setInt,
	reflect.Int8:       setInt8,
	reflect.Int16:      setInt16,
	reflect.Int32:      setInt32,
	reflect.Int64:      setInt64,
	reflect.Uint:       setUint,
	reflect.Uint8:      setUint8,
	reflect.Uint16:     setUint16,
	reflect.Uint32:     setUint32,
	reflect.Uint64:     setUint64,
	reflect.String:     setString,
	reflect.Float32:    setFloat32,
	reflect.Float64:    setFloat64,
	reflect.Complex64:  setComplex64,
	reflect.Complex128: setComplex128,
}

func getSetFunc(t reflect.Kind) setFunc {
	f, ok := copyTab[t]
	if !ok {
		panic(fmt.Sprintf("not support type:%T", t))
	}
	return f
}

func setBool(dstAddr, srcAddr unsafe.Pointer) {
	*(*bool)(dstAddr) = *(*bool)(srcAddr)
}

func setInt(dstAddr, srcAddr unsafe.Pointer) {
	*(*int)(dstAddr) = *(*int)(srcAddr)
}

func setInt8(dstAddr, srcAddr unsafe.Pointer) {
	*(*int8)(dstAddr) = *(*int8)(srcAddr)
}

func setInt16(dstAddr, srcAddr unsafe.Pointer) {
	*(*int16)(dstAddr) = *(*int16)(srcAddr)
}

func setInt32(dstAddr, srcAddr unsafe.Pointer) {
	*(*int32)(dstAddr) = *(*int32)(srcAddr)
}

func setInt64(dstAddr, srcAddr unsafe.Pointer) {
	*(*int64)(dstAddr) = *(*int64)(srcAddr)
}

func setUint(dstAddr, srcAddr unsafe.Pointer) {
	*(*uint)(dstAddr) = *(*uint)(srcAddr)
}

func setUint8(dstAddr, srcAddr unsafe.Pointer) {
	*(*uint8)(dstAddr) = *(*uint8)(srcAddr)
}

func setUint16(dstAddr, srcAddr unsafe.Pointer) {
	*(*uint16)(dstAddr) = *(*uint16)(srcAddr)
}

func setUint32(dstAddr, srcAddr unsafe.Pointer) {
	*(*uint32)(dstAddr) = *(*uint32)(srcAddr)
}

func setUint64(dstAddr, srcAddr unsafe.Pointer) {
	*(*uint64)(dstAddr) = *(*uint64)(srcAddr)
}

func setString(dstAddr, srcAddr unsafe.Pointer) {
	*(*string)(dstAddr) = *(*string)(srcAddr)
}

func setFloat32(dstAddr, srcAddr unsafe.Pointer) {
	*(*float32)(dstAddr) = *(*float32)(srcAddr)
}

func setFloat64(dstAddr, srcAddr unsafe.Pointer) {
	*(*float64)(dstAddr) = *(*float64)(srcAddr)
}

func setComplex64(dstAddr, srcAddr unsafe.Pointer) {
	*(*complex64)(dstAddr) = *(*complex64)(srcAddr)
}

func setComplex128(dstAddr, srcAddr unsafe.Pointer) {
	*(*complex128)(dstAddr) = *(*complex128)(srcAddr)
}

var copyBaseSliceTab = setFuncTab{
	reflect.Bool:       setBaseSliceBool,
	reflect.Int:        setBaseSliceInt,
	reflect.Int8:       setBaseSliceInt8,
	reflect.Int16:      setBaseSliceInt16,
	reflect.Int32:      setBaseSliceInt32,
	reflect.Int64:      setBaseSliceInt64,
	reflect.Uint:       setBaseSliceUint,
	reflect.Uint8:      setBaseSliceUint8,
	reflect.Uint16:     setBaseSliceUint16,
	reflect.Uint32:     setBaseSliceUint32,
	reflect.Uint64:     setBaseSliceUint64,
	reflect.String:     setBaseSliceString,
	reflect.Float32:    setBaseSliceFloat32,
	reflect.Float64:    setBaseSliceFloat64,
	reflect.Complex64:  setBaseSliceComplex64,
	reflect.Complex128: setBaseSliceComplex128,
}

func getSetBaseSliceFunc(t reflect.Kind) setFunc {
	f, ok := copyBaseSliceTab[t]
	if !ok {
		panic(fmt.Sprintf("not support type:%T", t))
	}
	return f
}

func setBaseSliceBool(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]bool)(dstAddr)
	srcSlice := *(*[]bool)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int)(dstAddr)
	srcSlice := *(*[]int)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt8(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int8)(dstAddr)
	srcSlice := *(*[]int8)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt16(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int16)(dstAddr)
	srcSlice := *(*[]int16)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt32(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int32)(dstAddr)
	srcSlice := *(*[]int32)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int64)(dstAddr)
	srcSlice := *(*[]int64)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint)(dstAddr)
	srcSlice := *(*[]uint)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint8(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint8)(dstAddr)
	srcSlice := *(*[]uint8)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint16(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint16)(dstAddr)
	srcSlice := *(*[]uint16)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint32(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint32)(dstAddr)
	srcSlice := *(*[]uint32)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint64)(dstAddr)
	srcSlice := *(*[]uint64)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceString(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]string)(dstAddr)
	srcSlice := *(*[]string)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceFloat32(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]float32)(dstAddr)
	srcSlice := *(*[]float32)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceFloat64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]float64)(dstAddr)
	srcSlice := *(*[]float64)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceComplex64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]complex64)(dstAddr)
	srcSlice := *(*[]complex64)(srcAddr)
	copy(dstSlice, srcSlice)
}

func setBaseSliceComplex128(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]complex128)(dstAddr)
	srcSlice := *(*[]complex128)(srcAddr)
	copy(dstSlice, srcSlice)
}
