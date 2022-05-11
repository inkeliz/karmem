
import karmem

public typealias EnumColor = UInt8
public let EnumColorRed : EnumColor = 0
public let EnumColorGreen : EnumColor = 1
public let EnumColorBlue : EnumColor = 2

public typealias EnumTeam = UInt8
public let EnumTeamHumans : EnumTeam = 0
public let EnumTeamOrcs : EnumTeam = 1
public let EnumTeamZombies : EnumTeam = 2
public let EnumTeamRobots : EnumTeam = 3
public let EnumTeamAliens : EnumTeam = 4

public struct Vec3 {
    public var X: Float = 0
    public var Y: Float = 0
    public var Z: Float = 0

    public init() {}

    public mutating func Reset() -> () {
        self.X = 0
        self.Y = 0
        self.Z = 0
    }

    @inline(__always)
    public mutating func WriteAsRoot(_ writer: inout karmem.Writer)  -> Bool {
        return self.Write(&writer, 0)
    }

    public mutating func Write(_ writer: inout karmem.Writer, _ start: UInt32) -> Bool {
        var offset = start
        let size: UInt32 = 16
        if (offset == 0) {
            offset = writer.Alloc(size)
            if (offset == 0xFFFFFFFF) {
                return false
            }
        }
        let __XOffset: UInt32 = offset + 0
        writer.WriteAt(__XOffset, self.X)
        let __YOffset: UInt32 = offset + 4
        writer.WriteAt(__YOffset, self.Y)
        let __ZOffset: UInt32 = offset + 8
        writer.WriteAt(__ZOffset, self.Z)

        return true
    }

    @inline(__always)
    public mutating func ReadAsRoot(_ reader: karmem.Reader) -> () {
        self.Read(NewVec3Viewer(reader, 0), reader)
    }

    @inline(__always)
    public mutating func Read(_ viewer: Vec3Viewer, _ reader: karmem.Reader) -> () {
    self.X = viewer.X()
    self.Y = viewer.Y()
    self.Z = viewer.Z()
    }

}

public func NewVec3() -> Vec3 {
    return Vec3()
}

public struct WeaponData {
    public var Damage: Int32 = 0
    public var Range: Int32 = 0

    public init() {}

    public mutating func Reset() -> () {
        self.Damage = 0
        self.Range = 0
    }

    @inline(__always)
    public mutating func WriteAsRoot(_ writer: inout karmem.Writer)  -> Bool {
        return self.Write(&writer, 0)
    }

    public mutating func Write(_ writer: inout karmem.Writer, _ start: UInt32) -> Bool {
        var offset = start
        let size: UInt32 = 16
        if (offset == 0) {
            offset = writer.Alloc(size)
            if (offset == 0xFFFFFFFF) {
                return false
            }
        }
        writer.WriteAt(offset, size)
        let __DamageOffset: UInt32 = offset + 4
        writer.WriteAt(__DamageOffset, self.Damage)
        let __RangeOffset: UInt32 = offset + 8
        writer.WriteAt(__RangeOffset, self.Range)

        return true
    }

    @inline(__always)
    public mutating func ReadAsRoot(_ reader: karmem.Reader) -> () {
        self.Read(NewWeaponDataViewer(reader, 0), reader)
    }

    @inline(__always)
    public mutating func Read(_ viewer: WeaponDataViewer, _ reader: karmem.Reader) -> () {
    self.Damage = viewer.Damage()
    self.Range = viewer.Range()
    }

}

public func NewWeaponData() -> WeaponData {
    return WeaponData()
}

public struct Weapon {
    public var Data: WeaponData = WeaponData()

    public init() {}

    public mutating func Reset() -> () {
        self.Data.Reset()
    }

    @inline(__always)
    public mutating func WriteAsRoot(_ writer: inout karmem.Writer)  -> Bool {
        return self.Write(&writer, 0)
    }

    public mutating func Write(_ writer: inout karmem.Writer, _ start: UInt32) -> Bool {
        var offset = start
        let size: UInt32 = 8
        if (offset == 0) {
            offset = writer.Alloc(size)
            if (offset == 0xFFFFFFFF) {
                return false
            }
        }
        let __DataSize: UInt32 = 16
        let __DataOffset = writer.Alloc(__DataSize)
        if (__DataOffset == 0) {
            return false
        }
        writer.WriteAt(offset + 0, __DataOffset)
        if (!self.Data.Write(&writer, __DataOffset)) {
            return false
        }

        return true
    }

    @inline(__always)
    public mutating func ReadAsRoot(_ reader: karmem.Reader) -> () {
        self.Read(NewWeaponViewer(reader, 0), reader)
    }

    @inline(__always)
    public mutating func Read(_ viewer: WeaponViewer, _ reader: karmem.Reader) -> () {
    self.Data.Read(viewer.Data(reader), reader)
    }

}

public func NewWeapon() -> Weapon {
    return Weapon()
}

public struct MonsterData {
    public var Pos: Vec3 = Vec3()
    public var Mana: Int16 = 0
    public var Health: Int16 = 0
    public var Name: [UInt8] = [UInt8]()
    public var Team: EnumTeam = 0
    public var Inventory: [UInt8] = [UInt8]()
    public var Color: EnumColor = 0
    public var Hitbox: [Double] = Array(repeating: Double(0), count: 5)
    public var Status: [Int32] = [Int32]()
    public var Weapons: [Weapon] = Array(repeating: Weapon(), count: 4)
    public var Path: [Vec3] = [Vec3]()

    public init() {}

    public mutating func Reset() -> () {
        self.Pos.Reset()
        self.Mana = 0
        self.Health = 0
        self.Name.removeAll()
        self.Team = 0
        self.Inventory.removeAll()
        self.Color = 0
        let __HitboxLen = self.Hitbox.count
        var __HitboxIndex = 0
        while (__HitboxIndex < __HitboxLen) {
                self.Hitbox[__HitboxIndex] = 0
            __HitboxIndex = __HitboxIndex + 1
        }
        self.Status.removeAll()
        let __WeaponsLen = self.Weapons.count
        var __WeaponsIndex = 0
        while (__WeaponsIndex < __WeaponsLen) {
                self.Weapons[__WeaponsIndex].Reset()
            __WeaponsIndex = __WeaponsIndex + 1
        }
        let __PathLen = self.Path.count
        var __PathIndex = 0
        while (__PathIndex < __PathLen) {
            self.Path[__PathIndex].Reset()
            __PathIndex = __PathIndex + 1
        }
        self.Path.removeAll()
    }

    @inline(__always)
    public mutating func WriteAsRoot(_ writer: inout karmem.Writer)  -> Bool {
        return self.Write(&writer, 0)
    }

    public mutating func Write(_ writer: inout karmem.Writer, _ start: UInt32) -> Bool {
        var offset = start
        let size: UInt32 = 152
        if (offset == 0) {
            offset = writer.Alloc(size)
            if (offset == 0xFFFFFFFF) {
                return false
            }
        }
        writer.WriteAt(offset, size)
        let __PosOffset: UInt32 = offset + 4
        if (!self.Pos.Write(&writer, __PosOffset)) {
            return false
        }
        let __ManaOffset: UInt32 = offset + 20
        writer.WriteAt(__ManaOffset, self.Mana)
        let __HealthOffset: UInt32 = offset + 22
        writer.WriteAt(__HealthOffset, self.Health)
        let __NameSize: UInt32 = 1 * UInt32(self.Name.count)
        let __NameOffset = writer.Alloc(__NameSize)
        if (__NameOffset == 0) {
            return false
        }
        writer.WriteAt(offset + 24, __NameOffset)
        writer.WriteAt(offset + 24 + 4, __NameSize)
        writer.WriteAt(offset + 24 + 4 + 4, 1)
        writer.WriteArrayAt(__NameOffset, self.Name, 1)
        let __TeamOffset: UInt32 = offset + 36
        writer.WriteAt(__TeamOffset, self.Team)
        let __InventorySize: UInt32 = 1 * UInt32(self.Inventory.count)
        let __InventoryOffset = writer.Alloc(__InventorySize)
        if (__InventoryOffset == 0) {
            return false
        }
        writer.WriteAt(offset + 37, __InventoryOffset)
        writer.WriteAt(offset + 37 + 4, __InventorySize)
        writer.WriteAt(offset + 37 + 4 + 4, 1)
        writer.WriteArrayAt(__InventoryOffset, self.Inventory, 1)
        let __ColorOffset: UInt32 = offset + 49
        writer.WriteAt(__ColorOffset, self.Color)
        let __HitboxOffset: UInt32 = offset + 50
        writer.WriteArrayAt(__HitboxOffset, self.Hitbox, 8)
        let __StatusSize: UInt32 = 4 * UInt32(self.Status.count)
        let __StatusOffset = writer.Alloc(__StatusSize)
        if (__StatusOffset == 0) {
            return false
        }
        writer.WriteAt(offset + 90, __StatusOffset)
        writer.WriteAt(offset + 90 + 4, __StatusSize)
        writer.WriteAt(offset + 90 + 4 + 4, 4)
        writer.WriteArrayAt(__StatusOffset, self.Status, 4)
        let __WeaponsOffset: UInt32 = offset + 102
        let __WeaponsLen = self.Weapons.count
        var __WeaponsIndex = 0
        while (__WeaponsIndex < __WeaponsLen) {
            if (!self.Weapons[__WeaponsIndex].Write(&writer, __WeaponsOffset + (UInt32(__WeaponsIndex) * 8))) {
                return false
            }
            __WeaponsIndex = __WeaponsIndex + 1
        }
        let __PathSize: UInt32 = 16 * UInt32(self.Path.count)
        let __PathOffset = writer.Alloc(__PathSize)
        if (__PathOffset == 0) {
            return false
        }
        writer.WriteAt(offset + 134, __PathOffset)
        writer.WriteAt(offset + 134 + 4, __PathSize)
        writer.WriteAt(offset + 134 + 4 + 4, 16)
        let __PathLen = self.Path.count
        var __PathIndex = 0
        while (__PathIndex < __PathLen) {
            if (!self.Path[__PathIndex].Write(&writer, __PathOffset + (UInt32(__PathIndex) * 16))) {
                return false
            }
            __PathIndex = __PathIndex + 1
        }

        return true
    }

    @inline(__always)
    public mutating func ReadAsRoot(_ reader: karmem.Reader) -> () {
        self.Read(NewMonsterDataViewer(reader, 0), reader)
    }

    @inline(__always)
    public mutating func Read(_ viewer: MonsterDataViewer, _ reader: karmem.Reader) -> () {
    self.Pos.Read(viewer.Pos(), reader)
    self.Mana = viewer.Mana()
    self.Health = viewer.Health()
    var __NameSlice = viewer.Name(reader)
    let __NameLen = __NameSlice.count
    self.Name.removeAll()
    if (__NameLen > self.Name.count) {
        self.Name.reserveCapacity(__NameLen)
        var __NameIndexClear = self.Name.count
        while(__NameIndexClear < __NameLen) {
            self.Name.append(0)
            __NameIndexClear = __NameIndexClear + 1
        }
    }
    var __NameIndex = 0
    while (__NameIndex < self.Name.count) {
        if (__NameIndex >= __NameLen) {
            self.Name[__NameIndex] = UInt8(0)
        } else {
            self.Name[__NameIndex] = __NameSlice[__NameIndex]
        }
        __NameIndex = __NameIndex + 1
    }
    self.Team = viewer.Team()
    var __InventorySlice = viewer.Inventory(reader)
    let __InventoryLen = __InventorySlice.count
    self.Inventory.removeAll()
    if (__InventoryLen > self.Inventory.count) {
        self.Inventory.reserveCapacity(__InventoryLen)
        var __InventoryIndexClear = self.Inventory.count
        while(__InventoryIndexClear < __InventoryLen) {
            self.Inventory.append(0)
            __InventoryIndexClear = __InventoryIndexClear + 1
        }
    }
    var __InventoryIndex = 0
    while (__InventoryIndex < self.Inventory.count) {
        if (__InventoryIndex >= __InventoryLen) {
            self.Inventory[__InventoryIndex] = 0
        } else {
            self.Inventory[__InventoryIndex] = __InventorySlice[__InventoryIndex]
        }
        __InventoryIndex = __InventoryIndex + 1
    }
    self.Color = viewer.Color()
    var __HitboxSlice = viewer.Hitbox()
    let __HitboxLen = __HitboxSlice.count
    var __HitboxIndex = 0
    while (__HitboxIndex < self.Hitbox.count) {
        if (__HitboxIndex >= __HitboxLen) {
            self.Hitbox[__HitboxIndex] = 0
        } else {
            self.Hitbox[__HitboxIndex] = __HitboxSlice[__HitboxIndex]
        }
        __HitboxIndex = __HitboxIndex + 1
    }
    var __StatusSlice = viewer.Status(reader)
    let __StatusLen = __StatusSlice.count
    self.Status.removeAll()
    if (__StatusLen > self.Status.count) {
        self.Status.reserveCapacity(__StatusLen)
        var __StatusIndexClear = self.Status.count
        while(__StatusIndexClear < __StatusLen) {
            self.Status.append(0)
            __StatusIndexClear = __StatusIndexClear + 1
        }
    }
    var __StatusIndex = 0
    while (__StatusIndex < self.Status.count) {
        if (__StatusIndex >= __StatusLen) {
            self.Status[__StatusIndex] = 0
        } else {
            self.Status[__StatusIndex] = __StatusSlice[__StatusIndex]
        }
        __StatusIndex = __StatusIndex + 1
    }
    var __WeaponsSlice = viewer.Weapons()
    let __WeaponsLen = __WeaponsSlice.count
    var __WeaponsIndex = 0
    while (__WeaponsIndex < self.Weapons.count) {
        if (__WeaponsIndex >= __WeaponsLen) {
            self.Weapons[__WeaponsIndex].Reset()
        } else {
            self.Weapons[__WeaponsIndex].Read(__WeaponsSlice[__WeaponsIndex], reader)
        }
        __WeaponsIndex = __WeaponsIndex + 1
    }
    var __PathSlice = viewer.Path(reader)
    let __PathLen = __PathSlice.count
    self.Path.removeAll()
    if (__PathLen > self.Path.count) {
        self.Path.reserveCapacity(__PathLen)
        var __PathIndexClear = self.Path.count
        while(__PathIndexClear < __PathLen) {
            self.Path.append(NewVec3())
            __PathIndexClear = __PathIndexClear + 1
        }
    }
    var __PathIndex = 0
    while (__PathIndex < self.Path.count) {
        if (__PathIndex >= __PathLen) {
            self.Path[__PathIndex].Reset()
        } else {
            self.Path[__PathIndex].Read(__PathSlice[__PathIndex], reader)
        }
        __PathIndex = __PathIndex + 1
    }
    }

}

public func NewMonsterData() -> MonsterData {
    return MonsterData()
}

public struct Monster {
    public var Data: MonsterData = MonsterData()

    public init() {}

    public mutating func Reset() -> () {
        self.Data.Reset()
    }

    @inline(__always)
    public mutating func WriteAsRoot(_ writer: inout karmem.Writer)  -> Bool {
        return self.Write(&writer, 0)
    }

    public mutating func Write(_ writer: inout karmem.Writer, _ start: UInt32) -> Bool {
        var offset = start
        let size: UInt32 = 8
        if (offset == 0) {
            offset = writer.Alloc(size)
            if (offset == 0xFFFFFFFF) {
                return false
            }
        }
        let __DataSize: UInt32 = 152
        let __DataOffset = writer.Alloc(__DataSize)
        if (__DataOffset == 0) {
            return false
        }
        writer.WriteAt(offset + 0, __DataOffset)
        if (!self.Data.Write(&writer, __DataOffset)) {
            return false
        }

        return true
    }

    @inline(__always)
    public mutating func ReadAsRoot(_ reader: karmem.Reader) -> () {
        self.Read(NewMonsterViewer(reader, 0), reader)
    }

    @inline(__always)
    public mutating func Read(_ viewer: MonsterViewer, _ reader: karmem.Reader) -> () {
    self.Data.Read(viewer.Data(reader), reader)
    }

}

public func NewMonster() -> Monster {
    return Monster()
}

public struct Monsters {
    public var Monsters: [Monster] = [Monster]()

    public init() {}

    public mutating func Reset() -> () {
        let __MonstersLen = self.Monsters.count
        var __MonstersIndex = 0
        while (__MonstersIndex < __MonstersLen) {
            self.Monsters[__MonstersIndex].Reset()
            __MonstersIndex = __MonstersIndex + 1
        }
        self.Monsters.removeAll()
    }

    @inline(__always)
    public mutating func WriteAsRoot(_ writer: inout karmem.Writer)  -> Bool {
        return self.Write(&writer, 0)
    }

    public mutating func Write(_ writer: inout karmem.Writer, _ start: UInt32) -> Bool {
        var offset = start
        let size: UInt32 = 24
        if (offset == 0) {
            offset = writer.Alloc(size)
            if (offset == 0xFFFFFFFF) {
                return false
            }
        }
        writer.WriteAt(offset, size)
        let __MonstersSize: UInt32 = 8 * UInt32(self.Monsters.count)
        let __MonstersOffset = writer.Alloc(__MonstersSize)
        if (__MonstersOffset == 0) {
            return false
        }
        writer.WriteAt(offset + 4, __MonstersOffset)
        writer.WriteAt(offset + 4 + 4, __MonstersSize)
        writer.WriteAt(offset + 4 + 4 + 4, 8)
        let __MonstersLen = self.Monsters.count
        var __MonstersIndex = 0
        while (__MonstersIndex < __MonstersLen) {
            if (!self.Monsters[__MonstersIndex].Write(&writer, __MonstersOffset + (UInt32(__MonstersIndex) * 8))) {
                return false
            }
            __MonstersIndex = __MonstersIndex + 1
        }

        return true
    }

    @inline(__always)
    public mutating func ReadAsRoot(_ reader: karmem.Reader) -> () {
        self.Read(NewMonstersViewer(reader, 0), reader)
    }

    @inline(__always)
    public mutating func Read(_ viewer: MonstersViewer, _ reader: karmem.Reader) -> () {
    var __MonstersSlice = viewer.Monsters(reader)
    let __MonstersLen = __MonstersSlice.count
    self.Monsters.removeAll()
    if (__MonstersLen > self.Monsters.count) {
        self.Monsters.reserveCapacity(__MonstersLen)
        var __MonstersIndexClear = self.Monsters.count
        while(__MonstersIndexClear < __MonstersLen) {
            self.Monsters.append(NewMonster())
            __MonstersIndexClear = __MonstersIndexClear + 1
        }
    }
    var __MonstersIndex = 0
    while (__MonstersIndex < self.Monsters.count) {
        if (__MonstersIndex >= __MonstersLen) {
            self.Monsters[__MonstersIndex].Reset()
        } else {
            self.Monsters[__MonstersIndex].Read(__MonstersSlice[__MonstersIndex], reader)
        }
        __MonstersIndex = __MonstersIndex + 1
    }
    }

}

public func NewMonsters() -> Monsters {
    return Monsters()
}

public struct Vec3Viewer {
    var _0: UInt64 = 0
    var _1: UInt64 = 0

    public init() {}

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return 16
    }
    @inline(__always)
    public func X() -> Float {
        var v : Float = Float(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 0, &v)
        }
        return v
    }
    @inline(__always)
    public func Y() -> Float {
        var v : Float = Float(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 4, &v)
        }
        return v
    }
    @inline(__always)
    public func Z() -> Float {
        var v : Float = Float(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 8, &v)
        }
        return v
    }
}

@inline(__always) public func NewVec3Viewer(_ reader: karmem.Reader, _ offset: UInt32) -> Vec3Viewer {
    if (!reader.IsValidOffset(offset, 16)) {
        return Vec3Viewer()
    }

    var v = Vec3Viewer()
    v = karmem.Load(reader.pointer, Int(offset), &v)
    return v
}
public struct WeaponDataViewer {
    var _0: UInt64 = 0
    var _1: UInt64 = 0

    public init() {}

    @inline(__always)
    public func SizeOf() -> UInt32 {
        var size = UInt32(0)
        return withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 0, &size)
        }
    }
    @inline(__always)
    public func Damage() -> Int32 {
        if ((UInt32(4) + UInt32(4)) > self.SizeOf()) {
            return 0
        }
        var v : Int32 = Int32(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 4, &v)
        }
        return v
    }
    @inline(__always)
    public func Range() -> Int32 {
        if ((UInt32(8) + UInt32(4)) > self.SizeOf()) {
            return 0
        }
        var v : Int32 = Int32(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 8, &v)
        }
        return v
    }
}

@inline(__always) public func NewWeaponDataViewer(_ reader: karmem.Reader, _ offset: UInt32) -> WeaponDataViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return WeaponDataViewer()
    }

    var v = WeaponDataViewer()
    v = karmem.Load(reader.pointer, Int(offset), &v)
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return WeaponDataViewer()
    }
    return v
}
public struct WeaponViewer {
    var _0: UInt64 = 0

    public init() {}

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return 8
    }
    @inline(__always)
    public func Data(_ reader: karmem.Reader) -> WeaponDataViewer {
        var offset = UInt32(0)
        offset = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 0, &offset)
        }
        return NewWeaponDataViewer(reader, offset)
    }
}

@inline(__always) public func NewWeaponViewer(_ reader: karmem.Reader, _ offset: UInt32) -> WeaponViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return WeaponViewer()
    }

    var v = WeaponViewer()
    v = karmem.Load(reader.pointer, Int(offset), &v)
    return v
}
public struct MonsterDataViewer {
    var _0: UInt64 = 0
    var _1: UInt64 = 0
    var _2: UInt64 = 0
    var _3: UInt64 = 0
    var _4: UInt64 = 0
    var _5: UInt64 = 0
    var _6: UInt64 = 0
    var _7: UInt64 = 0
    var _8: UInt64 = 0
    var _9: UInt64 = 0
    var _10: UInt64 = 0
    var _11: UInt64 = 0
    var _12: UInt64 = 0
    var _13: UInt64 = 0
    var _14: UInt64 = 0
    var _15: UInt64 = 0
    var _16: UInt64 = 0
    var _17: UInt64 = 0
    var _18: UInt64 = 0

    public init() {}

    @inline(__always)
    public func SizeOf() -> UInt32 {
        var size = UInt32(0)
        return withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 0, &size)
        }
    }
    @inline(__always)
    public func Pos() -> Vec3Viewer {
        if ((UInt32(4) + UInt32(16)) > self.SizeOf()) {
            return Vec3Viewer()
        }
        var v : Vec3Viewer = Vec3Viewer()
        return withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 4, &v)
        }
    }
    @inline(__always)
    public func Mana() -> Int16 {
        if ((UInt32(20) + UInt32(2)) > self.SizeOf()) {
            return 0
        }
        var v : Int16 = Int16(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 20, &v)
        }
        return v
    }
    @inline(__always)
    public func Health() -> Int16 {
        if ((UInt32(22) + UInt32(2)) > self.SizeOf()) {
            return 0
        }
        var v : Int16 = Int16(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 22, &v)
        }
        return v
    }
    @inline(__always)
    public func Name(_ reader: karmem.Reader) -> karmem.Slice<UInt8> {
        if ((UInt32(24) + UInt32(12)) > self.SizeOf()) {
            return withUnsafePointer(to: self) {
                return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, UInt8(0))
            }
        }
        var offset = UInt32(0)
        offset = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 24, &offset)
        }
        var size = UInt32(0)
        size = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 24 + 4, &size)
        }
        if (!reader.IsValidOffset(offset, size)) {
            return withUnsafePointer(to: self) {
               return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, UInt8(0))
            }
        }
        return karmem.NewSlice(UnsafeRawPointer(reader.pointer + Int(offset)), size / 1, 1, UInt8(0))
    }
    @inline(__always)
    public func Team() -> EnumTeam {
        if ((UInt32(36) + UInt32(1)) > self.SizeOf()) {
            return 0
        }
        var v : EnumTeam = EnumTeam(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 36, &v)
        }
        return EnumTeam(v)
    }
    @inline(__always)
    public func Inventory(_ reader: karmem.Reader) -> karmem.Slice<UInt8> {
        if ((UInt32(37) + UInt32(12)) > self.SizeOf()) {
            return withUnsafePointer(to: self) {
                return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, 0)
            }
        }
        var offset = UInt32(0)
        offset = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 37, &offset)
        }
        var size = UInt32(0)
        size = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 37 + 4, &size)
        }
        if (!reader.IsValidOffset(offset, size)) {
            return withUnsafePointer(to: self) {
               return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, 0)
            }
        }
        return karmem.NewSlice(UnsafeRawPointer(reader.pointer + Int(offset)), size / 1, 1, 0)
    }
    @inline(__always)
    public func Color() -> EnumColor {
        if ((UInt32(49) + UInt32(1)) > self.SizeOf()) {
            return 0
        }
        var v : EnumColor = EnumColor(0)
        v = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 49, &v)
        }
        return EnumColor(v)
    }
    @inline(__always)
    public func Hitbox() -> karmem.Slice<Double> {
        if ((UInt32(50) + UInt32(40)) > self.SizeOf()) {
            return withUnsafePointer(to: self) {
                return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, 0)
            }
        }
        return withUnsafePointer(to: self) {
            return karmem.NewSlice(UnsafeRawPointer(UnsafeRawPointer($0) + Int(50)), 5, 8, 0)
        }
    }
    @inline(__always)
    public func Status(_ reader: karmem.Reader) -> karmem.Slice<Int32> {
        if ((UInt32(90) + UInt32(12)) > self.SizeOf()) {
            return withUnsafePointer(to: self) {
                return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, 0)
            }
        }
        var offset = UInt32(0)
        offset = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 90, &offset)
        }
        var size = UInt32(0)
        size = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 90 + 4, &size)
        }
        if (!reader.IsValidOffset(offset, size)) {
            return withUnsafePointer(to: self) {
               return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, 0)
            }
        }
        return karmem.NewSlice(UnsafeRawPointer(reader.pointer + Int(offset)), size / 4, 4, 0)
    }
    @inline(__always)
    public func Weapons() -> karmem.Slice<WeaponViewer> {
        if ((UInt32(102) + UInt32(32)) > self.SizeOf()) {
            return withUnsafePointer(to: self) {
                return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, WeaponViewer())
            }
        }
        return withUnsafePointer(to: self) {
            return karmem.NewSlice(UnsafeRawPointer(UnsafeRawPointer($0) + Int(102)), 4, 8, WeaponViewer())
        }
    }
    @inline(__always)
    public func Path(_ reader: karmem.Reader) -> karmem.Slice<Vec3Viewer> {
        if ((UInt32(134) + UInt32(12)) > self.SizeOf()) {
            return withUnsafePointer(to: self) {
                return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, Vec3Viewer())
            }
        }
        var offset = UInt32(0)
        offset = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 134, &offset)
        }
        var size = UInt32(0)
        size = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 134 + 4, &size)
        }
        if (!reader.IsValidOffset(offset, size)) {
            return withUnsafePointer(to: self) {
               return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, Vec3Viewer())
            }
        }
        return karmem.NewSlice(UnsafeRawPointer(reader.pointer + Int(offset)), size / 16, 16, Vec3Viewer())
    }
}

@inline(__always) public func NewMonsterDataViewer(_ reader: karmem.Reader, _ offset: UInt32) -> MonsterDataViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return MonsterDataViewer()
    }

    var v = MonsterDataViewer()
    v = karmem.Load(reader.pointer, Int(offset), &v)
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return MonsterDataViewer()
    }
    return v
}
public struct MonsterViewer {
    var _0: UInt64 = 0

    public init() {}

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return 8
    }
    @inline(__always)
    public func Data(_ reader: karmem.Reader) -> MonsterDataViewer {
        var offset = UInt32(0)
        offset = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 0, &offset)
        }
        return NewMonsterDataViewer(reader, offset)
    }
}

@inline(__always) public func NewMonsterViewer(_ reader: karmem.Reader, _ offset: UInt32) -> MonsterViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return MonsterViewer()
    }

    var v = MonsterViewer()
    v = karmem.Load(reader.pointer, Int(offset), &v)
    return v
}
public struct MonstersViewer {
    var _0: UInt64 = 0
    var _1: UInt64 = 0
    var _2: UInt64 = 0

    public init() {}

    @inline(__always)
    public func SizeOf() -> UInt32 {
        var size = UInt32(0)
        return withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 0, &size)
        }
    }
    @inline(__always)
    public func Monsters(_ reader: karmem.Reader) -> karmem.Slice<MonsterViewer> {
        if ((UInt32(4) + UInt32(12)) > self.SizeOf()) {
            return withUnsafePointer(to: self) {
                return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, MonsterViewer())
            }
        }
        var offset = UInt32(0)
        offset = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 4, &offset)
        }
        var size = UInt32(0)
        size = withUnsafePointer(to: self) {
            return karmem.Load(UnsafeRawPointer($0), 4 + 4, &size)
        }
        if (!reader.IsValidOffset(offset, size)) {
            return withUnsafePointer(to: self) {
               return karmem.NewSlice(UnsafeRawPointer($0), 0, 0, MonsterViewer())
            }
        }
        return karmem.NewSlice(UnsafeRawPointer(reader.pointer + Int(offset)), size / 8, 8, MonsterViewer())
    }
}

@inline(__always) public func NewMonstersViewer(_ reader: karmem.Reader, _ offset: UInt32) -> MonstersViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return MonstersViewer()
    }

    var v = MonstersViewer()
    v = karmem.Load(reader.pointer, Int(offset), &v)
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return MonstersViewer()
    }
    return v
}
