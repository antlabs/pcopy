// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
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

	set setFunc
}

// 打印cacheAllFunc
func printCacheAllFunc() {
	rdlock.RLock()
	defer rdlock.RUnlock()

	for k, v := range cacheAllFunc {
		println("key:", k.dst.String(), k.src.String())
		for _, vv := range v.fieldFuncs {
			println("val:", vv.dstType.String(), vv.srcType.String(), vv.dstOffset, vv.srcOffset)
		}
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

	cacheFunc.do(0, dstAddr, srcAddr)
	return true
}

func add(addr unsafe.Pointer, offset int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(addr) + uintptr(offset))
}

func sub(base uintptr, addr uintptr) int {
	return int(base - uintptr(addr))
}

func newAllFieldFunc() *allFieldFunc {
	return &allFieldFunc{fieldFuncs: make([]*offsetAndFunc, 0, 8)}
}

func (af *allFieldFunc) append(fieldFunc *offsetAndFunc) {
	af.fieldFuncs = append(af.fieldFuncs, fieldFunc)
}

func baseType(kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.String,
		reflect.Bool,
		reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}

func (c *allFieldFunc) do(start int, dstAddr, srcAddr unsafe.Pointer) {
	for _, v := range c.fieldFuncs[start:] {
		// for k, v := range c.fieldFuncs[start:] {
		kind := v.srcType.Kind()
		switch {
		// case reflect.Array, reflect.Slice:
		// 	c.do(k+1, add(dstAddr, v.dstOffset), add(srcAddr, v.srcOffset))
		// 	return
		case baseType(kind):
			v.set(add(dstAddr, v.dstOffset), add(srcAddr, v.srcOffset))
		}
	}
}
