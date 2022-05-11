//go:build km
// +build km

package main

import (
	"testing"
)

func FuzzNRandom(f *testing.F) {
	_Struct = KStruct
	f.Add(encode())

	f.Fuzz(func(t *testing.T, b []byte) {
		KBenchmarkDecodeObjectAPIFrom(b)
	})
}
