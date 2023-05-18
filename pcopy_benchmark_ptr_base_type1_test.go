// Copyright [2020-2023] [guonaihong]
package pcopy

import "testing"

func Benchmark_Ptr_BaseType1_Unsafe_Pcopy(b *testing.B) {
	var dst test_BaseType_ptr_Dst
	err := Preheat(&dst, &local_test_BaseType_ptr_Dst)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst test_BaseType_ptr_Dst
		err := Copy(&dst, &local_test_BaseType_ptr_Dst, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_Ptr_BaseType1_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_BaseType_ptr_Dst
		dst.A = newDef(*local_test_BaseType_ptr_Dst.A)
		dst.B = newDef(*local_test_BaseType_ptr_Dst.B)
		dst.C = newDef(*local_test_BaseType_ptr_Dst.C)
		dst.D = newDef(*local_test_BaseType_ptr_Dst.D)
		dst.E = newDef(*local_test_BaseType_ptr_Dst.E)
		dst.F = newDef(*local_test_BaseType_ptr_Dst.F)
		dst.G = newDef(*local_test_BaseType_ptr_Dst.G)
		dst.H = newDef(*local_test_BaseType_ptr_Dst.H)
		dst.I = newDef(*local_test_BaseType_ptr_Dst.I)
		dst.J = newDef(*local_test_BaseType_ptr_Dst.J)
		dst.K = newDef(*local_test_BaseType_ptr_Dst.K)
		dst.L = newDef(*local_test_BaseType_ptr_Dst.L)
		dst.M = newDef(*local_test_BaseType_ptr_Dst.M)
		dst.N = newDef(*local_test_BaseType_ptr_Dst.N)
		dst.O = newDef(*local_test_BaseType_ptr_Dst.O)
		dst.P = newDef(*local_test_BaseType_ptr_Dst.P)
		dst.Q = newDef(*local_test_BaseType_ptr_Dst.Q)
		// dst.R = newDef(*local_test_BaseType_ptr_Dst.R)
	}
}

func Benchmark_Ptr_BaseType1_miniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_BaseType_ptr_Dst
		err := miniCopy(&dst, &local_test_BaseType_ptr_Dst)
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_Ptr_BaseType1_Reflect_Pcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst test_BaseType_ptr_Dst
		err := Copy(&dst, &local_test_BaseType_ptr_Dst)
		if err != nil {
			b.Fatal(err)
		}

	}
}
