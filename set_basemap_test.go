// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 生成set_basemap_func.go文件
// func Test_Gen_setBasemap(t *testing.T) {
// 	err := saveBaseMapFuncToFile("./set_basemap_func.go")
// 	assert.NoError(t, err)
// }

type DCopyDst_BaseMap struct {
	A map[string]string
	B map[string]int
	C map[string]float64
	D map[string]uint
	E map[string]uint8
	F map[string]uint16
	G map[string]uint32
	H map[string]uint64
	I map[string]bool
	J map[string]int8
	K map[string]int16
	L map[string]int32
	M map[string]int64
	// N map[string]struct{}
	O map[string]float32
	// P map[string]complex64
	BA int
	BB int8
	BC int16
	BD int32
	BE int64
	BF uint
	BG uint8
	BH uint16
	BI uint32
	BJ uint64
	BK bool
	BL float32
	BM float64
	// BN complex64
	// BO complex128

	CA []int
	CB []int8
	CC []int16
	CD []int32
	CE []int64
	CF []uint
	CG []uint8
	CH []uint16
	CI []uint32
	CJ []uint64
	CK []bool
	CL []float32
	CM []float64
	// CN []complex64
	// CO []complex128
}

type DCopySrc_BaseMap DCopyDst_BaseMap

var testSrc_BaseMap = DCopySrc_BaseMap{
	A: map[string]string{
		"1": "1",
		"2": "2",
	},
	B: map[string]int{
		"1": 1,
		"2": 2,
	},
	C: map[string]float64{
		"1": 1,
		"2": 2,
	},
	D: map[string]uint{
		"1": 1,
		"2": 2,
	},
	E: map[string]uint8{
		"1": 1,
		"2": 2,
	},
	F: map[string]uint16{
		"1": 1,
		"2": 2,
	},
	G: map[string]uint32{
		"1": 1,
		"2": 2,
	},
	H: map[string]uint64{
		"1": 1,
		"2": 2,
	},
	I: map[string]bool{
		"1": true,
		"2": false,
	},
	J: map[string]int8{
		"1": 1,
		"2": 2,
	},
	K: map[string]int16{
		"1": 1,
		"2": 2,
	},
	L: map[string]int32{
		"1": 1,
		"2": 2,
	},
	M: map[string]int64{
		"1": 1,
		"2": 2,
	},
	// N: map[string]struct{}{
	// 	"1": struct{}{},
	// 	"2": struct{}{},
	// },
	O: map[string]float32{
		"1": 1,
		"2": 2,
	},
	// P: map[string]complex64{
	// 	"1": 1,
	// 	"2": 2,
	// },
	BA: 1,
	BB: 1,
	BC: 1,
	BD: 1,
	BE: 1,
	BF: 1,
	BG: 1,
	BH: 1,
	BI: 1,
	BJ: 1,
	BK: true,
	BL: 1,
	BM: 1,
	// BN: 1,
	// BO: 1,

	CA: []int{1, 2},
	CB: []int8{1, 2},
	CC: []int16{1, 2},
	CD: []int32{1, 2},
	CE: []int64{1, 2},
	CF: []uint{1, 2},
	CG: []uint8{1, 2},
	CH: []uint16{1, 2},
	CI: []uint32{1, 2},
	CJ: []uint64{1, 2},
	CK: []bool{true, false},
	CL: []float32{1, 2},
	CM: []float64{1, 2},
	// CN: []complex64{1, 2},
	// CO: []complex128{1, 2},

}

func Test_SetBaseMap(t *testing.T) {
	var dst DCopyDst_BaseMap

	err := Preheat(&dst, &testSrc_BaseMap)
	dst = DCopyDst_BaseMap{}
	assert.NoError(t, err)

	err = Copy(&dst, &testSrc_BaseMap, WithUsePreheat())
	assert.NoError(t, err)

	var dst2 DCopyDst_BaseMap = DCopyDst_BaseMap(testSrc_BaseMap)
	assert.Equal(t, dst, dst2)
}
