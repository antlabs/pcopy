// Copyright [2020-2023] [guonaihong]
package deepcopy

import (
	"encoding/json"
	"testing"
)

type testData struct {
	Int64  int64    `json:"int_64"`
	Int32  int32    `json:"int_32"`
	Int16  int8     `json:"int_16"`
	Int8   int8     `json:"int_8"`
	UInt8  int8     `json:"u_int_8"`
	UInt64 uint64   `json:"u_int_64"`
	UInt32 uint32   `json:"u_int_32"`
	UInt16 uint8    `json:"u_int_16"`
	S      string   `json:"s"`
	Slice  []string `json:"slice"`
	Array  []int    `json:"array"`
}

var td = testData{
	Int64:  64,
	Int32:  32,
	Int16:  16,
	Int8:   8,
	UInt8:  18,
	UInt64: 164,
	UInt32: 132,
	UInt16: 116,
	S:      "test deepcopy",
	Slice:  []string{"123", "456", "789"},
	Array:  []int{0x33, 0x44, 0x55, 0x66, 0x77, 0x88},
}

func miniCopy(dst, src interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dst)
}

func Benchmark_MiniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst testData
		miniCopy(&dst, &td)
	}
}

func Benchmark_DeepCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst testData
		Copy(&dst, &td).Do()
	}
}
