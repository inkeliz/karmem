//go:build !tinygo && (km || struct)

package main

import (
	"benchmark.karmem.org/km"
	karmem "karmem.org/golang"
)

var IsKarmem = true
var KStruct km.Monsters
var KBuilder = karmem.NewWriter(1000)
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
	KStruct = km.Monsters{
		Monsters: make([]km.Monster, 1000),
	}
	var sum km.Vec3
	for i := range KStruct.Monsters {
		KStruct.Monsters[i] = km.Monster{
			Data: km.MonsterData{
				Pos:       km.Vec3{X: 1, Y: 2, Z: 3},
				Mana:      int16(1 * i),
				Health:    int16(2 * i),
				Name:      "Ǩ翐p±ƀ鵁灉Ĭ冖ɠctqo酂",
				Team:      km.TeamAliens,
				Inventory: make([]uint8, 100),
				Color:     0,
				Hitbox:    [5]float64{4.2, 313.12, 4.02, 90, -2942.123},
				Status:    []int32{23123, 83859, 30123123, 34123, -23111, -923, 93123, 11123},
				Weapons: [4]km.Weapon{
					{
						Data: km.WeaponData{
							Damage: 100,
							Range:  int32(i),
						},
					},
					{
						Data: km.WeaponData{
							Damage: 101,
							Range:  int32(i),
						},
					},
				},
				Path:    make([]km.Vec3, _pathLen),
				IsAlive: i&1 == 0,
			},
		}
		for j := range KStruct.Monsters[i].Data.Path {
			v := km.Vec3{
				X: float32(1),
				Y: float32(i),
				Z: float32(1),
			}
			sum.X += v.X
			sum.Y += v.Y
			sum.Z += v.Z
			KStruct.Monsters[i].Data.Path[j] = v
		}
	}
	KSumExpected += sum.X + sum.Y + sum.Z
}

func encode() []byte {
	KBuilder.Reset()
	if _, err := KStruct.Write(KBuilder, 0); err != nil {
		panic(err)
	}
	encoded := KBuilder.Bytes()
	return encoded
}

func decode(b []byte) (v km.Monsters) {
	v.ReadAsRoot(karmem.NewReader(b))
	return v
}
