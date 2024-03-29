// Copyright [2020-2023] [guonaihong]
package pcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试结构体内嵌的情况
func Test_Struct_Embed(t *testing.T) {
	type core struct {
		ID   int
		Name string
	}

	type dst struct {
		id   int
		name string
		core
	}

	type src dst

	var d dst
	s := src{core: core{ID: 3, Name: "name"}}
	err := Copy(&d, &s)
	assert.NoError(t, err)
	assert.Equal(t, d, dst(s))
}

// 测试结构体特殊情况
func Test_Struct_Special(t *testing.T) {
	for _, tc := range []testCase{
		// 测试结构体里面有小写成员变量
		func() testCase {
			type dst struct {
				id   int
				name string
			}

			type src dst

			Copy(&dst{}, &src{id: 3})
			return testCase{}
		}(),
		// 测试内嵌结构体的情况
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

func Test_Struct_SpecialCopyEx(t *testing.T) {
	for _, tc := range []testCase{
		// 测试结构体里面有小写成员变量
		func() testCase {
			type dst struct {
				id   int
				name string
			}

			type src dst

			Copy(&dst{}, &src{id: 3})
			return testCase{}
		}(),
		// 测试内嵌结构体的情况
		func() testCase {
			type core struct {
				ID   int
				Name string
			}

			type dst struct {
				id   int
				name string
				core
			}

			type src dst

			var d dst
			s := src{core: core{ID: 3, Name: "name"}}
			err := Copy(&d, &s)
			assert.NoError(t, err)
			return testCase{got: d, need: (dst)(s)}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
