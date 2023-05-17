// Copyright [2020-2023] [guonaihong]
package pcopy

import "testing"

func Benchmark_SliceWithStruct_Unsafe_dcopy(b *testing.B) {
	var dst test_SliceWithStruct_Src

	err := Preheat(&dst, &test_SliceWithStruct_Src{})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst test_SliceWithStruct_Src
		// err := Copy(&dst, &local_SliceWithStruct_Src)
		err := Copy(&dst, &local_SliceWithStruct_Src, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_SliceWithStruct_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_SliceWithStruct_Dst
		dst.A = make([]test_SliceWithStruct_Dst_Item1, len(local_SliceWithStruct_Src.A))
		copy(dst.A, local_SliceWithStruct_Src.A)
		dst.B = append(dst.B, local_SliceWithStruct_Src.B...)
	}
}

func Benchmark_SliceWithStruct_miniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_SliceWithStruct_Src
		err := miniCopy(&dst, &local_SliceWithStruct_Src)
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_SliceWithStruct_Reflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_SliceWithStruct_Src
		err := Copy(&dst, &local_SliceWithStruct_Src)
		if err != nil {
			b.Fatal(err)
		}

	}
}
