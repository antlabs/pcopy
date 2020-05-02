package deepcopy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Func(t *testing.T) {

	type fn struct {
		Add func()
	}

	for _, tc := range []testCase{
		// dst 里面没有Add成员变量
		func() testCase {

			Copy(new(int), fn{Add: func() {}}).Do()
			return testCase{true, true}
		}(),
		func() testCase {
			// dst里面成员变量 与 src不一致
			type dstFn struct {
				Add int
			}

			Copy(&dstFn{}, &fn{Add: func() {}}).Do()
			return testCase{true, true}
		}(),
	} {
		assert.Equal(t, tc.need, tc.got)
	}
}
