package deepcopy

import (
	"reflect"
	"unsafe"
)

type (
	setFunc           func(dstAddr, srcAddr unsafe.Pointer)
	setFuncTab        map[reflect.Kind]setFunc
	setBaseMapFuncTab map[baseMapKind]setFunc
)

type baseMapKind struct {
	key reflect.Kind
	val reflect.Kind
}

type baseMapTypeTmpl struct {
	TypeName []string
}
