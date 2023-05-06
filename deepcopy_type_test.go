// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试所有类型(除了指针和inteface{})
func Test_TypeCase(t *testing.T) {
	type allType struct {
		Bool       bool
		Int8       int8
		Int16      int16
		Int32      int32
		Int64      int64
		Uint8      uint8
		Uint16     uint16
		Uint32     uint32
		Uint64     uint64
		Uintptr    uintptr
		Float32    float32
		Float64    float64
		Complex64  complex64
		Complex128 complex128
		ArrayInt   [3]int
		// TODO       展开所有类型array
		SliceInt []int
		C        chan int

		// CB func() error
		S string
	}

	for _, tc := range []testCase{
		func() testCase {
			d := allType{}
			src := allType{
				Bool:       true,
				Int8:       1,
				Int16:      2,
				Int32:      3,
				Int64:      4,
				Uint8:      1,
				Uint16:     2,
				Uint32:     3,
				Uint64:     4,
				Uintptr:    0x123456,
				Float32:    3.1,
				Float64:    3.2,
				Complex64:  1.1,
				Complex128: 2.2,
				ArrayInt:   [3]int{1, 2, 3},
				SliceInt:   []int{4, 5, 6},
				C:          make(chan int),
				// CB:         func() error { return nil },
				S: "test all",
			}

			Copy(&d, &src)
			return testCase{got: d, need: src}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
