// Copyright [2020-2023] [guonaihong]
package pcopy

import "testing"

type interfaceBaseSliceSrcTest struct {
	A interface{}
	B interface{}
	C interface{}
	D interface{}
	E interface{}
	F interface{}
	G interface{}
	H interface{}
	I interface{}
	J interface{}
	K interface{}
	L interface{}
	M interface{}
	N interface{}
}

type interfaceSliceDstTest interfaceBaseSliceSrcTest

var localInterfaceSliceSrcTest = interfaceBaseSliceSrcTest{
	A: []int{1, 2, 3, 4, 5},
	B: []int8{1, 2, 3, 4, 5},
	C: []int16{1, 2, 3, 4, 5},
	D: []int32{1, 2, 3, 4, 5},
	E: []int64{1, 2, 3, 4, 5},
	F: []uint{1, 2, 3, 4, 5},
	G: []uint8{1, 2, 3, 4, 5},
	H: []uint16{1, 2, 3, 4, 5},
	I: []uint32{1, 2, 3, 4, 5},
	J: []uint64{1, 2, 3, 4, 5},
	K: []float32{1, 2, 3, 4, 5},
	L: []float64{1, 2, 3, 4, 5},
	M: []string{"1", "2", "3", "4", "5"},
	N: []bool{true, false, true, false, true},
}

func Benchmark_Interface_BaseSlice_Unsafe_Pcopy(b *testing.B) {
	var dst interfaceSliceDstTest
	err := Preheat(&dst, &localInterfaceSliceSrcTest)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst interfaceSliceDstTest
		err := Copy(&dst, &localInterfaceSliceSrcTest, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_Interface_BaseSlice_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst interfaceSliceDstTest
		dst.A = append([]int{}, localInterfaceSliceSrcTest.A.([]int)...)
		dst.B = append([]int8{}, localInterfaceSliceSrcTest.B.([]int8)...)
		dst.C = append([]int16{}, localInterfaceSliceSrcTest.C.([]int16)...)
		dst.D = append([]int32{}, localInterfaceSliceSrcTest.D.([]int32)...)
		dst.E = append([]int64{}, localInterfaceSliceSrcTest.E.([]int64)...)
		dst.F = append([]uint{}, localInterfaceSliceSrcTest.F.([]uint)...)
		dst.G = append([]uint8{}, localInterfaceSliceSrcTest.G.([]uint8)...)
		dst.H = append([]uint16{}, localInterfaceSliceSrcTest.H.([]uint16)...)
		dst.I = append([]uint32{}, localInterfaceSliceSrcTest.I.([]uint32)...)
		dst.J = append([]uint64{}, localInterfaceSliceSrcTest.J.([]uint64)...)
		dst.K = append([]float32{}, localInterfaceSliceSrcTest.K.([]float32)...)
		dst.L = append([]float64{}, localInterfaceSliceSrcTest.L.([]float64)...)
		dst.M = append([]string{}, localInterfaceSliceSrcTest.M.([]string)...)
		dst.N = append([]bool{}, localInterfaceSliceSrcTest.N.([]bool)...)

		a := &dst
		b := a
		_ = b
	}
}

func Benchmark_Interface_BaseSlice_MiniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst interfaceSliceDstTest
		err := miniCopy(&dst, &localInterfaceSliceSrcTest)
		if err != nil {
			b.Fatal(err)
		}
		// miniCopy(&dst, &td)
	}
}

func Benchmark_Interface_BaseSlice_Pcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst interfaceSliceDstTest
		err := Copy(&dst, &localInterfaceSliceSrcTest)
		if err != nil {
			b.Fatal(err)
		}

	}
}
