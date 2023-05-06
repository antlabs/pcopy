package dcopy

import (
	"reflect"
	"unsafe"
)

type (
	setUnsafeFunc     func(dstAddr, srcAddr unsafe.Pointer)
	setReflectFunc    func(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt *options) error
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
