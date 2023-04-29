package deepcopy

import "testing"

func Benchmark_Use_CachePtr_BaseSlice_Deepcopy(b *testing.B) {
	var dst FastCopyDst_BaseSlice
	err := Preheat(&dst, &testSrc_BaseSlice)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst FastCopyDst_BaseSlice
		err := CopyEx(&dst, &testSrc_BaseSlice, WithUsePreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_RawCopy_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst FastCopyDst_BaseSlice
		dst.Bool = testSrc_BaseSlice.Bool
		dst.Int = testSrc_BaseSlice.Int
		dst.Int8 = testSrc_BaseSlice.Int8
		dst.Int16 = testSrc_BaseSlice.Int16
		dst.Int32 = testSrc_BaseSlice.Int32
		dst.Int64 = testSrc_BaseSlice.Int64
		dst.Uint = testSrc_BaseSlice.Uint
		dst.Uint8 = testSrc_BaseSlice.Uint8
		dst.Uint16 = testSrc_BaseSlice.Uint16
		dst.Uint32 = testSrc_BaseSlice.Uint32
		dst.Uint64 = testSrc_BaseSlice.Uint64
		dst.Float32 = testSrc_BaseSlice.Float32
		dst.Float64 = testSrc_BaseSlice.Float64
		// dst.Complex64 = testSrc_BaseSlice.Complex64
		// dst.Complex128 = testSrc_BaseSlice.Complex128
		dst.SliceBool = append(dst.SliceBool, testSrc_BaseSlice.SliceBool...)
		dst.SliceInt = append(dst.SliceInt, testSrc_BaseSlice.SliceInt...)
		dst.SliceInt8 = append(dst.SliceInt8, testSrc_BaseSlice.SliceInt8...)
		dst.SliceInt16 = append(dst.SliceInt16, testSrc_BaseSlice.SliceInt16...)
		dst.SliceInt32 = append(dst.SliceInt32, testSrc_BaseSlice.SliceInt32...)
		dst.SliceInt64 = append(dst.SliceInt64, testSrc_BaseSlice.SliceInt64...)
		dst.SliceUint = append(dst.SliceUint, testSrc_BaseSlice.SliceUint...)
		dst.SliceUint8 = append(dst.SliceUint8, testSrc_BaseSlice.SliceUint8...)
		dst.SliceUint16 = append(dst.SliceUint16, testSrc_BaseSlice.SliceUint16...)
		dst.SliceUint32 = append(dst.SliceUint32, testSrc_BaseSlice.SliceUint32...)
		dst.SliceUint64 = append(dst.SliceUint64, testSrc_BaseSlice.SliceUint64...)
		dst.SliceFloat32 = append(dst.SliceFloat32, testSrc_BaseSlice.SliceFloat32...)
		dst.SliceFloat64 = append(dst.SliceFloat64, testSrc_BaseSlice.SliceFloat64...)
		// dst.SliceComplex64 = append(dst.SliceComplex64, testSrc_BaseSlice.SliceComplex64...)
		// dst.SliceComplex128 = append(dst.SliceComplex128, testSrc_BaseSlice.SliceComplex128...)
		dst.String = testSrc_BaseSlice.String
		dst.SliceString = append(dst.SliceString, testSrc_BaseSlice.SliceString...)
	}
}

func Benchmark_miniCopy_Deepcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst FastCopyDst_BaseSlice
		err := miniCopy(&dst, &testSrc_BaseSlice)
		if err != nil {
			b.Fatal(err)
		}

	}
}

func Benchmark_Reflect_Deepcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst FastCopyDst_BaseSlice
		err := CopyEx(&dst, &testSrc_BaseSlice)
		if err != nil {
			b.Fatal(err)
		}

	}
}
