// Copyright [2020-2023] [guonaihong]
package pcopy

import "testing"

func Benchmark_BaseMap_Unsafe_dcopy(b *testing.B) {
	var dst DCopyDst_BaseMap
	err := Preheat(&dst, &testSrc_BaseMap)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst DCopyDst_BaseMap
		err := Copy(&dst, &testSrc_BaseMap, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_BaseMap_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst DCopyDst_BaseMap
		dst.A = make(map[string]string)
		for k, v := range testSrc_BaseMap.A {
			dst.A[k] = v
		}
		dst.B = make(map[string]int)
		for k, v := range testSrc_BaseMap.B {
			dst.B[k] = v
		}
		dst.C = make(map[string]float64)
		for k, v := range testSrc_BaseMap.C {
			dst.C[k] = v
		}
		dst.D = make(map[string]uint)
		for k, v := range testSrc_BaseMap.D {
			dst.D[k] = v
		}
		dst.E = make(map[string]uint8)
		for k, v := range testSrc_BaseMap.E {
			dst.E[k] = v
		}
		dst.F = make(map[string]uint16)
		for k, v := range testSrc_BaseMap.F {
			dst.F[k] = v
		}
		dst.G = make(map[string]uint32)
		for k, v := range testSrc_BaseMap.G {
			dst.G[k] = v
		}
		dst.H = make(map[string]uint64)
		for k, v := range testSrc_BaseMap.H {
			dst.H[k] = v
		}
		dst.I = make(map[string]bool)
		for k, v := range testSrc_BaseMap.I {
			dst.I[k] = v
		}
		dst.J = make(map[string]int8)
		for k, v := range testSrc_BaseMap.J {
			dst.J[k] = v
		}
		dst.K = make(map[string]int16)
		for k, v := range testSrc_BaseMap.K {
			dst.K[k] = v
		}
		dst.L = make(map[string]int32)
		for k, v := range testSrc_BaseMap.L {
			dst.L[k] = v
		}
		dst.M = make(map[string]int64)
		for k, v := range testSrc_BaseMap.M {
			dst.M[k] = v
		}

		dst.O = make(map[string]float32)
		for k, v := range testSrc_BaseMap.O {
			dst.O[k] = v
		}

		dst.BA = testSrc_BaseMap.BA
		dst.BB = testSrc_BaseMap.BB
		dst.BC = testSrc_BaseMap.BC
		dst.BD = testSrc_BaseMap.BD
		dst.BE = testSrc_BaseMap.BE
		dst.BF = testSrc_BaseMap.BF
		dst.BG = testSrc_BaseMap.BG
		dst.BH = testSrc_BaseMap.BH
		dst.BI = testSrc_BaseMap.BI
		dst.BJ = testSrc_BaseMap.BJ
		dst.BK = testSrc_BaseMap.BK
		dst.BL = testSrc_BaseMap.BL
		dst.BM = testSrc_BaseMap.BM

		dst.CA = append(dst.CA, testSrc_BaseMap.CA...)
		dst.CB = append(dst.CB, testSrc_BaseMap.CB...)
		dst.CC = append(dst.CC, testSrc_BaseMap.CC...)
		dst.CD = append(dst.CD, testSrc_BaseMap.CD...)
		dst.CE = append(dst.CE, testSrc_BaseMap.CE...)
		dst.CF = append(dst.CF, testSrc_BaseMap.CF...)
		dst.CG = append(dst.CG, testSrc_BaseMap.CG...)
		dst.CH = append(dst.CH, testSrc_BaseMap.CH...)
		dst.CI = append(dst.CI, testSrc_BaseMap.CI...)
		dst.CJ = append(dst.CJ, testSrc_BaseMap.CJ...)
		dst.CK = append(dst.CK, testSrc_BaseMap.CK...)
		dst.CL = append(dst.CL, testSrc_BaseMap.CL...)
		dst.CM = append(dst.CM, testSrc_BaseMap.CM...)

	}
}

func Benchmark_BaseMap_miniCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst DCopyDst_BaseMap
		err := miniCopy(&dst, &testSrc_BaseMap)
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_BaseMap_Reflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst DCopyDst_BaseMap
		err := Copy(&dst, &testSrc_BaseMap)
		if err != nil {
			b.Fatal(err)
		}
	}
}
