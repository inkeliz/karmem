//go:build (wazero || wasmer) && !nobench
// +build wazero wasmer
// +build !nobench

package main

import (
	"testing"
	"unsafe"
)

func init() {
	initEncode()
}

func BenchmarkDecodeSumVec3(b *testing.B) {
	encoded := encode()
	m := initWasm(b, "KBenchmarkDecodeSumVec3", "KBenchmarkDecodeObjectAPI")
	defer m.Close()

	if !m.Write(encoded) {
		b.Fatal("invalid memory ptr")
	}

	if _, err := m.Run("KBenchmarkDecodeObjectAPI", uint64(len(encoded))); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	l := uint64(len(encoded))
	for i := 0; i < b.N; i++ {
		sum, err := m.Run("KBenchmarkDecodeSumVec3", l)
		if err != nil {
			b.Fatal(err)
		}
		if *(*float32)(unsafe.Pointer(&sum[0])) != KSumExpected {
			b.Fatal("invalid sum", KSumExpected, *(*float32)(unsafe.Pointer(&sum[0])))
		}
	}
}

func BenchmarkDecodeObjectAPI(b *testing.B) {
	encoded := encode()
	m := initWasm(b, "KBenchmarkDecodeObjectAPI")
	defer m.Close()

	b.ResetTimer()
	b.ReportAllocs()
	l := uint64(len(encoded))
	for i := 0; i < b.N; i++ {
		if !m.Write(encoded) {
			b.Fatal("invalid memory ptr")
		}
		_, err := m.Run("KBenchmarkDecodeObjectAPI", l)
		if err != nil {
			b.Fatal("x", err)
		}
	}
}

func BenchmarkEncodeObjectAPI(b *testing.B) {
	m := initWasm(b, "KBenchmarkEncodeObjectAPI", "KBenchmarkDecodeObjectAPI")
	defer m.Close()
	encoded := encode()
	m.Write(encoded)
	m.Run("KBenchmarkDecodeObjectAPI", uint64(len(encoded)))

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Run("KBenchmarkEncodeObjectAPI")
		if err != nil {
			b.Fatal("x", err)
		}
	}
}
