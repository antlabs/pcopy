// Copyright [2020-2023] [guonaihong]
package pcopy

import (
	"errors"
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// 不支持的类型
var ErrUnsupportedType = errors.New("Unsupported type")

// 不支持nil类型
var ErrUnsupportedNil = errors.New("Unsupported nil type")

// dst 和 src必须是指针类型
var ErrNotPointer = errors.New("dst and src must be pointer")

// 不能获取指针地址
var ErrNotAddr = errors.New("dst or src type can not get address")

var zeroUintptr unsafe.Pointer

// pcopy结构体
type pcopy struct {
	options
}

func Copy[T any, U any](dst *T, src *U, opts ...Option) error {
	if dst == nil || src == nil {
		return ErrUnsupportedNil
	}
	var opt options
	for _, o := range opts {
		o(&opt)
	}
	var dstI interface{} = dst
	var srcI interface{} = src

	return pcopyInner(dstI, srcI, opt)
}

func pcopyInner(dst, src interface{}, opt options) error {
	if dst == nil || src == nil {
		return ErrUnsupportedNil
	}

	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)
	// 开启预热逻辑
	dstAddr := zeroUintptr
	srcAddr := zeroUintptr

	// 预热逻辑和预热cache逻辑都要走
	var of offsetAndFunc

	var all *allFieldFunc
	if opt.preheat || opt.usePreheat {
		if dstValue.Kind() != reflect.Ptr || srcValue.Kind() != reflect.Ptr {
			return ErrNotPointer
		}

		if !dstValue.Elem().CanAddr() {
			return fmt.Errorf("dst %w", ErrNotAddr)
		}

		if !srcValue.Elem().CanAddr() {
			return fmt.Errorf("src %w", ErrNotAddr)
		}
		dstAddr = unsafe.Pointer(dstValue.Elem().UnsafeAddr())
		srcAddr = unsafe.Pointer(srcValue.Elem().UnsafeAddr())

		dstValue = dstValue.Elem()
		srcValue = srcValue.Elem()

		// 从cache load出类型直接执行
		exist, err := getFromCacheSetAndRun(dstSrcType{dst: dstValue.Type(), src: srcValue.Type()}, dstAddr, srcAddr, opt)
		if exist || err != nil {
			return err
		}

		of.srcType = srcValue.Type()
		of.dstType = dstValue.Type()

		if opt.preheat {
			all = newAllFieldFunc()
			// 如果顶层对象是结构体，肯定会有预热结构
			// 但是是其它类型可能没有预热结构, 所以这里要判断一下
			if dstValue.Type().Kind() != reflect.Struct && srcValue.Type().Kind() != reflect.Struct {
				saveToCache(dstSrcType{
					dst: dstValue.Type(),
					src: srcValue.Type(),
				},
					all)
			}
		}

		// 按道理已经预热过的类型, 不会走到这里
		// 但这里有个特殊情况，比如多级不对称指针。为了支持这种情况，就不报错了
		// if opt.usePreheat {
		// 	return errors.New("usePreheat must be used with preheat")
		// }
	}

	d := pcopy{
		options: opt,
	}
	return d.pcopy(dstValue, srcValue, dstAddr, srcAddr, of, all)
}

// // 需要的tag name
// func haveTagName(curTabName string) bool {
// 	return len(curTabName) > 0
// }

// 判断是array或slice类型
func isArraySlice(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

// 拷贝slice array
func (d *pcopy) cpySliceArray(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, of offsetAndFunc, all *allFieldFunc) error {
	// dst只能是slice和array类型
	if !isArraySlice(dst) {
		return nil
	}

	// 保存类型缓存
	l := src.Len()
	if dst.Len() > 0 && dst.Len() < src.Len() {
		l = dst.Len()
	}

	// 被拷贝dst类型是slice, 长度是0, src的长度有值
	if dst.Kind() == reflect.Slice && dst.Len() == 0 && src.Len() > 0 {
		dstElemType := dst.Type().Elem()
		newDst := reflect.MakeSlice(reflect.SliceOf(dstElemType), l, l)
		dst.Set(newDst)
	}

	if d.preheat {
		of.srcType = src.Type()
		of.dstType = dst.Type()
		// 如果是基础类型的slice, []int []string 这种
		if isBaseType(dst.Type().Elem().Kind()) && isBaseType(src.Type().Elem().Kind()) {
			of.unsafeSet = getSetBaseSliceFunc(dst.Type().Elem().Kind())
			all.append(of)
			of.createFlag = baseSliceTypeSet
			return nil
		}
		// 如果任一类型是基础类型的指针类型, *[]int *[]string 这种，和基础类型的slice的组合
		var k reflect.Kind
		if isBaseSliceOrBaseSlicePtr(dst.Type(), &k) && isBaseSliceOrBaseSlicePtr(src.Type(), &k) {
			of.reflectSet = getSetSliceElemIsBaseTypeOrPtrFunc(k, false)
			of.createFlag = baseMapTypeSet
			all.append(of)
			return nil
		}

		// 处理复合类型的slice
		of.reflectSet = getSetCompositeFunc(dst.Type().Kind())
		all.append(of)

		dstElemType := dst.Type().Elem()
		srcElemType := src.Type().Elem()
		exits := hasSetFromCache(dstSrcType{dst: dstElemType, src: srcElemType})
		// 已经存在穿上类型的转换表，直接返回
		if exits {
			return nil
		}

		// 元素不存在转换表，需要创建
		return pcopyInner(reflect.New(dstElemType).Interface(), reflect.New(srcElemType).Interface(), d.options)
	}

	for i := 0; i < l; i++ {
		if err := d.pcopy(dst.Index(i), src.Index(i), dstBase, srcBase, of, all); err != nil {
			return err
		}
	}
	return nil
}

// 拷贝map
func (d *pcopy) cpyMap(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, of offsetAndFunc, all *allFieldFunc) error {
	if dst.Kind() == reflect.Ptr {
		return d.cpyPtr(dst, src, dstBase, srcBase, of, all)
	}

	if dst.Kind() != src.Kind() {
		return nil
	}

	if !dst.CanSet() {
		return nil
	}

	// 检查value的类型
	if dst.Type().Elem().Kind() != src.Type().Elem().Kind() {
		return nil
	}

	// 检查key的类型
	if dst.Type().Key().Kind() != src.Type().Key().Kind() {
		return nil
	}

	if dst.IsNil() {

		newMap := reflect.MakeMap(src.Type())
		dst.Set(newMap)
	}

	if d.preheat {
		// 如果是基础类型的slice
		of.srcType = src.Type()
		of.dstType = dst.Type()
		if isBaseType(dst.Type().Elem().Kind()) && isBaseType(src.Type().Elem().Kind()) {
			of.unsafeSet = getSetBaseMapFunc(src.Type().Key().Kind(), src.Type().Elem().Kind(), true)
			of.createFlag = baseMapTypeSet
			all.append(of)
			return nil
		}

		of.reflectSet = getSetCompositeFunc(dst.Type().Kind())

		all.append(of)

		dstElemType := dst.Type().Elem()
		srcElemType := src.Type().Elem()
		exits := hasSetFromCache(dstSrcType{dst: dstElemType, src: srcElemType})
		// 已经存在穿上类型的转换表，直接返回
		if exits {
			return nil
		}

		// 元素不存在转换表，需要创建
		return pcopyInner(reflect.New(dstElemType).Interface(), reflect.New(srcElemType).Interface(), d.options)
	}

	iter := src.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		newVal := reflect.New(v.Type()).Elem()
		if err := d.pcopy(newVal, v, zeroUintptr, zeroUintptr, of, all); err != nil {
			return err
		}

		dst.SetMapIndex(k, newVal)
	}
	return nil
}

// 拷贝函数
func (d *pcopy) cpyFunc(dst, src reflect.Value) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	dst.Set(src)
	return nil
}

// 拷贝结构体
func (d *pcopy) cpyStruct(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, of offsetAndFunc, all *allFieldFunc) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	if dst.CanSet() {
		if _, ok := src.Interface().(time.Time); ok {
			if d.preheat {
				of.srcType = src.Type()
				of.dstType = dst.Type()
				of.unsafeSet = setTime
				of.createFlag = baseTypeSet
				all.append(of)
				return nil
			}

			dst.Set(src)
		}
	}

	if d.usePreheat {
		// 从cache load出类型直接执行
		exist, err := getFromCacheSetAndRun(dstSrcType{dst: dst.Type(), src: src.Type()}, dstBase, srcBase, d.options)
		if exist || err != nil {
			return nil
		}
	}

	typ := src.Type()
	for i, n := 0, src.NumField(); i < n; i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		// 使用src的字段名在dst里面取出reflect.Value值
		dstValue := dst.FieldByName(sf.Name)

		// dst没有src里面所有的字段，跳过
		if !dstValue.IsValid() {
			continue
		}

		// 检查结构体里面的字段是否有循环引用
		sField := src.Field(i)

		if d.preheat {
			// 更新这个类型的offset
			of.dstOffset = sub(dstValue.UnsafeAddr(), uintptr(dstBase))
			of.srcOffset = sub(sField.UnsafeAddr(), uintptr(srcBase))
		}

		if err := d.pcopy(dstValue, sField, dstBase, srcBase, of, all); err != nil {
			return err
		}
	}

	return nil
}

// 拷贝interface{}
func (d *pcopy) cpyInterface(dst, src reflect.Value, of offsetAndFunc, all *allFieldFunc) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	src = src.Elem()
	newDst := reflect.New(src.Type()).Elem()

	newDstAddr := zeroUintptr
	newSrcAddr := zeroUintptr
	if d.usePreheat {
		newDstAddr = unsafe.Pointer(newDst.UnsafeAddr())
		newSrcAddr = unsafe.Pointer(src.UnsafeAddr())
	}

	if err := d.pcopy(newDst, src, newDstAddr, newSrcAddr, of, all); err != nil {
		return err
	}

	dst.Set(newDst)
	return nil
}

func (d *pcopy) preheatPtr(dst, src reflect.Value, of offsetAndFunc, all *allFieldFunc) error {
	bkDst := dst
	bkSrc := src
	for {
		if dst.Kind() == reflect.Ptr && dst.IsNil() {
			// dst.CanSet必须放到dst.IsNil判断里面
			// 不然会影响到struct或者map类型的指针
			if !dst.CanSet() {
				return nil
			}
			p := reflect.New(dst.Type().Elem())
			dst.Set(p)
		}
		if src.Kind() == reflect.Ptr && src.IsNil() {
			// dst.CanSet必须放到dst.IsNil判断里面
			// 不然会影响到struct或者map类型的指针
			if !src.CanSet() {
				return nil
			}
			p := reflect.New(src.Type().Elem())
			src.Set(p)
		}

		if src.Kind() == reflect.Ptr {
			src = src.Elem()
		}

		if dst.Kind() == reflect.Ptr {
			dst = dst.Elem()
		}

		if src.Kind() != reflect.Ptr && dst.Kind() != reflect.Ptr {
			break
		}
	}

	if src.Kind() != dst.Kind() {
		return nil
	}

	if isBaseType(src.Kind()) {
		of.unsafeSet = getSetBaseFunc(src.Kind())
	} else if src.Kind() == reflect.Slice && isBaseType(src.Type().Elem().Kind()) {
		of.unsafeSet = getSetBaseSliceFunc(src.Type().Elem().Kind())
	} else if src.Kind() == reflect.Map {
		of.unsafeSet = getSetBaseMapFunc(src.Type().Key().Kind(), src.Type().Elem().Kind(), false)
	}
	of.dstType = bkDst.Type()
	of.srcType = bkSrc.Type()
	of.reflectSet = getSetCompositeFunc(reflect.Ptr)
	all.append(of)

	if src.Kind() == reflect.Struct {
		// 预热下这个结构体
		exits := hasSetFromCache(dstSrcType{dst: dst.Type(), src: src.Type()})
		if exits {
			return nil
		}
		return pcopyInner(dst.Addr().Interface(), src.Addr().Interface(), d.options)
	}
	return nil
}

// 拷贝指针
func (d *pcopy) cpyPtr(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, of offsetAndFunc, all *allFieldFunc) error {
	// 解引用之后的类型如果不一样，直接返回
	if d.preheat {
		return d.preheatPtr(dst, src, of, all)
	}

	if dst.Kind() == reflect.Ptr && dst.IsNil() {
		// dst.CanSet必须放到dst.IsNil判断里面
		// 不然会影响到struct或者map类型的指针
		if !dst.CanSet() {
			return nil
		}
		p := reflect.New(dst.Type().Elem())
		dst.Set(p)
	}

	if src.Kind() == reflect.Ptr {
		src = src.Elem()
		srcBase = unsafe.Pointer(src.UnsafeAddr())
	}

	if dst.Kind() == reflect.Ptr {
		dst = dst.Elem()
		dstBase = unsafe.Pointer(dst.UnsafeAddr())
	}

	return d.pcopy(dst, src, dstBase, srcBase, of, all)
}

// 其他类型
func (d *pcopy) cpyDefault(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, of offsetAndFunc, all *allFieldFunc) error {
	if dst.Kind() != src.Kind() {
		// panic(fmt.Sprintf("%v, %v", dst.Kind(), src.Kind()))
		return nil
	}

	if d.preheat {
		of.srcType = src.Type()
		of.dstType = dst.Type()
		of.unsafeSet = getSetBaseFunc(src.Kind())
		of.createFlag = baseTypeSet
		all.append(of)
		return nil
	}

	switch src.Kind() {
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		dst.SetInt(src.Int())
		return nil
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		dst.SetUint(src.Uint())
		return nil
	case reflect.String:
		dst.SetString(src.String())
		return nil
	case reflect.Bool:
		dst.SetBool(src.Bool())
		return nil
	case reflect.Float32, reflect.Float64:

		dst.SetFloat(src.Float())
		return nil
	}

	// 如果这里是枚举类型(type newType oldType)，哪怕底层的数据类型(oldType)一样，set也会报错, 所以在前面加个前置判断保护下
	dst.Set(src)
	return nil
}

func (d *pcopy) pcopy(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, of offsetAndFunc, all *allFieldFunc) error {
	// 预热的时候一定要绕开这个判断, 不管src有值没值都要继续往下走
	// 寻找和dst匹配的字段
	if !(d.preheat || d.usePreheat) {
		if src.IsZero() {
			return nil
		}
	}

	if dst.Kind() == reflect.Ptr {
		return d.cpyPtr(dst, src, dstBase, srcBase, of, all)
	}

	switch src.Kind() {
	case reflect.Slice, reflect.Array:
		return d.cpySliceArray(dst, src, dstBase, srcBase, of, all)

	case reflect.Map:
		return d.cpyMap(dst, src, dstBase, srcBase, of, all)

	case reflect.Func:
		return d.cpyFunc(dst, src)

	case reflect.Struct:
		// 保存类型缓存
		if d.preheat {
			all = addNextComposite(all, of)
			defer saveToCache(dstSrcType{dst.Type(), src.Type()}, all)
		}
		return d.cpyStruct(dst, src, dstBase, srcBase, of, all)

	case reflect.Interface:
		if d.preheat {
			of.srcType = src.Type()
			of.dstType = dst.Type()
			of.reflectSet = getSetCompositeFunc(src.Kind())
			all.append(of)
			// interface是可变类型。cache加速是把类型关系固化，所以这里只需要知道这个offset是interface就行
			// 后需要的类型是可变的，就没有必要分析现在interface存放的类型
			return nil
		}
		return d.cpyInterface(dst, src, of, all)

	case reflect.Ptr:
		return d.cpyPtr(dst, src, dstBase, srcBase, of, all)

	default:
		return d.cpyDefault(dst, src, dstBase, srcBase, of, all)
	}
}

func addNextComposite(all *allFieldFunc, of offsetAndFunc) (newAll *allFieldFunc) {
	of.nextComposite = newAllFieldFunc()
	of.createFlag = sliceTypeSet
	all.append(of)
	// 重置all
	of.createFlag = debugTypeSet
	newAll = of.nextComposite
	of.nextComposite = nil
	return
}
