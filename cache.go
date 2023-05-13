// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"reflect"
	"sync"
	"unsafe"
)

type (
	newBaseTypeFunc func(interface{}) interface{}
	setUnsafeFunc   func(dstAddr, srcAddr unsafe.Pointer)
	setReflectFunc  func(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) error
	// setReflectFunc    func(dstType, srcType reflect.Type, dstValType, srcValType reflect.Type, dst, src unsafe.Pointer, opt options) error
	setUnsafeFuncTab  map[reflect.Kind]setUnsafeFunc
	setReflectFuncTab map[reflect.Kind]setReflectFunc
	setBaseMapFuncTab map[baseMapKind]setUnsafeFunc
)

type baseMapKind struct {
	key reflect.Kind
	val reflect.Kind
}

type baseMapTypeTmpl struct {
	TypeName []string
}

var cacheAllFunc sync.Map

type dstSrcType struct {
	dst reflect.Type
	src reflect.Type
}

type allFieldFunc struct {
	fieldFuncs []offsetAndFunc
}

type flag int

const (
	// 空值类型
	emptyTypeSet     flag = iota
	baseTypeSet           // 基础类型
	sliceTypeSet          // slice类型
	baseSliceTypeSet      // 基础类型的slice
	structTypeSet         // 结构体类型
	baseMapTypeSet        // 基础类型的map
	debugTypeSet          // debug类型
)

type offsetAndFunc struct {
	dstType    reflect.Type
	srcType    reflect.Type
	dstOffset  int
	srcOffset  int
	unsafeSet  setUnsafeFunc
	reflectSet setReflectFunc

	// 找到一个复合类型
	nextComposite *allFieldFunc
	isBaseType    bool
	createFlag    flag // 记录offsetAndFunc这个对象生成的触发点, debug时用
}

func (o *offsetAndFunc) resetFlag() {
}

func saveToCache(a dstSrcType, fieldFunc *allFieldFunc) {
	// fmt.Printf("saveToCache: dst = %v, src = %v\n", a.dst, a.src)
	cacheAllFunc.LoadOrStore(a, fieldFunc)
}

func hasSetFromCache(a dstSrcType) (exist bool) {
	_, ok := cacheAllFunc.Load(a)
	return ok
}

func getFromCacheSetAndRun(a dstSrcType, dstAddr, srcAddr unsafe.Pointer, opt options) (exist bool, err error) {
	v, ok := cacheAllFunc.Load(a)
	if !ok {
		return false, nil
	}

	err = v.(*allFieldFunc).do(dstAddr, srcAddr, opt)
	// cacheFunc.do(dstAddr, srcAddr)
	return true, err
}

func newAllFieldFunc() (rv *allFieldFunc) {
	rv = &allFieldFunc{fieldFuncs: make([]offsetAndFunc, 0, 5)}
	return
}

func (af *allFieldFunc) append(fieldFunc offsetAndFunc) {
	af.fieldFuncs = append(af.fieldFuncs, fieldFunc)
}

func (c *allFieldFunc) do(dstBaseAddr, srcBaseAddr unsafe.Pointer, opt options) (err error) {
	for _, v := range c.fieldFuncs {
		var kind reflect.Kind

		if v.unsafeSet == nil && v.reflectSet == nil {
			goto next
		}

		kind = v.srcType.Kind()
		switch {
		case kind == reflect.Ptr:
			v.unsafeSet(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))
		// 处理slice
		case kind == reflect.Slice:
			// 由基础类型组成的slice
			if v.unsafeSet != nil {
				// 基础类型的slice直接一把函数搞定
				v.unsafeSet(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))
				continue
			}

			// 复合类型的slice
			if v.reflectSet != nil {
				if err := v.reflectSet(v.dstType, v.srcType, add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset), opt, &v); err != nil {
					return err
				}
			}
			// 处理map
		case kind == reflect.Map:
			// 基础类型的map直接一把函数搞定
			if v.unsafeSet != nil {
				v.unsafeSet(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))
				continue
			}

			// 复合类型的map
			if v.reflectSet != nil {
				if err := v.reflectSet(v.dstType, v.srcType, add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset), opt, &v); err != nil {
					return err
				}
			}
			// 基础类型
		case isBaseType(kind):
			v.unsafeSet(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))

			// 处理interface
		case kind == reflect.Interface:

			if v.reflectSet != nil {
				//  interface{}里面存放基础类型还是复合类型都由reflectSet函数处理
				if err := v.reflectSet(v.dstType, v.srcType, add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset), opt, &v); err != nil {
					return err
				}
			}

		}

	next:
		if v.nextComposite != nil {
			if err := v.nextComposite.do(dstBaseAddr, srcBaseAddr, opt); err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

func addOffset(addr unsafe.Pointer, offset uintptr, i int) unsafe.Pointer {
	return add(addr, int(offset)*i)
}

// 基址+offset
func add(addr unsafe.Pointer, offset int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(addr) + uintptr(offset))
}

// 基址-当前字段地址 求 offset
func sub(base uintptr, addr uintptr) int {
	return int(base - uintptr(addr))
}
