// Copyright [2020-2024] [guonaihong]
package dcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 自动生成的代码， 不要修改
// 生成的代码位于，如下位置
// set_basemap_tmpl.go
// 生成命令位于 set_basemap_test.go

func setSliceElemIsBaseTypeOrPtrInt(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*int)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Int && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]int)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Int {
		srcSlice := *(*[]*int)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]int)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]int, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrInt8(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*int8)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int8)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int8, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int8
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Int8 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]int8)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int8)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int8, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Int8 {
		srcSlice := *(*[]*int8)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]int8)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]int8, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int8
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrInt16(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*int16)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int16)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int16, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int16
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Int16 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]int16)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int16)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int16, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Int16 {
		srcSlice := *(*[]*int16)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]int16)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]int16, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int16
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrInt32(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*int32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int32
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Int32 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]int32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Int32 {
		srcSlice := *(*[]*int32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]int32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]int32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int32
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrInt64(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*int64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Int64 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]int64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*int64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*int64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Int64 {
		srcSlice := *(*[]*int64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]int64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]int64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv int64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrUint(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*uint)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Uint && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]uint)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Uint {
		srcSlice := *(*[]*uint)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]uint)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]uint, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrUint8(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*uint8)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint8)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint8, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint8
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Uint8 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]uint8)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint8)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint8, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Uint8 {
		srcSlice := *(*[]*uint8)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]uint8)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]uint8, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint8
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrUint16(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*uint16)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint16)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint16, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint16
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Uint16 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]uint16)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint16)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint16, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Uint16 {
		srcSlice := *(*[]*uint16)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]uint16)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]uint16, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint16
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrUint32(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*uint32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint32
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Uint32 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]uint32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Uint32 {
		srcSlice := *(*[]*uint32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]uint32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]uint32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint32
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrUint64(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*uint64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Uint64 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]uint64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uint64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uint64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Uint64 {
		srcSlice := *(*[]*uint64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]uint64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]uint64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uint64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrFloat32(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*float32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*float32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*float32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv float32
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Float32 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]float32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*float32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*float32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Float32 {
		srcSlice := *(*[]*float32)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]float32)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]float32, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv float32
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrFloat64(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*float64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*float64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*float64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv float64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Float64 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]float64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*float64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*float64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Float64 {
		srcSlice := *(*[]*float64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]float64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]float64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv float64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrString(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*string)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*string)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*string, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv string
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.String && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]string)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*string)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*string, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.String {
		srcSlice := *(*[]*string)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]string)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]string, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv string
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrBool(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*bool)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*bool)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*bool, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv bool
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Bool && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]bool)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*bool)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*bool, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Bool {
		srcSlice := *(*[]*bool)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]bool)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]bool, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv bool
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrUintptr(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*uintptr)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uintptr)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uintptr, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uintptr
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Uintptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]uintptr)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*uintptr)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*uintptr, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Uintptr {
		srcSlice := *(*[]*uintptr)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]uintptr)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]uintptr, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv uintptr
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrComplex64(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*complex64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*complex64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*complex64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv complex64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Complex64 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]complex64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*complex64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*complex64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Complex64 {
		srcSlice := *(*[]*complex64)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]complex64)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]complex64, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv complex64
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

func setSliceElemIsBaseTypeOrPtrComplex128(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]*complex128)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*complex128)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*complex128, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv complex128
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Complex128 && dstType.Elem().Kind() == reflect.Ptr {
		srcSlice := *(*[]complex128)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]*complex128)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]*complex128, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			dv := v
			*dstSlice = append(*dstSlice, &dv)
		}
		return nil
	}

	if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Complex128 {
		srcSlice := *(*[]*complex128)(src)
		if len(srcSlice) == 0 {
			return nil
		}
		dstSlice := (*[]complex128)(dst)
		if len(*dstSlice) < len(srcSlice) {
			*dstSlice = make([]complex128, 0, len(srcSlice))
		}
		for _, v := range srcSlice {
			var dv complex128
			if v != nil {
				dv = *v
			}
			*dstSlice = append(*dstSlice, dv)
		}
		return nil
	}
	return nil
}

var copySliceElemIsBaseTypeOrPtrTab = map[reflect.Kind]setReflectFunc{
	reflect.Int: setSliceElemIsBaseTypeOrPtrInt,

	reflect.Int8: setSliceElemIsBaseTypeOrPtrInt8,

	reflect.Int16: setSliceElemIsBaseTypeOrPtrInt16,

	reflect.Int32: setSliceElemIsBaseTypeOrPtrInt32,

	reflect.Int64: setSliceElemIsBaseTypeOrPtrInt64,

	reflect.Uint: setSliceElemIsBaseTypeOrPtrUint,

	reflect.Uint8: setSliceElemIsBaseTypeOrPtrUint8,

	reflect.Uint16: setSliceElemIsBaseTypeOrPtrUint16,

	reflect.Uint32: setSliceElemIsBaseTypeOrPtrUint32,

	reflect.Uint64: setSliceElemIsBaseTypeOrPtrUint64,

	reflect.Float32: setSliceElemIsBaseTypeOrPtrFloat32,

	reflect.Float64: setSliceElemIsBaseTypeOrPtrFloat64,

	reflect.String: setSliceElemIsBaseTypeOrPtrString,

	reflect.Bool: setSliceElemIsBaseTypeOrPtrBool,

	reflect.Uintptr: setSliceElemIsBaseTypeOrPtrUintptr,

	reflect.Complex64: setSliceElemIsBaseTypeOrPtrComplex64,

	reflect.Complex128: setSliceElemIsBaseTypeOrPtrComplex128,
}

func getSetSliceElemIsBaseTypeOrPtrFunc(src reflect.Kind, p bool) setReflectFunc {
	f, ok := copySliceElemIsBaseTypeOrPtrTab[src]
	if p && !ok {
		panic(fmt.Sprintf("not support type:dst %v ", src))
	}
	return f
}
