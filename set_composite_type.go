package deepcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

var copyCompositeSliceTab = setReflectFuncTab{
	reflect.Struct: setCompositeSliceStruct,
	reflect.Map:    setCompositeSliceMap,
}

func setCompositeSliceMap(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt *options) {
	srcVal := reflect.NewAt(srcType, src)
	if srcVal.Len() == 0 {
		return
	}

	dstVal := reflect.NewAt(dstType, dst)
	if dstVal.Cap() < srcVal.Len() {
	}
}

func setCompositeSliceStruct(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt *options) {
	srcVal := reflect.NewAt(srcType, src)
	if srcVal.Len() == 0 {
		return
	}

	dstVal := reflect.NewAt(dstType, dst)
	if dstVal.Cap() < srcVal.Len() {
		dstVal.Set(reflect.MakeSlice(dstType, srcVal.Len(), srcVal.Len()))
	}

	// dstVal.Elem().Type().Size()
	subDstType := dstVal.Elem().Type().Elem()
	subSrcType := srcVal.Elem().Type().Elem()
	key := dstSrcType{dst: subDstType, src: subSrcType}

	exits := getSetFromCacheAndRun(key, dst, src, opt)
	if exits {
		return
	}

	for i := 0; i < srcVal.Len(); i++ {
		dstVal.Index(i).Set(srcVal.Index(i))
	}
}

func getSetCompositeSliceFunc(t reflect.Kind) setUnsafeFunc {
	f, ok := copyBaseSliceTab[t]
	if !ok {
		panic(fmt.Sprintf("not support type:%T", t))
	}
	return f
}
