// Copyright [2020-2023] [guonaihong]
package pcopy

import "testing"

func Benchmark_Ptr_BaseSlice_Unsafe_Pcopy(b *testing.B) {
	err := Preheat(&test_BaseSliceType_ptr_Dst{}, &test_BaseSliceType_ptr_Src{})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst test_BaseSliceType_ptr_Dst
		err := Copy(&dst, &local_test_BaseSliceType_ptr_Src, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_Ptr_BaseSlice_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_BaseSliceType_ptr_Dst
		dst.A = newDef(append([]int{}, local_test_BaseSliceType_ptr_Src.A...))
		dst.B = newDef(append([]int8{}, local_test_BaseSliceType_ptr_Src.B...))
		dst.C = newDef(append([]int16{}, local_test_BaseSliceType_ptr_Src.C...))
		dst.D = newDef(append([]int32{}, local_test_BaseSliceType_ptr_Src.D...))
		dst.E = newDef(append([]int64{}, local_test_BaseSliceType_ptr_Src.E...))
		dst.F = newDef(append([]uint{}, local_test_BaseSliceType_ptr_Src.F...))
		dst.G = newDef(append([]uint8{}, local_test_BaseSliceType_ptr_Src.G...))
		dst.H = newDef(append([]uint16{}, local_test_BaseSliceType_ptr_Src.H...))
		dst.I = newDef(append([]uint32{}, local_test_BaseSliceType_ptr_Src.I...))
		dst.J = newDef(append([]uint64{}, local_test_BaseSliceType_ptr_Src.J...))
		dst.K = newDef(append([]uintptr{}, local_test_BaseSliceType_ptr_Src.K...))
		dst.L = newDef(append([]float32{}, local_test_BaseSliceType_ptr_Src.L...))
		dst.M = newDef(append([]float64{}, local_test_BaseSliceType_ptr_Src.M...))
		dst.N = newDef(append([]complex64{}, local_test_BaseSliceType_ptr_Src.N...))
		dst.O = newDef(append([]complex128{}, local_test_BaseSliceType_ptr_Src.O...))
		dst.P = newDef(append([]bool{}, local_test_BaseSliceType_ptr_Src.P...))
		dst.Q = newDef(append([]string{}, local_test_BaseSliceType_ptr_Src.Q...))
	}
}

func Benchmark_Ptr_BaseSlice_miniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_BaseSliceType_ptr_Dst
		err := miniCopy(&dst, &local_test_BaseSliceType_ptr_Src)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Ptr_BaseSlice_Reflect_Pcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_BaseSliceType_ptr_Dst
		err := Copy(&dst, &local_test_BaseSliceType_ptr_Src)
		if err != nil {
			b.Fatal(err)
		}

	}
}
