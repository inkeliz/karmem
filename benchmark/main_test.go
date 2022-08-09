//go:build !wasi && !struct && !tcp
// +build !wasi,!struct,!tcp

package main

import (
	"fmt"
	"reflect"
	"testing"
)

func init() {
	initEncode()
}

func TestDecodeObjectAPI(t *testing.T) {
	if !IsKarmem {
		return
	}
	encode()
	_Struct = KStruct
	decoded := decode(encode())
	if !reflect.DeepEqual(decoded, KStruct) {
		fmt.Println(decoded)
		t.Error("not match")
	}
}

func BenchmarkEncodeObjectAPI(b *testing.B) {
	encode()
	_Struct = KStruct
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		KBenchmarkEncodeObjectAPI()
	}
}

func BenchmarkDecodeObjectAPI(b *testing.B) {
	encoded := encode()
	copy(InputMemory[:], encoded)

	b.ResetTimer()
	b.ReportAllocs()
	l := uint32(len(encoded))
	for i := 0; i < b.N; i++ {
		KBenchmarkDecodeObjectAPI(l)
	}
	if !reflect.DeepEqual(_Struct, KStruct) {
		b.Error("not match")
	}
}

func BenchmarkDecodeSumVec3(b *testing.B) {
	encoded := encode()
	copy(InputMemory[:], encoded)

	b.ReportAllocs()
	b.ResetTimer()
	l := uint32(len(encoded))
	for i := 0; i < b.N; i++ {
		if n := KBenchmarkDecodeSumVec3(l); n != KSumExpected {
			b.Error("invalid sum", n, KSumExpected)
		}
	}
}
