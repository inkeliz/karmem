//go:build struct
// +build struct

package main

import (
	"benchmark.karmem.org/km"
	karmem "karmem.org/golang"
)

var (
	_Struct km.Monsters
	_Writer = karmem.NewFixedWriter(OutputMemory[:])
	_Reader = karmem.NewReader(InputMemory[:])
)

//export KBenchmarkEncodeObjectAPI
func KBenchmarkEncodeObjectAPI() {

}

//export KBenchmarkDecodeObjectAPI
func KBenchmarkDecodeObjectAPI() {
	_Struct.ReadAsRoot(_Reader)
}

//export KBenchmarkDecodeSumVec3
func KBenchmarkDecodeSumVec3(_ uint32) float32 {
	var sum km.Vec3
	monsters := _Struct.Monsters
	for i := range monsters {
		path := monsters[i].Data.Path
		for i := range path {
			sum.X += path[i].X
			sum.Y += path[i].Y
			sum.Z += path[i].Z
		}
	}

	return sum.X + sum.Y + sum.Z
}
