//go:build !fbs && (wazero || wasmer)
// +build !fbs
// +build wazero wasmer

package main

import (
	"github.com/r3labs/diff/v3"
	"testing"
)

func init() {
	initEncode()
}

func TestEncodeObjectAPI(t *testing.T) {
	m := initWasm(t, "KBenchmarkEncodeObjectAPI", "KBenchmarkDecodeObjectAPI")
	defer func() {
		if err := m.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	encoded := encode()
	l := uint64(len(encoded))
	m.Write(encoded)
	if _, err := m.Run("KBenchmarkDecodeObjectAPI", l); err != nil {
		t.Fatal(err)
	}

	_, err := m.Run("KBenchmarkEncodeObjectAPI")
	if err != nil {
		t.Fatal(err)
	}

	var expected = KStruct
	result := decode(m.Reader(uint32(len(OutputMemory))))

	di, err := diff.NewDiffer()
	if err != nil {
		t.Errorf("diff lib error: %v", err)
		return
	}

	changes, err := di.Diff(expected, result)
	if err != nil {
		t.Errorf("diff lib error: %v", err)
		return
	}

	if len(changes) > 0 {
		for _, x := range changes {
			t.Error(x.Type, x.Path)
		}
		t.Fatal("changes detected")
	}

}
