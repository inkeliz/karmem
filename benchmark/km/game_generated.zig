
const std = @import("std");
const karmem = @import("karmem");
const testing = std.testing;
const Allocator = std.mem.Allocator;
const mem = @import("std").mem;


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

pub const Vec3 = struct {
    X: f32,
    Y: f32,
    Z: f32,

    pub fn Reset(x: *Vec3) void {
        if (x == undefined) {
            return;
        }
        x.X = 0;
        x.Y = 0;
        x.Z = 0;
    }

    pub fn WriteAsRoot(x: *Vec3, writer: *karmem.Writer) Allocator.Error!u32 {
        return Vec3.Write(x, writer, 0);
    }

    pub fn Write(x: *Vec3, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset : u32 = @intCast(u32, start);
        var size : u32 = 16;
        if (offset == 0) {
            offset = try writer.Alloc(size);
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
        if (x == undefined) {
            return;
        }
        if (reader == undefined) {
            return;
        }
        x.X = Vec3Viewer.X(viewer);
        x.Y = Vec3Viewer.Y(viewer);
        x.Z = Vec3Viewer.Z(viewer);
    }

};


pub fn NewVec3() Vec3 {
    return Vec3 {
    .X = 0,
    .Y = 0,
    .Z = 0,
    };
}
pub const WeaponData = struct {
    Damage: i32,
    Range: i32,

    pub fn Reset(x: *WeaponData) void {
        if (x == undefined) {
            return;
        }
        x.Damage = 0;
        x.Range = 0;
    }

    pub fn WriteAsRoot(x: *WeaponData, writer: *karmem.Writer) Allocator.Error!u32 {
        return WeaponData.Write(x, writer, 0);
    }

    pub fn Write(x: *WeaponData, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset : u32 = @intCast(u32, start);
        var size : u32 = 16;
        if (offset == 0) {
            offset = try writer.Alloc(size);
        }
        karmem.Writer.WriteAt(writer, offset, @ptrCast([*]const u8, &size), 4);
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
        if (x == undefined) {
            return;
        }
        if (reader == undefined) {
            return;
        }
        x.Damage = WeaponDataViewer.Damage(viewer);
        x.Range = WeaponDataViewer.Range(viewer);
    }

};


pub fn NewWeaponData() WeaponData {
    return WeaponData {
    .Damage = 0,
    .Range = 0,
    };
}
pub const Weapon = struct {
    Data: WeaponData,

    pub fn Reset(x: *Weapon) void {
        if (x == undefined) {
            return;
        }
        WeaponData.Reset(&x.Data);
    }

    pub fn WriteAsRoot(x: *Weapon, writer: *karmem.Writer) Allocator.Error!u32 {
        return Weapon.Write(x, writer, 0);
    }

    pub fn Write(x: *Weapon, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset : u32 = @intCast(u32, start);
        var size : u32 = 8;
        if (offset == 0) {
            offset = try writer.Alloc(size);
        }
        var __DataSize : usize = 16;
        var __DataOffset = try writer.Alloc(__DataSize);

        karmem.Writer.WriteAt(writer, offset+0, @ptrCast([*]const u8, &__DataOffset), 4);
        _ = try WeaponData.Write(&x.Data, writer, __DataOffset);

        return offset;
    }

    pub fn ReadAsRoot(x: *Weapon, reader: *karmem.Reader) Allocator.Error!void {
        return Weapon.Read(x, NewWeaponViewer(reader, 0), reader);
    }

    pub fn Read(x: *Weapon, viewer: *const WeaponViewer, reader: *karmem.Reader) Allocator.Error!void {
        if (x == undefined) {
            return;
        }
        if (reader == undefined) {
            return;
        }
        try WeaponData.Read(&x.Data, WeaponViewer.Data(viewer,reader), reader);
    }

};


pub fn NewWeapon() Weapon {
    return Weapon {
    .Data = NewWeaponData(),
    };
}
pub const MonsterData = struct {
    Pos: Vec3,
    Mana: i16,
    Health: i16,
    Name: []u8,
    _NameCapacity: usize,
    Team: EnumTeam,
    Inventory: []u8,
    _InventoryCapacity: usize,
    Color: EnumColor,
    Hitbox: [5]f64,
    Status: []i32,
    _StatusCapacity: usize,
    Weapons: [4]Weapon,
    Path: []Vec3,
    _PathCapacity: usize,

    pub fn Reset(x: *MonsterData) void {
        if (x == undefined) {
            return;
        }
        Vec3.Reset(&x.Pos);
        x.Mana = 0;
        x.Health = 0;
    x.Name.len = 0;
        x.Team = DefaultEnumTeam;
    x.Inventory.len = 0;
        x.Color = DefaultEnumColor;
        x.Hitbox = [_]f64{ 0 } ** 5;
    x.Status.len = 0;
        var __WeaponsIndex : usize = 0;
        while (__WeaponsIndex < x.Weapons.len) {
            Weapon.Reset(&x.Weapons[__WeaponsIndex]);
            __WeaponsIndex = __WeaponsIndex + 1;
        }
        var __PathIndex : usize = 0;
        while (__PathIndex < x.Path.len) {
            Vec3.Reset(&x.Path[__PathIndex]);
            __PathIndex = __PathIndex + 1;
        }
    x.Path.len = 0;
    }

    pub fn WriteAsRoot(x: *MonsterData, writer: *karmem.Writer) Allocator.Error!u32 {
        return MonsterData.Write(x, writer, 0);
    }

    pub fn Write(x: *MonsterData, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset : u32 = @intCast(u32, start);
        var size : u32 = 152;
        if (offset == 0) {
            offset = try writer.Alloc(size);
        }
        karmem.Writer.WriteAt(writer, offset, @ptrCast([*]const u8, &size), 4);
        var __PosOffset = offset + 4;
        _ = try Vec3.Write(&x.Pos, writer, __PosOffset);
        var __ManaOffset = offset + 20;
        karmem.Writer.WriteAt(writer, __ManaOffset, @ptrCast([*]const u8, &x.Mana), 2);
        var __HealthOffset = offset + 22;
        karmem.Writer.WriteAt(writer, __HealthOffset, @ptrCast([*]const u8, &x.Health), 2);
        var __NameSize : usize = 1 * x.Name.len;
        var __NameOffset = try writer.Alloc(__NameSize);

        karmem.Writer.WriteAt(writer, offset+24, @ptrCast([*]const u8, &__NameOffset), 4);
        karmem.Writer.WriteAt(writer, offset+24+4, @ptrCast([*]const u8, &__NameSize), 4);
        var __NameSizeEach : u32 = 1;
        karmem.Writer.WriteAt(writer, offset+24+4+4, @ptrCast([*]const u8, &__NameSizeEach), 4);
        karmem.Writer.WriteAt(writer, __NameOffset, x.Name.ptr, __NameSize);
        var __TeamOffset = offset + 36;
        karmem.Writer.WriteAt(writer, __TeamOffset, @ptrCast([*]const u8, &x.Team), 1);
        var __InventorySize : usize = 1 * x.Inventory.len;
        var __InventoryOffset = try writer.Alloc(__InventorySize);

        karmem.Writer.WriteAt(writer, offset+37, @ptrCast([*]const u8, &__InventoryOffset), 4);
        karmem.Writer.WriteAt(writer, offset+37+4, @ptrCast([*]const u8, &__InventorySize), 4);
        var __InventorySizeEach : u32 = 1;
        karmem.Writer.WriteAt(writer, offset+37+4+4, @ptrCast([*]const u8, &__InventorySizeEach), 4);
        karmem.Writer.WriteAt(writer, __InventoryOffset, @ptrCast(*[]const u8, &x.Inventory).ptr, __InventorySize);
        var __ColorOffset = offset + 49;
        karmem.Writer.WriteAt(writer, __ColorOffset, @ptrCast([*]const u8, &x.Color), 1);
        var __HitboxSize : usize = 8 * x.Hitbox.len;
        var __HitboxOffset = offset + 50;
        karmem.Writer.WriteAt(writer, __HitboxOffset, @ptrCast([*]const u8, &x.Hitbox), __HitboxSize);
        var __StatusSize : usize = 4 * x.Status.len;
        var __StatusOffset = try writer.Alloc(__StatusSize);

        karmem.Writer.WriteAt(writer, offset+90, @ptrCast([*]const u8, &__StatusOffset), 4);
        karmem.Writer.WriteAt(writer, offset+90+4, @ptrCast([*]const u8, &__StatusSize), 4);
        var __StatusSizeEach : u32 = 4;
        karmem.Writer.WriteAt(writer, offset+90+4+4, @ptrCast([*]const u8, &__StatusSizeEach), 4);
        karmem.Writer.WriteAt(writer, __StatusOffset, @ptrCast(*[]const u8, &x.Status).ptr, __StatusSize);
        var __WeaponsSize : usize = 8 * x.Weapons.len;
        var __WeaponsOffset = offset + 102;
        var __WeaponsIndex : usize = 0;
        var __WeaponsEnd : usize = __WeaponsOffset + __WeaponsSize;
        while (__WeaponsOffset < __WeaponsEnd) {
            _ = try Weapon.Write(&x.Weapons[__WeaponsIndex], writer, __WeaponsOffset);
            __WeaponsOffset = __WeaponsOffset + 8;
            __WeaponsIndex = __WeaponsIndex + 1;
        }
        var __PathSize : usize = 16 * x.Path.len;
        var __PathOffset = try writer.Alloc(__PathSize);

        karmem.Writer.WriteAt(writer, offset+134, @ptrCast([*]const u8, &__PathOffset), 4);
        karmem.Writer.WriteAt(writer, offset+134+4, @ptrCast([*]const u8, &__PathSize), 4);
        var __PathSizeEach : u32 = 16;
        karmem.Writer.WriteAt(writer, offset+134+4+4, @ptrCast([*]const u8, &__PathSizeEach), 4);
        var __PathIndex : usize = 0;
        var __PathEnd : usize = __PathOffset + __PathSize;
        while (__PathOffset < __PathEnd) {
            _ = try Vec3.Write(&x.Path[__PathIndex], writer, __PathOffset);
            __PathOffset = __PathOffset + 16;
            __PathIndex = __PathIndex + 1;
        }

        return offset;
    }

    pub fn ReadAsRoot(x: *MonsterData, reader: *karmem.Reader) Allocator.Error!void {
        return MonsterData.Read(x, NewMonsterDataViewer(reader, 0), reader);
    }

    pub fn Read(x: *MonsterData, viewer: *const MonsterDataViewer, reader: *karmem.Reader) Allocator.Error!void {
        if (x == undefined) {
            return;
        }
        if (reader == undefined) {
            return;
        }
        try Vec3.Read(&x.Pos, MonsterDataViewer.Pos(viewer,), reader);
        x.Mana = MonsterDataViewer.Mana(viewer);
        x.Health = MonsterDataViewer.Health(viewer);
        var __NameSlice : []u8 = MonsterDataViewer.Name(viewer, reader);
        var __NameLen : usize = __NameSlice.len;
        if (__NameLen > x._NameCapacity) {
            var __NameCapacityTarget : usize = __NameLen;
            var __NameAlloc = try reader.allocator.reallocAtLeast(x.Name.ptr[0..x._NameCapacity], __NameCapacityTarget);
            var __NameNewIndex : usize = x._NameCapacity;
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
        var __NameIndex : usize = 0;
        while (__NameIndex < __NameLen) {
            x.Name[__NameIndex] = __NameSlice[__NameIndex];
            __NameIndex = __NameIndex + 1;
        }
        while (__NameIndex < x.Name.len) {
            x.Name[__NameIndex] = 0;
            __NameIndex = __NameIndex + 1;
        }
        x.Name.len = __NameLen;
        x.Team = MonsterDataViewer.Team(viewer);
        var __InventorySlice : []const u8 = MonsterDataViewer.Inventory(viewer, reader);
        var __InventoryLen : usize = __InventorySlice.len;
        if (__InventoryLen > x._InventoryCapacity) {
            var __InventoryCapacityTarget : usize = __InventoryLen;
            var __InventoryAlloc = try reader.allocator.reallocAtLeast(x.Inventory.ptr[0..x._InventoryCapacity], __InventoryCapacityTarget);
            var __InventoryNewIndex : usize = x._InventoryCapacity;
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
        var __InventoryIndex : usize = 0;
        while (__InventoryIndex < __InventoryLen) {
            x.Inventory[__InventoryIndex] = __InventorySlice[__InventoryIndex];
            __InventoryIndex = __InventoryIndex + 1;
        }
        while (__InventoryIndex < x.Inventory.len) {
            x.Inventory[__InventoryIndex] = 0;
            __InventoryIndex = __InventoryIndex + 1;
        }
        x.Inventory.len = __InventoryLen;
        x.Color = MonsterDataViewer.Color(viewer);
        var __HitboxSlice : []const f64 = MonsterDataViewer.Hitbox(viewer);
        var __HitboxLen : usize = __HitboxSlice.len;
        if (__HitboxLen > x.Hitbox.len) {
            __HitboxLen = x.Hitbox.len;
        }
        var __HitboxIndex : usize = 0;
        while (__HitboxIndex < __HitboxLen) {
            x.Hitbox[__HitboxIndex] = __HitboxSlice[__HitboxIndex];
            __HitboxIndex = __HitboxIndex + 1;
        }
        while (__HitboxIndex < x.Hitbox.len) {
            x.Hitbox[__HitboxIndex] = 0;
            __HitboxIndex = __HitboxIndex + 1;
        }
        var __StatusSlice : []const i32 = MonsterDataViewer.Status(viewer, reader);
        var __StatusLen : usize = __StatusSlice.len;
        if (__StatusLen > x._StatusCapacity) {
            var __StatusCapacityTarget : usize = __StatusLen;
            var __StatusAlloc = try reader.allocator.reallocAtLeast(x.Status.ptr[0..x._StatusCapacity], __StatusCapacityTarget);
            var __StatusNewIndex : usize = x._StatusCapacity;
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
        var __StatusIndex : usize = 0;
        while (__StatusIndex < __StatusLen) {
            x.Status[__StatusIndex] = __StatusSlice[__StatusIndex];
            __StatusIndex = __StatusIndex + 1;
        }
        while (__StatusIndex < x.Status.len) {
            x.Status[__StatusIndex] = 0;
            __StatusIndex = __StatusIndex + 1;
        }
        x.Status.len = __StatusLen;
        var __WeaponsSlice : []const WeaponViewer = MonsterDataViewer.Weapons(viewer);
        var __WeaponsLen : usize = __WeaponsSlice.len;
        if (__WeaponsLen > x.Weapons.len) {
            __WeaponsLen = x.Weapons.len;
        }
        var __WeaponsIndex : usize = 0;
        while (__WeaponsIndex < __WeaponsLen) {
            try Weapon.Read(&x.Weapons[__WeaponsIndex], &__WeaponsSlice[__WeaponsIndex], reader);
             __WeaponsIndex = __WeaponsIndex + 1;
        }
        while (__WeaponsIndex < x.Weapons.len) {
            Weapon.Reset(&x.Weapons[__WeaponsIndex]);
            __WeaponsIndex = __WeaponsIndex + 1;
        }
        var __PathSlice : []const Vec3Viewer = MonsterDataViewer.Path(viewer, reader);
        var __PathLen : usize = __PathSlice.len;
        if (__PathLen > x._PathCapacity) {
            var __PathCapacityTarget : usize = __PathLen;
            var __PathAlloc = try reader.allocator.reallocAtLeast(x.Path.ptr[0..x._PathCapacity], __PathCapacityTarget);
            var __PathNewIndex : usize = x._PathCapacity;
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
        var __PathIndex : usize = 0;
        while (__PathIndex < __PathLen) {
            try Vec3.Read(&x.Path[__PathIndex], &__PathSlice[__PathIndex], reader);
             __PathIndex = __PathIndex + 1;
        }
        while (__PathIndex < x.Path.len) {
            Vec3.Reset(&x.Path[__PathIndex]);
            __PathIndex = __PathIndex + 1;
        }
        x.Path.len = __PathLen;
    }

};


pub fn NewMonsterData() MonsterData {
    return MonsterData {
    .Pos = NewVec3(),
    .Mana = 0,
    .Health = 0,
    .Name = &[_]u8{},
    ._NameCapacity = 0,
    .Team = DefaultEnumTeam,
    .Inventory = &[_]u8{},
    ._InventoryCapacity = 0,
    .Color = DefaultEnumColor,
    .Hitbox = [_]f64{0} ** 5,
    .Status = &[_]i32{},
    ._StatusCapacity = 0,
    .Weapons = [_]Weapon{NewWeapon()} ** 4,
    .Path = &[_]Vec3{},
    ._PathCapacity = 0,
    };
}
pub const Monster = struct {
    Data: MonsterData,

    pub fn Reset(x: *Monster) void {
        if (x == undefined) {
            return;
        }
        MonsterData.Reset(&x.Data);
    }

    pub fn WriteAsRoot(x: *Monster, writer: *karmem.Writer) Allocator.Error!u32 {
        return Monster.Write(x, writer, 0);
    }

    pub fn Write(x: *Monster, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset : u32 = @intCast(u32, start);
        var size : u32 = 8;
        if (offset == 0) {
            offset = try writer.Alloc(size);
        }
        var __DataSize : usize = 152;
        var __DataOffset = try writer.Alloc(__DataSize);

        karmem.Writer.WriteAt(writer, offset+0, @ptrCast([*]const u8, &__DataOffset), 4);
        _ = try MonsterData.Write(&x.Data, writer, __DataOffset);

        return offset;
    }

    pub fn ReadAsRoot(x: *Monster, reader: *karmem.Reader) Allocator.Error!void {
        return Monster.Read(x, NewMonsterViewer(reader, 0), reader);
    }

    pub fn Read(x: *Monster, viewer: *const MonsterViewer, reader: *karmem.Reader) Allocator.Error!void {
        if (x == undefined) {
            return;
        }
        if (reader == undefined) {
            return;
        }
        try MonsterData.Read(&x.Data, MonsterViewer.Data(viewer,reader), reader);
    }

};


pub fn NewMonster() Monster {
    return Monster {
    .Data = NewMonsterData(),
    };
}
pub const Monsters = struct {
    Monsters: []Monster,
    _MonstersCapacity: usize,

    pub fn Reset(x: *Monsters) void {
        if (x == undefined) {
            return;
        }
        var __MonstersIndex : usize = 0;
        while (__MonstersIndex < x.Monsters.len) {
            Monster.Reset(&x.Monsters[__MonstersIndex]);
            __MonstersIndex = __MonstersIndex + 1;
        }
    x.Monsters.len = 0;
    }

    pub fn WriteAsRoot(x: *Monsters, writer: *karmem.Writer) Allocator.Error!u32 {
        return Monsters.Write(x, writer, 0);
    }

    pub fn Write(x: *Monsters, writer: *karmem.Writer, start: usize) Allocator.Error!u32 {
        var offset : u32 = @intCast(u32, start);
        var size : u32 = 24;
        if (offset == 0) {
            offset = try writer.Alloc(size);
        }
        karmem.Writer.WriteAt(writer, offset, @ptrCast([*]const u8, &size), 4);
        var __MonstersSize : usize = 8 * x.Monsters.len;
        var __MonstersOffset = try writer.Alloc(__MonstersSize);

        karmem.Writer.WriteAt(writer, offset+4, @ptrCast([*]const u8, &__MonstersOffset), 4);
        karmem.Writer.WriteAt(writer, offset+4+4, @ptrCast([*]const u8, &__MonstersSize), 4);
        var __MonstersSizeEach : u32 = 8;
        karmem.Writer.WriteAt(writer, offset+4+4+4, @ptrCast([*]const u8, &__MonstersSizeEach), 4);
        var __MonstersIndex : usize = 0;
        var __MonstersEnd : usize = __MonstersOffset + __MonstersSize;
        while (__MonstersOffset < __MonstersEnd) {
            _ = try Monster.Write(&x.Monsters[__MonstersIndex], writer, __MonstersOffset);
            __MonstersOffset = __MonstersOffset + 8;
            __MonstersIndex = __MonstersIndex + 1;
        }

        return offset;
    }

    pub fn ReadAsRoot(x: *Monsters, reader: *karmem.Reader) Allocator.Error!void {
        return Monsters.Read(x, NewMonstersViewer(reader, 0), reader);
    }

    pub fn Read(x: *Monsters, viewer: *const MonstersViewer, reader: *karmem.Reader) Allocator.Error!void {
        if (x == undefined) {
            return;
        }
        if (reader == undefined) {
            return;
        }
        var __MonstersSlice : []const MonsterViewer = MonstersViewer.Monsters(viewer, reader);
        var __MonstersLen : usize = __MonstersSlice.len;
        if (__MonstersLen > x._MonstersCapacity) {
            var __MonstersCapacityTarget : usize = __MonstersLen;
            var __MonstersAlloc = try reader.allocator.reallocAtLeast(x.Monsters.ptr[0..x._MonstersCapacity], __MonstersCapacityTarget);
            var __MonstersNewIndex : usize = x._MonstersCapacity;
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
        var __MonstersIndex : usize = 0;
        while (__MonstersIndex < __MonstersLen) {
            try Monster.Read(&x.Monsters[__MonstersIndex], &__MonstersSlice[__MonstersIndex], reader);
             __MonstersIndex = __MonstersIndex + 1;
        }
        while (__MonstersIndex < x.Monsters.len) {
            Monster.Reset(&x.Monsters[__MonstersIndex]);
            __MonstersIndex = __MonstersIndex + 1;
        }
        x.Monsters.len = __MonstersLen;
    }

};


pub fn NewMonsters() Monsters {
    return Monsters {
    .Monsters = &[_]Monster{},
    ._MonstersCapacity = 0,
    };
}


pub const Vec3Viewer = struct {
    _data: [16]u8,

    pub fn Size(x: *const Vec3Viewer) u32 {
        if (x == undefined) {
            return 0xFFFFFFFF;
        }
        return 16;
    }
    pub fn X(x: *const Vec3Viewer) f32 {
        if (x == undefined) {
            return 0;
        }
        return @ptrCast(*align(1) const f32, x._data[0..0+@sizeOf(f32)]).*;
    }
    pub fn Y(x: *const Vec3Viewer) f32 {
        if (x == undefined) {
            return 0;
        }
        return @ptrCast(*align(1) const f32, x._data[4..4+@sizeOf(f32)]).*;
    }
    pub fn Z(x: *const Vec3Viewer) f32 {
        if (x == undefined) {
            return 0;
        }
        return @ptrCast(*align(1) const f32, x._data[8..8+@sizeOf(f32)]).*;
    }

};

var NullVec3Viewer = Vec3Viewer {
    ._data = [_]u8{0} ** 16
};

pub fn NewVec3Viewer(reader: *karmem.Reader, offset: u32) *const Vec3Viewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 16)) {
        return &NullVec3Viewer;
    }
    var v = @ptrCast(*align(1) const Vec3Viewer, reader.memory[offset..offset+16]);
    return v;
}

pub const WeaponDataViewer = struct {
    _data: [16]u8,

    pub fn Size(x: *const WeaponDataViewer) u32 {
        if (x == undefined) {
            return 0xFFFFFFFF;
        }
        return @ptrCast(*align(1) const u32, x._data[0..4]).*;
    }
    pub fn Damage(x: *const WeaponDataViewer) i32 {
        if (x == undefined) {
            return 0;
        }
        if ((4 + 4) > WeaponDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i32, x._data[4..4+@sizeOf(i32)]).*;
    }
    pub fn Range(x: *const WeaponDataViewer) i32 {
        if (x == undefined) {
            return 0;
        }
        if ((8 + 4) > WeaponDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i32, x._data[8..8+@sizeOf(i32)]).*;
    }

};

var NullWeaponDataViewer = WeaponDataViewer {
    ._data = [_]u8{0} ** 16
};

pub fn NewWeaponDataViewer(reader: *karmem.Reader, offset: u32) *const WeaponDataViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 8)) {
        return &NullWeaponDataViewer;
    }
    var v = @ptrCast(*align(1) const WeaponDataViewer, reader.memory[offset..offset+8]);
    if (!karmem.Reader.IsValidOffset(reader, offset, WeaponDataViewer.Size(v))) {
        return &NullWeaponDataViewer;
    }
    return v;
}

pub const WeaponViewer = struct {
    _data: [8]u8,

    pub fn Size(x: *const WeaponViewer) u32 {
        if (x == undefined) {
            return 0xFFFFFFFF;
        }
        return 8;
    }
    pub fn Data(x: *const WeaponViewer, reader: *karmem.Reader) *const WeaponDataViewer {
        if (x == undefined) {
            return &NullWeaponDataViewer;
        }
        var offset = @ptrCast(*align(1) const u32, x._data[0..0+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, 16)) {
            return &NullWeaponDataViewer;
        }
        var v = @ptrCast(*align(1) const WeaponDataViewer, reader.memory[offset..offset+16]);
        if (!karmem.Reader.IsValidOffset(reader, offset, WeaponDataViewer.Size(v))) {
            return &NullWeaponDataViewer;
        }
        return v;
    }

};

var NullWeaponViewer = WeaponViewer {
    ._data = [_]u8{0} ** 8
};

pub fn NewWeaponViewer(reader: *karmem.Reader, offset: u32) *const WeaponViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 8)) {
        return &NullWeaponViewer;
    }
    var v = @ptrCast(*align(1) const WeaponViewer, reader.memory[offset..offset+8]);
    return v;
}

pub const MonsterDataViewer = struct {
    _data: [152]u8,

    pub fn Size(x: *const MonsterDataViewer) u32 {
        if (x == undefined) {
            return 0xFFFFFFFF;
        }
        return @ptrCast(*align(1) const u32, x._data[0..4]).*;
    }
    pub fn Pos(x: *const MonsterDataViewer) *const Vec3Viewer {
        if (x == undefined) {
            return &NullVec3Viewer;
        }
        if ((4 + 16) > MonsterDataViewer.Size(x)) {
            return &NullVec3Viewer;
        }
        return @ptrCast(*const Vec3Viewer, x._data[4..4+@sizeOf(*const Vec3Viewer)]);
    }
    pub fn Mana(x: *const MonsterDataViewer) i16 {
        if (x == undefined) {
            return 0;
        }
        if ((20 + 2) > MonsterDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i16, x._data[20..20+@sizeOf(i16)]).*;
    }
    pub fn Health(x: *const MonsterDataViewer) i16 {
        if (x == undefined) {
            return 0;
        }
        if ((22 + 2) > MonsterDataViewer.Size(x)) {
            return 0;
        }
        return @ptrCast(*align(1) const i16, x._data[22..22+@sizeOf(i16)]).*;
    }
    pub fn Name(x: *const MonsterDataViewer, reader: *karmem.Reader) []u8 {
        if (x == undefined) {
            return &[_]u8{};
        }
        if ((24 + 12) > MonsterDataViewer.Size(x)) {
            return &[_]u8{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[24..24+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[24+4..24+4+4]).*;
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
        if (x == undefined) {
            return DefaultEnumTeam;
        }
        if ((36 + 1) > MonsterDataViewer.Size(x)) {
            return DefaultEnumTeam;
        }
        return @ptrCast(*align(1) const EnumTeam, x._data[36..36+@sizeOf(EnumTeam)]).*;
    }
    pub fn Inventory(x: *const MonsterDataViewer, reader: *karmem.Reader) []const u8 {
        if (x == undefined) {
            return &[_]u8{};
        }
        if ((37 + 12) > MonsterDataViewer.Size(x)) {
            return &[_]u8{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[37..37+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[37+4..37+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return&[_]u8{};
        }
        var length = size / 1;
        if (length > 128) {
            length = 128;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]const u8, &slice).*;
    }
    pub fn Color(x: *const MonsterDataViewer) EnumColor {
        if (x == undefined) {
            return DefaultEnumColor;
        }
        if ((49 + 1) > MonsterDataViewer.Size(x)) {
            return DefaultEnumColor;
        }
        return @ptrCast(*align(1) const EnumColor, x._data[49..49+@sizeOf(EnumColor)]).*;
    }
    pub fn Hitbox(x: *const MonsterDataViewer) []const f64 {
        if (x == undefined) {
            return &[_]f64{};
        }
        if ((50 + 40) > MonsterDataViewer.Size(x)) {
            return &[_]f64{};
        }
        var slice = [2]usize{@ptrToInt(x)+50, 5};
        return @ptrCast(*align(1) const []const f64, &slice).*;
    }
    pub fn Status(x: *const MonsterDataViewer, reader: *karmem.Reader) []const i32 {
        if (x == undefined) {
            return &[_]i32{};
        }
        if ((90 + 12) > MonsterDataViewer.Size(x)) {
            return &[_]i32{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[90..90+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[90+4..90+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return&[_]i32{};
        }
        var length = size / 4;
        if (length > 10) {
            length = 10;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]const i32, &slice).*;
    }
    pub fn Weapons(x: *const MonsterDataViewer) []const WeaponViewer {
        if (x == undefined) {
            return &[_]WeaponViewer{};
        }
        if ((102 + 32) > MonsterDataViewer.Size(x)) {
            return &[_]WeaponViewer{};
        }
        var slice = [2]usize{@ptrToInt(x)+102, 4};
        return @ptrCast(*align(1) const []WeaponViewer, &slice).*;
    }
    pub fn Path(x: *const MonsterDataViewer, reader: *karmem.Reader) []const Vec3Viewer {
        if (x == undefined) {
            return &[_]Vec3Viewer{};
        }
        if ((134 + 12) > MonsterDataViewer.Size(x)) {
            return &[_]Vec3Viewer{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[134..134+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[134+4..134+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return &[_]Vec3Viewer{};
        }
        var length = size / 16;
        if (length > 2000) {
            length = 2000;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]const Vec3Viewer, &slice).*;
    }

};

var NullMonsterDataViewer = MonsterDataViewer {
    ._data = [_]u8{0} ** 152
};

pub fn NewMonsterDataViewer(reader: *karmem.Reader, offset: u32) *const MonsterDataViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 8)) {
        return &NullMonsterDataViewer;
    }
    var v = @ptrCast(*align(1) const MonsterDataViewer, reader.memory[offset..offset+8]);
    if (!karmem.Reader.IsValidOffset(reader, offset, MonsterDataViewer.Size(v))) {
        return &NullMonsterDataViewer;
    }
    return v;
}

pub const MonsterViewer = struct {
    _data: [8]u8,

    pub fn Size(x: *const MonsterViewer) u32 {
        if (x == undefined) {
            return 0xFFFFFFFF;
        }
        return 8;
    }
    pub fn Data(x: *const MonsterViewer, reader: *karmem.Reader) *const MonsterDataViewer {
        if (x == undefined) {
            return &NullMonsterDataViewer;
        }
        var offset = @ptrCast(*align(1) const u32, x._data[0..0+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, 152)) {
            return &NullMonsterDataViewer;
        }
        var v = @ptrCast(*align(1) const MonsterDataViewer, reader.memory[offset..offset+152]);
        if (!karmem.Reader.IsValidOffset(reader, offset, MonsterDataViewer.Size(v))) {
            return &NullMonsterDataViewer;
        }
        return v;
    }

};

var NullMonsterViewer = MonsterViewer {
    ._data = [_]u8{0} ** 8
};

pub fn NewMonsterViewer(reader: *karmem.Reader, offset: u32) *const MonsterViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 8)) {
        return &NullMonsterViewer;
    }
    var v = @ptrCast(*align(1) const MonsterViewer, reader.memory[offset..offset+8]);
    return v;
}

pub const MonstersViewer = struct {
    _data: [24]u8,

    pub fn Size(x: *const MonstersViewer) u32 {
        if (x == undefined) {
            return 0xFFFFFFFF;
        }
        return @ptrCast(*align(1) const u32, x._data[0..4]).*;
    }
    pub fn Monsters(x: *const MonstersViewer, reader: *karmem.Reader) []const MonsterViewer {
        if (x == undefined) {
            return &[_]MonsterViewer{};
        }
        if ((4 + 12) > MonstersViewer.Size(x)) {
            return &[_]MonsterViewer{};
        }
        var offset = @ptrCast(*align(1) const u32, x._data[4..4+4]).*;
        var size = @ptrCast(*align(1) const u32, x._data[4+4..4+4+4]).*;
        if (!karmem.Reader.IsValidOffset(reader, offset, size)) {
            return &[_]MonsterViewer{};
        }
        var length = size / 8;
        if (length > 2000) {
            length = 2000;
        }
        var slice = [2]usize{@ptrToInt(reader.memory.ptr)+offset, length};
        return @ptrCast(*[]const MonsterViewer, &slice).*;
    }

};

var NullMonstersViewer = MonstersViewer {
    ._data = [_]u8{0} ** 24
};

pub fn NewMonstersViewer(reader: *karmem.Reader, offset: u32) *const MonstersViewer {
    if (!karmem.Reader.IsValidOffset(reader, offset, 8)) {
        return &NullMonstersViewer;
    }
    var v = @ptrCast(*align(1) const MonstersViewer, reader.memory[offset..offset+8]);
    if (!karmem.Reader.IsValidOffset(reader, offset, MonstersViewer.Size(v))) {
        return &NullMonstersViewer;
    }
    return v;
}
