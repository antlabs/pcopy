// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

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

	// 如果指向的是同一个地址的函数，注释的方法是不行的
	// assert.Equal(t, d.A, s.A)
	assert.Equal(t, *(*uintptr)(unsafe.Pointer(&d.A)), *(*uintptr)(unsafe.Pointer(&s.A)))
}

// 测试特殊情况
func Test_Func_Special(t *testing.T) {
	type fn struct {
		Add func()
	}

	for _, tc := range []testCase{
		// dst 里面没有Add成员变量
		func() testCase {
			Copy(new(int), fn{Add: func() {}}).Do()
			return testCase{true, true}
		}(),
		func() testCase {
			// dst里面成员变量 与 src不一致
			type dstFn struct {
				Add int
			}

			Copy(&dstFn{}, &fn{Add: func() {}}).Do()
			return testCase{true, true}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
