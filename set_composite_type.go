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

type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func setCompositeInterface(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options) error {
	if dstType.Kind() != reflect.Interface {
		panic("dstType is not interface")
	}
	// 暂不支持带方法的interface
	if srcType.NumMethod() > 0 {
		return nil
	}

	// 暂不支持带方法的interface
	srcUnsafe := (*emptyInterface)(src)
	if srcUnsafe.typ == nil || srcUnsafe.word == nil {
		return nil
	}

	srcInter := (*interface{})(src)
	srcType = reflect.TypeOf(*srcInter)
	srcTypeKind := srcType.Kind()

	if isBaseType(srcType.Kind()) {
		// unsafeSet = getSetBaseFunc(srcTypeKind)
		*(*interface{})(dst) = getNewBaseType(srcTypeKind, *(*interface{})(src))
		return nil
	}

	if srcTypeKind == reflect.Slice {
		sh := (*reflect.SliceHeader)(srcUnsafe.word)
		if sh.Len == 0 {
			return nil
		}
		elemKind := srcType.Elem().Kind()
		if isBaseType(elemKind) {
			*(*interface{})(dst) = getNewBaseSliceType(elemKind, sh)
			return nil
		}
	}

	// TODO 这里也可以细化出map的情况
	// 以后再优化

	return copyInner((*interface{})(dst), (*interface{})(src), opt)
}

func setCompositeMap(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options) (err error) {
	srcMapVal := reflect.NewAt(srcType, src).Elem()
	if srcMapVal.Len() == 0 {
		return nil
	}

	dstMapVal := reflect.NewAt(dstType, dst).Elem()
	if dstMapVal.IsNil() {
		dstMapVal.Set(reflect.MakeMapWithSize(dstType, srcMapVal.Len()))
	}

	// fmt.Printf("type %T", dstMapVal)
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
			keyExits = getFromCacheSetAndRun(key, dst, src, opt)
		}

		if !keyExits {
			if err = copyInner(newVal.Interface(), srcVal.Interface(), opt); err != nil {
				return err
			}
		}

		if valExits {
			src = unsafe.Pointer(newKey.UnsafeAddr())
			dst = unsafe.Pointer(srcKey.UnsafeAddr())
			key := dstSrcType{dst: mapValueType, src: srcVal.Type()}
			valExits = getFromCacheSetAndRun(key, dst, src, opt)
		}

		if !valExits {
			if err = copyInner(newKey.Interface(), srcKey.Interface(), opt); err != nil {
				return err
			}
		}
	}
	return err
}

// func setCompositeSlice(dstType, srcType reflect.Type, dstValType, srcValType reflect.Type, dst, src unsafe.Pointer, opt options) error {
func setCompositeSlice(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options) error {
	// src转成reflect.Value
	srcVal := reflect.NewAt(srcType, src).Elem()
	if srcVal.Len() == 0 {
		return nil
	}

	// dst转成reflect.Value
	dstVal := reflect.NewAt(dstType, dst).Elem()
	if dstVal.Cap() < srcVal.Len() {
		dstVal.Set(reflect.MakeSlice(dstType, srcVal.Len(), srcVal.Len()))
	}

	offset := dstVal.Type().Elem().Size()
	subDstType := dstVal.Type().Elem()
	subSrcType := srcVal.Type().Elem()

	key := dstSrcType{dst: subDstType, src: subSrcType}

	dstSliceAddr := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(dst)).Data)
	srcSliceAddr := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(src)).Data)

	exits := getFromCacheSetAndRun(key, dstSliceAddr, srcSliceAddr, opt)
	if exits {
		for i := 1; i < srcVal.Len(); i++ {
			exits = getFromCacheSetAndRun(key, addOffset(dstSliceAddr, offset, i), addOffset(srcSliceAddr, offset, i), opt)
			if !exits {
				// 这里不可能出现exits为false的情况， 除非进程空间被unsafe.Pointer指针写坏了
				panic(fmt.Sprintf("not support type:subDstType(%v) subSrcType(%v)", key.dst, key.src))
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
