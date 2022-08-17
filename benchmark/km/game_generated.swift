
import karmem

var _Null : [UInt8] = Array(repeating: 0, count: 152)
var _NullReader = karmem.NewReader(_Null)

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

public typealias EnumPacketIdentifier = UInt64
public let EnumPacketIdentifierVec3 : EnumPacketIdentifier = 10268726485798425099
public let EnumPacketIdentifierWeaponData : EnumPacketIdentifier = 15342010214468761012
public let EnumPacketIdentifierWeapon : EnumPacketIdentifier = 8029074423243608167
public let EnumPacketIdentifierMonsterData : EnumPacketIdentifier = 12254962724431809041
public let EnumPacketIdentifierMonster : EnumPacketIdentifier = 5593793986513565154
public let EnumPacketIdentifierMonsters : EnumPacketIdentifier = 14096677544474027661



public struct Vec3 {
    public var X: Float = 0
    public var Y: Float = 0
    public var Z: Float = 0

    public init() {}

    public func PacketIdentifier() -> EnumPacketIdentifier {
        return EnumPacketIdentifierVec3
    }

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
        writer.memory.storeBytes(of: self.X, toByteOffset: Int(__XOffset), as:Float.self)
        let __YOffset: UInt32 = offset + 4
        writer.memory.storeBytes(of: self.Y, toByteOffset: Int(__YOffset), as:Float.self)
        let __ZOffset: UInt32 = offset + 8
        writer.memory.storeBytes(of: self.Z, toByteOffset: Int(__ZOffset), as:Float.self)

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

    public func PacketIdentifier() -> EnumPacketIdentifier {
        return EnumPacketIdentifierWeaponData
    }

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
        writer.memory.storeBytes(of: UInt32(12), toByteOffset: Int(offset), as: UInt32.self)
        let __DamageOffset: UInt32 = offset + 4
        writer.memory.storeBytes(of: self.Damage, toByteOffset: Int(__DamageOffset), as:Int32.self)
        let __RangeOffset: UInt32 = offset + 8
        writer.memory.storeBytes(of: self.Range, toByteOffset: Int(__RangeOffset), as:Int32.self)

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

    public func PacketIdentifier() -> EnumPacketIdentifier {
        return EnumPacketIdentifierWeapon
    }

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
        writer.memory.storeBytes(of: UInt32(__DataOffset), toByteOffset: Int(offset + 0), as: UInt32.self)
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
    public var IsAlive: Bool = false

    public init() {}

    public func PacketIdentifier() -> EnumPacketIdentifier {
        return EnumPacketIdentifierMonsterData
    }

    public mutating func Reset() -> () {
        self.Pos.Reset()
        self.Mana = 0
        self.Health = 0
        self.Name.removeAll(keepingCapacity: true)
        self.Team = 0
        self.Inventory.removeAll(keepingCapacity: true)
        self.Color = 0
        self.Hitbox.removeAll(keepingCapacity: true)
        self.Status.removeAll(keepingCapacity: true)
        self.Weapons.removeAll(keepingCapacity: true)
        self.Path.removeAll(keepingCapacity: true)
        self.IsAlive = false
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
        writer.memory.storeBytes(of: UInt32(147), toByteOffset: Int(offset), as: UInt32.self)
        let __PosOffset: UInt32 = offset + 4
        if (!self.Pos.Write(&writer, __PosOffset)) {
            return false
        }
        let __ManaOffset: UInt32 = offset + 20
        writer.memory.storeBytes(of: self.Mana, toByteOffset: Int(__ManaOffset), as:Int16.self)
        let __HealthOffset: UInt32 = offset + 22
        writer.memory.storeBytes(of: self.Health, toByteOffset: Int(__HealthOffset), as:Int16.self)
        let __NameSize: UInt32 = 1 * UInt32(self.Name.count)
        let __NameOffset = writer.Alloc(__NameSize)
        if (__NameOffset == 0) {
            return false
        }
        writer.memory.storeBytes(of: UInt32(__NameOffset), toByteOffset: Int(offset + 24), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(__NameSize), toByteOffset: Int(offset + 24 + 4), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(1), toByteOffset: Int(offset + 24 + 4 + 4), as: UInt32.self)
        var __NameIndex = 0
        var __NameCurrentOffset = __NameOffset
        while(__NameIndex < self.Name.count) {
            writer.memory.storeBytes(of: self.Name[__NameIndex], toByteOffset: Int(__NameCurrentOffset), as: UInt8.self )
            __NameIndex = __NameIndex + 1
            __NameCurrentOffset = __NameCurrentOffset + 1
        }
        let __TeamOffset: UInt32 = offset + 36
        writer.memory.storeBytes(of: self.Team, toByteOffset: Int(__TeamOffset), as:EnumTeam.self)
        let __InventorySize: UInt32 = 1 * UInt32(self.Inventory.count)
        let __InventoryOffset = writer.Alloc(__InventorySize)
        if (__InventoryOffset == 0) {
            return false
        }
        writer.memory.storeBytes(of: UInt32(__InventoryOffset), toByteOffset: Int(offset + 37), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(__InventorySize), toByteOffset: Int(offset + 37 + 4), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(1), toByteOffset: Int(offset + 37 + 4 + 4), as: UInt32.self)
        var __InventoryIndex = 0
        var __InventoryCurrentOffset = __InventoryOffset
        while(__InventoryIndex < self.Inventory.count) {
            writer.memory.storeBytes(of: self.Inventory[__InventoryIndex], toByteOffset: Int(__InventoryCurrentOffset), as:UInt8.self)
            __InventoryIndex = __InventoryIndex + 1
            __InventoryCurrentOffset = __InventoryCurrentOffset + 1
        }
        let __ColorOffset: UInt32 = offset + 49
        writer.memory.storeBytes(of: self.Color, toByteOffset: Int(__ColorOffset), as:EnumColor.self)
        let __HitboxOffset: UInt32 = offset + 50
        var __HitboxIndex = 0
        var __HitboxCurrentOffset = __HitboxOffset
        while(__HitboxIndex < self.Hitbox.count) {
            writer.memory.storeBytes(of: self.Hitbox[__HitboxIndex], toByteOffset: Int(__HitboxCurrentOffset), as:Double.self)
            __HitboxIndex = __HitboxIndex + 1
            __HitboxCurrentOffset = __HitboxCurrentOffset + 8
        }
        let __StatusSize: UInt32 = 4 * UInt32(self.Status.count)
        let __StatusOffset = writer.Alloc(__StatusSize)
        if (__StatusOffset == 0) {
            return false
        }
        writer.memory.storeBytes(of: UInt32(__StatusOffset), toByteOffset: Int(offset + 90), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(__StatusSize), toByteOffset: Int(offset + 90 + 4), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(4), toByteOffset: Int(offset + 90 + 4 + 4), as: UInt32.self)
        var __StatusIndex = 0
        var __StatusCurrentOffset = __StatusOffset
        while(__StatusIndex < self.Status.count) {
            writer.memory.storeBytes(of: self.Status[__StatusIndex], toByteOffset: Int(__StatusCurrentOffset), as:Int32.self)
            __StatusIndex = __StatusIndex + 1
            __StatusCurrentOffset = __StatusCurrentOffset + 4
        }
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
        writer.memory.storeBytes(of: UInt32(__PathOffset), toByteOffset: Int(offset + 134), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(__PathSize), toByteOffset: Int(offset + 134 + 4), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(16), toByteOffset: Int(offset + 134 + 4 + 4), as: UInt32.self)
        let __PathLen = self.Path.count
        var __PathIndex = 0
        while (__PathIndex < __PathLen) {
            if (!self.Path[__PathIndex].Write(&writer, __PathOffset + (UInt32(__PathIndex) * 16))) {
                return false
            }
            __PathIndex = __PathIndex + 1
        }
        let __IsAliveOffset: UInt32 = offset + 146
        writer.memory.storeBytes(of: self.IsAlive, toByteOffset: Int(__IsAliveOffset), as:Bool.self)

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
    self.Name.removeAll(keepingCapacity: true)
    if (__NameLen > self.Name.count) {
        self.Name.reserveCapacity(__NameLen)
        var __NameIndexClear = self.Name.count
        while(__NameIndexClear < __NameLen) {
            self.Name.append(0)
            __NameIndexClear = __NameIndexClear + 1
        }
    }
    var __NameIndex = 0
    while (__NameIndex < __NameLen) {
        self.Name[__NameIndex] = __NameSlice[__NameIndex]
        __NameIndex = __NameIndex + 1
    }
    self.Team = viewer.Team()
    var __InventorySlice = viewer.Inventory(reader)
    let __InventoryLen = __InventorySlice.count
    self.Inventory.removeAll(keepingCapacity: true)
    if (__InventoryLen > self.Inventory.count) {
        self.Inventory.reserveCapacity(__InventoryLen)
        var __InventoryIndexClear = self.Inventory.count
        while(__InventoryIndexClear < __InventoryLen) {
            self.Inventory.append(0)
            __InventoryIndexClear = __InventoryIndexClear + 1
        }
    }
    var __InventoryIndex = 0
    while (__InventoryIndex < __InventoryLen) {
        self.Inventory[__InventoryIndex] = __InventorySlice[__InventoryIndex]
        __InventoryIndex = __InventoryIndex + 1
    }
    self.Color = viewer.Color()
    var __HitboxSlice = viewer.Hitbox()
    var __HitboxLen = __HitboxSlice.count
    if (__HitboxLen > 5) {
        __HitboxLen = 5
    }
    var __HitboxIndex = 0
    while (__HitboxIndex < __HitboxLen) {
        self.Hitbox[__HitboxIndex] = __HitboxSlice[__HitboxIndex]
        __HitboxIndex = __HitboxIndex + 1
    }
    while (__HitboxIndex < self.Hitbox.count) {
        self.Hitbox[__HitboxIndex] = 0
        __HitboxIndex = __HitboxIndex + 1
    }
    var __StatusSlice = viewer.Status(reader)
    let __StatusLen = __StatusSlice.count
    self.Status.removeAll(keepingCapacity: true)
    if (__StatusLen > self.Status.count) {
        self.Status.reserveCapacity(__StatusLen)
        var __StatusIndexClear = self.Status.count
        while(__StatusIndexClear < __StatusLen) {
            self.Status.append(0)
            __StatusIndexClear = __StatusIndexClear + 1
        }
    }
    var __StatusIndex = 0
    while (__StatusIndex < __StatusLen) {
        self.Status[__StatusIndex] = __StatusSlice[__StatusIndex]
        __StatusIndex = __StatusIndex + 1
    }
    var __WeaponsSlice = viewer.Weapons()
    var __WeaponsLen = __WeaponsSlice.count
    if (__WeaponsLen > 4) {
        __WeaponsLen = 4
    }
    var __WeaponsIndex = 0
    while (__WeaponsIndex < __WeaponsLen) {
        self.Weapons[__WeaponsIndex].Read(__WeaponsSlice[__WeaponsIndex], reader)
        __WeaponsIndex = __WeaponsIndex + 1
    }
    while (__WeaponsIndex < self.Weapons.count) {
        self.Weapons[__WeaponsIndex].Reset()
        __WeaponsIndex = __WeaponsIndex + 1
    }
    var __PathSlice = viewer.Path(reader)
    let __PathLen = __PathSlice.count
    self.Path.removeAll(keepingCapacity: true)
    if (__PathLen > self.Path.count) {
        self.Path.reserveCapacity(__PathLen)
        var __PathIndexClear = self.Path.count
        while(__PathIndexClear < __PathLen) {
            self.Path.append(NewVec3())
            __PathIndexClear = __PathIndexClear + 1
        }
    }
    var __PathIndex = 0
    while (__PathIndex < __PathLen) {
        self.Path[__PathIndex].Read(__PathSlice[__PathIndex], reader)
        __PathIndex = __PathIndex + 1
    }
    self.IsAlive = viewer.IsAlive()
    }

}

public func NewMonsterData() -> MonsterData {
    return MonsterData()
}

public struct Monster {
    public var Data: MonsterData = MonsterData()

    public init() {}

    public func PacketIdentifier() -> EnumPacketIdentifier {
        return EnumPacketIdentifierMonster
    }

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
        writer.memory.storeBytes(of: UInt32(__DataOffset), toByteOffset: Int(offset + 0), as: UInt32.self)
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

    public func PacketIdentifier() -> EnumPacketIdentifier {
        return EnumPacketIdentifierMonsters
    }

    public mutating func Reset() -> () {
        self.Monsters.removeAll(keepingCapacity: true)
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
        writer.memory.storeBytes(of: UInt32(16), toByteOffset: Int(offset), as: UInt32.self)
        let __MonstersSize: UInt32 = 8 * UInt32(self.Monsters.count)
        let __MonstersOffset = writer.Alloc(__MonstersSize)
        if (__MonstersOffset == 0) {
            return false
        }
        writer.memory.storeBytes(of: UInt32(__MonstersOffset), toByteOffset: Int(offset + 4), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(__MonstersSize), toByteOffset: Int(offset + 4 + 4), as: UInt32.self)
        writer.memory.storeBytes(of: UInt32(8), toByteOffset: Int(offset + 4 + 4 + 4), as: UInt32.self)
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
    self.Monsters.removeAll(keepingCapacity: true)
    if (__MonstersLen > self.Monsters.count) {
        self.Monsters.reserveCapacity(__MonstersLen)
        var __MonstersIndexClear = self.Monsters.count
        while(__MonstersIndexClear < __MonstersLen) {
            self.Monsters.append(NewMonster())
            __MonstersIndexClear = __MonstersIndexClear + 1
        }
    }
    var __MonstersIndex = 0
    while (__MonstersIndex < __MonstersLen) {
        self.Monsters[__MonstersIndex].Read(__MonstersSlice[__MonstersIndex], reader)
        __MonstersIndex = __MonstersIndex + 1
    }
    }

}

public func NewMonsters() -> Monsters {
    return Monsters()
}


public struct Vec3Viewer : karmem.StructureViewer {
    var _ptr : UnsafeRawPointer
    public var karmemPointer: UnsafeRawPointer { get { return self._ptr } set(newValue) { self._ptr = newValue } }

    public init(ptr: UnsafeRawPointer) {
        self._ptr = ptr
        self.karmemPointer = ptr
    }

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return 16
    }
    @inline(__always)
    public func X() -> Float {
        return (self.karmemPointer + 0).loadUnaligned(as: Float.self)
    }
    @inline(__always)
    public func Y() -> Float {
        return (self.karmemPointer + 4).loadUnaligned(as: Float.self)
    }
    @inline(__always)
    public func Z() -> Float {
        return (self.karmemPointer + 8).loadUnaligned(as: Float.self)
    }
}

@inline(__always) public func NewVec3Viewer(_ reader: karmem.Reader, _ offset: UInt32) -> Vec3Viewer {
    if (!reader.IsValidOffset(offset, 16)) {
        return Vec3Viewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) }))
    }

    let v = Vec3Viewer(ptr: UnsafeRawPointer(reader.pointer + Int(offset)))
    return v
}

public struct WeaponDataViewer : karmem.StructureViewer {
    var _ptr : UnsafeRawPointer
    public var karmemPointer: UnsafeRawPointer { get { return self._ptr } set(newValue) { self._ptr = newValue } }

    public init(ptr: UnsafeRawPointer) {
        self._ptr = ptr
        self.karmemPointer = ptr
    }

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return self.karmemPointer.loadUnaligned(as: UInt32.self)
    }
    @inline(__always)
    public func Damage() -> Int32 {
        if ((UInt32(4) + UInt32(4)) > self.SizeOf()) {
            return 0
        }
        return (self.karmemPointer + 4).loadUnaligned(as: Int32.self)
    }
    @inline(__always)
    public func Range() -> Int32 {
        if ((UInt32(8) + UInt32(4)) > self.SizeOf()) {
            return 0
        }
        return (self.karmemPointer + 8).loadUnaligned(as: Int32.self)
    }
}

@inline(__always) public func NewWeaponDataViewer(_ reader: karmem.Reader, _ offset: UInt32) -> WeaponDataViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return WeaponDataViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) }))
    }

    let v = WeaponDataViewer(ptr: UnsafeRawPointer(reader.pointer + Int(offset)))
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return WeaponDataViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }))
    }
    return v
}

public struct WeaponViewer : karmem.StructureViewer {
    var _ptr : UnsafeRawPointer
    public var karmemPointer: UnsafeRawPointer { get { return self._ptr } set(newValue) { self._ptr = newValue } }

    public init(ptr: UnsafeRawPointer) {
        self._ptr = ptr
        self.karmemPointer = ptr
    }

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return 8
    }
    @inline(__always)
    public func Data(_ reader: karmem.Reader) -> WeaponDataViewer {
        let offset = (self.karmemPointer + 0).loadUnaligned(as: UInt32.self)
        return NewWeaponDataViewer(reader, offset)
    }
}

@inline(__always) public func NewWeaponViewer(_ reader: karmem.Reader, _ offset: UInt32) -> WeaponViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return WeaponViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) }))
    }

    let v = WeaponViewer(ptr: UnsafeRawPointer(reader.pointer + Int(offset)))
    return v
}

public struct MonsterDataViewer : karmem.StructureViewer {
    var _ptr : UnsafeRawPointer
    public var karmemPointer: UnsafeRawPointer { get { return self._ptr } set(newValue) { self._ptr = newValue } }

    public init(ptr: UnsafeRawPointer) {
        self._ptr = ptr
        self.karmemPointer = ptr
    }

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return self.karmemPointer.loadUnaligned(as: UInt32.self)
    }
    @inline(__always)
    public func Pos() -> Vec3Viewer {
        if ((UInt32(4) + UInt32(16)) > self.SizeOf()) {
            return Vec3Viewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }))
        }
        return  Vec3Viewer(ptr: self.karmemPointer + 4)
    }
    @inline(__always)
    public func Mana() -> Int16 {
        if ((UInt32(20) + UInt32(2)) > self.SizeOf()) {
            return 0
        }
        return (self.karmemPointer + 20).loadUnaligned(as: Int16.self)
    }
    @inline(__always)
    public func Health() -> Int16 {
        if ((UInt32(22) + UInt32(2)) > self.SizeOf()) {
            return 0
        }
        return (self.karmemPointer + 22).loadUnaligned(as: Int16.self)
    }
    @inline(__always)
    public func Name(_ reader: karmem.Reader) -> karmem.Slice<UInt8> {
        if ((UInt32(24) + UInt32(12)) > self.SizeOf()) {
                return karmem.NewSliceUnaligned(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as: UInt8.self)
        }
        let offset = (self.karmemPointer + 24).loadUnaligned(as: UInt32.self)
        let size = (self.karmemPointer + 24 + 4).loadUnaligned(as: UInt32.self)
        if (!reader.IsValidOffset(offset, size)) {
                return karmem.NewSliceUnaligned(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as: UInt8.self)
        }

        var length = size / 1
        if (length > 512) {
            length = 512
        }
        return karmem.NewSliceUnaligned(UnsafeRawPointer(reader.pointer + Int(offset)), length, 1, as: UInt8.self)
    }
    @inline(__always)
    public func Team() -> EnumTeam {
        if ((UInt32(36) + UInt32(1)) > self.SizeOf()) {
            return 0
        }
        return (self.karmemPointer + 36).loadUnaligned(as: EnumTeam.self)
    }
    @inline(__always)
    public func Inventory(_ reader: karmem.Reader) -> karmem.Slice<UInt8> {
        if ((UInt32(37) + UInt32(12)) > self.SizeOf()) {
                return karmem.NewSliceUnaligned(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as: UInt8.self)
        }
        let offset = (self.karmemPointer + 37).loadUnaligned(as: UInt32.self)
        let size = (self.karmemPointer + 37 + 4).loadUnaligned(as: UInt32.self)
        if (!reader.IsValidOffset(offset, size)) {
                return karmem.NewSliceUnaligned(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as: UInt8.self)
        }

        var length = size / 1
        if (length > 128) {
            length = 128
        }
        return karmem.NewSliceUnaligned(UnsafeRawPointer(reader.pointer + Int(offset)), length, 1, as: UInt8.self)
    }
    @inline(__always)
    public func Color() -> EnumColor {
        if ((UInt32(49) + UInt32(1)) > self.SizeOf()) {
            return 0
        }
        return (self.karmemPointer + 49).loadUnaligned(as: EnumColor.self)
    }
    @inline(__always)
    public func Hitbox() -> karmem.Slice<Double> {
        if ((UInt32(50) + UInt32(40)) > self.SizeOf()) {
                return karmem.NewSliceUnaligned(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as: Double.self)
        }
        return karmem.NewSliceUnaligned(UnsafeRawPointer(self.karmemPointer + Int(50)), 5, 8, as: Double.self)
    }
    @inline(__always)
    public func Status(_ reader: karmem.Reader) -> karmem.Slice<Int32> {
        if ((UInt32(90) + UInt32(12)) > self.SizeOf()) {
                return karmem.NewSliceUnaligned(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as: Int32.self)
        }
        let offset = (self.karmemPointer + 90).loadUnaligned(as: UInt32.self)
        let size = (self.karmemPointer + 90 + 4).loadUnaligned(as: UInt32.self)
        if (!reader.IsValidOffset(offset, size)) {
                return karmem.NewSliceUnaligned(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as: Int32.self)
        }

        var length = size / 4
        if (length > 10) {
            length = 10
        }
        return karmem.NewSliceUnaligned(UnsafeRawPointer(reader.pointer + Int(offset)), length, 4, as: Int32.self)
    }
    @inline(__always)
    public func Weapons() -> karmem.SliceStructure<WeaponViewer> {
        if ((UInt32(102) + UInt32(32)) > self.SizeOf()) {
                return karmem.NewSliceStructure(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as:WeaponViewer.self)
        }
        return karmem.NewSliceStructure(UnsafeRawPointer(self.karmemPointer + Int(102)), 4, 8, as:WeaponViewer.self)
    }
    @inline(__always)
    public func Path(_ reader: karmem.Reader) -> karmem.SliceStructure<Vec3Viewer> {
        if ((UInt32(134) + UInt32(12)) > self.SizeOf()) {
                return karmem.NewSliceStructure(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as:Vec3Viewer.self)
        }
        let offset = (self.karmemPointer + 134).loadUnaligned(as: UInt32.self)
        let size = (self.karmemPointer + 134 + 4).loadUnaligned(as: UInt32.self)
        if (!reader.IsValidOffset(offset, size)) {
                return karmem.NewSliceStructure(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as:Vec3Viewer.self)
        }

        var length = size / 16
        if (length > 2000) {
            length = 2000
        }
        return karmem.NewSliceStructure(UnsafeRawPointer(reader.pointer + Int(offset)), length, 16, as:Vec3Viewer.self)
    }
    @inline(__always)
    public func IsAlive() -> Bool {
        if ((UInt32(146) + UInt32(1)) > self.SizeOf()) {
            return false
        }
        return (self.karmemPointer + 146).loadUnaligned(as: Bool.self)
    }
}

@inline(__always) public func NewMonsterDataViewer(_ reader: karmem.Reader, _ offset: UInt32) -> MonsterDataViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return MonsterDataViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) }))
    }

    let v = MonsterDataViewer(ptr: UnsafeRawPointer(reader.pointer + Int(offset)))
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return MonsterDataViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }))
    }
    return v
}

public struct MonsterViewer : karmem.StructureViewer {
    var _ptr : UnsafeRawPointer
    public var karmemPointer: UnsafeRawPointer { get { return self._ptr } set(newValue) { self._ptr = newValue } }

    public init(ptr: UnsafeRawPointer) {
        self._ptr = ptr
        self.karmemPointer = ptr
    }

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return 8
    }
    @inline(__always)
    public func Data(_ reader: karmem.Reader) -> MonsterDataViewer {
        let offset = (self.karmemPointer + 0).loadUnaligned(as: UInt32.self)
        return NewMonsterDataViewer(reader, offset)
    }
}

@inline(__always) public func NewMonsterViewer(_ reader: karmem.Reader, _ offset: UInt32) -> MonsterViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return MonsterViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) }))
    }

    let v = MonsterViewer(ptr: UnsafeRawPointer(reader.pointer + Int(offset)))
    return v
}

public struct MonstersViewer : karmem.StructureViewer {
    var _ptr : UnsafeRawPointer
    public var karmemPointer: UnsafeRawPointer { get { return self._ptr } set(newValue) { self._ptr = newValue } }

    public init(ptr: UnsafeRawPointer) {
        self._ptr = ptr
        self.karmemPointer = ptr
    }

    @inline(__always)
    public func SizeOf() -> UInt32 {
        return self.karmemPointer.loadUnaligned(as: UInt32.self)
    }
    @inline(__always)
    public func Monsters(_ reader: karmem.Reader) -> karmem.SliceStructure<MonsterViewer> {
        if ((UInt32(4) + UInt32(12)) > self.SizeOf()) {
                return karmem.NewSliceStructure(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as:MonsterViewer.self)
        }
        let offset = (self.karmemPointer + 4).loadUnaligned(as: UInt32.self)
        let size = (self.karmemPointer + 4 + 4).loadUnaligned(as: UInt32.self)
        if (!reader.IsValidOffset(offset, size)) {
                return karmem.NewSliceStructure(_Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }), 0, 0, as:MonsterViewer.self)
        }

        var length = size / 8
        if (length > 2000) {
            length = 2000
        }
        return karmem.NewSliceStructure(UnsafeRawPointer(reader.pointer + Int(offset)), length, 8, as:MonsterViewer.self)
    }
}

@inline(__always) public func NewMonstersViewer(_ reader: karmem.Reader, _ offset: UInt32) -> MonstersViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return MonstersViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) }))
    }

    let v = MonstersViewer(ptr: UnsafeRawPointer(reader.pointer + Int(offset)))
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return MonstersViewer(ptr: _Null.withUnsafeBufferPointer({ return UnsafeRawPointer($0.baseAddress!) }))
    }
    return v
}
