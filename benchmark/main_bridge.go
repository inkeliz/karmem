//go:build wazero || wasmer || tcp
// +build wazero wasmer tcp

package main

type Bridge interface {
	Write(b []byte) bool
	Reader(l uint32) []byte
	ReaderReset(b []byte)
	Run(s Functions, v ...uint64) ([]uint64, error)
	Close() error
}

type Functions uint32

func (f Functions) String() string {
	switch f {
	case FunctionKBenchmarkEncodeObjectAPI:
		return "KBenchmarkEncodeObjectAPI"
	case FunctionKBenchmarkDecodeObjectAPI:
		return "KBenchmarkDecodeObjectAPI"
	case FunctionKBenchmarkDecodeSumVec3:
		return "KBenchmarkDecodeSumVec3"
	default:
		panic("invalid function")
	}
}

const (
	FunctionKBenchmarkDecodeObjectAPI Functions = 1 + iota
	FunctionKBenchmarkEncodeObjectAPI
	FunctionKBenchmarkDecodeSumVec3
)
