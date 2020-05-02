package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试interface{}
func Test_Inteface(t *testing.T) {
	type interfaceTest struct {
		I interface{}
		S interface{}
	}

	for _, tc := range []testCase{
		func() testCase {
			d := interfaceTest{}
			src := interfaceTest{
				I: 5,
				S: "hello",
			}

			Copy(&d, &src).Do()
			return testCase{got: d, need: src}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 测试interface{}特殊情况
// 只要不崩溃就是对的
func Test_Interface_Special(t *testing.T) {
	for _, tc := range []testCase{
		// src, dst里面同名成员变量类型不一样
		func() testCase {
			type dst struct {
				I int
			}

			type src struct {
				I interface{}
			}

			Copy(&dst{}, &src{I: "hello"}).Do()
			return testCase{}
		}(),

		// src有nil成员变量
		func() testCase {
			type dst struct {
				I int
			}

			type src struct {
				I interface{}
			}

			Copy((*int)(nil), (*int)(nil)).Do()
			return testCase{}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
