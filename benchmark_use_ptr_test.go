package deepcopy

import (
	"testing"
)

func Benchmark_Use_CachePtr_Deepcopy(b *testing.B) {
	var dst FastCopyDst
	err := Preheat(&dst, &testSrc)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst FastCopyDst
		err := CopyEx(&dst, &testSrc, WithPreheat())
		if err != nil {
			b.Fatal(err)
		}

	}
}
