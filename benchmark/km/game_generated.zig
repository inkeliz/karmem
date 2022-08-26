
const std = @import("std");
const karmem = @import("karmem");
const testing = std.testing;
const Allocator = std.mem.Allocator;
const mem = @import("std").mem;

var _Null: [111]u8 = [_]u8{ 0 } ** 111;
var _NullReader: karmem.Reader = karmem.NewReader(std.heap.page_allocator, _Null[0..111]);


pub const EnumColor = enum(u8) {
    Red = 0,
    Green = 1,
    Blue = 2,
};

const DefaultEnumColor = EnumColor.Red;

pub const EnumTeam = enum(u8) {
    Humans = 0,
    Orcs = 1,
    Zombies = 2,
    Robots = 3,
    Aliens = 4,
};

const DefaultEnumTeam = EnumTeam.Humans;
pub const EnumPacketIdentifier = enum(u64) {
    Vec3 = 10268726485798425099,
    WeaponData = 15342010214468761012,
    Weapon = 8029074423243608167,
    MonsterData = 12254962724431809041,
    Monster = 5593793986513565154,
    Monsters = 14096677544474027661,
};

    
pub const Vec3 = struct {
    X: f32 = 0,
    Y: f32 = 0,
    Z: f32 = 0,

    pub fn PacketIdentifier() EnumPacketIdentifier {
        return EnumPacketIdentifier.Vec3;
    }  

    pub fn Reset(x: *Vec3) void {
        Vec3.Read(x, @ptrCast(*Vec3Viewer, &_Null[0]), &_NullReader) catch unreachable;
    }

    pub fn WriteAsRoot(x: *Vec3, writer: *karmem.Writer) Allocator.Error!u32 {
        return Vec3.Write(x, writer, 0);
    }

    pub fn Write(x: *Vec3, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset: u32 = @intCast(u32, start);
        var size: u32 = 12;
        if (offset == 0) {
            offset = try karmem.Writer.Alloc(writer, size);
        }
        var __XOffset = offset + 0;
        karmem.Writer.WriteAt(writer, __XOffset, @ptrCast([*]const u8, &x.X), 4);
        var __YOffset = offset + 4;
        karmem.Writer.WriteAt(writer, __YOffset, @ptrCast([*]const u8, &x.Y), 4);
        var __ZOffset = offset + 8;
        karmem.Writer.WriteAt(writer, __ZOffset, @ptrCast([*]const u8, &x.Z), 4);

        return offset;
    }

    pub fn ReadAsRoot(x: *Vec3, reader: *karmem.Reader) Allocator.Error!void {
        return Vec3.Read(x, NewVec3Viewer(reader, 0), reader);
    }

    pub fn Read(x: *Vec3, viewer: *const Vec3Viewer, reader: *karmem.Reader) Allocator.Error!void {
        _ = x;
        _ = reader;
        x.X = Vec3Viewer.X(viewer);
        x.Y = Vec3Viewer.Y(viewer);
        x.Z = Vec3Viewer.Z(viewer);
    }

};


pub fn NewVec3() Vec3 {
    var r: Vec3 = Vec3 {};
    return r;
}
pub const WeaponData = struct {
    Damage: i32 = 0,
    Range: i32 = 0,

    pub fn PacketIdentifier() EnumPacketIdentifier {
        return EnumPacketIdentifier.WeaponData;
    }  

    pub fn Reset(x: *WeaponData) void {
        WeaponData.Read(x, @ptrCast(*WeaponDataViewer, &_Null[0]), &_NullReader) catch unreachable;
    }

    pub fn WriteAsRoot(x: *WeaponData, writer: *karmem.Writer) Allocator.Error!u32 {
        return WeaponData.Write(x, writer, 0);
    }

    pub fn Write(x: *WeaponData, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset: u32 = @intCast(u32, start);
        var size: u32 = 12;
        if (offset == 0) {
            offset = try karmem.Writer.Alloc(writer, size);
        }
        var sizeData: u32 = 12;
        karmem.Writer.WriteAt(writer, offset, @ptrCast([*]const u8, &sizeData), 4);
        var __DamageOffset = offset + 4;
        karmem.Writer.WriteAt(writer, __DamageOffset, @ptrCast([*]const u8, &x.Damage), 4);
        var __RangeOffset = offset + 8;
        karmem.Writer.WriteAt(writer, __RangeOffset, @ptrCast([*]const u8, &x.Range), 4);

        return offset;
    }

    pub fn ReadAsRoot(x: *WeaponData, reader: *karmem.Reader) Allocator.Error!void {
        return WeaponData.Read(x, NewWeaponDataViewer(reader, 0), reader);
    }

    pub fn Read(x: *WeaponData, viewer: *const WeaponDataViewer, reader: *karmem.Reader) Allocator.Error!void {
        _ = x;
        _ = reader;
        x.Damage = WeaponDataViewer.Damage(viewer);
        x.Range = WeaponDataViewer.Range(viewer);
    }

};


pub fn NewWeaponData() WeaponData {
    var r: WeaponData = WeaponData {};
    return r;
}
pub const Weapon = struct {
    Data: WeaponData = NewWeaponData(),

    pub fn PacketIdentifier() EnumPacketIdentifier {
        return EnumPacketIdentifier.Weapon;
    }  

    pub fn Reset(x: *Weapon) void {
        Weapon.Read(x, @ptrCast(*WeaponViewer, &_Null[0]), &_NullReader) catch unreachable;
    }

    pub fn WriteAsRoot(x: *Weapon, writer: *karmem.Writer) Allocator.Error!u32 {
        return Weapon.Write(x, writer, 0);
    }

    pub fn Write(x: *Weapon, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset: u32 = @intCast(u32, start);
        var size: u32 = 4;
        if (offset == 0) {
            offset = try karmem.Writer.Alloc(writer, size);
        }
        var __DataSize: usize = 12;
        var __DataOffset = try karmem.Writer.Alloc(writer, __DataSize);

        karmem.Writer.WriteAt(writer, offset+0, @ptrCast([*]const u8, &__DataOffset), 4);
        _ = try WeaponData.Write(&x.Data, writer, __DataOffset);

        return offset;
    }

    pub fn ReadAsRoot(x: *Weapon, reader: *karmem.Reader) Allocator.Error!void {
        return Weapon.Read(x, NewWeaponViewer(reader, 0), reader);
    }

    pub fn Read(x: *Weapon, viewer: *const WeaponViewer, reader: *karmem.Reader) Allocator.Error!void {
        _ = x;
        _ = reader;
        try WeaponData.Read(&x.Data, WeaponViewer.Data(viewer,reader), reader);
    }

};


pub fn NewWeapon() Weapon {
    var r: Weapon = Weapon {};
    return r;
}
pub const MonsterData = struct {
    Pos: Vec3 = NewVec3(),
    Mana: i16 = 0,
    Health: i16 = 0,
    Name: []u8 = &[_]u8{},
    _NameCapacity: usize = 0,
    Team: EnumTeam = DefaultEnumTeam,
    Inventory: []u8 = &[_]u8{},
    _InventoryCapacity: usize = 0,
    Color: EnumColor = DefaultEnumColor,
    Hitbox: [5]f64 = [_]f64{0} ** 5,
    Status: []i32 = &[_]i32{},
    _StatusCapacity: usize = 0,
    Weapons: [4]Weapon = [_]Weapon{NewWeapon()} ** 4,
    Path: []Vec3 = &[_]Vec3{},
    _PathCapacity: usize = 0,
    IsAlive: bool = false,

    pub fn PacketIdentifier() EnumPacketIdentifier {
        return EnumPacketIdentifier.MonsterData;
    }  

    pub fn Reset(x: *MonsterData) void {
        MonsterData.Read(x, @ptrCast(*MonsterDataViewer, &_Null[0]), &_NullReader) catch unreachable;
    }

    pub fn WriteAsRoot(x: *MonsterData, writer: *karmem.Writer) Allocator.Error!u32 {
        return MonsterData.Write(x, writer, 0);
    }

    pub fn Write(x: *MonsterData, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset: u32 = @intCast(u32, start);
        var size: u32 = 111;
        if (offset == 0) {
            offset = try karmem.Writer.Alloc(writer, size);
        }
        var sizeData: u32 = 111;
        karmem.Writer.WriteAt(writer, offset, @ptrCast([*]const u8, &sizeData), 4);
        var __PosOffset = offset + 4;
        _ = try Vec3.Write(&x.Pos, writer, __PosOffset);
        var __ManaOffset = offset + 16;
        karmem.Writer.WriteAt(writer, __ManaOffset, @ptrCast([*]const u8, &x.Mana), 2);
        var __HealthOffset = offset + 18;
        karmem.Writer.WriteAt(writer, __HealthOffset, @ptrCast([*]const u8, &x.Health), 2);
        var __NameSize: usize = 1 * x.Name.len;
        var __NameOffset = try karmem.Writer.Alloc(writer, __NameSize);

        karmem.Writer.WriteAt(writer, offset+20, @ptrCast([*]const u8, &__NameOffset), 4);
        karmem.Writer.WriteAt(writer, offset+20+4, @ptrCast([*]const u8, &__NameSize), 4);
        karmem.Writer.WriteAt(writer, __NameOffset, x.Name.ptr, __NameSize);
        var __TeamOffset = offset + 28;
        karmem.Writer.WriteAt(writer, __TeamOffset, @ptrCast([*]const u8, &x.Team), 1);
        var __InventorySize: usize = 1 * x.Inventory.len;
        var __InventoryOffset = try karmem.Writer.Alloc(writer, __InventorySize);

        karmem.Writer.WriteAt(writer, offset+29, @ptrCast([*]const u8, &__InventoryOffset), 4);
        karmem.Writer.WriteAt(writer, offset+29+4, @ptrCast([*]const u8, &__InventorySize), 4);
        karmem.Writer.WriteAt(writer, __InventoryOffset, @ptrCast(*[]const u8, &x.Inventory).ptr, __InventorySize);
        var __ColorOffset = offset + 37;
        karmem.Writer.WriteAt(writer, __ColorOffset, @ptrCast([*]const u8, &x.Color), 1);
        var __HitboxSize: usize = 8 * x.Hitbox.len;
        var __HitboxOffset = offset + 38;
        karmem.Writer.WriteAt(writer, __HitboxOffset, @ptrCast([*]const u8, &x.Hitbox), __HitboxSize);
        var __StatusSize: usize = 4 * x.Status.len;
        var __StatusOffset = try karmem.Writer.Alloc(writer, __StatusSize);

        karmem.Writer.WriteAt(writer, offset+78, @ptrCast([*]const u8, &__StatusOffset), 4);
        karmem.Writer.WriteAt(writer, offset+78+4, @ptrCast([*]const u8, &__StatusSize), 4);
        karmem.Writer.WriteAt(writer, __StatusOffset, @ptrCast(*[]const u8, &x.Status).ptr, __StatusSize);
        var __WeaponsSize: usize = 4 * x.Weapons.len;
        var __WeaponsOffset = offset + 86;
        var __WeaponsIndex: usize = 0;
        var __WeaponsEnd: usize = __WeaponsOffset + __WeaponsSize;
        while (__WeaponsOffset < __WeaponsEnd) {
            _ = try Weapon.Write(&x.Weapons[__WeaponsIndex], writer, __WeaponsOffset);
            __WeaponsOffset = __WeaponsOffset + 4;
            __WeaponsIndex = __WeaponsIndex + 1;
        }
        var __PathSize: usize = 12 * x.Path.len;
        var __PathOffset = try karmem.Writer.Alloc(writer, __PathSize);

        karmem.Writer.WriteAt(writer, offset+102, @ptrCast([*]const u8, &__PathOffset), 4);
        karmem.Writer.WriteAt(writer, offset+102+4, @ptrCast([*]const u8, &__PathSize), 4);
        var __PathIndex: usize = 0;
        var __PathEnd: usize = __PathOffset + __PathSize;
        while (__PathOffset < __PathEnd) {
            _ = try Vec3.Write(&x.Path[__PathIndex], writer, __PathOffset);
            __PathOffset = __PathOffset + 12;
            __PathIndex = __PathIndex + 1;
        }
        var __IsAliveOffset = offset + 110;
        karmem.Writer.WriteAt(writer, __IsAliveOffset, @ptrCast([*]const u8, &x.IsAlive), 1);

        return offset;
    }

    pub fn ReadAsRoot(x: *MonsterData, reader: *karmem.Reader) Allocator.Error!void {
        return MonsterData.Read(x, NewMonsterDataViewer(reader, 0), reader);
    }

    pub fn Read(x: *MonsterData, viewer: *const MonsterDataViewer, reader: *karmem.Reader) Allocator.Error!void {
        _ = x;
        _ = reader;
        try Vec3.Read(&x.Pos, MonsterDataViewer.Pos(viewer,), reader);
        x.Mana = MonsterDataViewer.Mana(viewer);
        x.Health = MonsterDataViewer.Health(viewer);
        var __NameSlice: []u8 = MonsterDataViewer.Name(viewer, reader);
        var __NameLen: usize = __NameSlice.len;
        if (__NameLen > x._NameCapacity) {
            var __NameCapacityTarget: usize = __NameLen;
            var __NameAlloc = try reader.allocator.reallocAtLeast(x.Name.ptr[0..x._NameCapacity], __NameCapacityTarget);
            var __NameNewIndex: usize = x._NameCapacity;
            while (__NameNewIndex < __NameAlloc.len) {
                __NameAlloc[__NameNewIndex] = 0;
                __NameNewIndex = __NameNewIndex + 1;
            }
            x.Name.ptr = __NameAlloc.ptr;
            x.Name.len = __NameLen;
            x._NameCapacity = __NameAlloc.len;
        }
        if (__NameLen > x.Name.len) {
            x.Name.len = __NameLen;
        }
        var __NameIndex: usize = 0;
        while (__NameIndex < __NameLen) {
            x.Name[__NameIndex] = __NameSlice[__NameIndex];
            __NameIndex = __NameIndex + 1;
        }
        x.Name.len = __NameLen;
        x.Team = MonsterDataViewer.Team(viewer);
        var __InventorySlice: []u8 = MonsterDataViewer.Inventory(viewer, reader);
        var __InventoryLen: usize = __InventorySlice.len;
        if (__InventoryLen > x._InventoryCapacity) {
            var __InventoryCapacityTarget: usize = __InventoryLen;
            var __InventoryAlloc = try reader.allocator.reallocAtLeast(x.Inventory.ptr[0..x._InventoryCapacity], __InventoryCapacityTarget);
            var __InventoryNewIndex: usize = x._InventoryCapacity;
            while (__InventoryNewIndex < __InventoryAlloc.len) {
                __InventoryAlloc[__InventoryNewIndex] = 0;
                __InventoryNewIndex = __InventoryNewIndex + 1;
            }
            x.Inventory.ptr = __InventoryAlloc.ptr;
            x.Inventory.len = __InventoryLen;
            x._InventoryCapacity = __InventoryAlloc.len;
        }
        if (__InventoryLen > x.Inventory.len) {
            x.Inventory.len = __InventoryLen;
        }
        var __InventoryIndex: usize = 0;
        while (__InventoryIndex < __InventoryLen) {
            x.Inventory[__InventoryIndex] = __InventorySlice[__InventoryIndex];
            __InventoryIndex = __InventoryIndex + 1;
        }
        x.Inventory.len = __InventoryLen;
        x.Color = MonsterDataViewer.Color(viewer);
        var __HitboxSlice: []f64 = MonsterDataViewer.Hitbox(viewer);
        var __HitboxLen: usize = __HitboxSlice.len;
        if (__HitboxLen > 5) {
            __HitboxLen = 5;
        }
        var __HitboxIndex: usize = 0;
        while (__HitboxIndex < __HitboxLen) {
            x.Hitbox[__HitboxIndex] = __HitboxSlice[__HitboxIndex];
            __HitboxIndex = __HitboxIndex + 1;
        }
        while (__HitboxIndex < x.Hitbox.len) {
            x.Hitbox[__HitboxIndex] = 0;
            __HitboxIndex = __HitboxIndex + 1;
        }
        var __StatusSlice: []i32 = MonsterDataViewer.Status(viewer, reader);
        var __StatusLen: usize = __StatusSlice.len;
        if (__StatusLen > x._StatusCapacity) {
            var __StatusCapacityTarget: usize = __StatusLen;
            var __StatusAlloc = try reader.allocator.reallocAtLeast(x.Status.ptr[0..x._StatusCapacity], __StatusCapacityTarget);
            var __StatusNewIndex: usize = x._StatusCapacity;
            while (__StatusNewIndex < __StatusAlloc.len) {
                __StatusAlloc[__StatusNewIndex] = 0;
                __StatusNewIndex = __StatusNewIndex + 1;
            }
            x.Status.ptr = __StatusAlloc.ptr;
            x.Status.len = __StatusLen;
            x._StatusCapacity = __StatusAlloc.len;
        }
        if (__StatusLen > x.Status.len) {
            x.Status.len = __StatusLen;
        }
        var __StatusIndex: usize = 0;
        while (__StatusIndex < __StatusLen) {
            x.Status[__StatusIndex] = __StatusSlice[__StatusIndex];
            __StatusIndex = __StatusIndex + 1;
        }
        x.Status.len = __StatusLen;
        var __WeaponsSlice: []const WeaponViewer = MonsterDataViewer.Weapons(viewer);
        var __WeaponsLen: usize = __WeaponsSlice.len;
        if (__WeaponsLen > 4) {
            __WeaponsLen = 4;
        }
        var __WeaponsIndex: usize = 0;
        while (__WeaponsIndex < __WeaponsLen) {
            try Weapon.Read(&x.Weapons[__WeaponsIndex], &__WeaponsSlice[__WeaponsIndex], reader);
             __WeaponsIndex = __WeaponsIndex + 1;
        }
        while (__WeaponsIndex < x.Weapons.len) {
            Weapon.Reset(&x.Weapons[__WeaponsIndex]);
            __WeaponsIndex = __WeaponsIndex + 1;
        }
        var __PathSlice: []const Vec3Viewer = MonsterDataViewer.Path(viewer, reader);
        var __PathLen: usize = __PathSlice.len;
        if (__PathLen > x._PathCapacity) {
            var __PathCapacityTarget: usize = __PathLen;
            var __PathAlloc = try reader.allocator.reallocAtLeast(x.Path.ptr[0..x._PathCapacity], __PathCapacityTarget);
            var __PathNewIndex: usize = x._PathCapacity;
            while (__PathNewIndex < __PathAlloc.len) {
                __PathAlloc[__PathNewIndex] = NewVec3();
                __PathNewIndex = __PathNewIndex + 1;
            }
            x.Path.ptr = __PathAlloc.ptr;
            x.Path.len = __PathLen;
            x._PathCapacity = __PathAlloc.len;
        }
        if (__PathLen > x.Path.len) {
            x.Path.len = __PathLen;
        }
        var __PathIndex: usize = 0;
        while (__PathIndex < __PathLen) {
            try Vec3.Read(&x.Path[__PathIndex], &__PathSlice[__PathIndex], reader);
             __PathIndex = __PathIndex + 1;
        }
        x.Path.len = __PathLen;
        x.IsAlive = MonsterDataViewer.IsAlive(viewer);
    }

};


pub fn NewMonsterData() MonsterData {
    var r: MonsterData = MonsterData {};
    return r;
}
pub const Monster = struct {
    Data: MonsterData = NewMonsterData(),

    pub fn PacketIdentifier() EnumPacketIdentifier {
        return EnumPacketIdentifier.Monster;
    }  

    pub fn Reset(x: *Monster) void {
        Monster.Read(x, @ptrCast(*MonsterViewer, &_Null[0]), &_NullReader) catch unreachable;
    }

    pub fn WriteAsRoot(x: *Monster, writer: *karmem.Writer) Allocator.Error!u32 {
        return Monster.Write(x, writer, 0);
    }

    pub fn Write(x: *Monster, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset: u32 = @intCast(u32, start);
        var size: u32 = 4;
        if (offset == 0) {
            offset = try karmem.Writer.Alloc(writer, size);
        }
        var __DataSize: usize = 111;
        var __DataOffset = try karmem.Writer.Alloc(writer, __DataSize);

        karmem.Writer.WriteAt(writer, offset+0, @ptrCast([*]const u8, &__DataOffset), 4);
        _ = try MonsterData.Write(&x.Data, writer, __DataOffset);

        return offset;
    }

    pub fn ReadAsRoot(x: *Monster, reader: *karmem.Reader) Allocator.Error!void {
        return Monster.Read(x, NewMonsterViewer(reader, 0), reader);
    }

    pub fn Read(x: *Monster, viewer: *const MonsterViewer, reader: *karmem.Reader) Allocator.Error!void {
        _ = x;
        _ = reader;
        try MonsterData.Read(&x.Data, MonsterViewer.Data(viewer,reader), reader);
    }

};


pub fn NewMonster() Monster {
    var r: Monster = Monster {};
    return r;
}
pub const Monsters = struct {
    Monsters: []Monster = &[_]Monster{},
    _MonstersCapacity: usize = 0,

    pub fn PacketIdentifier() EnumPacketIdentifier {
        return EnumPacketIdentifier.Monsters;
    }  

    pub fn Reset(x: *Monsters) void {
        Monsters.Read(x, @ptrCast(*MonstersViewer, &_Null[0]), &_NullReader) catch unreachable;
    }

    pub fn WriteAsRoot(x: *Monsters, writer: *karmem.Writer) Allocator.Error!u32 {
        return Monsters.Write(x, writer, 0);
    }

    pub fn Write(x: *Monsters, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset: u32 = @intCast(u32, start);
        var size: u32 = 12;
        if (offset == 0) {
            offset = try karmem.Writer.Alloc(writer, size);
        }
        var sizeData: u32 = 12;
        karmem.Writer.WriteAt(writer, offset, @ptrCast([*]const u8, &sizeData), 4);
        var __MonstersSize: usize = 4 * x.Monsters.len;
        var __MonstersOffset = try karmem.Writer.Alloc(writer, __MonstersSize);

        karmem.Writer.WriteAt(writer, offset+4, @ptrCast([*]const u8, &__MonstersOffset), 4);
        karmem.Writer.WriteAt(writer, offset+4+4, @ptrCast([*]const u8, &__MonstersSize), 4);
        var __MonstersIndex: usize = 0;
        var __MonstersEnd: usize = __MonstersOffset + __MonstersSize;
        while (__MonstersOffset < __MonstersEnd) {
            _ = try Monster.Write(&x.Monsters[__MonstersIndex], writer, __MonstersOffset);
            __MonstersOffset = __MonstersOffset + 4;
            __MonstersIndex = __MonstersIndex + 1;
        }

        return offset;
    }

    pub fn ReadAsRoot(x: *Monsters, reader: *karmem.Reader) Allocator.Error!void {
        return Monsters.Read(x, NewMonstersViewer(reader, 0), reader);
    }

    pub fn Read(x: *Monsters, viewer: *const MonstersViewer, reader: *karmem.Reader) Allocator.Error!void {
        _ = x;
        _ = reader;
        var __MonstersSlice: []const MonsterViewer = MonstersViewer.Monsters(viewer, reader);
        var __MonstersLen: usize = __MonstersSlice.len;
        if (__MonstersLen > x._MonstersCapacity) {
            var __MonstersCapacityTarget: usize = __MonstersLen;
            var __MonstersAlloc = try reader.allocator.reallocAtLeast(x.Monsters.ptr[0..x._MonstersCapacity], __MonstersCapacityTarget);
            var __MonstersNewIndex: usize = x._MonstersCapacity;
            while (__MonstersNewIndex < __MonstersAlloc.len) {
                __MonstersAlloc[__MonstersNewIndex] = NewMonster();
                __MonstersNewIndex = __MonstersNewIndex + 1;
            }
            x.Monsters.ptr = __MonstersAlloc.ptr;
            x.Monsters.len = __MonstersLen;
            x._MonstersCapacity = __MonstersAlloc.len;
        }
        if (__MonstersLen > x.Monsters.len) {
            x.Monsters.len = __MonstersLen;
        }
        var __MonstersIndex: usize = 0;
        while (__MonstersIndex < __MonstersLen) {
            try Monster.Read(&x.Monsters[__MonstersIndex], &__MonstersSlice[__MonstersIndex], reader);
             __MonstersIndex = __MonstersIndex + 1;
        }
        x.Monsters.len = __MonstersLen;
    }

};


pub fn NewMonsters() Monsters {
    var r: Monsters = Monsters {};
    return r;
}


pub const Vec3Viewer = extern struct {
    _data: [12]u8,

    pub fn Size(x: *const Vec3Viewer) u32 {
    _ = x;
        return 12;
    }
    pub fn X(x: *const Vec3Viewer) f32 {
        return @ptrCast(*align(1) const f32, x._data[0..0+@sizeOf(f32)]).*;
    }
    pub fn Y(x: *const Vec3Viewer) f32 {
        return @ptrCast(*align(1) const f32, x._data[4..4+@sizeOf(f32)]).*;
    }
    pub fn Z(x: *const Vec3Viewer) f32 {
        return @ptrCast(*align(1) const f32, x._data[8..8+@sizeOf(f32)]).*;
    }

};

pub fn NewVec3Viewer(reader: *karmem.Reader, offset: u32) *const Vec3Viewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 12)) {
        return @ptrCast(*Vec3Viewer, &_Null[0]);
    }
    var v = @ptrCast(*align(1) const Vec3Viewer, reader.memory[offset..offset+12]);
    return v;
}

pub const WeaponDataViewer = extern struct {
    _data: [12]u8,

    pub fn Size(x: *const WeaponDataViewer) u32 {
    _ = x;
        return @ptrCast(*align(1) const u32, x._data[0..4]).*;
    }
    pub fn Damage(x: *const WeaponDataViewer) i32 {
        if ((4 + 4) > WeaponDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i32, x._data[4..4+@sizeOf(i32)]).*;
    }
    pub fn Range(x: *const WeaponDataViewer) i32 {
        if ((8 + 4) > WeaponDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i32, x._data[8..8+@sizeOf(i32)]).*;
    }

};

pub fn NewWeaponDataViewer(reader: *karmem.Reader, offset: u32) *const WeaponDataViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 4)) {
        return @ptrCast(*WeaponDataViewer, &_Null[0]);
    }
    var v = @ptrCast(*align(1) const WeaponDataViewer, reader.memory[offset..offset+4]);
    if (!karmem.Reader.IsValidOffset(reader, offset, WeaponDataViewer.Size(v))) {
        return @ptrCast(*WeaponDataViewer, &_Null[0]);
    }
    return v;
}

pub const WeaponViewer = extern struct {
    _data: [4]u8,

    pub fn Size(x: *const WeaponViewer) u32 {
    _ = x;
        return 4;
    }
    pub fn Data(x: *const WeaponViewer, reader: *karmem.Reader) *const WeaponDataViewer {
        var offset = @ptrCast(*align(1) const u32, x._data[0..0+4]).*;
         return NewWeaponDataViewer(reader, offset);
    }

};

pub fn NewWeaponViewer(reader: *karmem.Reader, offset: u32) *const WeaponViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 4)) {
        return @ptrCast(*WeaponViewer, &_Null[0]);
    }
    var v = @ptrCast(*align(1) const WeaponViewer, reader.memory[offset..offset+4]);
    return v;
}

pub const MonsterDataViewer = extern struct {
    _data: [111]u8,

    pub fn Size(x: *const MonsterDataViewer) u32 {
    _ = x;
        return @ptrCast(*align(1) const u32, x._data[0..4]).*;
    }
    pub fn Pos(x: *const MonsterDataViewer) *const Vec3Viewer {
        if ((4 + 12) > MonsterDataViewer.Size(x)) {
            return @ptrCast(*Vec3Viewer, &_Null[0]);
        }
        return @ptrCast(*const Vec3Viewer, x._data[4..4+@sizeOf(*const Vec3Viewer)]);
    }
    pub fn Mana(x: *const MonsterDataViewer) i16 {
        if ((16 + 2) > MonsterDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i16, x._data[16..16+@sizeOf(i16)]).*;
    }
    pub fn Health(x: *const MonsterDataViewer) i16 {
        if ((18 + 2) > MonsterDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i16, x._data[18..18+@sizeOf(i16)]).*;
    }
    pub fn Name(x: *const MonsterDataViewer, reader: *karmem.Reader) []u8 {
        if ((20 + 8) > MonsterDataViewer.Size(x)) {
            return &[_]u8{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[20..20+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[20+4..20+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return&[_]u8{};
        }
        var length = size / 1;
        if (length > 512) {
            length = 512;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]u8, &slice).*;
    }
    pub fn Team(x: *const MonsterDataViewer) EnumTeam {
        if ((28 + 1) > MonsterDataViewer.Size(x)) {
            return DefaultEnumTeam;
        }
        return @ptrCast(*align(1) const EnumTeam, x._data[28..28+@sizeOf(EnumTeam)]).*;
    }
    pub fn Inventory(x: *const MonsterDataViewer, reader: *karmem.Reader) []u8 {
        if ((29 + 8) > MonsterDataViewer.Size(x)) {
            return &[_]u8{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[29..29+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[29+4..29+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return&[_]u8{};
        }
        var length = size / 1;
        if (length > 128) {
            length = 128;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]u8, &slice).*;
    }
    pub fn Color(x: *const MonsterDataViewer) EnumColor {
        if ((37 + 1) > MonsterDataViewer.Size(x)) {
            return DefaultEnumColor;
        }
        return @ptrCast(*align(1) const EnumColor, x._data[37..37+@sizeOf(EnumColor)]).*;
    }
    pub fn Hitbox(x: *const MonsterDataViewer) []f64 {
        if ((38 + 40) > MonsterDataViewer.Size(x)) {
            return &[_]f64{};
        }
        var slice = [2]usize{@ptrToInt(x)+38, 5};
        return @ptrCast(*align(1) const []f64, &slice).*;
    }
    pub fn Status(x: *const MonsterDataViewer, reader: *karmem.Reader) []i32 {
        if ((78 + 8) > MonsterDataViewer.Size(x)) {
            return &[_]i32{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[78..78+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[78+4..78+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return&[_]i32{};
        }
        var length = size / 4;
        if (length > 10) {
            length = 10;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]i32, &slice).*;
    }
    pub fn Weapons(x: *const MonsterDataViewer) []const WeaponViewer {
        if ((86 + 16) > MonsterDataViewer.Size(x)) {
            return &[_]WeaponViewer{};
        }
        var slice = [2]usize{@ptrToInt(x)+86, 4};
        return @ptrCast(*align(1) const []WeaponViewer, &slice).*;
    }
    pub fn Path(x: *const MonsterDataViewer, reader: *karmem.Reader) []const Vec3Viewer {
        if ((102 + 8) > MonsterDataViewer.Size(x)) {
            return &[_]Vec3Viewer{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[102..102+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[102+4..102+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return &[_]Vec3Viewer{};
        }
        var length = size / 12;
        if (length > 2000) {
            length = 2000;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]const Vec3Viewer, &slice).*;
    }
    pub fn IsAlive(x: *const MonsterDataViewer) bool {
        if ((110 + 1) > MonsterDataViewer.Size(x)) {
            return false;
        }
        return @ptrCast(*align(1) const bool, x._data[110..110+@sizeOf(bool)]).*;
    }

};

pub fn NewMonsterDataViewer(reader: *karmem.Reader, offset: u32) *const MonsterDataViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 4)) {
        return @ptrCast(*MonsterDataViewer, &_Null[0]);
    }
    var v = @ptrCast(*align(1) const MonsterDataViewer, reader.memory[offset..offset+4]);
    if (!karmem.Reader.IsValidOffset(reader, offset, MonsterDataViewer.Size(v))) {
        return @ptrCast(*MonsterDataViewer, &_Null[0]);
    }
    return v;
}

pub const MonsterViewer = extern struct {
    _data: [4]u8,

    pub fn Size(x: *const MonsterViewer) u32 {
    _ = x;
        return 4;
    }
    pub fn Data(x: *const MonsterViewer, reader: *karmem.Reader) *const MonsterDataViewer {
        var offset = @ptrCast(*align(1) const u32, x._data[0..0+4]).*;
         return NewMonsterDataViewer(reader, offset);
    }

};

pub fn NewMonsterViewer(reader: *karmem.Reader, offset: u32) *const MonsterViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 4)) {
        return @ptrCast(*MonsterViewer, &_Null[0]);
    }
    var v = @ptrCast(*align(1) const MonsterViewer, reader.memory[offset..offset+4]);
    return v;
}

pub const MonstersViewer = extern struct {
    _data: [12]u8,

    pub fn Size(x: *const MonstersViewer) u32 {
    _ = x;
        return @ptrCast(*align(1) const u32, x._data[0..4]).*;
    }
    pub fn Monsters(x: *const MonstersViewer, reader: *karmem.Reader) []const MonsterViewer {
        if ((4 + 8) > MonstersViewer.Size(x)) {
            return &[_]MonsterViewer{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[4..4+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[4+4..4+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return &[_]MonsterViewer{};
        }
        var length = size / 4;
        if (length > 2000) {
            length = 2000;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]const MonsterViewer, &slice).*;
    }

};

pub fn NewMonstersViewer(reader: *karmem.Reader, offset: u32) *const MonstersViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 4)) {
        return @ptrCast(*MonstersViewer, &_Null[0]);
    }
    var v = @ptrCast(*align(1) const MonstersViewer, reader.memory[offset..offset+4]);
    if (!karmem.Reader.IsValidOffset(reader, offset, MonstersViewer.Size(v))) {
        return @ptrCast(*MonstersViewer, &_Null[0]);
    }
    return v;
}
