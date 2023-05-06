package dcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

var copyCompositeSliceTab = setReflectFuncTab{
	reflect.Struct: setCompositeSliceStruct,
	reflect.Map:    setCompositeSliceMap,
}

func setCompositeSliceMap(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt *options) error {
	srcVal := reflect.NewAt(srcType, src)
	if srcVal.Len() == 0 {
		return nil
	}

	dstVal := reflect.NewAt(dstType, dst)
	if dstVal.Cap() < srcVal.Len() {
	}
	return nil
}

func setCompositeSliceStruct(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt *options) error {
	srcVal := reflect.NewAt(srcType, src)
	if srcVal.Len() == 0 {
		return nil
	}

	dstVal := reflect.NewAt(dstType, dst)
	if dstVal.Cap() < srcVal.Len() {
		dstVal.Set(reflect.MakeSlice(dstType, srcVal.Len(), srcVal.Len()))
	}

	offset := dstVal.Elem().Type().Size()
	subDstType := dstVal.Elem().Type().Elem()
	subSrcType := srcVal.Elem().Type().Elem()
	key := dstSrcType{dst: subDstType, src: subSrcType}

	// 子元素项目如果是已经缓存的类型，直接调用缓存的函数
	exits := getSetFromCacheAndRun(key, dst, src, opt)
	if exits {
		for i := 1; i < srcVal.Len(); i++ {
			exits = getSetFromCacheAndRun(key, addOffset(dst, offset, i), addOffset(src, offset, i), opt)
			if !exits {
				panic(fmt.Sprintf("not support type:%v", key))
			}
		}
		return nil
	}

	return copyInner(dstVal.Interface(), srcVal.Interface(), opt)
}

func getSetCompositeSliceFunc(t reflect.Kind) setUnsafeFunc {
	f, ok := copyBaseSliceTab[t]
	if !ok {
		panic(fmt.Sprintf("not support type:%T", t))
	}
	return f
}
