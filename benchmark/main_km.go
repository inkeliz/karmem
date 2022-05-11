//go:build km
// +build km

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
	_Writer.Reset()
	if _, err := _Struct.WriteAsRoot(_Writer); err != nil {
		panic(err)
	}
}

//export KBenchmarkDecodeObjectAPI
func KBenchmarkDecodeObjectAPI(size uint32) {
	_Reader.SetSize(size)
	_Struct.ReadAsRoot(_Reader)
}

//export KBenchmarkDecodeObjectAPIFrom
func KBenchmarkDecodeObjectAPIFrom(b []byte) {
	_Struct.ReadAsRoot(karmem.NewReader(b))
}

//export KBenchmarkDecodeSumVec3
func KBenchmarkDecodeSumVec3(size uint32) float32 {
	_Reader.SetSize(size)

	monsters := km.NewMonstersViewer(_Reader, 0)
	monstersList := monsters.Monsters(_Reader)

	var sum km.Vec3
	for i := range monstersList {
		path := monstersList[i].Data(_Reader).Path(_Reader)

		for p := range path {
			sum.X += path[p].X()
			sum.Y += path[p].Y()
			sum.Z += path[p].Z()
		}
	}

	return sum.X + sum.Y + sum.Z
}
