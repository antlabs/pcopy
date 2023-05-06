// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"reflect"
	"sync"
	"unsafe"
)

// cacheAllFunc map[dstSrcType]*allFieldFunc = make(map[dstSrcType]*allFieldFunc)
var cacheAllFunc sync.Map // rdlock       sync.RWMutex

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
	dstType   reflect.Type
	srcType   reflect.Type
	dstOffset int
	srcOffset int
	set       setUnsafeFunc

	// 找到一个复合类型
	nextComposite *allFieldFunc
	baseSlice     bool // 是否是基础类型的slice
	baseMap       bool // 是否是基础类型的map
	createFlag    flag // 记录offsetAndFunc这个对象生成的触发点
}

func saveToCache(fieldFunc *allFieldFunc, a dstSrcType) {
	cacheAllFunc.LoadOrStore(a, fieldFunc)
}

func getSetFromCacheAndRun(a dstSrcType, dstAddr, srcAddr unsafe.Pointer, opt *options) (exist bool) {
	v, ok := cacheAllFunc.Load(a)
	if !ok {
		return false
	}

	v.(*allFieldFunc).do(dstAddr, srcAddr, opt)
	// cacheFunc.do(dstAddr, srcAddr)
	return true
}

func newAllFieldFunc() (rv *allFieldFunc) {
	rv = &allFieldFunc{fieldFuncs: make([]offsetAndFunc, 0, 5)}
	return
}

func (af *allFieldFunc) append(fieldFunc offsetAndFunc) {
	af.fieldFuncs = append(af.fieldFuncs, fieldFunc)
}

func (c *allFieldFunc) do(dstBaseAddr, srcBaseAddr unsafe.Pointer, opt *options) {
	for _, v := range c.fieldFuncs {
		var kind reflect.Kind

		if v.set == nil {
			goto next
		}

		kind = v.srcType.Kind()
		switch {

		case kind == reflect.Slice:
			if v.baseSlice {
				// 基础类型的slice直接一把函数搞定
				v.set(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))
				continue
			}
		case kind == reflect.Map:
			if v.baseMap {
				// 基础类型的map直接一把函数搞定
				v.set(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))
				continue
			}

		case isBaseType(kind):
			v.set(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))
		}

	next:
		if v.nextComposite != nil {
			v.nextComposite.do(dstBaseAddr, srcBaseAddr, opt)
			continue
		}
	}
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
