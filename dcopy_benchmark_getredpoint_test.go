package dcopy

import (
	"testing"
)

func Benchmark_GetRedPoint_Unsafe_dcopy(b *testing.B) {
	err := Preheat(&GetRedPointResp{}, &GetRedPointRespData{})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rv := GetRedPointResp{}

		err := Copy(&rv.Data, &local_GetRedPointRespData, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_GetRedPoint_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst GetRedPointResp
		dst.Data = &GetRedPointResp_GetRedPointRespData{}
		dst.Data.Point = make(map[uint32]uint32)
		for k, v := range local_GetRedPointRespData.Point {
			dst.Data.Point[k] = v
		}
	}
}

func Benchmark_GetRedPoint_MiniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		rv := GetRedPointResp{}
		err := miniCopy(&rv.Data, local_GetRedPointRespData)
		if err != nil {
			b.Fatal(err)
		}
		// miniCopy(&dst, &td)
	}
}

func Benchmark_GetRedPoint_dcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		rv := GetRedPointResp{}
		err := Copy(&rv.Data, local_GetRedPointRespData)
		if err != nil {
			b.Fatal(err)
		}

	}
}
