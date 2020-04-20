package deepcopy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	DeepCopy(&d, &s)
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

	DeepCopy(&d, &s)
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

	DeepCopy(&d, &s)
	assert.Equal(t, d, dst{A: a, B: b})
}
