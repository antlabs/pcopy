// Copyright [2020-2023] [guonaihong]
package dcopy

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

var zeroUintptr = unsafe.Pointer(uintptr(0))

// 优化下面的代码，让性能变得更高
const (
	noDepthLimited = -1
	noTagLimit     = ""
)

// dcopy结构体
type dcopy struct {
	options
}

func Copy(dst, src interface{}, opts ...Option) error {
	if dst == nil || src == nil {
		return ErrUnsupportedNil
	}
	var opt options
	for _, o := range opts {
		o(&opt)
	}
	return copyInner(dst, src, &opt)
}

func copyInner(dst, src interface{}, opt *options) error {
	if dst == nil || src == nil {
		return ErrUnsupportedNil
	}

	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)
	// 开启预热逻辑
	dstAddr := zeroUintptr
	srcAddr := zeroUintptr

	// 预热逻辑和预热cache逻辑都要走
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
		exist := getSetFromCacheAndRun(dstSrcType{dst: dstValue.Type(), src: srcValue.Type()}, dstAddr, srcAddr, opt)
		if exist {
			return nil
		}

		if opt.preheat {
			all = newAllFieldFunc()
		}
	}

	d := dcopy{
		options: *opt,
	}
	if opt.maxDepth == 0 {
		d.maxDepth = noDepthLimited
	}

	return d.dcopy(dstValue, srcValue, dstAddr, srcAddr, 0, offsetAndFunc{}, all)
}

// 需要的tag name
func haveTagName(curTabName string) bool {
	return len(curTabName) > 0
}

// 判断是array或slice类型
func isArraySlice(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

// 拷贝slice array
func (d *dcopy) cpySliceArray(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, depth int, of offsetAndFunc, all *allFieldFunc) error {
	// dst只能是slice和array类型
	if !isArraySlice(dst) {
		return nil
	}

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
		// 如果是基础类型的slice
		if isBaseType(dst.Type().Elem().Kind()) && isBaseType(src.Type().Elem().Kind()) {
			// of.dstOffset = sub(dst.UnsafeAddr(), uintptr(dstBase))
			// of.srcOffset = sub(src.UnsafeAddr(), uintptr(srcBase))
			of.srcType = src.Type()
			of.dstType = dst.Type()
			of.set = getSetBaseSliceFunc(dst.Type().Elem().Kind())
			of.baseSlice = true
			of.createFlag = baseSliceTypeSet
			all.append(of)
			return nil
		}

		// 如果是结构体类型的slice, 也即是复合slice
		of.srcType = src.Type()
		of.dstType = dst.Type()
		// of.set = getSetCompositeSliceFunc(dst.Type().Elem())
		of.baseSlice = false
		all.append(of)
		return nil
	}

	for i := 0; i < l; i++ {
		// 基地址+ i*类型大小
		// TODO
		if err := d.dcopy(dst.Index(i), src.Index(i), dstBase, srcBase, depth, of, all); err != nil {
			return err
		}
	}
	return nil
}

// 拷贝map
func (d *dcopy) cpyMap(dst, src reflect.Value, depth int, of offsetAndFunc, all *allFieldFunc) error {
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
		if isBaseType(dst.Type().Elem().Kind()) && isBaseType(src.Type().Elem().Kind()) {
			// of.dstOffset = sub(dst.UnsafeAddr(), uintptr(dstBase))
			// of.srcOffset = sub(src.UnsafeAddr(), uintptr(srcBase))
			of.srcType = src.Type()
			of.dstType = dst.Type()
			of.set = getSetBaseMapFunc(src.Type().Key().Kind(), src.Type().Elem().Kind())
			of.baseMap = true
			of.createFlag = baseMapTypeSet
			all.append(of)
			return nil
		}
	}

	iter := src.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		newVal := reflect.New(v.Type()).Elem()
		if err := d.dcopy(newVal, v, zeroUintptr, zeroUintptr, depth, of, all); err != nil {
			return err
		}

		dst.SetMapIndex(k, newVal)
	}
	return nil
}

// 拷贝函数
func (d *dcopy) cpyFunc(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	/*
		结构体成员如果是一个函数变量, dst.IsNil()会返回true
		if dst.IsNil() {
			fmt.Printf("hell world\n")
			newFunc := reflect.New(src.Type())
			dst = newFunc.Elem()
		}
	*/

	dst.Set(src)
	return nil
}

// 检查循环引用
/*
func (d *dcopy) checkCycle(sField reflect.Value) error {
	if sField.CanAddr() {

		addr := sField.UnsafeAddr()
		v := visit{addr: addr, typ: sField.Type()}

		if _, ok := d.visited[v]; ok {
			return ErrCircularReference
		}

		d.visited[v] = struct{}{}
	}

	return nil
}
*/

// 拷贝结构体
func (d *dcopy) cpyStruct(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, depth int, of offsetAndFunc, all *allFieldFunc) error {
	if dst.Kind() != src.Kind() {
		if dst.Kind() == reflect.Ptr {
			// 不是空指针，直接解引用
			if !dst.IsNil() {
				return d.cpyStruct(dst.Elem(), src, dstBase, srcBase, depth, of, all)
			}

			// 被拷贝结构体是指针类型，值是空，
			// 并且dst可以被设置值
			// 直接malloc一块内存
			if dst.Type().Elem().Kind() == reflect.Struct {
				if dst.CanSet() {
					p := reflect.New(dst.Type().Elem())
					dst.Set(p)
					return d.cpyStruct(dst.Elem(), src, dstBase, srcBase, depth, of, all)
				}
			}
		}
		return nil
	}

	// time.Time类型
	if dst.CanSet() {
		if _, ok := src.Interface().(time.Time); ok {
			dst.Set(src)
			return nil
		}
	}

	if d.usePreheat {
		// 从cache load出类型直接执行
		exist := getSetFromCacheAndRun(dstSrcType{dst: dst.Type(), src: src.Type()}, dstBase, srcBase, &d.options)
		if exist {
			return nil
		}
	}

	typ := src.Type()
	for i, n := 0, src.NumField(); i < n; i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		// 检查是否注册tag
		if len(d.tagName) > 0 && !haveTagName(sf.Tag.Get(d.tagName)) {
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
		/*
			if err := d.checkCycle(sField); err != nil {
				return err
			}
		*/

		if d.preheat {
			// 更新这个类型的offset
			of.dstOffset = sub(dstValue.UnsafeAddr(), uintptr(dstBase))
			of.srcOffset = sub(sField.UnsafeAddr(), uintptr(srcBase))
		}

		if err := d.dcopy(dstValue, sField, dstBase, srcBase, depth+1, of, all); err != nil {
			return err
		}
	}

	return nil
}

// 拷贝interface{}
func (d *dcopy) cpyInterface(dst, src reflect.Value, depth int, of offsetAndFunc, all *allFieldFunc) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	src = src.Elem()
	newDst := reflect.New(src.Type()).Elem()

	if err := d.dcopy(newDst, src, zeroUintptr, zeroUintptr, depth, of, all); err != nil {
		return err
	}

	dst.Set(newDst)
	return nil
}

// 拷贝指针
func (d *dcopy) cpyPtr(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, depth int, of offsetAndFunc, all *allFieldFunc) error {
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
	}

	if dst.Kind() == reflect.Ptr {
		dst = dst.Elem()
	}

	if d.preheat {
		// dstPtr = unsafe.Pointer(dst.UnsafeAddr())
		// srcPtr = unsafe.Pointer(src.UnsafeAddr())
		// fmt.Printf("#### %p:%p\n", dstPtr, srcPtr)
	}
	return d.dcopy(dst, src, dstBase, srcBase, depth, of, all)
}

// 其他类型
func (d *dcopy) cpyDefault(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, depth int, of offsetAndFunc, all *allFieldFunc) error {
	if dst.Kind() != src.Kind() {
		// panic(fmt.Sprintf("%v, %v", dst.Kind(), src.Kind()))
		return nil
	}

	if d.preheat {
		of.srcType = src.Type()
		of.dstType = dst.Type()
		of.set = getSetFunc(src.Kind())
		of.createFlag = baseTypeSet
		all.append(of)
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

func (d *dcopy) dcopy(dst, src reflect.Value, dstBase, srcBase unsafe.Pointer, depth int, of offsetAndFunc, all *allFieldFunc) error {
	// 预热的时间一定要绕开这个判断, 默认src都有值
	// 寻找和dst匹配的字段
	if !(d.preheat || d.usePreheat) {
		if src.IsZero() {
			return nil
		}
	} else {
		// 预热逻辑，先走白名单， 等稳定了，再走黑名单
		if !isBaseType(src.Kind()) && src.Kind() != reflect.Struct && src.Kind() != reflect.Ptr && src.Kind() != reflect.Slice && src.Kind() != reflect.Map {
			panic(fmt.Sprintf("遇到未知的类型:%v", src.Kind()))
			return nil
		}
	}

	// 检查递归深度
	if d.maxDepth != noDepthLimited && depth > d.maxDepth {
		return nil
	}

	switch src.Kind() {
	case reflect.Slice, reflect.Array:
		// 保存类型缓存
		if d.preheat {
			all = addNextComposite(all, of)
			defer saveToCache(all, dstSrcType{dst.Type(), src.Type()})
		}
		return d.cpySliceArray(dst, src, dstBase, srcBase, depth, of, all)

	case reflect.Map:
		if d.preheat {
			all = addNextComposite(all, of)
			defer saveToCache(all, dstSrcType{dst.Type(), src.Type()})
		}
		return d.cpyMap(dst, src, depth, of, all)

	case reflect.Func:
		return d.cpyFunc(dst, src, depth)

	case reflect.Struct:
		// 保存类型缓存
		// var all *allFieldFunc
		if d.preheat {
			all = addNextComposite(all, of)
			defer saveToCache(all, dstSrcType{dst.Type(), src.Type()})
		}
		return d.cpyStruct(dst, src, dstBase, srcBase, depth, of, all)

	case reflect.Interface:
		return d.cpyInterface(dst, src, depth, of, all)

	case reflect.Ptr:
		return d.cpyPtr(dst, src, dstBase, srcBase, depth, of, all)

	default:
		return d.cpyDefault(dst, src, dstBase, srcBase, depth, of, all)
	}
}

func addNextComposite(all *allFieldFunc, of offsetAndFunc) (newAll *allFieldFunc) {
	of.nextComposite = newAllFieldFunc()
	of.createFlag = sliceTypeSet
	all.append(of)
	of.createFlag = debugTypeSet
	// 重置all
	newAll = of.nextComposite
	of.nextComposite = nil
	return
}
