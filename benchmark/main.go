package main

import (
	_ "embed"
	"unsafe"
)

var InputMemory [8_000_000]byte
var OutputMemory [8_000_000]byte

func main() {}

//export InputMemoryPointer
func InputMemoryPointer() uint32 {
	return uint32(uintptr(unsafe.Pointer(&InputMemory[0])))
}

//export OutputMemoryPointer
func OutputMemoryPointer() uint32 {
	return uint32(uintptr(unsafe.Pointer(&OutputMemory[0])))
}
