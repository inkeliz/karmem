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

	var sum float32
	for i := range monstersList {
		path := monstersList[i].Data(_Reader).Path(_Reader)

		for p := range path {
			pp := &path[p]
			sum += pp.X() + pp.Y() + pp.Z()
		}
	}
	return sum
}

//export KBenchmarkDecodeSumUint8
func KBenchmarkDecodeSumUint8(size uint32) uint32 {
	_Reader.SetSize(size)

	monsters := km.NewMonstersViewer(_Reader, 0)
	monstersList := monsters.Monsters(_Reader)

	var sum uint32
	for i := range monstersList {
		inv := monstersList[i].Data(_Reader).Inventory(_Reader)

		for j := range inv {
			sum += uint32(inv[j])
		}
	}

	return sum
}

//export KBenchmarkDecodeSumStats
func KBenchmarkDecodeSumStats(size uint32) uint32 {
	_Reader.SetSize(size)

	monsters := km.NewMonstersViewer(_Reader, 0)
	monstersList := monsters.Monsters(_Reader)

	var sum km.WeaponData
	for i := range monstersList {
		weapons := monstersList[i].Data(_Reader).Weapons()

		for j := range weapons {
			data := weapons[j].Data(_Reader)
			sum.Ammo += data.Ammo()
			sum.Damage += data.Damage()
			sum.ClipSize += data.ClipSize()
			sum.ReloadTime += data.ReloadTime()
			sum.Range += data.Range()
		}
	}

	_Writer.Reset()
	if _, err := sum.WriteAsRoot(_Writer); err != nil {
		panic(err)
	}
	return uint32(len(_Writer.Bytes()))
}

//export KNOOP
func KNOOP() uint32 {
	return 42
}
