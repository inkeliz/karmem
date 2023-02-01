// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"
)

type Color int8

const (
	ColorRed   Color = 0
	ColorGreen Color = 1
	ColorBlue  Color = 2
)

var EnumNamesColor = map[Color]string{
	ColorRed:   "Red",
	ColorGreen: "Green",
	ColorBlue:  "Blue",
}

var EnumValuesColor = map[string]Color{
	"Red":   ColorRed,
	"Green": ColorGreen,
	"Blue":  ColorBlue,
}

func (v Color) String() string {
	if s, ok := EnumNamesColor[v]; ok {
		return s
	}
	return "Color(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Team int8

const (
	TeamHumans  Team = 0
	TeamOrcs    Team = 1
	TeamZombies Team = 2
	TeamRobots  Team = 3
	TeamAliens  Team = 4
)

var EnumNamesTeam = map[Team]string{
	TeamHumans:  "Humans",
	TeamOrcs:    "Orcs",
	TeamZombies: "Zombies",
	TeamRobots:  "Robots",
	TeamAliens:  "Aliens",
}

var EnumValuesTeam = map[string]Team{
	"Humans":  TeamHumans,
	"Orcs":    TeamOrcs,
	"Zombies": TeamZombies,
	"Robots":  TeamRobots,
	"Aliens":  TeamAliens,
}

func (v Team) String() string {
	if s, ok := EnumNamesTeam[v]; ok {
		return s
	}
	return "Team(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Vec3T struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func (t *Vec3T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreateVec3(builder, t.X, t.Y, t.Z)
}
func (rcv *Vec3) UnPackTo(t *Vec3T) {
	t.X = rcv.X()
	t.Y = rcv.Y()
	t.Z = rcv.Z()
}

func (rcv *Vec3) UnPack() *Vec3T {
	if rcv == nil { return nil }
	t := &Vec3T{}
	rcv.UnPackTo(t)
	return t
}

type Vec3 struct {
	_tab flatbuffers.Struct
}

func (rcv *Vec3) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec3) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Vec3) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Vec3) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Vec3) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Vec3) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Vec3) Z() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Vec3) MutateZ(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateVec3(builder *flatbuffers.Builder, x float32, y float32, z float32) flatbuffers.UOffsetT {
	builder.Prep(4, 12)
	builder.PrependFloat32(z)
	builder.PrependFloat32(y)
	builder.PrependFloat32(x)
	return builder.Offset()
}
type WeaponT struct {
	Damage int32 `json:"damage"`
	Ammo uint16 `json:"ammo"`
	ClipSize byte `json:"clip_size"`
	ReloadTime float32 `json:"reload_time"`
	Range int32 `json:"range"`
}

func (t *WeaponT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	WeaponStart(builder)
	WeaponAddDamage(builder, t.Damage)
	WeaponAddAmmo(builder, t.Ammo)
	WeaponAddClipSize(builder, t.ClipSize)
	WeaponAddReloadTime(builder, t.ReloadTime)
	WeaponAddRange(builder, t.Range)
	return WeaponEnd(builder)
}

func (rcv *Weapon) UnPackTo(t *WeaponT) {
	t.Damage = rcv.Damage()
	t.Ammo = rcv.Ammo()
	t.ClipSize = rcv.ClipSize()
	t.ReloadTime = rcv.ReloadTime()
	t.Range = rcv.Range()
}

func (rcv *Weapon) UnPack() *WeaponT {
	if rcv == nil { return nil }
	t := &WeaponT{}
	rcv.UnPackTo(t)
	return t
}

type Weapon struct {
	_tab flatbuffers.Table
}

func GetRootAsWeapon(buf []byte, offset flatbuffers.UOffsetT) *Weapon {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Weapon{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWeapon(buf []byte, offset flatbuffers.UOffsetT) *Weapon {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Weapon{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Weapon) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Weapon) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Weapon) Damage() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Weapon) MutateDamage(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Weapon) Ammo() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Weapon) MutateAmmo(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *Weapon) ClipSize() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Weapon) MutateClipSize(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *Weapon) ReloadTime() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Weapon) MutateReloadTime(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *Weapon) Range() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Weapon) MutateRange(n int32) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

func WeaponStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func WeaponAddDamage(builder *flatbuffers.Builder, damage int32) {
	builder.PrependInt32Slot(0, damage, 0)
}
func WeaponAddAmmo(builder *flatbuffers.Builder, ammo uint16) {
	builder.PrependUint16Slot(1, ammo, 0)
}
func WeaponAddClipSize(builder *flatbuffers.Builder, clipSize byte) {
	builder.PrependByteSlot(2, clipSize, 0)
}
func WeaponAddReloadTime(builder *flatbuffers.Builder, reloadTime float32) {
	builder.PrependFloat32Slot(3, reloadTime, 0.0)
}
func WeaponAddRange(builder *flatbuffers.Builder, range_ int32) {
	builder.PrependInt32Slot(4, range_, 0)
}
func WeaponEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type MonstersT struct {
	Monsters []*MonsterT `json:"monsters"`
}

func (t *MonstersT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	monstersOffset := flatbuffers.UOffsetT(0)
	if t.Monsters != nil {
		monstersLength := len(t.Monsters)
		monstersOffsets := make([]flatbuffers.UOffsetT, monstersLength)
		for j := 0; j < monstersLength; j++ {
			monstersOffsets[j] = t.Monsters[j].Pack(builder)
		}
		MonstersStartMonstersVector(builder, monstersLength)
		for j := monstersLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(monstersOffsets[j])
		}
		monstersOffset = builder.EndVector(monstersLength)
	}
	MonstersStart(builder)
	MonstersAddMonsters(builder, monstersOffset)
	return MonstersEnd(builder)
}

func (rcv *Monsters) UnPackTo(t *MonstersT) {
	monstersLength := rcv.MonstersLength()
	t.Monsters = make([]*MonsterT, monstersLength)
	for j := 0; j < monstersLength; j++ {
		x := Monster{}
		rcv.Monsters(&x, j)
		t.Monsters[j] = x.UnPack()
	}
}

func (rcv *Monsters) UnPack() *MonstersT {
	if rcv == nil { return nil }
	t := &MonstersT{}
	rcv.UnPackTo(t)
	return t
}

type Monsters struct {
	_tab flatbuffers.Table
}

func GetRootAsMonsters(buf []byte, offset flatbuffers.UOffsetT) *Monsters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Monsters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMonsters(buf []byte, offset flatbuffers.UOffsetT) *Monsters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Monsters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Monsters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Monsters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Monsters) Monsters(obj *Monster, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monsters) MonstersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MonstersStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MonstersAddMonsters(builder *flatbuffers.Builder, monsters flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(monsters), 0)
}
func MonstersStartMonstersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MonstersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type MonsterT struct {
	Pos *Vec3T `json:"pos"`
	Mana int16 `json:"mana"`
	Health int16 `json:"health"`
	Name string `json:"name"`
	Team Team `json:"team"`
	Inventory []byte `json:"inventory"`
	Color Color `json:"color"`
	Hitbox []float64 `json:"hitbox"`
	Status []int32 `json:"status"`
	Weapons []*WeaponT `json:"weapons"`
	Path []*Vec3T `json:"path"`
	IsAlive bool `json:"is_alive"`
}

func (t *MonsterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	inventoryOffset := flatbuffers.UOffsetT(0)
	if t.Inventory != nil {
		inventoryOffset = builder.CreateByteString(t.Inventory)
	}
	hitboxOffset := flatbuffers.UOffsetT(0)
	if t.Hitbox != nil {
		hitboxLength := len(t.Hitbox)
		MonsterStartHitboxVector(builder, hitboxLength)
		for j := hitboxLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Hitbox[j])
		}
		hitboxOffset = builder.EndVector(hitboxLength)
	}
	statusOffset := flatbuffers.UOffsetT(0)
	if t.Status != nil {
		statusLength := len(t.Status)
		MonsterStartStatusVector(builder, statusLength)
		for j := statusLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.Status[j])
		}
		statusOffset = builder.EndVector(statusLength)
	}
	weaponsOffset := flatbuffers.UOffsetT(0)
	if t.Weapons != nil {
		weaponsLength := len(t.Weapons)
		weaponsOffsets := make([]flatbuffers.UOffsetT, weaponsLength)
		for j := 0; j < weaponsLength; j++ {
			weaponsOffsets[j] = t.Weapons[j].Pack(builder)
		}
		MonsterStartWeaponsVector(builder, weaponsLength)
		for j := weaponsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(weaponsOffsets[j])
		}
		weaponsOffset = builder.EndVector(weaponsLength)
	}
	pathOffset := flatbuffers.UOffsetT(0)
	if t.Path != nil {
		pathLength := len(t.Path)
		MonsterStartPathVector(builder, pathLength)
		for j := pathLength - 1; j >= 0; j-- {
			t.Path[j].Pack(builder)
		}
		pathOffset = builder.EndVector(pathLength)
	}
	MonsterStart(builder)
	posOffset := t.Pos.Pack(builder)
	MonsterAddPos(builder, posOffset)
	MonsterAddMana(builder, t.Mana)
	MonsterAddHealth(builder, t.Health)
	MonsterAddName(builder, nameOffset)
	MonsterAddTeam(builder, t.Team)
	MonsterAddInventory(builder, inventoryOffset)
	MonsterAddColor(builder, t.Color)
	MonsterAddHitbox(builder, hitboxOffset)
	MonsterAddStatus(builder, statusOffset)
	MonsterAddWeapons(builder, weaponsOffset)
	MonsterAddPath(builder, pathOffset)
	MonsterAddIsAlive(builder, t.IsAlive)
	return MonsterEnd(builder)
}

func (rcv *Monster) UnPackTo(t *MonsterT) {
	t.Pos = rcv.Pos(nil).UnPack()
	t.Mana = rcv.Mana()
	t.Health = rcv.Health()
	t.Name = string(rcv.Name())
	t.Team = rcv.Team()
	t.Inventory = rcv.InventoryBytes()
	t.Color = rcv.Color()
	hitboxLength := rcv.HitboxLength()
	t.Hitbox = make([]float64, hitboxLength)
	for j := 0; j < hitboxLength; j++ {
		t.Hitbox[j] = rcv.Hitbox(j)
	}
	statusLength := rcv.StatusLength()
	t.Status = make([]int32, statusLength)
	for j := 0; j < statusLength; j++ {
		t.Status[j] = rcv.Status(j)
	}
	weaponsLength := rcv.WeaponsLength()
	t.Weapons = make([]*WeaponT, weaponsLength)
	for j := 0; j < weaponsLength; j++ {
		x := Weapon{}
		rcv.Weapons(&x, j)
		t.Weapons[j] = x.UnPack()
	}
	pathLength := rcv.PathLength()
	t.Path = make([]*Vec3T, pathLength)
	for j := 0; j < pathLength; j++ {
		x := Vec3{}
		rcv.Path(&x, j)
		t.Path[j] = x.UnPack()
	}
	t.IsAlive = rcv.IsAlive()
}

func (rcv *Monster) UnPack() *MonsterT {
	if rcv == nil { return nil }
	t := &MonsterT{}
	rcv.UnPackTo(t)
	return t
}

type Monster struct {
	_tab flatbuffers.Table
}

func GetRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Monster{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Monster{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Monster) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Monster) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Monster) Pos(obj *Vec3) *Vec3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vec3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Mana() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateMana(n int16) bool {
	return rcv._tab.MutateInt16Slot(6, n)
}

func (rcv *Monster) Health() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateHealth(n int16) bool {
	return rcv._tab.MutateInt16Slot(8, n)
}

func (rcv *Monster) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Team() Team {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return Team(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Monster) MutateTeam(n Team) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

func (rcv *Monster) Inventory(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Monster) InventoryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) InventoryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) MutateInventory(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Monster) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Monster) MutateColor(n Color) bool {
	return rcv._tab.MutateInt8Slot(16, int8(n))
}

func (rcv *Monster) Hitbox(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *Monster) HitboxLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) MutateHitbox(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *Monster) Status(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Monster) StatusLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) MutateStatus(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *Monster) Weapons(obj *Weapon, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) WeaponsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Path(obj *Vec3, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 12
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) PathLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) IsAlive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Monster) MutateIsAlive(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

func MonsterStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func MonsterAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(pos), 0)
}
func MonsterAddMana(builder *flatbuffers.Builder, mana int16) {
	builder.PrependInt16Slot(1, mana, 0)
}
func MonsterAddHealth(builder *flatbuffers.Builder, health int16) {
	builder.PrependInt16Slot(2, health, 0)
}
func MonsterAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(name), 0)
}
func MonsterAddTeam(builder *flatbuffers.Builder, team Team) {
	builder.PrependInt8Slot(4, int8(team), 0)
}
func MonsterAddInventory(builder *flatbuffers.Builder, inventory flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(inventory), 0)
}
func MonsterStartInventoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MonsterAddColor(builder *flatbuffers.Builder, color Color) {
	builder.PrependInt8Slot(6, int8(color), 0)
}
func MonsterAddHitbox(builder *flatbuffers.Builder, hitbox flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(hitbox), 0)
}
func MonsterStartHitboxVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func MonsterAddStatus(builder *flatbuffers.Builder, status flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(status), 0)
}
func MonsterStartStatusVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MonsterAddWeapons(builder *flatbuffers.Builder, weapons flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(weapons), 0)
}
func MonsterStartWeaponsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MonsterAddPath(builder *flatbuffers.Builder, path flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(path), 0)
}
func MonsterStartPathVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(12, numElems, 4)
}
func MonsterAddIsAlive(builder *flatbuffers.Builder, isAlive bool) {
	builder.PrependBoolSlot(11, isAlive, false)
}
func MonsterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
