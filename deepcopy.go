package deepcopy

import (
	"errors"
	"reflect"
)

const (
	noDepthLimited = -1
	noTagLimit     = ""
)

type visit struct {
	addr uintptr
	typ  reflect.Type
}

type copyFunc func(dst, src reflect.Value, depth int) error

var ErrCircularReference = errors.New("deepcopy.Copy:Circular reference")

// deepCopy结构体
type deepCopy struct {
	dst      interface{}
	src      interface{}
	tagName  string
	maxDepth int
	visited  map[visit]struct{}
	//fieldMapping map[string]string
	tab map[reflect.Kind]copyFunc
}

// 设置dst, src数据源
func Copy(dst, src interface{}) *deepCopy {
	if dst == nil || src == nil {
		return nil
	}

	d := deepCopy{
		maxDepth: noDepthLimited,
		dst:      dst,
		src:      src,
		visited:  make(map[visit]struct{}),
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

// 设置支持的类型
func (d *deepCopy) OnlyType(kind ...reflect.Kind) *deepCopy {
	if d.tab == nil && len(kind) > 0 {
		d.tab = make(map[reflect.Kind]copyFunc, len(kind))
		d.tab[reflect.Struct] = d.cpyStruct
		d.tab[reflect.Ptr] = d.cpyPtr
	}

	for _, v := range kind {
		switch v {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:

			d.tab[v] = d.cpyDefault

		case reflect.Slice, reflect.Array:
			d.tab[v] = d.cpySliceArray
		case reflect.Map:
			d.tab[v] = d.cpyMap
		case reflect.Func:
			d.tab[v] = d.cpyFunc
		case reflect.Struct:
			d.tab[v] = d.cpyStruct
		case reflect.Interface:
			d.tab[v] = d.cpyInterface
		case reflect.Ptr:
			d.tab[v] = d.cpyPtr
		case reflect.String:
			d.tab[v] = d.cpyDefault
		}
	}

	return d
}

// 只设置把src的某个字段拷贝到dst的某个字段，只支持同层次拷贝
/* 暂时不实现，好像不是那么必须
func (d *deepCopy) OnlyField(dstSrcField ...string) *deepCopy {
	return d
}
*/

// 需要的tag name
func haveTagName(curTabName string) bool {
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

// 判断是array或slice类型
func isArraySlice(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

// 拷贝slice array
func (d *deepCopy) cpySliceArray(dst, src reflect.Value, depth int) error {
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
	return nil
}

// 拷贝map
func (d *deepCopy) cpyMap(dst, src reflect.Value, depth int) error {
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
	return nil
}

// 拷贝函数
func (d *deepCopy) cpyFunc(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	/*
		结构体成员如果是一个函数变量, dst.IsNil()会返回true
		if dst.IsNil() {
			fmt.Printf("hell world\n")
			newFunc := reflect.New(src.Type())
			dst = newFunc.Elem()
		}
	*/

	dst.Set(src)
	return nil
}

// 拷贝结构体
func (d *deepCopy) cpyStruct(dst, src reflect.Value, depth int) error {

	typ := src.Type()
	for i, n := 0, src.NumField(); i < n; i++ {
		sf := typ.Field(i)
		if len(d.tagName) > 0 && !haveTagName(sf.Tag.Get(d.tagName)) {
			continue
		}

		dstValue := dst.FieldByName(sf.Name)
		if !dstValue.IsValid() {
			continue
		}

		if dstValue.Kind() == reflect.Ptr && dstValue.IsNil() {
			p := reflect.New(src.Field(i).Type().Elem())
			dstValue.Set(p)
		}

		// 检查循环引用
		sField := src.Field(i)
		if src.Field(i).CanAddr() {

			addr := sField.UnsafeAddr()
			v := visit{addr: addr, typ: src.Field(i).Type()}

			if _, ok := d.visited[v]; ok {
				return ErrCircularReference
			}

			d.visited[v] = struct{}{}
		}

		if err := d.deepCopy(dstValue, sField, depth+1); err != nil {
			return err
		}
	}

	return nil
}

// 拷贝interface{}
func (d *deepCopy) cpyInterface(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}
	/*

		TODO dst如果是空指针的行为
	*/
	return d.deepCopy(dst.Elem(), src.Elem(), depth)
}

// 拷贝指针
func (d *deepCopy) cpyPtr(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	if src.Kind() == reflect.Ptr {
		src = src.Elem()
		dst = dst.Elem()
	}

	return d.deepCopy(dst, src, depth)
}

// 其他类型
func (d *deepCopy) cpyDefault(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}
	dst.Set(src)
	return nil
}

func (d *deepCopy) deepCopy(dst, src reflect.Value, depth int) error {

	if src.IsZero() {
		return nil
	}

	if d.maxDepth != noDepthLimited && depth > d.maxDepth {
		return nil
	}

	if d.tab != nil {
		cpy, ok := d.tab[src.Kind()]
		if !ok {
			return nil
		}
		return cpy(dst, src, depth)
	}

	switch src.Kind() {
	case reflect.Slice, reflect.Array:
		return d.cpySliceArray(dst, src, depth)

	case reflect.Map:
		return d.cpyMap(dst, src, depth)

	case reflect.Func:
		return d.cpyFunc(dst, src, depth)

	case reflect.Struct:
		return d.cpyStruct(dst, src, depth)

	case reflect.Interface:
		return d.cpyInterface(dst, src, depth)

	case reflect.Ptr:
		return d.cpyPtr(dst, src, depth)

	default:
		return d.cpyDefault(dst, src, depth)
	}

	return nil
}
