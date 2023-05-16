package dcopy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SliceStructPtrElemA struct {
	Info     *MessageWhoAndTime
	LikeType LikeType
	Name     string
	Age      int
}

type SliceStructPtrElemB struct {
	Info     *MessageWhoAndTime
	LikeType LikeType
	Name     string
	Age      int
}

type SliceStructPtrElemC struct {
	Info     *MessageWhoAndTime
	LikeType LikeType
	Name     string
	Age      int
}

type SliceStructPtrElemD struct {
	Info     *MessageWhoAndTime
	LikeType LikeType
	Name     string
	Age      int
}

type SliceStructPtrTest_Src struct {
	A []*SliceStructPtrElemA
	B []*SliceStructPtrElemB
}

type SliceStructPtrTest_Dst struct {
	A []SliceStructPtrElemA
	B []SliceStructPtrElemB
}

var local_SliceStructPtrTest_Src = SliceStructPtrTest_Src{
	A: []*SliceStructPtrElemA{
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "A1",
				Nickname: "A1",
				Time:     "A1",
				DID:      1,
				Seq:      1,
			},
			LikeType: LikeType_COLLECT,
			Name:     "A1",
			Age:      1,
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "A1",
				Nickname: "A1",
				Time:     "A1",
				DID:      1,
				Seq:      1,
			},
			LikeType: LikeType_COLLECT,
			Name:     "A2",
			Age:      2,
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "A1",
				Nickname: "A1",
				Time:     "A1",
				DID:      1,
				Seq:      1,
			},
			LikeType: LikeType_COLLECT,
			Name:     "A3",
			Age:      3,
		},
	},
	B: []*SliceStructPtrElemB{
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "A1",
				Nickname: "A1",
				Time:     "A1",
				DID:      1,
				Seq:      1,
			},
			LikeType: LikeType_COLLECT,
			Name:     "B1",
			Age:      1,
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "A1",
				Nickname: "A1",
				Time:     "A1",
				DID:      1,
				Seq:      1,
			},
			LikeType: LikeType_COLLECT,
			Name:     "B2",
			Age:      2,
		},
		{
			Info: &MessageWhoAndTime{
				HeadPic:  "A1",
				Nickname: "A1",
				Time:     "A1",
				DID:      1,
				Seq:      1,
			},
			LikeType: LikeType_COLLECT,
			Name:     "B3",
			Age:      3,
		},
	},
}

var local_SliceStructPtrTest_Dst = SliceStructPtrTest_Dst{
	A: []SliceStructPtrElemA{
		{
			Name: "A1",
			Age:  1,
		},
		{
			Name: "A2",
			Age:  2,
		},
		{
			Name: "A3",
			Age:  3,
		},
	},
	B: []SliceStructPtrElemB{
		{
			Name: "B1",
			Age:  1,
		},
		{
			Name: "B2",
			Age:  2,
		},
		{
			Name: "B3",
			Age:  3,
		},
	},
}

func Test_SliceStructPtr_Src(t *testing.T) {
	err := Preheat(&SliceStructPtrTest_Src{}, &SliceStructPtrTest_Src{})
	assert.NoError(t, err, fmt.Sprintf("Test_SliceStructPtr_Src %+v", err))

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst SliceStructPtrTest_Src
		err := Copy(&dst, &local_SliceStructPtrTest_Src, opts...)
		assert.Equal(t, local_SliceStructPtrTest_Src, dst)
		assert.NoError(t, err)
	}
}

func Test_SliceStructPtr_Dst(t *testing.T) {
	err := Preheat(&SliceStructPtrTest_Dst{}, &SliceStructPtrTest_Dst{})
	assert.NoError(t, err, fmt.Sprintf("Test_SliceStructPtr_Src %+v", err))

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst SliceStructPtrTest_Dst
		err := Copy(&dst, &local_SliceStructPtrTest_Dst, opts...)
		assert.NoError(t, err)
		assert.Equal(t, local_SliceStructPtrTest_Dst, dst)
	}
}

func Test_SliceStructPtr_Dst_Src(t *testing.T) {
	err := Preheat(&SliceStructPtrTest_Dst{}, &SliceStructPtrTest_Src{})
	assert.NoError(t, err, fmt.Sprintf("Test_SliceStructPtr_Src %+v", err))

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst SliceStructPtrTest_Dst
		err := Copy(&dst, &local_SliceStructPtrTest_Src, opts...)
		assert.NoError(t, err)

		for i, v := range local_SliceStructPtrTest_Src.A {
			assert.Equal(t, dst.A[i], *v)
		}

		for i, v := range local_SliceStructPtrTest_Src.B {
			assert.Equal(t, dst.B[i], *v)
		}
	}
}

type SliceStructPtrElemA2 struct {
	Name string
}

type SliceStructPtrTest_Src2 struct {
	A []*SliceStructPtrElemA2
}

type SliceStructPtrTest_Src_Resp struct {
	Data *SliceStructPtrTest_Src2
}

var local_SliceStructPtrTest_Src2 = SliceStructPtrTest_Src2{
	A: []*SliceStructPtrElemA2{
		{
			Name: "A1",
		},
	},
}

func Test_SliceStructPtr_Dst_Src2(t *testing.T) {
	err := Preheat(&SliceStructPtrTest_Src2{}, &SliceStructPtrTest_Src2{})
	assert.NoError(t, err, fmt.Sprintf("Test_SliceStructPtr_Src %+v", err))

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}
		var dst SliceStructPtrTest_Src_Resp
		err := Copy(&dst.Data, &local_SliceStructPtrTest_Src2, opts...)
		assert.NoError(t, err)

		assert.NotNil(t, dst.Data)
		assert.NotEqual(t, len(dst.Data.A), 0)
		for i, v := range local_SliceStructPtrTest_Src2.A {
			assert.Equal(t, *dst.Data.A[i], *v)
		}

	}
}
