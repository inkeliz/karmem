//go:build !tinygo && fbs

package main

import (
	"benchmark.karmem.org/fbs"
	"benchmark.karmem.org/km"
	flatbuffers "github.com/google/flatbuffers/go"
)

var IsKarmem = false
var KStruct *fbs.MonstersT
var KBuilder = flatbuffers.NewBuilder(1000)
var KStarted = false
var KSumExpected = float32(0.0)

func init() {
	initEncode()
}

var _pathLen = 100

func initEncode() {
	if KStarted {
		return
	}
	KStarted = true
	KStruct = &fbs.MonstersT{
		Monsters: make([]*fbs.MonsterT, 1000),
	}
	var sum km.Vec3
	for i := range KStruct.Monsters {
		KStruct.Monsters[i] = &fbs.MonsterT{
			Pos:       &fbs.Vec3T{X: 1, Y: 2, Z: 3},
			Mana:      int16(1 * i),
			Health:    int16(2 * i),
			Name:      "Ǩ翐p±ƀ鵁灉Ĭ冖ɠctqo酂",
			Team:      fbs.TeamAliens,
			Inventory: make([]uint8, 100),
			Color:     0,
			//	Hitbox:    [5]float64{4.2, 313.12, 4.02, 90, -2942.123},
			//	Status:    []int32{23123, 83859, 30123123, 34123, -23111, -923, 93123, 11123},
			Weapons: []*fbs.WeaponT{
				{
					Damage: 100,
					Range:  int32(i),
				},
				{
					Damage: 101,
					Range:  int32(i),
				},
			},
			Path: make([]*fbs.Vec3T, _pathLen),
		}
		for j := range KStruct.Monsters[i].Path {
			v := &fbs.Vec3T{
				X: float32(1),
				Y: float32(i),
				Z: float32(1),
			}
			sum.X += v.X
			sum.Y += v.Y
			sum.Z += v.Z
			KStruct.Monsters[i].Path[j] = v
		}
	}
	KSumExpected += sum.X + sum.Y + sum.Z
}

func encode() []byte {
	KBuilder.Reset()
	KBuilder.Finish(KStruct.Pack(KBuilder))
	encoded := KBuilder.FinishedBytes()
	return encoded
}

func decode(b []byte) (v *fbs.MonstersT) {
	x := new(fbs.MonstersT)
	fbs.GetRootAsMonsters(b, 0).UnPackTo(x)
	return x
}
