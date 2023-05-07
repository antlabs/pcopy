// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	need interface{}
	got  interface{}
}

// 最大深度
// func Test_MaxDepth(t *testing.T) {
// 	type depth struct {
// 		First string
// 		Data  struct {
// 			Result string
// 		}
// 		Err struct {
// 			ErrMsg struct {
// 				Message string
// 			}
// 		}
// 	}

// 	src := depth{}
// 	src.First = "first"
// 	src.Data.Result = "test"
// 	src.Err.ErrMsg.Message = "good"

// 	for _, tc := range []testCase{
// 		func() testCase {
// 			d := depth{}
// 			err := Copy(&d, &src, WithMaxDepth(2))
// 			assert.NoError(t, err)
// 			if err != nil {
// 				return testCase{}
// 			}
// 			need := depth{}
// 			need.First = "first"
// 			need.Data.Result = "test"
// 			return testCase{got: d, need: need}
// 		}(),
// 		func() testCase {
// 			d := depth{}
// 			Copy(&d, &src, WithMaxDepth(1))
// 			need := depth{}
// 			need.First = "first"
// 			return testCase{got: d, need: need}
// 		}(),
// 		func() testCase {
// 			d := depth{}
// 			Copy(&d, &src, WithMaxDepth(3))
// 			need := depth{}
// 			need.First = "first"
// 			need.Data.Result = "test"
// 			need.Err.ErrMsg.Message = "good"
// 			return testCase{got: d, need: need}
// 		}(),
// 	} {
// 		assert.Equal(t, tc.need, tc.got)
// 	}
// }

// 测试设置tag的情况
// func Test_TagName(t *testing.T) {
// 	type tagName struct {
// 		First string `copy:"first"`
// 		Data  struct {
// 			Result string
// 		}
// 	}

// 	src := tagName{}
// 	src.First = "first"
// 	src.Data.Result = "test"

// 	for _, tc := range []testCase{
// 		func() testCase {
// 			d := tagName{}
// 			Copy(&d, &src, WithTagName("copy"))
// 			need := tagName{}
// 			need.First = "first"
// 			return testCase{got: d, need: need}
// 		}(),
// 	} {
// 		assert.Equal(t, tc.need, tc.got)
// 	}
// }

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
			Copy(&d, &s)
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
			Copy(&d, &s)
			return testCase{got: d, need: d}
		}(),
		func() testCase {
			Copy(new(int), nil)
			return testCase{got: true, need: true}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 测试循环引用
/*
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
			return Copy(&d, &s)
		}(),
	} {
		assert.Error(t, e)
	}
}
*/
