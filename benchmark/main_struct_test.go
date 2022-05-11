//go:build struct && !wasi
// +build struct,!wasi

package main

import (
	"testing"
)

func init() {
	initEncode()
}

func BenchmarkDecodeSumVec3(b *testing.B) {
	initEncode()
	_Struct = KStruct
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if sum := KBenchmarkDecodeSumVec3(0); sum != KSumExpected {
			b.Error("invalid sum", sum, KSumExpected)
		}
	}
}
