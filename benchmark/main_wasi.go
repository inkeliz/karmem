//go:build wazero || wasmer
// +build wazero wasmer

package main

type Wasm interface {
	Write(b []byte) bool
	Reader(l uint32) []byte
	ReaderReset(b []byte)
	Run(s string, v ...uint64) ([]uint64, error)
	Close() error
}
