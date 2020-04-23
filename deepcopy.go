package deepcopy

import (
	"reflect"
)

const (
	noDepthLimited = -1
	noTagLimit     = ""
)

// deepCopy结构体
type deepCopy struct {
	maxDepth int
	tagName  string
	dst      interface{}
	src      interface{}
}

// 设置dst, src数据源
func Copy(dst, src interface{}) *deepCopy {
	if dst == nil || src == nil {
		return nil
	}

	d := deepCopy{maxDepth: noDepthLimited, dst: dst, src: src}
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

// 需要的tag name
func needTagName(curTabName string) bool {
	return len(curTabName) > 0
}

// Do() 开干
func (d *deepCopy) Do() error {
	return d.deepCopy(reflect.ValueOf(d.dst), reflect.ValueOf(d.src), 0)
}

// 取最小值
func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// 是array或slice类型
func isArraySlice(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

func (d *deepCopy) deepCopy(dst, src reflect.Value, depth int) error {

	if src.IsZero() {
		return nil
	}

	if d.maxDepth != noDepthLimited && depth > d.maxDepth {
		return nil
	}

	switch src.Kind() {
	case reflect.Slice, reflect.Array:

		if !isArraySlice(dst) {
			return nil
		}

		if dst.Kind() == reflect.Slice && dst.Len() == 0 && src.Len() > 0 {
			// MakeSlice的类型用reflect.SliceOf(src.Index(0).Type()),而不用src.Type()的原因
			// 这里src的类型可能是array和slice。而需要的是src元素T[0]的 slice类型, 使用src.Type()会拿到T的array和slice类型
			newDst := reflect.MakeSlice(reflect.SliceOf(src.Index(0).Type()), src.Len(), src.Len())
			dst.Set(newDst)
		}

		l := min(dst.Len(), src.Len())
		for i := 0; i < l; i++ {
			if err := d.deepCopy(dst.Index(i), src.Index(i), depth); err != nil {
				return err
			}
		}

	case reflect.Map:
		if dst.Kind() != reflect.Map {
			return nil
		}

		if dst.IsNil() {
			newMap := reflect.MakeMap(src.Type())
			dst.Set(newMap)
		}

		iter := src.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()

			newVal := reflect.New(v.Type()).Elem()
			if err := d.deepCopy(newVal, v, depth); err != nil {
				return err
			}

			dst.SetMapIndex(k, newVal)
		}

	case reflect.Func:
		/*
			结构体成员如果是一个函数变量, dst.IsNil()会返回true
			if dst.IsNil() {
				fmt.Printf("hell world\n")
				newFunc := reflect.New(src.Type())
				dst = newFunc.Elem()
			}
		*/

		dst.Set(src)
	case reflect.Struct:

		typ := src.Type()
		for i, n := 0, src.NumField(); i < n; i++ {
			sf := typ.Field(i)
			if !needTagName(sf.Tag.Get(d.tagName)) {
				continue
			}

			if err := d.deepCopy(dst.FieldByName(sf.Name), src.Field(i), depth+1); err != nil {
				return err
			}
		}

	case reflect.Interface:
		/*

			TODO dst如果是空指针的行为
		*/

		return d.deepCopy(dst.Elem(), src.Elem(), depth)
	case reflect.Ptr:
		if dst.Kind() == reflect.Ptr {
			dst = dst.Elem()
		}

		return d.deepCopy(dst, src.Elem(), depth)
	default:
		dst.Set(src)
	}

	return nil
}
