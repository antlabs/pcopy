package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OnlyBaseType(t *testing.T) {
	err := Preheat(new(int), new(int))
	assert.NoError(t, err)
	var i int
	err = Copy(&i, newDef(3), WithUsePreheat())
	assert.NoError(t, err)
	assert.Equal(t, 3, i)
}
