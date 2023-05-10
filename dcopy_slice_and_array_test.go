// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test_SliceWithStruct_Dst_Item1 struct {
	A int
	B int8
	C int16
	D int32
	E int64
	F uint
	G uint8
	H uint16
	I uint32
	J uint64
	K float32
	L float64
	M string
	N bool
}

type test_SliceWithStruct_Dst struct {
	A []test_SliceWithStruct_Dst_Item1
	B []int
}

type test_SliceWithStruct_Src test_SliceWithStruct_Dst

func Test_SliceWithStruct(t *testing.T) {
	for _, tc := range []test_SliceWithStruct_Src{
		{
			A: []test_SliceWithStruct_Dst_Item1{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11.1, 12.2, "13", true},
				{14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24.2, 25.3, "26", false},
			},
			B: []int{1, 2, 3, 4, 5},
		},
	} {
		var d test_SliceWithStruct_Dst
		Preheat(&d, &tc)
		d = test_SliceWithStruct_Dst{}
		Copy(&d, &tc)
		assert.Equal(t, tc, test_SliceWithStruct_Src(d))
	}
}

// 测试slice拷贝到array
func Test_SliceToArray(t *testing.T) {
	for _, tc := range []testCase{
		func() testCase {
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

			var need dst
			need.A = [3]int{1, 2, 3}

			Copy(&d, &s)
			return testCase{got: d, need: need}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

func Test_SliceToArrayEx(t *testing.T) {
	for _, tc := range []testCase{
		func() testCase {
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

			var need dst
			need.A = [3]int{1, 2, 3}

			err := Copy(&d, &s)
			assert.NoError(t, err)
			return testCase{got: d, need: need}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 测试array拷贝到slice
func Test_ArrayToSlice(t *testing.T) {
	for _, tc := range []testCase{
		func() testCase {
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

			var need dst
			need.A = s.A[:]

			Copy(&d, &s)
			return testCase{got: d, need: need}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 测试特殊情况
func Test_ArraySlice_Special(t *testing.T) {
	// 不崩溃就是对的
	for _, tc := range []testCase{
		func() testCase {
			type dst struct {
				A int
			}

			type src struct {
				A [3]int
			}

			var d dst
			s := src{
				A: [3]int{1, 2, 3},
			}
			Copy(&d, &s)
			return testCase{}
		}(),

		func() testCase {
			a1 := []int{1, 2, 3, 4, 5, 6}
			a2 := []string{}
			Copy(&a2, &a1)
			return testCase{}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 测试
func Test_ArraySlice_SpecialCopyEx(t *testing.T) {
	// 不崩溃就是对的
	for _, tc := range []testCase{
		func() testCase {
			type dst struct {
				A int
			}

			type src struct {
				A [3]int
			}

			var d dst
			s := src{
				A: [3]int{1, 2, 3},
			}
			Copy(&d, &s)
			return testCase{}
		}(),

		func() testCase {
			a1 := []int{1, 2, 3, 4, 5, 6}
			a2 := []string{}
			Copy(&a2, &a1)
			return testCase{}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
