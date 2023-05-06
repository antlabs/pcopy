package deepcopy

import (
	"encoding/json"
	"testing"
)

func Benchmark_BaseType_Unsafe_Deepcopy(b *testing.B) {
	var dst DCopyDst
	err := Preheat(&dst, &testSrc)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst DCopyDst
		err := Copy(&dst, &testSrc, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_BaseType_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst DCopyDst
		dst.Bool = testSrc.Bool
		dst.Int = testSrc.Int
		dst.Int8 = testSrc.Int8
		dst.Int16 = testSrc.Int16
		dst.Int32 = testSrc.Int32
		dst.Int64 = testSrc.Int64
		dst.Uint = testSrc.Uint
		dst.Uint8 = testSrc.Uint8
		dst.Uint16 = testSrc.Uint16
		dst.Uint32 = testSrc.Uint32
		dst.Uint64 = testSrc.Uint64
		dst.Float32 = testSrc.Float32
		dst.Float64 = testSrc.Float64
		// dst.Complex64 = testSrc.Complex64
		dst.String = testSrc.String
		a := &dst
		b := a
		_ = b
	}
}

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

func Benchmark_BaseType_MiniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst DCopyDst
		err := miniCopy(&dst, &testSrc)
		if err != nil {
			b.Fatal(err)
		}
		// miniCopy(&dst, &td)
	}
}

func Benchmark_BaseType_DeepCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst DCopyDst
		err := Copy(&dst, &testSrc)
		if err != nil {
			b.Fatal(err)
		}

	}
}
