//go:build (wazero || wasmer || tcp) && !nobench && km
// +build wazero wasmer tcp
// +build !nobench
// +build km

package main

import (
	"testing"

	"benchmark.karmem.org/km"
	karmem "karmem.org/golang"
)

func init() {
	initEncode()
}

func BenchmarkDecodeSumUint8(b *testing.B) {
	m := initBridge(b, "KBenchmarkEncodeObjectAPI", "KBenchmarkDecodeObjectAPI", "KBenchmarkDecodeSumUint8")
	defer m.Close()
	encoded := encode()
	m.Write(encoded)
	m.Run(FunctionKBenchmarkDecodeObjectAPI, uint64(len(encoded)))

	b.ResetTimer()
	b.ReportAllocs()
	l := uint64(len(encoded))
	for i := 0; i < b.N; i++ {
		n, err := m.Run(FunctionKBenchmarkDecodeSumUint8, l)
		if err != nil {
			b.Fatal("x", err)
		}
		if uint32(n[0]) != KSumInventoryExpected {
			b.Error("invalid sum", n, KSumInventoryExpected)
		}
	}
}

func BenchmarkDecodeSumStats(b *testing.B) {
	m := initBridge(b, "KBenchmarkEncodeObjectAPI", "KBenchmarkDecodeObjectAPI", "KBenchmarkDecodeSumStats")
	defer m.Close()
	encoded := encode()
	m.Write(encoded)
	m.Run(FunctionKBenchmarkDecodeObjectAPI, uint64(len(encoded)))

	b.ResetTimer()
	b.ReportAllocs()
	l := uint64(len(encoded))
	for i := 0; i < b.N; i++ {
		n, err := m.Run(FunctionKBenchmarkDecodeSumStats, l)
		if err != nil {
			b.Fatal("x", err)
		}

		result := km.NewWeaponDataViewer(karmem.NewReader(m.Reader(uint32(n[0]))), 0)

		if n := result.Damage(); n != KSumWeaponStatus.Damage {
			b.Error("invalid sum", n, KSumWeaponStatus.Damage)
		}

		if n := result.Range(); n != KSumWeaponStatus.Range {
			b.Error("invalid sum", n, KSumWeaponStatus.Range)
		}

		if n := result.Ammo(); n != KSumWeaponStatus.Ammo {
			b.Error("invalid sum", n, KSumWeaponStatus.Ammo)
		}

		if n := result.ClipSize(); n != KSumWeaponStatus.ClipSize {
			b.Error("invalid sum", n, KSumWeaponStatus.ClipSize)
		}
	}
}

func BenchmarkNoOp(b *testing.B) {
	m := initBridge(b, "KBenchmarkEncodeObjectAPI", "KNOOP")
	defer m.Close()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		n, err := m.Run(FunctionKNOOP)
		if err != nil {
			b.Fatal("x", err)
		}
		if uint32(n[0]) != 42 {
			b.Error("invalid no-op value")
		}
	}
}
