// Copyright [2020-2023] [guonaihong]
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
	// 这是一个复全类型, 最外层是ptr
	copyCompositeTab[reflect.Ptr] = setCompositePtr
}

type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func setCompositePtr(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	// 1.基础类型的指针
	// 2.基础slice的指针
	// 3.基础map的指针
	dstVal := reflect.NewAt(dstType, dst)
	srcVal := reflect.NewAt(srcType, src)
	for {

		if srcVal.Kind() == reflect.Ptr && srcVal.IsNil() {
			return nil
		}

		if dstVal.Kind() == reflect.Ptr && dstVal.IsNil() {
			dstVal.Set(reflect.New(dstVal.Type().Elem()))
		}

		if srcVal.Kind() == reflect.Ptr {
			srcVal = srcVal.Elem()
		}

		if dstVal.Kind() == reflect.Ptr {
			dstVal = dstVal.Elem()
		}

		if srcVal.Type().Kind() != reflect.Ptr && dstVal.Type().Kind() != reflect.Ptr {
			break
		}
	}

	if of.unsafeSet != nil {
		of.unsafeSet(unsafe.Pointer(dstVal.UnsafeAddr()), unsafe.Pointer(srcVal.UnsafeAddr()))
		return nil
	}

	if dstVal.Kind() == reflect.Interface && srcVal.Kind() == reflect.Interface {
		return setCompositeInterface(dstVal.Type(), srcVal.Type(), unsafe.Pointer(dstVal.UnsafeAddr()), unsafe.Pointer(srcVal.UnsafeAddr()), opt, of)
	} else if dstVal.Kind() == reflect.Slice && srcVal.Kind() == reflect.Slice {
		return setCompositeSlice(dstVal.Type(), srcVal.Type(), unsafe.Pointer(dstVal.UnsafeAddr()), unsafe.Pointer(srcVal.UnsafeAddr()), opt, of)
	} else if dstVal.Kind() == reflect.Map && srcVal.Kind() == reflect.Map {
		return setCompositeMap(dstVal.Type(), srcVal.Type(), unsafe.Pointer(dstVal.UnsafeAddr()), unsafe.Pointer(srcVal.UnsafeAddr()), opt, of)
	}

	exits, err := getFromCacheSetAndRun(dstSrcType{dst: dstVal.Type(), src: srcVal.Type()}, unsafe.Pointer(dstVal.UnsafeAddr()), unsafe.Pointer(srcVal.UnsafeAddr()), opt)
	if err != nil || exits {
		return err
	}

	return dcopyInner(dstVal.Interface(), srcVal.Interface(), opt)
}

func setCompositeInterface(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) error {
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

	return dcopyInner((*interface{})(dst), (*interface{})(src), opt)
}

// map 拿不到UnsafeAddr指针，所以只能用reflect.Value来做
func setCompositeMap(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
	srcMap := reflect.NewAt(srcType, src).Elem()
	if srcMap.Len() == 0 {
		return nil
	}

	dstMap := reflect.NewAt(dstType, dst).Elem()
	if dstMap.IsNil() {
		dstMap.Set(reflect.MakeMapWithSize(dstType, srcMap.Len()))
	}

	mapKeyType := dstMap.Type().Key()
	mapValType := dstMap.Type().Elem()

	// 获取src map的值
	iter := srcMap.MapRange()

	isBaseKeyType := isBaseType(mapKeyType.Kind())
	isBaseValType := isBaseType(mapValType.Kind())

	// 对map进行反射得到的值不能取地址，所以这里要关闭preheat和usePreheat
	opt.preheat = false
	opt.usePreheat = false

	for iter.Next() {
		srcMapKey := iter.Key()
		srcMapVal := iter.Value()

		newDstVal := reflect.New(mapValType)
		newDstKey := reflect.New(mapKeyType)

		if isBaseKeyType {
			newKey2 := getNewBaseType(mapKeyType.Kind(), srcMapKey.Interface())
			newDstKey.Elem().Set(reflect.ValueOf(newKey2))
		} else {
			if err = dcopyInner(newDstKey.Interface(), srcMapKey.Interface(), opt); err != nil {
				return err
			}
		}

		if isBaseValType {
			newVal2 := getNewBaseType(mapValType.Kind(), srcMapVal.Interface())
			newDstVal.Elem().Set(reflect.ValueOf(newVal2))
			goto next
		}

		if err = dcopyInner(newDstVal.Interface(), srcMapVal.Interface(), opt); err != nil {
			return err
		}
	next:
		dstMap.SetMapIndex(newDstKey.Elem(), newDstVal.Elem())
	}
	return err
}

func setCompositeSlice(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) error {
	srcValPtr := reflect.NewAt(srcType, src)
	// src转成reflect.Value
	srcVal := srcValPtr.Elem()
	if srcVal.Len() == 0 {
		return nil
	}

	dstValPtr := reflect.NewAt(dstType, dst)
	// dst转成reflect.Value
	dstVal := dstValPtr.Elem()
	if dstVal.Len() < srcVal.Len() {
		dstVal.Set(reflect.MakeSlice(dstType, srcVal.Len(), srcVal.Len()))
	}

	dstOffset := dstVal.Type().Elem().Size()
	srcOffset := srcVal.Type().Elem().Size()
	subDstType := dstVal.Type().Elem()
	subSrcType := srcVal.Type().Elem()

	key := dstSrcType{dst: subDstType, src: subSrcType}

	dstSliceAddr := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(dst)).Data)
	srcSliceAddr := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(src)).Data)

	exits, err := getFromCacheSetAndRun(key, dstSliceAddr, srcSliceAddr, opt)
	if err != nil {
		return err
	}
	if exits {
		for i := 1; i < srcVal.Len(); i++ {
			exits, err = getFromCacheSetAndRun(key, addOffset(dstSliceAddr, dstOffset, i), addOffset(srcSliceAddr, srcOffset, i), opt)
			if err != nil {
				return err
			}
			if !exits {
				// 这里不可能出现exits为false的情况， 除非进程空间被unsafe.Pointer指针写坏了
				panic(fmt.Sprintf("not support type:subDstType(%v) subSrcType(%v)", key.dst, key.src))
			}
		}
		return nil
	}

	if subDstType.Kind() == reflect.Ptr || subSrcType.Kind() == reflect.Ptr {

		for i := 0; i < srcVal.Len(); i++ {
			if err := setCompositePtr(subDstType, subSrcType, addOffset(dstSliceAddr, dstOffset, i), addOffset(srcSliceAddr, srcOffset, i), opt, of); err != nil {
				return err
			}
		}
		return nil
	}

	return dcopyInner(dstValPtr.Interface(), srcValPtr.Interface(), opt)
}

func getSetCompositeFunc(t reflect.Kind) setReflectFunc {
	f, ok := copyCompositeTab[t]
	if !ok {
		panic(fmt.Sprintf("not support type:%T", t))
	}
	return f
}
