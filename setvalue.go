package deepcopy

import (
	"reflect"
	"unsafe"
)

type (
	setFunc    func(dstAddr, srcAddr unsafe.Pointer)
	setFuncTab map[reflect.Kind]setFunc
)
