package pcopy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LetterInbox struct {
	ID      int32
	FromRid int64
	ToRid   int64
	Msg     string
}

var resp = []LetterInbox{
	{
		ID:      1,
		FromRid: 3,
		ToRid:   4,
		Msg:     "ww,",
	},
}

type ReadLetterResponseItem struct {
	Msg string
}

type ReadLetterResponse struct {
	List []*ReadLetterResponseItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func Test_ReadLetter(t *testing.T) {
	var rv ReadLetterResponse
	err := Preheat(&rv.List, &resp)
	assert.NoError(t, err)
	err = Copy(&rv.List, &resp, WithUsePreheat())
	assert.NoError(t, err)

	fmt.Printf("%#v\n", rv.List)
	fmt.Printf("%#v\n", rv.List[0])
	// fmt.Printf("%#v\n", rv.List[1])
}

func Test_ReadLetter2(t *testing.T) {
	var rv ReadLetterResponse
	err := Preheat(&rv.List, &LetterInbox{})
	assert.NoError(t, err)

	resp2 := []LetterInbox{
		{
			ID:      1,
			FromRid: 3,
			ToRid:   4,
			Msg:     "hello",
		},
	}
	err = Copy(&rv.List, &resp2, WithUsePreheat())
	assert.NoError(t, err)

	assert.Equal(t, rv.List[0].Msg, "hello")
}

func Test_ReadLetter3(t *testing.T) {
	var rv ReadLetterResponse

	resp2 := []LetterInbox{
		{
			ID:      1,
			FromRid: 3,
			ToRid:   4,
			Msg:     "hello",
		},
	}
	err := Copy(&rv.List, &resp2)
	assert.NoError(t, err)

	assert.Equal(t, rv.List[0].Msg, "hello")
}
