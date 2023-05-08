// Copyright [2020-2023] [guonaihong]
package dcopy

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

			Copy(&d, &src)
			return testCase{got: d, need: src}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// TODO 打开
// func Test_Inteface_DCopy(t *testing.T) {
// 	type interfaceTest struct {
// 		I interface{}
// 		S interface{}
// 	}

// 	for _, tc := range []interfaceTest{
// 		{
// 			I: 5,
// 			S: "hello",
// 		},
// 	} {

// 		var data interfaceTest

// 		err := Preheat(&data, &tc)
// 		assert.NoError(t, err)
// 		data = interfaceTest{}
// 		err = Copy(&data, &tc, WithUsePreheat())
// 		assert.NoError(t, err)
// 		fmt.Printf("#### %v\n", data)
// 		// assert.Equal(t, tc, data)
// 	}
// }

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

			Copy(&dst{}, &src{I: "hello"})
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

			Copy((*int)(nil), (*int)(nil))
			return testCase{}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
