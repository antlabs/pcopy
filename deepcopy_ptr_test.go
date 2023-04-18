package deepcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试指针
func Test_Ptr_OK(t *testing.T) {
	type interfaceTest struct {
		Iptr *int
		Fptr *float64
	}

	for _, tc := range []testCase{
		func() testCase {
			d := interfaceTest{}
			src := interfaceTest{
				Iptr: new(int),
				Fptr: new(float64),
			}

			*src.Iptr = 3
			*src.Fptr = 3.3
			Copy(&d, &src).Do()
			return testCase{got: d, need: src}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

func Test_Ptr_OKCopyEx(t *testing.T) {
	type interfaceTest struct {
		Iptr *int
		Fptr *float64
	}

	for _, tc := range []testCase{
		func() testCase {
			d := interfaceTest{}
			src := interfaceTest{
				Iptr: new(int),
				Fptr: new(float64),
			}

			*src.Iptr = 3
			*src.Fptr = 3.3
			err := CopyEx(&d, &src)
			assert.NoError(t, err)
			return testCase{got: d, need: src}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}

// 测试指针特殊情况
// 只要不崩溃就是对的
func Test_Ptr_Special(t *testing.T) {
	for _, tc := range []testCase{
		// dst 是空指针
		func() testCase {
			Copy((*int)(nil), new(int)).Do()
			return testCase{}
		}(),

		// dst, src是不同类型
		func() testCase {
			Copy("hello", new(int)).Do()
			return testCase{}
		}(),
		// dst 是双指针
		func() testCase {
			n := 3
			dst := 0
			dstPtr := &dst
			dstPtrPtr := &dstPtr
			err := Copy(dstPtrPtr, &n).Do()
			assert.NoError(t, err)
			return testCase{}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
