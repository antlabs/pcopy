package pcopy

import (
	"reflect"
	"testing"
	"unsafe"
)

func Test_IsInt64(t *testing.T) {
	if !isInt64(reflect.Int64) {
		t.Fatalf("reflect.Int64 is int64")
		i := int64(0)
		if unsafe.Sizeof(i) != 8 {
			t.Fatalf("int64 size is not 8")
		}
	}

	if !isInt64(reflect.Uint64) {
		t.Fatalf("reflect.Uint64 is int64")
		i := uint64(0)
		if unsafe.Sizeof(i) != 8 {
			t.Fatalf("uint64 size is not 8")
		}
	}

	if !isInt64(reflect.Int) {
		t.Fatalf("reflect.Int is int64")
		i := int(0)
		if unsafe.Sizeof(i) != 8 {
			t.Fatalf("int size is not 8")
		}
	}
}

type TestBaseTypeExDst struct {
	N64 uint64
	N32 uint32
	N16 uint16
	N8  uint8
}

type TestBaseTypeExSrc struct {
	N64 int64
	N32 int32
	N16 int16
	N8  int8
}

func Test_Base_Type_Ex(t *testing.T) {
	err := Preheat(&TestBaseTypeExDst{}, &TestBaseTypeExSrc{})
	if err != nil {
		t.Fatalf("preheat error: %v", err)
	}
	d := TestBaseTypeExDst{}
	s := TestBaseTypeExSrc{
		N64: 1,
		N32: 2,
		N16: 3,
		N8:  4,
	}
	need := TestBaseTypeExDst{
		N64: 1,
		N32: 2,
		N16: 3,
		N8:  4,
	}
	err = Copy(&d, &s, WithUsePreheat())
	if err != nil {
		t.Fatalf("copy error rv: %v", err)
	}
	if d != need {
		t.Fatalf("copy error: %v", d)
	}
}
