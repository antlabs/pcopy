// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"fmt"
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

func Test_Interface_DCopy2(t *testing.T) {
	// local := interfaceSrcTest{
	// 	A: uint(1),
	// 	B: uint8(2),
	// 	C: uint16(3),
	// 	D: uint32(4),
	// 	E: uint64(5),
	// 	F: int(6),
	// 	G: int8(7),
	// 	H: int16(8),
	// 	I: int32(9),
	// 	J: int64(10),
	// 	K: float32(11),
	// 	L: float64(12),
	// 	M: "13",
	// 	N: true,
	// }
	err := Preheat(&interfaceDstTest{}, &interfaceSrcTest{})
	// err := Preheat(&interfaceDstTest{}, &local)
	assert.NoError(t, err)
	for _, tc := range []interfaceSrcTest{
		{
			A: uint(1),
			B: uint8(2),
			C: uint16(3),
			D: uint32(4),
			E: uint64(5),
			F: int(6),
			G: int8(7),
			H: int16(8),
			I: int32(9),
			J: int64(10),
			K: float32(11),
			L: float64(12),
			M: "13",
			N: true,
		},
	} {

		var data interfaceDstTest

		err = Copy(&data, &tc, WithUsePreheat())
		assert.NoError(t, err)
		fmt.Print(tc.A.(uint))
		assert.Equal(t, interfaceDstTest(tc), data)
	}
}

// TODO 打开
func Test_Interface_DCopy(t *testing.T) {
	type interfaceTest struct {
		I interface{}
		S interface{}
	}

	err := Preheat(&interfaceTest{}, &interfaceTest{})
	assert.NoError(t, err)
	for _, tc := range []interfaceTest{
		{
			I: 5,
			S: "hello",
		},
		{
			I: "world",
			S: 5,
		},
	} {

		var data interfaceTest

		err = Copy(&data, &tc, WithUsePreheat())
		assert.NoError(t, err)
		fmt.Printf("%v, %v\n", data, tc)
		assert.Equal(t, tc, data)
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
