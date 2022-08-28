//go:build !wasi && !struct && !tcp && km && !nobench

package main

import (
	"testing"

	"benchmark.karmem.org/km"
	karmem "karmem.org/golang"
)

func BenchmarkDecodeSumUint8(b *testing.B) {
	encoded := encode()
	copy(InputMemory[:], encoded)

	b.ReportAllocs()
	b.ResetTimer()
	l := uint32(len(encoded))
	for i := 0; i < b.N; i++ {
		if n := KBenchmarkDecodeSumUint8(l); n != KSumInventoryExpected {
			b.Error("invalid sum", n, KSumExpected)
		}
	}
}

func BenchmarkDecodeSumStats(b *testing.B) {
	encoded := encode()
	copy(InputMemory[:], encoded)

	reader := karmem.NewReader(OutputMemory[:])
	b.ReportAllocs()
	b.ResetTimer()
	l := uint32(len(encoded))
	for i := 0; i < b.N; i++ {
		size := KBenchmarkDecodeSumStats(l)
		reader.SetSize(size)

		result := km.NewWeaponDataViewer(reader, 0)
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
