package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

			Copy(&d, &s).Do()
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

			Copy(&d, &s).Do()
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
			Copy(&d, &s).Do()
			return testCase{}
		}(),

		func() testCase {

			a1 := []int{1, 2, 3, 4, 5, 6}
			a2 := []string{}
			Copy(&a2, &a1).Do()
			return testCase{}
		}(),
	} {

		assert.Equal(t, tc.need, tc.got)
	}
}
