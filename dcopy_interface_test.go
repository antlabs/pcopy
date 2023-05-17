// Copyright [2020-2023] [guonaihong]
package pcopy

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

// 基础类型
func Test_Interface_DCopy2(t *testing.T) {
	err := Preheat(&interfaceDstTest{}, &interfaceBaseSrcTest{})
	// err := Preheat(&interfaceDstTest{}, &local)
	assert.NoError(t, err)
	for _, tc := range []interfaceBaseSrcTest{
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
			// M: "13",
			// N: true,
		},
	} {

		var data interfaceDstTest

		err = Copy(&data, &tc, WithUsePreheat())
		// (&data).printAddress()
		// fmt.Printf("data: %x: data.A: %T, data.A: %d\n", data, data.A, data.A)
		assert.NoError(t, err)
		assert.Equal(t, interfaceDstTest(tc), data)
	}
}

func Test_Interface_DCopy3(t *testing.T) {
	err := Preheat(&interfaceSliceDstTest{}, &interfaceBaseSliceSrcTest{})
	assert.NoError(t, err)
	for _, tc := range []interfaceBaseSliceSrcTest{
		{
			A: []uint{1, 2, 3, 4},
			B: []uint8{2, 3, 4, 5},
			C: []uint16{3, 4, 5, 6},
			D: []uint32{4, 5, 6, 7},
			E: []uint64{5, 6, 7, 8},
			F: []int{6, 7, 8, 9},
			G: []int8{7, 8, 9, 10},
			H: []int16{8, 9, 10, 11},
			I: []int32{9, 10, 11, 12},
			J: []int64{10, 11, 12, 13},
			K: []float32{11, 12, 13, 14},
			L: []float64{12, 13, 14, 15},
			M: []string{"13", "14", "15", "16"},
			N: []bool{true, false, true, false},
		},
	} {

		var data interfaceSliceDstTest

		err = Copy(&data, &tc, WithUsePreheat())
		// (&data).printAddress()
		// fmt.Printf("data: %x: data.A: %T, data.A: %d\n", data, data.A, data.A)
		assert.NoError(t, err)
		assert.Equal(t, interfaceSliceDstTest(tc), data)
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
