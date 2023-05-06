package dcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

var copyBaseSliceTab = setUnsafeFuncTab{
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

func getSetBaseSliceFunc(t reflect.Kind) setUnsafeFunc {
	f, ok := copyBaseSliceTab[t]
	if !ok {
		panic(fmt.Sprintf("not support type:%T", t))
	}
	return f
}

func setBaseSliceBool(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]bool)(dstAddr)
	srcSlice := *(*[]bool)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]bool)(dstAddr) = make([]bool, len(srcSlice))
		dstSlice = *(*[]bool)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int)(dstAddr)
	srcSlice := *(*[]int)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]int)(dstAddr) = make([]int, len(srcSlice))
		dstSlice = *(*[]int)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt8(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int8)(dstAddr)
	srcSlice := *(*[]int8)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]int8)(dstAddr) = make([]int8, len(srcSlice))
		dstSlice = *(*[]int8)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt16(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int16)(dstAddr)
	srcSlice := *(*[]int16)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]int16)(dstAddr) = make([]int16, len(srcSlice))
		dstSlice = *(*[]int16)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt32(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int32)(dstAddr)
	srcSlice := *(*[]int32)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]int32)(dstAddr) = make([]int32, len(srcSlice))
		dstSlice = *(*[]int32)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceInt64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]int64)(dstAddr)
	srcSlice := *(*[]int64)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]int64)(dstAddr) = make([]int64, len(srcSlice))
		dstSlice = *(*[]int64)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint)(dstAddr)
	srcSlice := *(*[]uint)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]uint)(dstAddr) = make([]uint, len(srcSlice))
		dstSlice = *(*[]uint)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint8(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint8)(dstAddr)
	srcSlice := *(*[]uint8)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]uint8)(dstAddr) = make([]uint8, len(srcSlice))
		dstSlice = *(*[]uint8)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint16(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint16)(dstAddr)
	srcSlice := *(*[]uint16)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]uint16)(dstAddr) = make([]uint16, len(srcSlice))
		dstSlice = *(*[]uint16)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint32(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint32)(dstAddr)
	srcSlice := *(*[]uint32)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]uint32)(dstAddr) = make([]uint32, len(srcSlice))
		dstSlice = *(*[]uint32)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceUint64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]uint64)(dstAddr)
	srcSlice := *(*[]uint64)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]uint64)(dstAddr) = make([]uint64, len(srcSlice))
		dstSlice = *(*[]uint64)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceString(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]string)(dstAddr)
	srcSlice := *(*[]string)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]string)(dstAddr) = make([]string, len(srcSlice))
		dstSlice = *(*[]string)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceFloat32(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]float32)(dstAddr)
	srcSlice := *(*[]float32)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]float32)(dstAddr) = make([]float32, len(srcSlice))
		dstSlice = *(*[]float32)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceFloat64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]float64)(dstAddr)
	srcSlice := *(*[]float64)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]float64)(dstAddr) = make([]float64, len(srcSlice))
		dstSlice = *(*[]float64)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceComplex64(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]complex64)(dstAddr)
	srcSlice := *(*[]complex64)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]complex64)(dstAddr) = make([]complex64, len(srcSlice))
		dstSlice = *(*[]complex64)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}

func setBaseSliceComplex128(dstAddr, srcAddr unsafe.Pointer) {
	dstSlice := *(*[]complex128)(dstAddr)
	srcSlice := *(*[]complex128)(srcAddr)
	if cap(dstSlice) < len(srcSlice) {
		*(*[]complex128)(dstAddr) = make([]complex128, len(srcSlice))
		dstSlice = *(*[]complex128)(dstAddr)
	}
	copy(dstSlice, srcSlice)
}
