package deepcopy

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	need interface{}
	got  interface{}
}

// 考虑dst为空指针情况

// 测试slice拷贝到array
// TODO更多的情况
func Test_SliceToArray(t *testing.T) {
	type dst struct {
		A [3]int
	}

	type src struct {
		A []int
	}

	var d dst
	s := src{
		A: []int{1, 2, 3, 4, 5},
	}

	Copy(&d, &s).Do()
	assert.Equal(t, d, dst{A: [3]int{1, 2, 3}})
}

// 测试array拷贝到slice
// TODO 更多的情况
func Test_ArrayToSlice(t *testing.T) {
	type dst struct {
		A []int
	}

	type src struct {
		A [3]int
	}

	var d dst
	s := src{
		A: [3]int{1, 2, 3},
	}

	Copy(&d, &s).Do()
	assert.Equal(t, d, dst{A: []int{1, 2, 3}})
}

func Test_MapToMap(t *testing.T) {
	type dst struct {
		A map[int]int
		B map[string]string
	}

	type src struct {
		B map[string]string
		A map[int]int
	}

	var d dst
	b := map[string]string{
		"testA": "testA",
		"testB": "testB",
	}

	a := map[int]int{
		1: 1,
		2: 2,
	}

	s := src{
		B: b,
		A: a,
	}

	Copy(&d, &s).Do()
	assert.Equal(t, d, dst{A: a, B: b})
}

func Test_FuncToFunc(t *testing.T) {
	type dst struct {
		A func()
	}

	type src struct {
		A func()
	}

	d := dst{}
	s := src{
		A: func() { fmt.Printf("hello") },
	}

	Copy(&d, &s).Do()

	//如果指向的是同一个地址的函数，注释的方法是不行的
	//assert.Equal(t, d.A, s.A)
	assert.Equal(t, *(*uintptr)(unsafe.Pointer(&d.A)), *(*uintptr)(unsafe.Pointer(&s.A)))
}

// 最大深度
func Test_MaxDepth(t *testing.T) {
	type depth struct {
		First string
		Data  struct {
			Result string
		}
		Err struct {
			ErrMsg struct {
				Message string
			}
		}
	}

	src := depth{}
	src.First = "first"
	src.Data.Result = "test"
	src.Err.ErrMsg.Message = "good"

	for _, tc := range []testCase{
		func() testCase {
			d := depth{}
			err := Copy(&d, &src).MaxDepth(2).Do()
			assert.NoError(t, err)
			if err != nil {
				return testCase{}
			}
			need := depth{}
			need.First = "first"
			need.Data.Result = "test"
			return testCase{got: d, need: need}
		}(),
		/*
			func() testCase {
				d := depth{}
				Copy(&d, &src).MaxDepth(1).Do()
				need := depth{}
				need.First = "first"
				return testCase{got: d, need: need}
			}(),
				func() testCase {
					d := depth{}
					Copy(&d, &src).MaxDepth(3).Do()
					need := depth{}
					need.First = "first"
					need.Data.Result = "test"
					need.Err.ErrMsg.Message = "good"
					return testCase{got: d, need: need}
				}(),
		*/
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

func Test_TagName(t *testing.T) {
	type tagName struct {
		First string `copy:"first"`
		Data  struct {
			Result string
		}
	}

	src := tagName{}
	src.First = "first"
	src.Data.Result = "test"

	for _, tc := range []testCase{
		func() testCase {
			d := tagName{}
			Copy(&d, &src).RegisterTagName("copy").Do()
			need := tagName{}
			need.First = "first"
			return testCase{got: d, need: need}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

func Test_Type(t *testing.T) {
	type typeTest struct {
		ID   int
		Data struct {
			Result string
		}
	}

	src := typeTest{}
	src.ID = 3
	src.Data.Result = "test"

	for _, tc := range []testCase{
		func() testCase {

			d := typeTest{}
			Copy(&d, &src).OnlyType(reflect.Int).Do()
			need := typeTest{}
			need.ID = 3
			return testCase{got: d, need: need}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 下面的test case 确保不panic
func Test_Special(t *testing.T) {
	for _, tc := range []testCase{
		func() testCase {
			// src有的字段, dst里面没有
			type src struct {
				Sex string
			}

			type dst struct {
				ID string
			}

			d := dst{}
			s := src{Sex: "m"}
			Copy(&d, &s).Do()
			return testCase{got: d, need: d}
		}(),
		func() testCase {
			// 同样的字段不同数据类型，不拷贝
			type src struct {
				Sex string
			}

			type dst struct {
				Sex int
			}

			d := dst{}
			s := src{Sex: "m"}
			Copy(&d, &s).Do()
			return testCase{got: d, need: d}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

func Test_Cycle(t *testing.T) {
	for _, e := range []error{
		func() error {
			type src2 struct {
				P1 *src2
				ID string
			}

			// p1指向自己，构造一个环
			s := src2{}
			s.P1 = &s

			d := src2{}
			return Copy(&d, &s).Do()
		}(),
	} {
		assert.Error(t, e)
	}
}
