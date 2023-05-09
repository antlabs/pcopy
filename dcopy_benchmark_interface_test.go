package dcopy

import (
	"testing"
)

type interfaceSrcTest struct {
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
	// O interface{}
	// P interface{}
}

type interfaceDstTest interfaceSrcTest

var localInterfaceSrcTest = interfaceSrcTest{
	A: uint(1),
	B: uint8(2),
	C: uint16(3),
	D: uint32(4),
	E: uint64(5),
	F: int(6),
	G: int8(7),
	H: int16(8),
	I: int32(9),
	J: int64(10),
	K: float32(11),
	L: float64(12),
	M: "13",
	N: true,
}

func Benchmark_Interface_Unsafe_dcopy(b *testing.B) {
	var dst interfaceDstTest
	err := Preheat(&dst, &localInterfaceSrcTest)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst interfaceDstTest
		err := Copy(&dst, &localInterfaceSrcTest, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_Interface_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst interfaceDstTest
		dst.A = localInterfaceSrcTest.A.(uint)
		dst.B = localInterfaceSrcTest.B.(uint8)
		dst.C = localInterfaceSrcTest.C.(uint16)
		dst.D = localInterfaceSrcTest.D.(uint32)
		dst.E = localInterfaceSrcTest.E.(uint64)
		dst.F = localInterfaceSrcTest.F.(int)
		dst.G = localInterfaceSrcTest.G.(int8)
		dst.H = localInterfaceSrcTest.H.(int16)
		dst.I = localInterfaceSrcTest.I.(int32)
		dst.J = localInterfaceSrcTest.J.(int64)
		dst.K = localInterfaceSrcTest.K.(float32)
		dst.L = localInterfaceSrcTest.L.(float64)
		dst.M = localInterfaceSrcTest.M.(string)
		dst.N = localInterfaceSrcTest.N.(bool)

		a := &dst
		b := a
		_ = b
	}
}

func Benchmark_Interface_MiniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst interfaceDstTest
		err := miniCopy(&dst, &localInterfaceSrcTest)
		if err != nil {
			b.Fatal(err)
		}
		// miniCopy(&dst, &td)
	}
}

func Benchmark_Interface_dcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var dst testData
		var dst interfaceDstTest
		err := Copy(&dst, &localInterfaceSrcTest)
		if err != nil {
			b.Fatal(err)
		}

	}
}
