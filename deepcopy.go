package deepcopy

import (
	"reflect"
)

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func isArraySlice(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

func deepCopy(dst, src reflect.Value) error {

	if src.IsZero() {
		return nil
	}

	switch src.Kind() {
	case reflect.Slice, reflect.Array:

		if !isArraySlice(dst) {
			return nil
		}

		if dst.Kind() == reflect.Slice && dst.Len() == 0 && src.Len() > 0 {
			// MakeSlice的类型用reflect.SliceOf(src.Index(0).Type()),不用src.Type()的原因
			// 这里src的类型可能是array和slice。我们要的是src元素T的 []T类型,就能不使用src.Type()
			newDst := reflect.MakeSlice(reflect.SliceOf(src.Index(0).Type()), src.Len(), src.Len())
			dst.Set(newDst)
		}

		l := min(dst.Len(), src.Len())
		for i := 0; i < l; i++ {
			if err := deepCopy(dst.Index(i), src.Index(i)); err != nil {
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

			dst.SetMapIndex(k, v)
		}

	case reflect.Func:
	case reflect.Struct:

		typ := src.Type()
		for i, n := 0, src.NumField(); i < n; i++ {
			sf := typ.Field(i)
			if err := deepCopy(dst.FieldByName(sf.Name), src.Field(i)); err != nil {
				return err
			}
		}

	case reflect.Interface:
	case reflect.Ptr:
		if dst.Kind() == reflect.Ptr {
			dst = dst.Elem()
		}
		return deepCopy(dst, src.Elem())
	default:
		dst.Set(src)
	}

	return nil
}

func DeepCopy(dst, src interface{}) error {
	if dst == nil || src == nil {
		return nil
	}

	return deepCopy(reflect.ValueOf(dst), reflect.ValueOf(src))
}
