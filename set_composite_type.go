package dcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

var copyCompositeTab = map[reflect.Kind]setReflectFunc{}

func init() {
	// 这是一个复合类型，最外层是slice
	copyCompositeTab[reflect.Slice] = setCompositeSlice
	// 这是一个复合类型，最外层是map
	copyCompositeTab[reflect.Map] = setCompositeMap
	// 这是一个复合类型，最外层是interface
	copyCompositeTab[reflect.Interface] = setCompositeInterface
}

// TODO 再优化
// func setCompositeInterface(dstType, srcType reflect.Type, dstValType, srcValType reflect.Type, dst, src unsafe.Pointer, opt options) error {
func setCompositeInterface(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options) error {
	srcVal := reflect.NewAt(srcType, src)
	if srcVal.IsNil() {
		return nil
	}

	dstVal := reflect.NewAt(dstType, dst)
	// exits := getSetFromCacheAndRun(key, dst, src, opt)
	// if exits {
	// }
	return copyInner(dstVal.Interface(), srcVal.Interface(), opt)
}

// func setCompositeMap(dstType, srcType reflect.Type, dstValType, srcValType reflect.Type, dst, src unsafe.Pointer, opt options) error {
func setCompositeMap(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options) (err error) {
	srcMapVal := reflect.NewAt(srcType, src)
	if srcMapVal.Len() == 0 {
		return nil
	}

	dstMapVal := reflect.NewAt(dstType, dst)
	if dstMapVal.IsNil() {
		dstMapVal.Set(reflect.MakeMap(dstType))
	}

	mapKeyType := dstMapVal.Elem().Type().Key()
	mapValueType := dstMapVal.Elem().Type().Elem()
	// key := dstSrcType{dst: mapKeyType, src: mapValueType}

	keyExits := true
	valExits := true
	// 获取src map的值
	for _, srcKey := range srcMapVal.MapKeys() {
		srcVal := srcMapVal.MapIndex(srcKey)

		newVal := reflect.New(mapValueType)
		newKey := reflect.New(mapKeyType)

		dstMapVal.SetMapIndex(newKey, newVal)

		if keyExits {
			src = unsafe.Pointer(newVal.UnsafeAddr())
			dst = unsafe.Pointer(srcVal.UnsafeAddr())
			key := dstSrcType{dst: mapValueType, src: srcVal.Type()}
			keyExits = getSetFromCacheAndRun(key, dst, src, opt)
		} else {
			err = copyInner(newVal.Interface(), srcVal.Interface(), opt)
		}

		if valExits {
			src = unsafe.Pointer(newKey.UnsafeAddr())
			dst = unsafe.Pointer(srcKey.UnsafeAddr())
			key := dstSrcType{dst: mapValueType, src: srcVal.Type()}
			keyExits = getSetFromCacheAndRun(key, dst, src, opt)
		} else {
			err = copyInner(newKey.Interface(), srcKey.Interface(), opt)
		}
	}
	return err
}

// func setCompositeSlice(dstType, srcType reflect.Type, dstValType, srcValType reflect.Type, dst, src unsafe.Pointer, opt options) error {
func setCompositeSlice(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options) error {
	// src转成reflect.Value
	srcVal := reflect.NewAt(srcType, src)
	if srcVal.Len() == 0 {
		return nil
	}

	// dst转成reflect.Value
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

func getSetCompositeFunc(t reflect.Kind) setReflectFunc {
	f, ok := copyCompositeTab[t]
	if !ok {
		panic(fmt.Sprintf("not support type:%T", t))
	}
	return f
}
