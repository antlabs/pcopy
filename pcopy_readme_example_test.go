package pcopy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dst struct {
	ID     int
	Result string
}

type src struct {
	ID   int
	Text string
}

func Test_ReadME_Struct_Example(t *testing.T) {
	d, s := dst{}, src{ID: 3}
	err := Preheat(&dst{}, &src{}) // 一对类型只要预热一次
	assert.Nil(t, err)
	err = Copy(&d, &s, WithUsePreheat())
	assert.Nil(t, err)
	assert.Equal(t, d.ID, s.ID)
}

func Test_ReadME_Slice_Example2(t *testing.T) {
	i := []int{1, 2, 3, 4, 5, 6}
	var o []int

	err := Preheat(&o, &i)
	assert.NoError(t, err)
	err = Copy(&o, &i, WithUsePreheat())
	assert.NoError(t, err)

	assert.Equal(t, i, o)
	fmt.Printf("%#v\n", o)
}

func Test_ReadMe_Map_Example3(t *testing.T) {
	i := map[string]int{
		"cat":  100,
		"head": 10,
		"tr":   3,
		"tail": 44,
	}

	var o map[string]int
	err := Preheat(&o, &i)
	assert.NoError(t, err)
	err = Copy(&o, &i, WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, i, o)
	fmt.Printf("%#v\n", o)
}
