// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"reflect"
	"sync"
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
	srcKind   reflect.Kind
	dstOffset int
	srcOffset int

	set setFunc
}
