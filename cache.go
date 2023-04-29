// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var (
	cacheAllFunc map[dstSrcType]*allFieldFunc = make(map[dstSrcType]*allFieldFunc)
	rdlock       sync.RWMutex
	OpenCache    bool
)

type dstSrcType struct {
	dst reflect.Type
	src reflect.Type
}

type allFieldFunc struct {
	fieldFuncs []*offsetAndFunc
}

type offsetAndFunc struct {
	dstType   reflect.Type
	srcType   reflect.Type
	dstOffset int
	srcOffset int

	set       setFunc
	baseSlice bool // 是否是基础类型的slice
	// 找到一个复合类型
	headComposite *allFieldFunc
}

// 打印cacheAllFunc
func printCacheAllFunc() {
	rdlock.RLock()
	defer rdlock.RUnlock()

	for k, v := range cacheAllFunc {
		println("map dst.string, src.string:", k.dst.String(), k.src.String())
		for i, vv := range v.fieldFuncs {
			if vv == nil {
				fmt.Printf("nil, type %v, index:%d\n", vv, i)
				continue
			}
			if vv.srcType == nil {
				// continue
			}
			fmt.Printf("(%d)dst val: %v,", i, vv.dstType)
			fmt.Printf("(%d)src val: %v ", i, vv.srcType)
			// vv.dstOffset, vv.srcOffset)
		}
		println()
	}
}

func saveToCache(fieldFunc *allFieldFunc, a dstSrcType) {
	rdlock.Lock()
	defer rdlock.Unlock()

	cacheAllFunc[a] = fieldFunc
}

func getSetFromCacheAndRun(a dstSrcType, dstAddr, srcAddr unsafe.Pointer) (exist bool) {
	rdlock.RLock()
	cacheFunc, ok := cacheAllFunc[a]
	if !ok {
		rdlock.RUnlock()
		return ok
	}
	rdlock.RUnlock()

	cacheFunc.do(dstAddr, srcAddr)
	return true
}

// 基址+offset
func add(addr unsafe.Pointer, offset int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(addr) + uintptr(offset))
}

// 基址-当前字段地址 求 offset
func sub(base uintptr, addr uintptr) int {
	return int(base - uintptr(addr))
}

func newAllFieldFunc() *allFieldFunc {
	return &allFieldFunc{fieldFuncs: make([]*offsetAndFunc, 0, 8)}
}

func (af *allFieldFunc) append(fieldFunc *offsetAndFunc) {
	af.fieldFuncs = append(af.fieldFuncs, fieldFunc)
}

func (c *allFieldFunc) do(dstBaseAddr, srcBaseAddr unsafe.Pointer) {
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
		case isBaseType(kind):
			v.set(add(dstBaseAddr, v.dstOffset), add(srcBaseAddr, v.srcOffset))
		}

	next:
		if v.headComposite != nil {
			fmt.Printf("%p, offset:%d, %d\n", c, v.dstOffset, v.srcOffset)
			v.headComposite.do(dstBaseAddr, srcBaseAddr)
			continue
		}
	}
}
