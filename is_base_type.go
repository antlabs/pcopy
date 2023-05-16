// Copyright [2020-2023] [guonaihong]
package dcopy

import "reflect"

func isBaseType(b reflect.Kind) bool {
	switch b {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.String, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Uintptr:
		return true
	}
	return false
}

func isBaseSliceOrBaseSlicePtr(t reflect.Type, k *reflect.Kind) bool {
	if t.Kind() != reflect.Slice {
		return false
	}

	if isBaseType(t.Elem().Kind()) {
		if k != nil {
			*k = t.Elem().Kind()
		}
		return true
	}

	if t.Elem().Kind() != reflect.Ptr {
		return false
	}

	if isBaseType(t.Elem().Elem().Kind()) {
		if k != nil {
			*k = t.Elem().Elem().Kind()
		}
		return true
	}
	return false
}
