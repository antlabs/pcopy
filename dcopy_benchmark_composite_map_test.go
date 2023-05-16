// Copyright [2020-2023] [guonaihong]
package dcopy

import "testing"

func Benchmark_CompositeMap_Unsafe_dcopy(b *testing.B) {
	var dst test_MapWithMap_Dst
	err := Preheat(&dst, &local_MapWithMap_Src)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst test_MapWithMap_Dst
		err := Copy(&dst, &local_MapWithMap_Src, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_CompositeMap_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_MapWithMap_Dst
		dst.A = make(map[string]string)
		for k, v := range local_MapWithMap_Src.A {
			dst.A[k] = v
		}

		dst.B = make(map[string]map[string]string)
		for k, v := range local_MapWithMap_Src.B {
			dst.B[k] = make(map[string]string)
			for k1, v1 := range v {
				dst.B[k][k1] = v1
			}
		}

		dst.C = make(map[string]test_MapWithMap_Item)
		for k, v := range local_MapWithMap_Src.C {
			dst.C[k] = v
		}

	}
}

func Benchmark_CompositeMap_miniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_MapWithMap_Dst
		err := miniCopy(&dst, &local_MapWithMap_Src)
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_CompositeMap_Reflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_MapWithMap_Dst
		err := Copy(&dst, &local_MapWithMap_Src)
		if err != nil {
			b.Fatal(err)
		}

	}
}
