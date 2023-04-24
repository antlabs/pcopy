// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"errors"
	"reflect"
)

/*
type visit struct {
	addr uintptr
	typ  reflect.Type
}
var ErrCircularReference = errors.New("deepcopy.Copy:Circular reference")
*/

// 设置dst, src数据源
func Copy(dst, src interface{}) *deepCopy {
	if dst == nil || src == nil {
		return &deepCopy{err: errors.New("Unsupported type:nil")}
	}

	d := deepCopy{
		maxDepth: noDepthLimited,
		dst:      reflect.ValueOf(dst),
		src:      reflect.ValueOf(src),
		// visited:  make(map[visit]struct{}, 8),
	}

	return &d
}

// 设置最多递归的层次
func (d *deepCopy) MaxDepth(maxDepth int) *deepCopy {
	d.maxDepth = maxDepth
	return d
}

// 设置tag name，结构体的tag等于RegisterTagName注册的tag，才会copy值
func (d *deepCopy) RegisterTagName(tagName string) *deepCopy {
	d.tagName = tagName
	return d
}

// 只设置把src的某个字段拷贝到dst的某个字段，只支持同层次拷贝
/* 暂时不实现，好像不是那么必须
func (d *deepCopy) OnlyField(dstSrcField ...string) *deepCopy {
	return d
}
*/
