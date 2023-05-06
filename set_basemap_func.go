package dcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 自动生成的代码， 不要修改
// 生成命令位于 set_basemap_test.go

func setBaseMapIntInt(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntInt8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntInt16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntInt32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntInt64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntUint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntUint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntUint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntUint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntUint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntFloat32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntFloat64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntString(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapIntBool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt8Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int8]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int8]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int8]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt16Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int16]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int16]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int16]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt32Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int32]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int32]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int32]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapInt64Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[int64]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[int64]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[int64]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintInt(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintInt8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintInt16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintInt32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintInt64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintUint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintUint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintUint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintUint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintUint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintFloat32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintFloat64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintString(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUintBool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint8Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint8]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint8]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint8]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint16Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint16]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint16]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint16]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint32Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint32]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint32]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint32]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapUint64Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[uint64]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[uint64]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[uint64]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat32Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float32]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float32]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float32]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Int(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Int8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Int16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Int32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Int64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Uint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Uint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Uint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Uint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Uint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Float32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Float64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64String(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapFloat64Bool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[float64]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[float64]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[float64]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringInt(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringInt8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringInt16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringInt32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringInt64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringUint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringUint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringUint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringUint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringUint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringFloat32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringFloat64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringString(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapStringBool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[string]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[string]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[string]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolInt(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]int)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]int)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]int, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolInt8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]int8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]int8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]int8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolInt16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]int16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]int16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]int16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolInt32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]int32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]int32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]int32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolInt64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]int64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]int64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]int64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolUint(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]uint)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]uint)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]uint, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolUint8(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]uint8)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]uint8)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]uint8, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolUint16(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]uint16)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]uint16)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]uint16, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolUint32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]uint32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]uint32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]uint32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolUint64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]uint64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]uint64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]uint64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolFloat32(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]float32)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]float32)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]float32, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolFloat64(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]float64)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]float64)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]float64, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolString(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]string)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]string)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]string, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

func setBaseMapBoolBool(dstAddr, srcAddr unsafe.Pointer) {
	src := (*map[bool]bool)(srcAddr)
	if len(*src) == 0 {
		return
	}
	dst := (*map[bool]bool)(dstAddr)
	if dst == nil || len(*dst) == 0 {
		*dst = make(map[bool]bool, len(*src))
	}

	for k, v := range *src {
		(*dst)[k] = v
	}
}

var baseMapTab = setBaseMapFuncTab{

	baseMapKind{key: reflect.Int, val: reflect.Int}: setBaseMapIntInt,

	baseMapKind{key: reflect.Int, val: reflect.Int8}: setBaseMapIntInt8,

	baseMapKind{key: reflect.Int, val: reflect.Int16}: setBaseMapIntInt16,

	baseMapKind{key: reflect.Int, val: reflect.Int32}: setBaseMapIntInt32,

	baseMapKind{key: reflect.Int, val: reflect.Int64}: setBaseMapIntInt64,

	baseMapKind{key: reflect.Int, val: reflect.Uint}: setBaseMapIntUint,

	baseMapKind{key: reflect.Int, val: reflect.Uint8}: setBaseMapIntUint8,

	baseMapKind{key: reflect.Int, val: reflect.Uint16}: setBaseMapIntUint16,

	baseMapKind{key: reflect.Int, val: reflect.Uint32}: setBaseMapIntUint32,

	baseMapKind{key: reflect.Int, val: reflect.Uint64}: setBaseMapIntUint64,

	baseMapKind{key: reflect.Int, val: reflect.Float32}: setBaseMapIntFloat32,

	baseMapKind{key: reflect.Int, val: reflect.Float64}: setBaseMapIntFloat64,

	baseMapKind{key: reflect.Int, val: reflect.String}: setBaseMapIntString,

	baseMapKind{key: reflect.Int, val: reflect.Bool}: setBaseMapIntBool,

	baseMapKind{key: reflect.Int8, val: reflect.Int}: setBaseMapInt8Int,

	baseMapKind{key: reflect.Int8, val: reflect.Int8}: setBaseMapInt8Int8,

	baseMapKind{key: reflect.Int8, val: reflect.Int16}: setBaseMapInt8Int16,

	baseMapKind{key: reflect.Int8, val: reflect.Int32}: setBaseMapInt8Int32,

	baseMapKind{key: reflect.Int8, val: reflect.Int64}: setBaseMapInt8Int64,

	baseMapKind{key: reflect.Int8, val: reflect.Uint}: setBaseMapInt8Uint,

	baseMapKind{key: reflect.Int8, val: reflect.Uint8}: setBaseMapInt8Uint8,

	baseMapKind{key: reflect.Int8, val: reflect.Uint16}: setBaseMapInt8Uint16,

	baseMapKind{key: reflect.Int8, val: reflect.Uint32}: setBaseMapInt8Uint32,

	baseMapKind{key: reflect.Int8, val: reflect.Uint64}: setBaseMapInt8Uint64,

	baseMapKind{key: reflect.Int8, val: reflect.Float32}: setBaseMapInt8Float32,

	baseMapKind{key: reflect.Int8, val: reflect.Float64}: setBaseMapInt8Float64,

	baseMapKind{key: reflect.Int8, val: reflect.String}: setBaseMapInt8String,

	baseMapKind{key: reflect.Int8, val: reflect.Bool}: setBaseMapInt8Bool,

	baseMapKind{key: reflect.Int16, val: reflect.Int}: setBaseMapInt16Int,

	baseMapKind{key: reflect.Int16, val: reflect.Int8}: setBaseMapInt16Int8,

	baseMapKind{key: reflect.Int16, val: reflect.Int16}: setBaseMapInt16Int16,

	baseMapKind{key: reflect.Int16, val: reflect.Int32}: setBaseMapInt16Int32,

	baseMapKind{key: reflect.Int16, val: reflect.Int64}: setBaseMapInt16Int64,

	baseMapKind{key: reflect.Int16, val: reflect.Uint}: setBaseMapInt16Uint,

	baseMapKind{key: reflect.Int16, val: reflect.Uint8}: setBaseMapInt16Uint8,

	baseMapKind{key: reflect.Int16, val: reflect.Uint16}: setBaseMapInt16Uint16,

	baseMapKind{key: reflect.Int16, val: reflect.Uint32}: setBaseMapInt16Uint32,

	baseMapKind{key: reflect.Int16, val: reflect.Uint64}: setBaseMapInt16Uint64,

	baseMapKind{key: reflect.Int16, val: reflect.Float32}: setBaseMapInt16Float32,

	baseMapKind{key: reflect.Int16, val: reflect.Float64}: setBaseMapInt16Float64,

	baseMapKind{key: reflect.Int16, val: reflect.String}: setBaseMapInt16String,

	baseMapKind{key: reflect.Int16, val: reflect.Bool}: setBaseMapInt16Bool,

	baseMapKind{key: reflect.Int32, val: reflect.Int}: setBaseMapInt32Int,

	baseMapKind{key: reflect.Int32, val: reflect.Int8}: setBaseMapInt32Int8,

	baseMapKind{key: reflect.Int32, val: reflect.Int16}: setBaseMapInt32Int16,

	baseMapKind{key: reflect.Int32, val: reflect.Int32}: setBaseMapInt32Int32,

	baseMapKind{key: reflect.Int32, val: reflect.Int64}: setBaseMapInt32Int64,

	baseMapKind{key: reflect.Int32, val: reflect.Uint}: setBaseMapInt32Uint,

	baseMapKind{key: reflect.Int32, val: reflect.Uint8}: setBaseMapInt32Uint8,

	baseMapKind{key: reflect.Int32, val: reflect.Uint16}: setBaseMapInt32Uint16,

	baseMapKind{key: reflect.Int32, val: reflect.Uint32}: setBaseMapInt32Uint32,

	baseMapKind{key: reflect.Int32, val: reflect.Uint64}: setBaseMapInt32Uint64,

	baseMapKind{key: reflect.Int32, val: reflect.Float32}: setBaseMapInt32Float32,

	baseMapKind{key: reflect.Int32, val: reflect.Float64}: setBaseMapInt32Float64,

	baseMapKind{key: reflect.Int32, val: reflect.String}: setBaseMapInt32String,

	baseMapKind{key: reflect.Int32, val: reflect.Bool}: setBaseMapInt32Bool,

	baseMapKind{key: reflect.Int64, val: reflect.Int}: setBaseMapInt64Int,

	baseMapKind{key: reflect.Int64, val: reflect.Int8}: setBaseMapInt64Int8,

	baseMapKind{key: reflect.Int64, val: reflect.Int16}: setBaseMapInt64Int16,

	baseMapKind{key: reflect.Int64, val: reflect.Int32}: setBaseMapInt64Int32,

	baseMapKind{key: reflect.Int64, val: reflect.Int64}: setBaseMapInt64Int64,

	baseMapKind{key: reflect.Int64, val: reflect.Uint}: setBaseMapInt64Uint,

	baseMapKind{key: reflect.Int64, val: reflect.Uint8}: setBaseMapInt64Uint8,

	baseMapKind{key: reflect.Int64, val: reflect.Uint16}: setBaseMapInt64Uint16,

	baseMapKind{key: reflect.Int64, val: reflect.Uint32}: setBaseMapInt64Uint32,

	baseMapKind{key: reflect.Int64, val: reflect.Uint64}: setBaseMapInt64Uint64,

	baseMapKind{key: reflect.Int64, val: reflect.Float32}: setBaseMapInt64Float32,

	baseMapKind{key: reflect.Int64, val: reflect.Float64}: setBaseMapInt64Float64,

	baseMapKind{key: reflect.Int64, val: reflect.String}: setBaseMapInt64String,

	baseMapKind{key: reflect.Int64, val: reflect.Bool}: setBaseMapInt64Bool,

	baseMapKind{key: reflect.Uint, val: reflect.Int}: setBaseMapUintInt,

	baseMapKind{key: reflect.Uint, val: reflect.Int8}: setBaseMapUintInt8,

	baseMapKind{key: reflect.Uint, val: reflect.Int16}: setBaseMapUintInt16,

	baseMapKind{key: reflect.Uint, val: reflect.Int32}: setBaseMapUintInt32,

	baseMapKind{key: reflect.Uint, val: reflect.Int64}: setBaseMapUintInt64,

	baseMapKind{key: reflect.Uint, val: reflect.Uint}: setBaseMapUintUint,

	baseMapKind{key: reflect.Uint, val: reflect.Uint8}: setBaseMapUintUint8,

	baseMapKind{key: reflect.Uint, val: reflect.Uint16}: setBaseMapUintUint16,

	baseMapKind{key: reflect.Uint, val: reflect.Uint32}: setBaseMapUintUint32,

	baseMapKind{key: reflect.Uint, val: reflect.Uint64}: setBaseMapUintUint64,

	baseMapKind{key: reflect.Uint, val: reflect.Float32}: setBaseMapUintFloat32,

	baseMapKind{key: reflect.Uint, val: reflect.Float64}: setBaseMapUintFloat64,

	baseMapKind{key: reflect.Uint, val: reflect.String}: setBaseMapUintString,

	baseMapKind{key: reflect.Uint, val: reflect.Bool}: setBaseMapUintBool,

	baseMapKind{key: reflect.Uint8, val: reflect.Int}: setBaseMapUint8Int,

	baseMapKind{key: reflect.Uint8, val: reflect.Int8}: setBaseMapUint8Int8,

	baseMapKind{key: reflect.Uint8, val: reflect.Int16}: setBaseMapUint8Int16,

	baseMapKind{key: reflect.Uint8, val: reflect.Int32}: setBaseMapUint8Int32,

	baseMapKind{key: reflect.Uint8, val: reflect.Int64}: setBaseMapUint8Int64,

	baseMapKind{key: reflect.Uint8, val: reflect.Uint}: setBaseMapUint8Uint,

	baseMapKind{key: reflect.Uint8, val: reflect.Uint8}: setBaseMapUint8Uint8,

	baseMapKind{key: reflect.Uint8, val: reflect.Uint16}: setBaseMapUint8Uint16,

	baseMapKind{key: reflect.Uint8, val: reflect.Uint32}: setBaseMapUint8Uint32,

	baseMapKind{key: reflect.Uint8, val: reflect.Uint64}: setBaseMapUint8Uint64,

	baseMapKind{key: reflect.Uint8, val: reflect.Float32}: setBaseMapUint8Float32,

	baseMapKind{key: reflect.Uint8, val: reflect.Float64}: setBaseMapUint8Float64,

	baseMapKind{key: reflect.Uint8, val: reflect.String}: setBaseMapUint8String,

	baseMapKind{key: reflect.Uint8, val: reflect.Bool}: setBaseMapUint8Bool,

	baseMapKind{key: reflect.Uint16, val: reflect.Int}: setBaseMapUint16Int,

	baseMapKind{key: reflect.Uint16, val: reflect.Int8}: setBaseMapUint16Int8,

	baseMapKind{key: reflect.Uint16, val: reflect.Int16}: setBaseMapUint16Int16,

	baseMapKind{key: reflect.Uint16, val: reflect.Int32}: setBaseMapUint16Int32,

	baseMapKind{key: reflect.Uint16, val: reflect.Int64}: setBaseMapUint16Int64,

	baseMapKind{key: reflect.Uint16, val: reflect.Uint}: setBaseMapUint16Uint,

	baseMapKind{key: reflect.Uint16, val: reflect.Uint8}: setBaseMapUint16Uint8,

	baseMapKind{key: reflect.Uint16, val: reflect.Uint16}: setBaseMapUint16Uint16,

	baseMapKind{key: reflect.Uint16, val: reflect.Uint32}: setBaseMapUint16Uint32,

	baseMapKind{key: reflect.Uint16, val: reflect.Uint64}: setBaseMapUint16Uint64,

	baseMapKind{key: reflect.Uint16, val: reflect.Float32}: setBaseMapUint16Float32,

	baseMapKind{key: reflect.Uint16, val: reflect.Float64}: setBaseMapUint16Float64,

	baseMapKind{key: reflect.Uint16, val: reflect.String}: setBaseMapUint16String,

	baseMapKind{key: reflect.Uint16, val: reflect.Bool}: setBaseMapUint16Bool,

	baseMapKind{key: reflect.Uint32, val: reflect.Int}: setBaseMapUint32Int,

	baseMapKind{key: reflect.Uint32, val: reflect.Int8}: setBaseMapUint32Int8,

	baseMapKind{key: reflect.Uint32, val: reflect.Int16}: setBaseMapUint32Int16,

	baseMapKind{key: reflect.Uint32, val: reflect.Int32}: setBaseMapUint32Int32,

	baseMapKind{key: reflect.Uint32, val: reflect.Int64}: setBaseMapUint32Int64,

	baseMapKind{key: reflect.Uint32, val: reflect.Uint}: setBaseMapUint32Uint,

	baseMapKind{key: reflect.Uint32, val: reflect.Uint8}: setBaseMapUint32Uint8,

	baseMapKind{key: reflect.Uint32, val: reflect.Uint16}: setBaseMapUint32Uint16,

	baseMapKind{key: reflect.Uint32, val: reflect.Uint32}: setBaseMapUint32Uint32,

	baseMapKind{key: reflect.Uint32, val: reflect.Uint64}: setBaseMapUint32Uint64,

	baseMapKind{key: reflect.Uint32, val: reflect.Float32}: setBaseMapUint32Float32,

	baseMapKind{key: reflect.Uint32, val: reflect.Float64}: setBaseMapUint32Float64,

	baseMapKind{key: reflect.Uint32, val: reflect.String}: setBaseMapUint32String,

	baseMapKind{key: reflect.Uint32, val: reflect.Bool}: setBaseMapUint32Bool,

	baseMapKind{key: reflect.Uint64, val: reflect.Int}: setBaseMapUint64Int,

	baseMapKind{key: reflect.Uint64, val: reflect.Int8}: setBaseMapUint64Int8,

	baseMapKind{key: reflect.Uint64, val: reflect.Int16}: setBaseMapUint64Int16,

	baseMapKind{key: reflect.Uint64, val: reflect.Int32}: setBaseMapUint64Int32,

	baseMapKind{key: reflect.Uint64, val: reflect.Int64}: setBaseMapUint64Int64,

	baseMapKind{key: reflect.Uint64, val: reflect.Uint}: setBaseMapUint64Uint,

	baseMapKind{key: reflect.Uint64, val: reflect.Uint8}: setBaseMapUint64Uint8,

	baseMapKind{key: reflect.Uint64, val: reflect.Uint16}: setBaseMapUint64Uint16,

	baseMapKind{key: reflect.Uint64, val: reflect.Uint32}: setBaseMapUint64Uint32,

	baseMapKind{key: reflect.Uint64, val: reflect.Uint64}: setBaseMapUint64Uint64,

	baseMapKind{key: reflect.Uint64, val: reflect.Float32}: setBaseMapUint64Float32,

	baseMapKind{key: reflect.Uint64, val: reflect.Float64}: setBaseMapUint64Float64,

	baseMapKind{key: reflect.Uint64, val: reflect.String}: setBaseMapUint64String,

	baseMapKind{key: reflect.Uint64, val: reflect.Bool}: setBaseMapUint64Bool,

	baseMapKind{key: reflect.Float32, val: reflect.Int}: setBaseMapFloat32Int,

	baseMapKind{key: reflect.Float32, val: reflect.Int8}: setBaseMapFloat32Int8,

	baseMapKind{key: reflect.Float32, val: reflect.Int16}: setBaseMapFloat32Int16,

	baseMapKind{key: reflect.Float32, val: reflect.Int32}: setBaseMapFloat32Int32,

	baseMapKind{key: reflect.Float32, val: reflect.Int64}: setBaseMapFloat32Int64,

	baseMapKind{key: reflect.Float32, val: reflect.Uint}: setBaseMapFloat32Uint,

	baseMapKind{key: reflect.Float32, val: reflect.Uint8}: setBaseMapFloat32Uint8,

	baseMapKind{key: reflect.Float32, val: reflect.Uint16}: setBaseMapFloat32Uint16,

	baseMapKind{key: reflect.Float32, val: reflect.Uint32}: setBaseMapFloat32Uint32,

	baseMapKind{key: reflect.Float32, val: reflect.Uint64}: setBaseMapFloat32Uint64,

	baseMapKind{key: reflect.Float32, val: reflect.Float32}: setBaseMapFloat32Float32,

	baseMapKind{key: reflect.Float32, val: reflect.Float64}: setBaseMapFloat32Float64,

	baseMapKind{key: reflect.Float32, val: reflect.String}: setBaseMapFloat32String,

	baseMapKind{key: reflect.Float32, val: reflect.Bool}: setBaseMapFloat32Bool,

	baseMapKind{key: reflect.Float64, val: reflect.Int}: setBaseMapFloat64Int,

	baseMapKind{key: reflect.Float64, val: reflect.Int8}: setBaseMapFloat64Int8,

	baseMapKind{key: reflect.Float64, val: reflect.Int16}: setBaseMapFloat64Int16,

	baseMapKind{key: reflect.Float64, val: reflect.Int32}: setBaseMapFloat64Int32,

	baseMapKind{key: reflect.Float64, val: reflect.Int64}: setBaseMapFloat64Int64,

	baseMapKind{key: reflect.Float64, val: reflect.Uint}: setBaseMapFloat64Uint,

	baseMapKind{key: reflect.Float64, val: reflect.Uint8}: setBaseMapFloat64Uint8,

	baseMapKind{key: reflect.Float64, val: reflect.Uint16}: setBaseMapFloat64Uint16,

	baseMapKind{key: reflect.Float64, val: reflect.Uint32}: setBaseMapFloat64Uint32,

	baseMapKind{key: reflect.Float64, val: reflect.Uint64}: setBaseMapFloat64Uint64,

	baseMapKind{key: reflect.Float64, val: reflect.Float32}: setBaseMapFloat64Float32,

	baseMapKind{key: reflect.Float64, val: reflect.Float64}: setBaseMapFloat64Float64,

	baseMapKind{key: reflect.Float64, val: reflect.String}: setBaseMapFloat64String,

	baseMapKind{key: reflect.Float64, val: reflect.Bool}: setBaseMapFloat64Bool,

	baseMapKind{key: reflect.String, val: reflect.Int}: setBaseMapStringInt,

	baseMapKind{key: reflect.String, val: reflect.Int8}: setBaseMapStringInt8,

	baseMapKind{key: reflect.String, val: reflect.Int16}: setBaseMapStringInt16,

	baseMapKind{key: reflect.String, val: reflect.Int32}: setBaseMapStringInt32,

	baseMapKind{key: reflect.String, val: reflect.Int64}: setBaseMapStringInt64,

	baseMapKind{key: reflect.String, val: reflect.Uint}: setBaseMapStringUint,

	baseMapKind{key: reflect.String, val: reflect.Uint8}: setBaseMapStringUint8,

	baseMapKind{key: reflect.String, val: reflect.Uint16}: setBaseMapStringUint16,

	baseMapKind{key: reflect.String, val: reflect.Uint32}: setBaseMapStringUint32,

	baseMapKind{key: reflect.String, val: reflect.Uint64}: setBaseMapStringUint64,

	baseMapKind{key: reflect.String, val: reflect.Float32}: setBaseMapStringFloat32,

	baseMapKind{key: reflect.String, val: reflect.Float64}: setBaseMapStringFloat64,

	baseMapKind{key: reflect.String, val: reflect.String}: setBaseMapStringString,

	baseMapKind{key: reflect.String, val: reflect.Bool}: setBaseMapStringBool,

	baseMapKind{key: reflect.Bool, val: reflect.Int}: setBaseMapBoolInt,

	baseMapKind{key: reflect.Bool, val: reflect.Int8}: setBaseMapBoolInt8,

	baseMapKind{key: reflect.Bool, val: reflect.Int16}: setBaseMapBoolInt16,

	baseMapKind{key: reflect.Bool, val: reflect.Int32}: setBaseMapBoolInt32,

	baseMapKind{key: reflect.Bool, val: reflect.Int64}: setBaseMapBoolInt64,

	baseMapKind{key: reflect.Bool, val: reflect.Uint}: setBaseMapBoolUint,

	baseMapKind{key: reflect.Bool, val: reflect.Uint8}: setBaseMapBoolUint8,

	baseMapKind{key: reflect.Bool, val: reflect.Uint16}: setBaseMapBoolUint16,

	baseMapKind{key: reflect.Bool, val: reflect.Uint32}: setBaseMapBoolUint32,

	baseMapKind{key: reflect.Bool, val: reflect.Uint64}: setBaseMapBoolUint64,

	baseMapKind{key: reflect.Bool, val: reflect.Float32}: setBaseMapBoolFloat32,

	baseMapKind{key: reflect.Bool, val: reflect.Float64}: setBaseMapBoolFloat64,

	baseMapKind{key: reflect.Bool, val: reflect.String}: setBaseMapBoolString,

	baseMapKind{key: reflect.Bool, val: reflect.Bool}: setBaseMapBoolBool,
}

func getSetBaseMapFunc(key reflect.Kind, val reflect.Kind) setUnsafeFunc {
	k := baseMapKind{key: key, val: val}
	f, ok := baseMapTab[k]
	if !ok {
		panic(fmt.Sprintf("not support type:key %v val: %v", key, val))
	}
	return f
}
