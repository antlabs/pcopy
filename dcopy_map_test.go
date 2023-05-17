// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test_MapWithMap_Item struct {
	A string
	B string
}

type test_MapWithMap_Dst struct {
	A map[string]string
	B map[string]map[string]string
	C map[string]test_MapWithMap_Item
}

type test_MapWithMap_Src test_MapWithMap_Dst

var local_MapWithMap_Src = test_MapWithMap_Src{
	A: map[string]string{
		"1": "1",
		"2": "2",
	},
	B: map[string]map[string]string{
		"1": {
			"1": "1",
			"2": "2",
		},
		"2": {
			"1": "1",
			"2": "2",
		},
	},
	C: map[string]test_MapWithMap_Item{
		"1": {
			A: "1",
			B: "2",
		},
		"2": {
			A: "1",
			B: "2",
		},
	},
}

func Test_MapWithMap(t *testing.T) {
	err := Preheat(&test_MapWithMap_Dst{}, &test_MapWithMap_Src{})
	assert.NoError(t, err)

	d := test_MapWithMap_Dst{}
	err = Copy(&d, &local_MapWithMap_Src, WithUsePreheat())
	assert.NoError(t, err)
	// Copy(&d, &local_MapWithMap_Src)
	assert.Equal(t, d, test_MapWithMap_Dst(local_MapWithMap_Src))
}

func Test_MapToMap2(t *testing.T) {
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

	err := Preheat(&dst{}, &src{})
	assert.NoError(t, err)
	Copy(&d, &s, WithUsePreheat())
	assert.Equal(t, d, dst{A: a, B: b})
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

	Copy(&d, &s)
	assert.Equal(t, d, dst{A: a, B: b})
}

func Test_Map_Special(t *testing.T) {
	type mVal struct {
		ID   int
		Name string
	}

	for _, tc := range []testCase{
		// map里面的值是结构体
		func() testCase {
			src := map[string]mVal{
				"1": {ID: 1, Name: "name:1"},
				"2": {ID: 2, Name: "name:2"},
				"3": {ID: 3, Name: "name:3"},
			}

			var dst map[string]mVal
			Copy(&dst, &src)
			return testCase{got: dst, need: src}
		}(),
		// dst的地址不是指针，没有发生panic
		func() testCase {
			src := map[string]mVal{
				"1": {ID: 1, Name: "name:1"},
			}

			var dst map[string]mVal
			Copy(&dst, &src)
			return testCase{}
		}(),
		func() testCase {
			src := map[string]mVal{
				"1": {ID: 1, Name: "name:1"},
			}

			Copy(new(int), &src)
			return testCase{}
		}(),
		// key相同,value不同
		func() testCase {
			dst := map[int]string{
				1: "hello",
			}
			src := map[int]int{
				1: 1,
			}

			Copy(&dst, &src)
			return testCase{}
		}(),
		// key不同，value不同
		func() testCase {
			dst := map[string]int64{
				"hello": 3,
			}
			src := map[int]int64{
				1: 64,
			}

			Copy(&dst, &src)
			return testCase{}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
