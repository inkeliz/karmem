//go:build !fbs && (wazero || wasmer || tcp)
// +build !fbs
// +build wazero wasmer tcp

package main

import (
	"testing"

	"github.com/r3labs/diff/v3"
)

func init() {
	initEncode()
}

func TestEncodeObjectAPI(t *testing.T) {
	m := initBridge(t, "KBenchmarkEncodeObjectAPI", "KBenchmarkDecodeObjectAPI")
	defer func() {
		if err := m.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	encoded := encode()
	l := uint64(len(encoded))
	m.Write(encoded)
	if _, err := m.Run(FunctionKBenchmarkDecodeObjectAPI, l); err != nil {
		t.Fatal(err)
	}

	_, err := m.Run(FunctionKBenchmarkEncodeObjectAPI)
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
