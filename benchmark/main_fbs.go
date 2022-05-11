//go:build fbs
// +build fbs

package main

import (
	"benchmark.karmem.org/fbs"
	flatbuffers "github.com/google/flatbuffers/go"
)

var (
	_Struct = new(fbs.MonstersT)
	_Writer *flatbuffers.Builder
)

func init() {
    _Writer = &flatbuffers.Builder{Bytes: OutputMemory[:]}
}

//export KBenchmarkEncodeObjectAPI
func KBenchmarkEncodeObjectAPI() {
	_Writer.Reset()
	_Writer.Finish(_Struct.Pack(_Writer))
}

//export KBenchmarkDecodeObjectAPI
func KBenchmarkDecodeObjectAPI(_ uint32) {
	fbs.GetRootAsMonsters(InputMemory[:], 0).UnPackTo(_Struct)
}

//export KBenchmarkDecodeSumVec3
func KBenchmarkDecodeSumVec3(_ uint32) float32 {
	_Reader := fbs.GetRootAsMonsters(InputMemory[:], 0)

	var monster fbs.Monster
	var path fbs.Vec3
	var sum fbs.Vec3T

	l := _Reader.MonstersLength()
	for i := 0; i < l; i++ {
		_Reader.Monsters(&monster, i)

		l := monster.PathLength()
		for p := 0; p < l; p++ {
			monster.Path(&path, p)
			sum.X += path.X()
			sum.Y += path.Y()
			sum.Z += path.Z()
		}
	}
	return sum.X + sum.Y + sum.Z
}
