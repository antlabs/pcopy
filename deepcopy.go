package deepcopy

import (
	"errors"
	"reflect"
	"time"
)

const (
	noDepthLimited = -1
	noTagLimit     = ""
)

/*
type visit struct {
	addr uintptr
	typ  reflect.Type
}
var ErrCircularReference = errors.New("deepcopy.Copy:Circular reference")
*/

// deepCopy结构体
type deepCopy struct {
	dst      interface{}
	src      interface{}
	tagName  string
	maxDepth int
	//visited  map[visit]struct{}

	err error
}

// 设置dst, src数据源
func Copy(dst, src interface{}) *deepCopy {
	if dst == nil || src == nil {
		return &deepCopy{err: errors.New("Unsupported type:nil")}
	}

	d := deepCopy{
		maxDepth: noDepthLimited,
		dst:      dst,
		src:      src,
		//visited:  make(map[visit]struct{}, 8),
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

// 需要的tag name
func haveTagName(curTabName string) bool {
	return len(curTabName) > 0
}

// Do() 开干
func (d *deepCopy) Do() error {
	return d.deepCopy(reflect.ValueOf(d.dst), reflect.ValueOf(d.src), 0)
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

	// dst只能是slice和array类型
	if !isArraySlice(dst) {
		return nil
	}

	l := src.Len()
	if dst.Len() > 0 && dst.Len() < src.Len() {
		l = dst.Len()
	}

	// 被拷贝dst类型是slice, 长度是0, src的长度有值
	if dst.Kind() == reflect.Slice && dst.Len() == 0 && src.Len() > 0 {

		dstElemType := dst.Type().Elem()
		newDst := reflect.MakeSlice(reflect.SliceOf(dstElemType), l, l)
		dst.Set(newDst)
	}

	for i := 0; i < l; i++ {
		if err := d.deepCopy(dst.Index(i), src.Index(i), depth); err != nil {
			return err
		}
	}
	return nil
}

// 拷贝map
func (d *deepCopy) cpyMap(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	if !dst.CanSet() {
		return nil
	}

	// 检查value的类型
	if dst.Type().Elem().Kind() != src.Type().Elem().Kind() {
		return nil
	}

	// 检查key的类型
	if dst.Type().Key().Kind() != src.Type().Key().Kind() {
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

// 检查循环引用
/*
func (d *deepCopy) checkCycle(sField reflect.Value) error {
	if sField.CanAddr() {

		addr := sField.UnsafeAddr()
		v := visit{addr: addr, typ: sField.Type()}

		if _, ok := d.visited[v]; ok {
			return ErrCircularReference
		}

		d.visited[v] = struct{}{}
	}

	return nil
}
*/

// 拷贝结构体
func (d *deepCopy) cpyStruct(dst, src reflect.Value, depth int) error {

	if dst.Kind() != src.Kind() {
		if dst.Kind() == reflect.Ptr {
			// 不是空指针，直接解引用
			if !dst.IsNil() {
				return d.cpyStruct(dst.Elem(), src, depth)
			}

			// 被拷贝结构体是指针类型，值是空，
			// 并且dst可以被设置值
			// 直接malloc一块内存
			if dst.Type().Elem().Kind() == reflect.Struct {
				if dst.CanSet() {
					p := reflect.New(dst.Type().Elem())
					dst.Set(p)
					return d.cpyStruct(dst.Elem(), src, depth)
				}
			}
		}
		return nil
	}

	// time.Time类型
	if dst.CanSet() {

		if _, ok := src.Interface().(time.Time); ok {
			dst.Set(src)
			return nil
		}
	}

	typ := src.Type()
	for i, n := 0, src.NumField(); i < n; i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		// 检查是否注册tag
		if len(d.tagName) > 0 && !haveTagName(sf.Tag.Get(d.tagName)) {
			continue
		}

		// 使用src的字段名在dst里面取出reflect.Value值
		dstValue := dst.FieldByName(sf.Name)

		// dst没有src里面所有的字段，跳过
		if !dstValue.IsValid() {
			continue
		}

		// 检查结构体里面的字段是否有循环引用
		sField := src.Field(i)
		/*
			if err := d.checkCycle(sField); err != nil {
				return err
			}
		*/

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

	src = src.Elem()
	newDst := reflect.New(src.Type()).Elem()

	if err := d.deepCopy(newDst, src, depth); err != nil {
		return err
	}

	dst.Set(newDst)
	return nil
}

// 拷贝指针
func (d *deepCopy) cpyPtr(dst, src reflect.Value, depth int) error {

	if dst.Kind() == reflect.Ptr && dst.IsNil() {
		// dst.CanSet必须放到dst.IsNil判断里面
		// 不然会影响到struct或者map类型的指针
		if !dst.CanSet() {
			return nil
		}
		p := reflect.New(dst.Type().Elem())
		dst.Set(p)
	}

	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}

	if dst.Kind() == reflect.Ptr {
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
	if d.err != nil {
		return d.err
	}

	if src.IsZero() {
		return nil
	}

	// 检查递归深度
	if d.maxDepth != noDepthLimited && depth > d.maxDepth {
		return nil
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

}
