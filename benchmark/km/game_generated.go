package km

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

var _Null = make([]byte, 111)
var _NullReader = karmem.NewReader(_Null)

type (
	Color uint8
)

const (
	ColorRed   Color = 0
	ColorGreen Color = 1
	ColorBlue  Color = 2
)

type (
	Team uint8
)

const (
	TeamHumans  Team = 0
	TeamOrcs    Team = 1
	TeamZombies Team = 2
	TeamRobots  Team = 3
	TeamAliens  Team = 4
)

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierVec3        = 10268726485798425099
	PacketIdentifierWeaponData  = 15342010214468761012
	PacketIdentifierWeapon      = 8029074423243608167
	PacketIdentifierMonsterData = 12254962724431809041
	PacketIdentifierMonster     = 5593793986513565154
	PacketIdentifierMonsters    = 14096677544474027661
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func NewVec3() Vec3 {
	return Vec3{}
}

func (x *Vec3) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierVec3
}

func (x *Vec3) Reset() {
	x.Read((*Vec3Viewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Vec3) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Vec3) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(12)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__XOffset := offset + 0
	writer.Write4At(__XOffset, *(*uint32)(unsafe.Pointer(&x.X)))
	__YOffset := offset + 4
	writer.Write4At(__YOffset, *(*uint32)(unsafe.Pointer(&x.Y)))
	__ZOffset := offset + 8
	writer.Write4At(__ZOffset, *(*uint32)(unsafe.Pointer(&x.Z)))

	return offset, nil
}

func (x *Vec3) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewVec3Viewer(reader, 0), reader)
}

func (x *Vec3) Read(viewer *Vec3Viewer, reader *karmem.Reader) {
	x.X = viewer.X()
	x.Y = viewer.Y()
	x.Z = viewer.Z()
}

type WeaponData struct {
	Damage int32
	Range  int32
}

func NewWeaponData() WeaponData {
	return WeaponData{}
}

func (x *WeaponData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierWeaponData
}

func (x *WeaponData) Reset() {
	x.Read((*WeaponDataViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *WeaponData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *WeaponData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(12)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(12))
	__DamageOffset := offset + 4
	writer.Write4At(__DamageOffset, *(*uint32)(unsafe.Pointer(&x.Damage)))
	__RangeOffset := offset + 8
	writer.Write4At(__RangeOffset, *(*uint32)(unsafe.Pointer(&x.Range)))

	return offset, nil
}

func (x *WeaponData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewWeaponDataViewer(reader, 0), reader)
}

func (x *WeaponData) Read(viewer *WeaponDataViewer, reader *karmem.Reader) {
	x.Damage = viewer.Damage()
	x.Range = viewer.Range()
}

type Weapon struct {
	Data WeaponData
}

func NewWeapon() Weapon {
	return Weapon{}
}

func (x *Weapon) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierWeapon
}

func (x *Weapon) Reset() {
	x.Read((*WeaponViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Weapon) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Weapon) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(12)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *Weapon) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewWeaponViewer(reader, 0), reader)
}

func (x *Weapon) Read(viewer *WeaponViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type MonsterData struct {
	Pos       Vec3
	Mana      int16
	Health    int16
	Name      string
	Team      Team
	Inventory []byte
	Color     Color
	Hitbox    [5]float64
	Status    []int32
	Weapons   [4]Weapon
	Path      []Vec3
	IsAlive   bool
}

func NewMonsterData() MonsterData {
	return MonsterData{}
}

func (x *MonsterData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierMonsterData
}

func (x *MonsterData) Reset() {
	x.Read((*MonsterDataViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *MonsterData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *MonsterData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(111)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(111))
	__PosOffset := offset + 4
	if _, err := x.Pos.Write(writer, __PosOffset); err != nil {
		return offset, err
	}
	__ManaOffset := offset + 16
	writer.Write2At(__ManaOffset, *(*uint16)(unsafe.Pointer(&x.Mana)))
	__HealthOffset := offset + 18
	writer.Write2At(__HealthOffset, *(*uint16)(unsafe.Pointer(&x.Health)))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__NameOffset))
	writer.Write4At(offset+20+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__TeamOffset := offset + 28
	writer.Write1At(__TeamOffset, *(*uint8)(unsafe.Pointer(&x.Team)))
	__InventorySize := uint(1 * len(x.Inventory))
	__InventoryOffset, err := writer.Alloc(__InventorySize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+29, uint32(__InventoryOffset))
	writer.Write4At(offset+29+4, uint32(__InventorySize))
	__InventorySlice := *(*[3]uint)(unsafe.Pointer(&x.Inventory))
	__InventorySlice[1] = __InventorySize
	__InventorySlice[2] = __InventorySize
	writer.WriteAt(__InventoryOffset, *(*[]byte)(unsafe.Pointer(&__InventorySlice)))
	__ColorOffset := offset + 37
	writer.Write1At(__ColorOffset, *(*uint8)(unsafe.Pointer(&x.Color)))
	__HitboxOffset := offset + 38
	writer.WriteAt(__HitboxOffset, (*[40]byte)(unsafe.Pointer(&x.Hitbox))[:])
	__StatusSize := uint(4 * len(x.Status))
	__StatusOffset, err := writer.Alloc(__StatusSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+78, uint32(__StatusOffset))
	writer.Write4At(offset+78+4, uint32(__StatusSize))
	__StatusSlice := *(*[3]uint)(unsafe.Pointer(&x.Status))
	__StatusSlice[1] = __StatusSize
	__StatusSlice[2] = __StatusSize
	writer.WriteAt(__StatusOffset, *(*[]byte)(unsafe.Pointer(&__StatusSlice)))
	__WeaponsOffset := offset + 86
	for i := range x.Weapons {
		if _, err := x.Weapons[i].Write(writer, __WeaponsOffset); err != nil {
			return offset, err
		}
		__WeaponsOffset += 4
	}
	__PathSize := uint(12 * len(x.Path))
	__PathOffset, err := writer.Alloc(__PathSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+102, uint32(__PathOffset))
	writer.Write4At(offset+102+4, uint32(__PathSize))
	for i := range x.Path {
		if _, err := x.Path[i].Write(writer, __PathOffset); err != nil {
			return offset, err
		}
		__PathOffset += 12
	}
	__IsAliveOffset := offset + 110
	writer.Write1At(__IsAliveOffset, *(*uint8)(unsafe.Pointer(&x.IsAlive)))

	return offset, nil
}

func (x *MonsterData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewMonsterDataViewer(reader, 0), reader)
}

func (x *MonsterData) Read(viewer *MonsterDataViewer, reader *karmem.Reader) {
	x.Pos.Read(viewer.Pos(), reader)
	x.Mana = viewer.Mana()
	x.Health = viewer.Health()
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	x.Team = Team(viewer.Team())
	__InventorySlice := viewer.Inventory(reader)
	__InventoryLen := len(__InventorySlice)
	if __InventoryLen > cap(x.Inventory) {
		x.Inventory = append(x.Inventory, make([]byte, __InventoryLen-len(x.Inventory))...)
	}
	if __InventoryLen > len(x.Inventory) {
		x.Inventory = x.Inventory[:__InventoryLen]
	}
	copy(x.Inventory, __InventorySlice)
	x.Inventory = x.Inventory[:__InventoryLen]
	x.Color = Color(viewer.Color())
	__HitboxSlice := viewer.Hitbox()
	__HitboxLen := len(__HitboxSlice)
	if __HitboxLen > 5 {
		__HitboxLen = 5
	}
	copy(x.Hitbox[:], __HitboxSlice)
	for i := __HitboxLen; i < len(x.Hitbox); i++ {
		x.Hitbox[i] = 0
	}
	__StatusSlice := viewer.Status(reader)
	__StatusLen := len(__StatusSlice)
	if __StatusLen > cap(x.Status) {
		x.Status = append(x.Status, make([]int32, __StatusLen-len(x.Status))...)
	}
	if __StatusLen > len(x.Status) {
		x.Status = x.Status[:__StatusLen]
	}
	copy(x.Status, __StatusSlice)
	x.Status = x.Status[:__StatusLen]
	__WeaponsSlice := viewer.Weapons()
	__WeaponsLen := len(__WeaponsSlice)
	if __WeaponsLen > 4 {
		__WeaponsLen = 4
	}
	for i := 0; i < __WeaponsLen; i++ {
		x.Weapons[i].Read(&__WeaponsSlice[i], reader)
	}
	for i := __WeaponsLen; i < len(x.Weapons); i++ {
		x.Weapons[i].Reset()
	}
	__PathSlice := viewer.Path(reader)
	__PathLen := len(__PathSlice)
	if __PathLen > cap(x.Path) {
		x.Path = append(x.Path, make([]Vec3, __PathLen-len(x.Path))...)
	}
	if __PathLen > len(x.Path) {
		x.Path = x.Path[:__PathLen]
	}
	for i := 0; i < __PathLen; i++ {
		x.Path[i].Read(&__PathSlice[i], reader)
	}
	x.Path = x.Path[:__PathLen]
	x.IsAlive = viewer.IsAlive()
}

type Monster struct {
	Data MonsterData
}

func NewMonster() Monster {
	return Monster{}
}

func (x *Monster) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierMonster
}

func (x *Monster) Reset() {
	x.Read((*MonsterViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Monster) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Monster) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(111)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *Monster) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewMonsterViewer(reader, 0), reader)
}

func (x *Monster) Read(viewer *MonsterViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type Monsters struct {
	Monsters []Monster
}

func NewMonsters() Monsters {
	return Monsters{}
}

func (x *Monsters) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierMonsters
}

func (x *Monsters) Reset() {
	x.Read((*MonstersViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Monsters) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Monsters) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(12)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(12))
	__MonstersSize := uint(4 * len(x.Monsters))
	__MonstersOffset, err := writer.Alloc(__MonstersSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__MonstersOffset))
	writer.Write4At(offset+4+4, uint32(__MonstersSize))
	for i := range x.Monsters {
		if _, err := x.Monsters[i].Write(writer, __MonstersOffset); err != nil {
			return offset, err
		}
		__MonstersOffset += 4
	}

	return offset, nil
}

func (x *Monsters) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewMonstersViewer(reader, 0), reader)
}

func (x *Monsters) Read(viewer *MonstersViewer, reader *karmem.Reader) {
	__MonstersSlice := viewer.Monsters(reader)
	__MonstersLen := len(__MonstersSlice)
	if __MonstersLen > cap(x.Monsters) {
		x.Monsters = append(x.Monsters, make([]Monster, __MonstersLen-len(x.Monsters))...)
	}
	if __MonstersLen > len(x.Monsters) {
		x.Monsters = x.Monsters[:__MonstersLen]
	}
	for i := 0; i < __MonstersLen; i++ {
		x.Monsters[i].Read(&__MonstersSlice[i], reader)
	}
	x.Monsters = x.Monsters[:__MonstersLen]
}

type Vec3Viewer [12]byte

func NewVec3Viewer(reader *karmem.Reader, offset uint32) (v *Vec3Viewer) {
	if !reader.IsValidOffset(offset, 12) {
		return (*Vec3Viewer)(unsafe.Pointer(&_Null))
	}
	v = (*Vec3Viewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *Vec3Viewer) size() uint32 {
	return 12
}
func (x *Vec3Viewer) X() (v float32) {
	return *(*float32)(unsafe.Add(unsafe.Pointer(x), 0))
}
func (x *Vec3Viewer) Y() (v float32) {
	return *(*float32)(unsafe.Add(unsafe.Pointer(x), 4))
}
func (x *Vec3Viewer) Z() (v float32) {
	return *(*float32)(unsafe.Add(unsafe.Pointer(x), 8))
}

type WeaponDataViewer [12]byte

func NewWeaponDataViewer(reader *karmem.Reader, offset uint32) (v *WeaponDataViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*WeaponDataViewer)(unsafe.Pointer(&_Null))
	}
	v = (*WeaponDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*WeaponDataViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *WeaponDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *WeaponDataViewer) Damage() (v int32) {
	if 4+4 > x.size() {
		return v
	}
	return *(*int32)(unsafe.Add(unsafe.Pointer(x), 4))
}
func (x *WeaponDataViewer) Range() (v int32) {
	if 8+4 > x.size() {
		return v
	}
	return *(*int32)(unsafe.Add(unsafe.Pointer(x), 8))
}

type WeaponViewer [4]byte

func NewWeaponViewer(reader *karmem.Reader, offset uint32) (v *WeaponViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*WeaponViewer)(unsafe.Pointer(&_Null))
	}
	v = (*WeaponViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *WeaponViewer) size() uint32 {
	return 4
}
func (x *WeaponViewer) Data(reader *karmem.Reader) (v *WeaponDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewWeaponDataViewer(reader, offset)
}

type MonsterDataViewer [111]byte

func NewMonsterDataViewer(reader *karmem.Reader, offset uint32) (v *MonsterDataViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*MonsterDataViewer)(unsafe.Pointer(&_Null))
	}
	v = (*MonsterDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*MonsterDataViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *MonsterDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *MonsterDataViewer) Pos() (v *Vec3Viewer) {
	if 4+12 > x.size() {
		return (*Vec3Viewer)(unsafe.Pointer(&_Null))
	}
	return (*Vec3Viewer)(unsafe.Add(unsafe.Pointer(x), 4))
}
func (x *MonsterDataViewer) Mana() (v int16) {
	if 16+2 > x.size() {
		return v
	}
	return *(*int16)(unsafe.Add(unsafe.Pointer(x), 16))
}
func (x *MonsterDataViewer) Health() (v int16) {
	if 18+2 > x.size() {
		return v
	}
	return *(*int16)(unsafe.Add(unsafe.Pointer(x), 18))
}
func (x *MonsterDataViewer) Name(reader *karmem.Reader) (v string) {
	if 20+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	if length > 512 {
		length = 512
	}
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *MonsterDataViewer) Team() (v Team) {
	if 28+1 > x.size() {
		return v
	}
	return *(*Team)(unsafe.Add(unsafe.Pointer(x), 28))
}
func (x *MonsterDataViewer) Inventory(reader *karmem.Reader) (v []byte) {
	if 29+8 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 29))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 29+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	if length > 128 {
		length = 128
	}
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *MonsterDataViewer) Color() (v Color) {
	if 37+1 > x.size() {
		return v
	}
	return *(*Color)(unsafe.Add(unsafe.Pointer(x), 37))
}
func (x *MonsterDataViewer) Hitbox() (v []float64) {
	if 38+40 > x.size() {
		return []float64{}
	}
	slice := [3]uintptr{
		uintptr(unsafe.Add(unsafe.Pointer(x), 38)), 5, 5,
	}
	return *(*[]float64)(unsafe.Pointer(&slice))
}
func (x *MonsterDataViewer) Status(reader *karmem.Reader) (v []int32) {
	if 78+8 > x.size() {
		return []int32{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 78))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 78+4))
	if !reader.IsValidOffset(offset, size) {
		return []int32{}
	}
	length := uintptr(size / 4)
	if length > 10 {
		length = 10
	}
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]int32)(unsafe.Pointer(&slice))
}
func (x *MonsterDataViewer) Weapons() (v []WeaponViewer) {
	if 86+16 > x.size() {
		return []WeaponViewer{}
	}
	slice := [3]uintptr{
		uintptr(unsafe.Add(unsafe.Pointer(x), 86)), 4, 4,
	}
	return *(*[]WeaponViewer)(unsafe.Pointer(&slice))
}
func (x *MonsterDataViewer) Path(reader *karmem.Reader) (v []Vec3Viewer) {
	if 102+8 > x.size() {
		return []Vec3Viewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 102))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 102+4))
	if !reader.IsValidOffset(offset, size) {
		return []Vec3Viewer{}
	}
	length := uintptr(size / 12)
	if length > 2000 {
		length = 2000
	}
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]Vec3Viewer)(unsafe.Pointer(&slice))
}
func (x *MonsterDataViewer) IsAlive() (v bool) {
	if 110+1 > x.size() {
		return v
	}
	return *(*bool)(unsafe.Add(unsafe.Pointer(x), 110))
}

type MonsterViewer [4]byte

func NewMonsterViewer(reader *karmem.Reader, offset uint32) (v *MonsterViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*MonsterViewer)(unsafe.Pointer(&_Null))
	}
	v = (*MonsterViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *MonsterViewer) size() uint32 {
	return 4
}
func (x *MonsterViewer) Data(reader *karmem.Reader) (v *MonsterDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewMonsterDataViewer(reader, offset)
}

type MonstersViewer [12]byte

func NewMonstersViewer(reader *karmem.Reader, offset uint32) (v *MonstersViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*MonstersViewer)(unsafe.Pointer(&_Null))
	}
	v = (*MonstersViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*MonstersViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *MonstersViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *MonstersViewer) Monsters(reader *karmem.Reader) (v []MonsterViewer) {
	if 4+8 > x.size() {
		return []MonsterViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []MonsterViewer{}
	}
	length := uintptr(size / 4)
	if length > 2000 {
		length = 2000
	}
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]MonsterViewer)(unsafe.Pointer(&slice))
}
