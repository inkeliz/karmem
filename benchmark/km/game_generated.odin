
package game

import "../../odin/karmem"
import "core:runtime"
import "core:mem"
import "core:reflect"

_Null := [152]byte{}
_NullReader := karmem.NewReaderArray(_Null[:])

EnumColor :: enum u8 {
    Red = 0,
    Green = 1,
    Blue = 2,
}

EnumTeam :: enum u8 {
    Humans = 0,
    Orcs = 1,
    Zombies = 2,
    Robots = 3,
    Aliens = 4,
}



EnumPacketIdentifier :: enum u64 {
    PacketIdentifierVec3 = 10268726485798425099,
    PacketIdentifierWeaponData = 15342010214468761012,
    PacketIdentifierWeapon = 8029074423243608167,
    PacketIdentifierMonsterData = 12254962724431809041,
    PacketIdentifierMonster = 5593793986513565154,
    PacketIdentifierMonsters = 14096677544474027661,
}
Vec3 :: struct #packed {
    X: f32,
    Y: f32,
    Z: f32,
}

NewVec3 :: #force_inline proc() -> Vec3 #no_bounds_check {
    return Vec3{}
}

Vec3Reset :: #force_inline proc(x: ^Vec3) #no_bounds_check {
    Vec3Read(x, (^Vec3Viewer)(&_Null), &_NullReader)
}

Vec3WriteAsRoot :: #force_inline proc(x: ^Vec3, writer: ^karmem.Writer) -> (uint, karmem.Error) #no_bounds_check {
    return Vec3Write(x, writer, 0)
}

Vec3Write :: proc(x: ^Vec3, writer: ^karmem.Writer, start: uint) -> (uint, karmem.Error) #no_bounds_check {
    offset := start
    size := u32(16)
    if offset == 0 {
        off, err := karmem.WriterAlloc(writer, size)
        if err != karmem.Error.ERR_NONE {
            return 0, err
        }
        offset = off
    }
    __XOffset := offset+0
    karmem.WriterWrite4At(writer, __XOffset, (cast(^u32)&x.X)^)
    __YOffset := offset+4
    karmem.WriterWrite4At(writer, __YOffset, (cast(^u32)&x.Y)^)
    __ZOffset := offset+8
    karmem.WriterWrite4At(writer, __ZOffset, (cast(^u32)&x.Z)^)

    return offset, nil
}

Vec3ReadAsRoot :: #force_inline proc(x: ^Vec3, reader: ^karmem.Reader) #no_bounds_check {
    Vec3Read(x, NewVec3Viewer(reader, 0), reader)
}

Vec3Read :: proc(x: ^Vec3, viewer: ^Vec3Viewer, reader: ^karmem.Reader) #no_bounds_check {
    x.X = Vec3ViewerX(viewer)
    x.Y = Vec3ViewerY(viewer)
    x.Z = Vec3ViewerZ(viewer)
}
WeaponData :: struct #packed {
    Damage: i32,
    Range: i32,
}

NewWeaponData :: #force_inline proc() -> WeaponData #no_bounds_check {
    return WeaponData{}
}

WeaponDataReset :: #force_inline proc(x: ^WeaponData) #no_bounds_check {
    WeaponDataRead(x, (^WeaponDataViewer)(&_Null), &_NullReader)
}

WeaponDataWriteAsRoot :: #force_inline proc(x: ^WeaponData, writer: ^karmem.Writer) -> (uint, karmem.Error) #no_bounds_check {
    return WeaponDataWrite(x, writer, 0)
}

WeaponDataWrite :: proc(x: ^WeaponData, writer: ^karmem.Writer, start: uint) -> (uint, karmem.Error) #no_bounds_check {
    offset := start
    size := u32(16)
    if offset == 0 {
        off, err := karmem.WriterAlloc(writer, size)
        if err != karmem.Error.ERR_NONE {
            return 0, err
        }
        offset = off
    }
    karmem.WriterWrite4At(writer, offset, u32(12))
    __DamageOffset := offset+4
    karmem.WriterWrite4At(writer, __DamageOffset, (cast(^u32)&x.Damage)^)
    __RangeOffset := offset+8
    karmem.WriterWrite4At(writer, __RangeOffset, (cast(^u32)&x.Range)^)

    return offset, nil
}

WeaponDataReadAsRoot :: #force_inline proc(x: ^WeaponData, reader: ^karmem.Reader) #no_bounds_check {
    WeaponDataRead(x, NewWeaponDataViewer(reader, 0), reader)
}

WeaponDataRead :: proc(x: ^WeaponData, viewer: ^WeaponDataViewer, reader: ^karmem.Reader) #no_bounds_check {
    x.Damage = WeaponDataViewerDamage(viewer)
    x.Range = WeaponDataViewerRange(viewer)
}
Weapon :: struct #packed {
    Data: WeaponData,
}

NewWeapon :: #force_inline proc() -> Weapon #no_bounds_check {
    return Weapon{}
}

WeaponReset :: #force_inline proc(x: ^Weapon) #no_bounds_check {
    WeaponRead(x, (^WeaponViewer)(&_Null), &_NullReader)
}

WeaponWriteAsRoot :: #force_inline proc(x: ^Weapon, writer: ^karmem.Writer) -> (uint, karmem.Error) #no_bounds_check {
    return WeaponWrite(x, writer, 0)
}

WeaponWrite :: proc(x: ^Weapon, writer: ^karmem.Writer, start: uint) -> (uint, karmem.Error) #no_bounds_check {
    offset := start
    size := u32(8)
    if offset == 0 {
        off, err := karmem.WriterAlloc(writer, size)
        if err != karmem.Error.ERR_NONE {
            return 0, err
        }
        offset = off
    }
    __DataSize := u32(16)
    __DataOffset, __DataErr := karmem.WriterAlloc(writer, __DataSize)
    if __DataErr != karmem.Error.ERR_NONE {
        return 0, __DataErr
    }
    karmem.WriterWrite4At(writer, offset+0, u32(__DataOffset))
    if _, err := WeaponDataWrite(&x.Data, writer, __DataOffset); err != nil {
        return offset, err
    }

    return offset, nil
}

WeaponReadAsRoot :: #force_inline proc(x: ^Weapon, reader: ^karmem.Reader) #no_bounds_check {
    WeaponRead(x, NewWeaponViewer(reader, 0), reader)
}

WeaponRead :: proc(x: ^Weapon, viewer: ^WeaponViewer, reader: ^karmem.Reader) #no_bounds_check {
    WeaponDataRead(&x.Data, WeaponViewerData(viewer,reader), reader)
}
MonsterData :: struct #packed {
    Pos: Vec3,
    Mana: i16,
    Health: i16,
    Name: string,
    Team: EnumTeam,
    Inventory: [dynamic]u8,
    Color: EnumColor,
    Hitbox: [5]f64,
    Status: [dynamic]i32,
    Weapons: [4]Weapon,
    Path: [dynamic]Vec3,
    IsAlive: bool,
}

NewMonsterData :: #force_inline proc() -> MonsterData #no_bounds_check {
    return MonsterData{}
}

MonsterDataReset :: #force_inline proc(x: ^MonsterData) #no_bounds_check {
    MonsterDataRead(x, (^MonsterDataViewer)(&_Null), &_NullReader)
}

MonsterDataWriteAsRoot :: #force_inline proc(x: ^MonsterData, writer: ^karmem.Writer) -> (uint, karmem.Error) #no_bounds_check {
    return MonsterDataWrite(x, writer, 0)
}

MonsterDataWrite :: proc(x: ^MonsterData, writer: ^karmem.Writer, start: uint) -> (uint, karmem.Error) #no_bounds_check {
    offset := start
    size := u32(152)
    if offset == 0 {
        off, err := karmem.WriterAlloc(writer, size)
        if err != karmem.Error.ERR_NONE {
            return 0, err
        }
        offset = off
    }
    karmem.WriterWrite4At(writer, offset, u32(147))
    __PosOffset := offset+4
    if _, err := Vec3Write(&x.Pos, writer, __PosOffset); err != nil {
        return offset, err
    }
    __ManaOffset := offset+20
    karmem.WriterWrite2At(writer, __ManaOffset, (cast(^u16)&x.Mana)^)
    __HealthOffset := offset+22
    karmem.WriterWrite2At(writer, __HealthOffset, (cast(^u16)&x.Health)^)
    __NameSize := u32(1 * len(x.Name))
    __NameOffset, __NameErr := karmem.WriterAlloc(writer, __NameSize)
    if __NameErr != karmem.Error.ERR_NONE {
        return 0, __NameErr
    }
    karmem.WriterWrite4At(writer, offset+24, u32(__NameOffset))
    karmem.WriterWrite4At(writer, offset+24 + 4, u32(__NameSize))
    karmem.WriterWrite4At(writer, offset+24 + 4 + 4, 1)
    if __NameSize > 0 {
        karmem.WriterWriteAt(writer, __NameOffset, rawptr((cast(^[^]u8)(&x.Name))^), __NameSize)
    }
    __TeamOffset := offset+36
    karmem.WriterWrite1At(writer, __TeamOffset, (cast(^u8)&x.Team)^)
    __InventorySize := u32(1 * len(x.Inventory))
    __InventoryOffset, __InventoryErr := karmem.WriterAlloc(writer, __InventorySize)
    if __InventoryErr != karmem.Error.ERR_NONE {
        return 0, __InventoryErr
    }
    karmem.WriterWrite4At(writer, offset+37, u32(__InventoryOffset))
    karmem.WriterWrite4At(writer, offset+37 + 4, u32(__InventorySize))
    karmem.WriterWrite4At(writer, offset+37 + 4 + 4, 1)
    if __InventorySize > 0 {
        karmem.WriterWriteAt(writer, __InventoryOffset, rawptr(&x.Inventory[0]), __InventorySize)
    }
    __ColorOffset := offset+49
    karmem.WriterWrite1At(writer, __ColorOffset, (cast(^u8)&x.Color)^)
    __HitboxOffset := offset+50
    __HitboxSize := u32(8 * len(x.Hitbox))
    karmem.WriterWriteAt(writer, __HitboxOffset, rawptr(&x.Hitbox), __HitboxSize)
    __StatusSize := u32(4 * len(x.Status))
    __StatusOffset, __StatusErr := karmem.WriterAlloc(writer, __StatusSize)
    if __StatusErr != karmem.Error.ERR_NONE {
        return 0, __StatusErr
    }
    karmem.WriterWrite4At(writer, offset+90, u32(__StatusOffset))
    karmem.WriterWrite4At(writer, offset+90 + 4, u32(__StatusSize))
    karmem.WriterWrite4At(writer, offset+90 + 4 + 4, 4)
    if __StatusSize > 0 {
        karmem.WriterWriteAt(writer, __StatusOffset, rawptr(&x.Status[0]), __StatusSize)
    }
    __WeaponsOffset := offset+102
    __WeaponsSize := u32(8 * len(x.Weapons))
    for _, i in x.Weapons {
        if _, err := WeaponWrite(&x.Weapons[i], writer, __WeaponsOffset); err != nil {
            return offset, err
        }
        __WeaponsOffset += 8
    }
    __PathSize := u32(16 * len(x.Path))
    __PathOffset, __PathErr := karmem.WriterAlloc(writer, __PathSize)
    if __PathErr != karmem.Error.ERR_NONE {
        return 0, __PathErr
    }
    karmem.WriterWrite4At(writer, offset+134, u32(__PathOffset))
    karmem.WriterWrite4At(writer, offset+134 + 4, u32(__PathSize))
    karmem.WriterWrite4At(writer, offset+134 + 4 + 4, 16)
    for _, i in x.Path {
        if _, err := Vec3Write(&x.Path[i], writer, __PathOffset); err != nil {
            return offset, err
        }
        __PathOffset += 16
    }
    __IsAliveOffset := offset+146
    karmem.WriterWrite1At(writer, __IsAliveOffset, (cast(^u8)&x.IsAlive)^)

    return offset, nil
}

MonsterDataReadAsRoot :: #force_inline proc(x: ^MonsterData, reader: ^karmem.Reader) #no_bounds_check {
    MonsterDataRead(x, NewMonsterDataViewer(reader, 0), reader)
}

MonsterDataRead :: proc(x: ^MonsterData, viewer: ^MonsterDataViewer, reader: ^karmem.Reader) #no_bounds_check {
    Vec3Read(&x.Pos, MonsterDataViewerPos(viewer,), reader)
    x.Mana = MonsterDataViewerMana(viewer)
    x.Health = MonsterDataViewerHealth(viewer)
    __NameString := MonsterDataViewerName(viewer, reader)
    if x.Name != __NameString {
        if len(__NameString) > 0 {
            if __NameString != x.Name {
                __NameStringCopy := make([]u8, len(__NameString))
                runtime.copy_from_string(__NameStringCopy[:], __NameString)
                delete(x.Name)
                x.Name = (cast(^string)(&__NameStringCopy))^
            }
        } else {
            x.Name = ""
        }
    }
    x.Team = EnumTeam(MonsterDataViewerTeam(viewer))
    __InventorySlice := MonsterDataViewerInventory(viewer, reader)
    __InventoryLen := len(__InventorySlice)
    __InventorySize := 1 * __InventoryLen
    if __InventoryLen > cap(x.Inventory) {
        __InventoryRealloc := make([dynamic]u8, __InventoryLen)
        if x.Inventory != nil {
            for _, i in  x.Inventory {
                __InventoryRealloc[i] = x.Inventory[i]
            }
            delete(x.Inventory)
        }
        x.Inventory = __InventoryRealloc
    }
    if x.Inventory != nil {
        (cast(^[3]int)(&x.Inventory))[1] = __InventoryLen
    }
    for i := 0; i < __InventoryLen; i += 1{
        x.Inventory[i] = __InventorySlice[i]
    }
    x.Color = EnumColor(MonsterDataViewerColor(viewer))
    __HitboxSlice := MonsterDataViewerHitbox(viewer)
    __HitboxLen := len(__HitboxSlice)
    if (__HitboxLen > 5) {
        __HitboxLen = 5
    }
    for i := 0; i < __HitboxLen; i += 1{
        x.Hitbox[i] = __HitboxSlice[i]
    }
    for i := __HitboxLen; i < len(x.Hitbox); i += 1 {
        x.Hitbox[i] = 0
    }
    __StatusSlice := MonsterDataViewerStatus(viewer, reader)
    __StatusLen := len(__StatusSlice)
    __StatusSize := 4 * __StatusLen
    if __StatusLen > cap(x.Status) {
        __StatusRealloc := make([dynamic]i32, __StatusLen)
        if x.Status != nil {
            for _, i in  x.Status {
                __StatusRealloc[i] = x.Status[i]
            }
            delete(x.Status)
        }
        x.Status = __StatusRealloc
    }
    if x.Status != nil {
        (cast(^[3]int)(&x.Status))[1] = __StatusLen
    }
    for i := 0; i < __StatusLen; i += 1{
        x.Status[i] = __StatusSlice[i]
    }
    __WeaponsSlice := MonsterDataViewerWeapons(viewer)
    __WeaponsLen := len(__WeaponsSlice)
    if (__WeaponsLen > 4) {
        __WeaponsLen = 4
    }
    for i := 0; i < __WeaponsLen; i += 1 {
        WeaponRead(&x.Weapons[i], &__WeaponsSlice[i], reader)
    }
    for i := __WeaponsLen; i < len(x.Weapons); i += 1 {
        WeaponReset(&x.Weapons[i])
    }
    __PathSlice := MonsterDataViewerPath(viewer, reader)
    __PathLen := len(__PathSlice)
    __PathSize := 16 * __PathLen
    if __PathLen > cap(x.Path) {
        __PathRealloc := make([dynamic]Vec3, __PathLen)
        if x.Path != nil {
            for _, i in  x.Path {
                __PathRealloc[i] = x.Path[i]
            }
            delete(x.Path)
        }
        x.Path = __PathRealloc
    }
    if x.Path != nil {
        (cast(^[3]int)(&x.Path))[1] = __PathLen
    }
    for i := 0; i < __PathLen; i += 1 {
        Vec3Read(&x.Path[i], &__PathSlice[i], reader)
    }
    x.IsAlive = MonsterDataViewerIsAlive(viewer)
}
Monster :: struct #packed {
    Data: MonsterData,
}

NewMonster :: #force_inline proc() -> Monster #no_bounds_check {
    return Monster{}
}

MonsterReset :: #force_inline proc(x: ^Monster) #no_bounds_check {
    MonsterRead(x, (^MonsterViewer)(&_Null), &_NullReader)
}

MonsterWriteAsRoot :: #force_inline proc(x: ^Monster, writer: ^karmem.Writer) -> (uint, karmem.Error) #no_bounds_check {
    return MonsterWrite(x, writer, 0)
}

MonsterWrite :: proc(x: ^Monster, writer: ^karmem.Writer, start: uint) -> (uint, karmem.Error) #no_bounds_check {
    offset := start
    size := u32(8)
    if offset == 0 {
        off, err := karmem.WriterAlloc(writer, size)
        if err != karmem.Error.ERR_NONE {
            return 0, err
        }
        offset = off
    }
    __DataSize := u32(152)
    __DataOffset, __DataErr := karmem.WriterAlloc(writer, __DataSize)
    if __DataErr != karmem.Error.ERR_NONE {
        return 0, __DataErr
    }
    karmem.WriterWrite4At(writer, offset+0, u32(__DataOffset))
    if _, err := MonsterDataWrite(&x.Data, writer, __DataOffset); err != nil {
        return offset, err
    }

    return offset, nil
}

MonsterReadAsRoot :: #force_inline proc(x: ^Monster, reader: ^karmem.Reader) #no_bounds_check {
    MonsterRead(x, NewMonsterViewer(reader, 0), reader)
}

MonsterRead :: proc(x: ^Monster, viewer: ^MonsterViewer, reader: ^karmem.Reader) #no_bounds_check {
    MonsterDataRead(&x.Data, MonsterViewerData(viewer,reader), reader)
}
Monsters :: struct #packed {
    Monsters: [dynamic]Monster,
}

NewMonsters :: #force_inline proc() -> Monsters #no_bounds_check {
    return Monsters{}
}

MonstersReset :: #force_inline proc(x: ^Monsters) #no_bounds_check {
    MonstersRead(x, (^MonstersViewer)(&_Null), &_NullReader)
}

MonstersWriteAsRoot :: #force_inline proc(x: ^Monsters, writer: ^karmem.Writer) -> (uint, karmem.Error) #no_bounds_check {
    return MonstersWrite(x, writer, 0)
}

MonstersWrite :: proc(x: ^Monsters, writer: ^karmem.Writer, start: uint) -> (uint, karmem.Error) #no_bounds_check {
    offset := start
    size := u32(24)
    if offset == 0 {
        off, err := karmem.WriterAlloc(writer, size)
        if err != karmem.Error.ERR_NONE {
            return 0, err
        }
        offset = off
    }
    karmem.WriterWrite4At(writer, offset, u32(16))
    __MonstersSize := u32(8 * len(x.Monsters))
    __MonstersOffset, __MonstersErr := karmem.WriterAlloc(writer, __MonstersSize)
    if __MonstersErr != karmem.Error.ERR_NONE {
        return 0, __MonstersErr
    }
    karmem.WriterWrite4At(writer, offset+4, u32(__MonstersOffset))
    karmem.WriterWrite4At(writer, offset+4 + 4, u32(__MonstersSize))
    karmem.WriterWrite4At(writer, offset+4 + 4 + 4, 8)
    for _, i in x.Monsters {
        if _, err := MonsterWrite(&x.Monsters[i], writer, __MonstersOffset); err != nil {
            return offset, err
        }
        __MonstersOffset += 8
    }

    return offset, nil
}

MonstersReadAsRoot :: #force_inline proc(x: ^Monsters, reader: ^karmem.Reader) #no_bounds_check {
    MonstersRead(x, NewMonstersViewer(reader, 0), reader)
}

MonstersRead :: proc(x: ^Monsters, viewer: ^MonstersViewer, reader: ^karmem.Reader) #no_bounds_check {
    __MonstersSlice := MonstersViewerMonsters(viewer, reader)
    __MonstersLen := len(__MonstersSlice)
    __MonstersSize := 8 * __MonstersLen
    if __MonstersLen > cap(x.Monsters) {
        __MonstersRealloc := make([dynamic]Monster, __MonstersLen)
        if x.Monsters != nil {
            for _, i in  x.Monsters {
                __MonstersRealloc[i] = x.Monsters[i]
            }
            delete(x.Monsters)
        }
        x.Monsters = __MonstersRealloc
    }
    if x.Monsters != nil {
        (cast(^[3]int)(&x.Monsters))[1] = __MonstersLen
    }
    for i := 0; i < __MonstersLen; i += 1 {
        MonsterRead(&x.Monsters[i], &__MonstersSlice[i], reader)
    }
}

Vec3Viewer :: struct #packed {
    _data: [16]byte
}

NewVec3Viewer :: #force_inline proc(reader: ^karmem.Reader, offset: u32) -> ^Vec3Viewer  #no_bounds_check {
    if karmem.ReaderIsValidOffset(reader, offset, 16) == false {
        return (^Vec3Viewer)(&_Null)
    }
    v := cast(^Vec3Viewer)(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return v
}

Vec3ViewerSize :: #force_inline proc(x: ^Vec3Viewer) -> u32 #no_bounds_check {
    return 16
}
Vec3ViewerX :: #force_inline proc(x: ^Vec3Viewer,) -> f32 #no_bounds_check {
    return ((^f32)(mem.ptr_offset(cast([^]u8)(x), 0)))^
}
Vec3ViewerY :: #force_inline proc(x: ^Vec3Viewer,) -> f32 #no_bounds_check {
    return ((^f32)(mem.ptr_offset(cast([^]u8)(x), 4)))^
}
Vec3ViewerZ :: #force_inline proc(x: ^Vec3Viewer,) -> f32 #no_bounds_check {
    return ((^f32)(mem.ptr_offset(cast([^]u8)(x), 8)))^
}
WeaponDataViewer :: struct #packed {
    _data: [16]byte
}

NewWeaponDataViewer :: #force_inline proc(reader: ^karmem.Reader, offset: u32) -> ^WeaponDataViewer  #no_bounds_check {
    if karmem.ReaderIsValidOffset(reader, offset, 8) == false {
        return (^WeaponDataViewer)(&_Null)
    }
    v := cast(^WeaponDataViewer)(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    if karmem.ReaderIsValidOffset(reader, offset, WeaponDataViewerSize(v)) == false {
        return (^WeaponDataViewer)(&_Null)
    }
    return v
}

WeaponDataViewerSize :: #force_inline proc(x: ^WeaponDataViewer) -> u32 #no_bounds_check {
    return ((^u32)(x))^
}
WeaponDataViewerDamage :: #force_inline proc(x: ^WeaponDataViewer,) -> i32 #no_bounds_check {
    if 4 + 4 > WeaponDataViewerSize(x) {
        return 0
    }
    return ((^i32)(mem.ptr_offset(cast([^]u8)(x), 4)))^
}
WeaponDataViewerRange :: #force_inline proc(x: ^WeaponDataViewer,) -> i32 #no_bounds_check {
    if 8 + 4 > WeaponDataViewerSize(x) {
        return 0
    }
    return ((^i32)(mem.ptr_offset(cast([^]u8)(x), 8)))^
}
WeaponViewer :: struct #packed {
    _data: [8]byte
}

NewWeaponViewer :: #force_inline proc(reader: ^karmem.Reader, offset: u32) -> ^WeaponViewer  #no_bounds_check {
    if karmem.ReaderIsValidOffset(reader, offset, 8) == false {
        return (^WeaponViewer)(&_Null)
    }
    v := cast(^WeaponViewer)(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return v
}

WeaponViewerSize :: #force_inline proc(x: ^WeaponViewer) -> u32 #no_bounds_check {
    return 8
}
WeaponViewerData :: #force_inline proc(x: ^WeaponViewer,reader: ^karmem.Reader) -> ^WeaponDataViewer #no_bounds_check {
    offset := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 0)))^
    return NewWeaponDataViewer(reader, offset)
}
MonsterDataViewer :: struct #packed {
    _data: [152]byte
}

NewMonsterDataViewer :: #force_inline proc(reader: ^karmem.Reader, offset: u32) -> ^MonsterDataViewer  #no_bounds_check {
    if karmem.ReaderIsValidOffset(reader, offset, 8) == false {
        return (^MonsterDataViewer)(&_Null)
    }
    v := cast(^MonsterDataViewer)(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    if karmem.ReaderIsValidOffset(reader, offset, MonsterDataViewerSize(v)) == false {
        return (^MonsterDataViewer)(&_Null)
    }
    return v
}

MonsterDataViewerSize :: #force_inline proc(x: ^MonsterDataViewer) -> u32 #no_bounds_check {
    return ((^u32)(x))^
}
MonsterDataViewerPos :: #force_inline proc(x: ^MonsterDataViewer,) -> ^Vec3Viewer #no_bounds_check {
    if 4 + 16 > MonsterDataViewerSize(x) {
        return (^Vec3Viewer)(&_Null)
    }
    return ((^Vec3Viewer)(mem.ptr_offset(cast([^]u8)(x), 4)))
}
MonsterDataViewerMana :: #force_inline proc(x: ^MonsterDataViewer,) -> i16 #no_bounds_check {
    if 20 + 2 > MonsterDataViewerSize(x) {
        return 0
    }
    return ((^i16)(mem.ptr_offset(cast([^]u8)(x), 20)))^
}
MonsterDataViewerHealth :: #force_inline proc(x: ^MonsterDataViewer,) -> i16 #no_bounds_check {
    if 22 + 2 > MonsterDataViewerSize(x) {
        return 0
    }
    return ((^i16)(mem.ptr_offset(cast([^]u8)(x), 22)))^
}
MonsterDataViewerName :: #force_inline proc(x: ^MonsterDataViewer,reader: ^karmem.Reader) -> string #no_bounds_check {
    if 24 + 12 > MonsterDataViewerSize(x) {
        return ""
    }
    offset := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 24)))^
    size := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 24 + 4)))^
    if karmem.ReaderIsValidOffset(reader, offset, size) == false {
        return ""
    }
    length := uint(size / 1)
    if length > 512 {
        length = 512
    }
    if length == 0 {
        return ""
    }
    slice := [2]uint{0, length}
    (cast(^rawptr)(&slice[0]))^ = rawptr(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return (cast(^string)(&slice))^
}
MonsterDataViewerTeam :: #force_inline proc(x: ^MonsterDataViewer,) -> EnumTeam #no_bounds_check {
    if 36 + 1 > MonsterDataViewerSize(x) {
        return EnumTeam(0)
    }
    return ((^EnumTeam)(mem.ptr_offset(cast([^]u8)(x), 36)))^
}
MonsterDataViewerInventory :: #force_inline proc(x: ^MonsterDataViewer,reader: ^karmem.Reader) -> []u8 #no_bounds_check {
    if 37 + 12 > MonsterDataViewerSize(x) {
        return []u8{}
    }
    offset := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 37)))^
    size := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 37 + 4)))^
    if karmem.ReaderIsValidOffset(reader, offset, size) == false {
        return []u8{}
    }
    length := uint(size / 1)
    if length > 128 {
        length = 128
    }
    if length == 0 {
        return []u8{}
    }
    slice := [2]uint{0, length}
    (cast(^rawptr)(&slice[0]))^ = rawptr(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return (cast(^[]u8)(&slice))^
}
MonsterDataViewerColor :: #force_inline proc(x: ^MonsterDataViewer,) -> EnumColor #no_bounds_check {
    if 49 + 1 > MonsterDataViewerSize(x) {
        return EnumColor(0)
    }
    return ((^EnumColor)(mem.ptr_offset(cast([^]u8)(x), 49)))^
}
MonsterDataViewerHitbox :: #force_inline proc(x: ^MonsterDataViewer,) -> []f64 #no_bounds_check {
    if 50 + 40 > MonsterDataViewerSize(x) {
        return []f64{}
    }
    slice := [2]int{ 0, 5}
    (cast(^rawptr)(&slice[0]))^ = rawptr(mem.ptr_offset(cast([^]u8)(x), 50))

    return ((^[]f64)(&slice))^
}
MonsterDataViewerStatus :: #force_inline proc(x: ^MonsterDataViewer,reader: ^karmem.Reader) -> []i32 #no_bounds_check {
    if 90 + 12 > MonsterDataViewerSize(x) {
        return []i32{}
    }
    offset := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 90)))^
    size := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 90 + 4)))^
    if karmem.ReaderIsValidOffset(reader, offset, size) == false {
        return []i32{}
    }
    length := uint(size / 4)
    if length > 10 {
        length = 10
    }
    if length == 0 {
        return []i32{}
    }
    slice := [2]uint{0, length}
    (cast(^rawptr)(&slice[0]))^ = rawptr(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return (cast(^[]i32)(&slice))^
}
MonsterDataViewerWeapons :: #force_inline proc(x: ^MonsterDataViewer,) -> []WeaponViewer #no_bounds_check {
    if 102 + 32 > MonsterDataViewerSize(x) {
        return []WeaponViewer{}
    }
    slice := [2]int{ 0, 4}
    (cast(^rawptr)(&slice[0]))^ = rawptr(mem.ptr_offset(cast([^]u8)(x), 102))

    return ((^[]WeaponViewer)(&slice))^
}
MonsterDataViewerPath :: #force_inline proc(x: ^MonsterDataViewer,reader: ^karmem.Reader) -> []Vec3Viewer #no_bounds_check {
    if 134 + 12 > MonsterDataViewerSize(x) {
        return []Vec3Viewer{}
    }
    offset := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 134)))^
    size := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 134 + 4)))^
    if karmem.ReaderIsValidOffset(reader, offset, size) == false {
        return []Vec3Viewer{}
    }
    length := uint(size / 16)
    if length > 2000 {
        length = 2000
    }
    if length == 0 {
        return []Vec3Viewer{}
    }
    slice := [2]uint{0, length}
    (cast(^rawptr)(&slice[0]))^ = rawptr(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return (cast(^[]Vec3Viewer)(&slice))^
}
MonsterDataViewerIsAlive :: #force_inline proc(x: ^MonsterDataViewer,) -> bool #no_bounds_check {
    if 146 + 1 > MonsterDataViewerSize(x) {
        return false
    }
    return ((^bool)(mem.ptr_offset(cast([^]u8)(x), 146)))^
}
MonsterViewer :: struct #packed {
    _data: [8]byte
}

NewMonsterViewer :: #force_inline proc(reader: ^karmem.Reader, offset: u32) -> ^MonsterViewer  #no_bounds_check {
    if karmem.ReaderIsValidOffset(reader, offset, 8) == false {
        return (^MonsterViewer)(&_Null)
    }
    v := cast(^MonsterViewer)(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return v
}

MonsterViewerSize :: #force_inline proc(x: ^MonsterViewer) -> u32 #no_bounds_check {
    return 8
}
MonsterViewerData :: #force_inline proc(x: ^MonsterViewer,reader: ^karmem.Reader) -> ^MonsterDataViewer #no_bounds_check {
    offset := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 0)))^
    return NewMonsterDataViewer(reader, offset)
}
MonstersViewer :: struct #packed {
    _data: [24]byte
}

NewMonstersViewer :: #force_inline proc(reader: ^karmem.Reader, offset: u32) -> ^MonstersViewer  #no_bounds_check {
    if karmem.ReaderIsValidOffset(reader, offset, 8) == false {
        return (^MonstersViewer)(&_Null)
    }
    v := cast(^MonstersViewer)(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    if karmem.ReaderIsValidOffset(reader, offset, MonstersViewerSize(v)) == false {
        return (^MonstersViewer)(&_Null)
    }
    return v
}

MonstersViewerSize :: #force_inline proc(x: ^MonstersViewer) -> u32 #no_bounds_check {
    return ((^u32)(x))^
}
MonstersViewerMonsters :: #force_inline proc(x: ^MonstersViewer,reader: ^karmem.Reader) -> []MonsterViewer #no_bounds_check {
    if 4 + 12 > MonstersViewerSize(x) {
        return []MonsterViewer{}
    }
    offset := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 4)))^
    size := ((^u32)(mem.ptr_offset(cast([^]u8)(x), 4 + 4)))^
    if karmem.ReaderIsValidOffset(reader, offset, size) == false {
        return []MonsterViewer{}
    }
    length := uint(size / 8)
    if length > 2000 {
        length = 2000
    }
    if length == 0 {
        return []MonsterViewer{}
    }
    slice := [2]uint{0, length}
    (cast(^rawptr)(&slice[0]))^ = rawptr(mem.ptr_offset(cast([^]u8)(reader.pointer), offset))
    return (cast(^[]MonsterViewer)(&slice))^
}
