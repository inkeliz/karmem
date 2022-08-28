const std = @import("std");
const km = @import("km/game_generated.zig");
const karmem = @import("karmem");
const print = @import("std").debug.print;

var InputMemory: []u8 = undefined;
var OutputMemory: []u8 = undefined;

var KarmemStruct: km.Monsters = km.NewMonsters();
var KarmemWriter: karmem.Writer = undefined;
var KarmemReader: karmem.Reader = undefined;

pub fn main() void {}

export fn _start() void {
    InputMemory = std.heap.page_allocator.alloc(u8, 8_000_000) catch unreachable;
    OutputMemory = std.heap.page_allocator.alloc(u8, 8_000_000) catch unreachable;

    KarmemReader = karmem.NewReader(std.heap.page_allocator, InputMemory[0..OutputMemory.len]);
    KarmemWriter = karmem.NewFixedWriter(OutputMemory[0..OutputMemory.len]);
}

export fn InputMemoryPointer() u32 {
    return @intCast(u32, @ptrToInt(InputMemory.ptr));
}

export fn OutputMemoryPointer() u32 {
    return @intCast(u32, @ptrToInt(OutputMemory.ptr));
}

export fn KBenchmarkEncodeObjectAPI() void {
    karmem.Writer.Reset(&KarmemWriter);
    _ = km.Monsters.WriteAsRoot(&KarmemStruct, &KarmemWriter) catch unreachable;
}

export fn KBenchmarkDecodeObjectAPI(size: u32) void {
    _ = KarmemReader.SetSize(size);
    _ = km.Monsters.ReadAsRoot(&KarmemStruct, &KarmemReader) catch unreachable;
}

export fn KBenchmarkDecodeSumVec3(size: u32) f32 {
    _ = KarmemReader.SetSize(size);

    var monsters = km.NewMonstersViewer(&KarmemReader, 0);
    var monsterList = monsters.Monsters(&KarmemReader);

    var i: usize = 0;
    var sum: f32 = 0;
    while (i < monsterList.len) {
        var path = monsterList[i].Data(&KarmemReader).Path(&KarmemReader);

        var p: usize = 0;
        while (p < path.len) {
            var pp = &path[p];
            sum += pp.X() + pp.Y() + pp.Z();
            p += 1;
        }
        i += 1;
    }

    return sum;
}

export fn KBenchmarkDecodeSumUint8(size: u32) u32 {
    _ = KarmemReader.SetSize(size);

    var monsters = km.NewMonstersViewer(&KarmemReader, 0);
    var monsterList = monsters.Monsters(&KarmemReader);

    var i: usize = 0;
    var sum: u32 = 0;
    while (i < monsterList.len) {
        var inv = monsterList[i].Data(&KarmemReader).Inventory(&KarmemReader);

        var j: usize = 0;
        while (j < inv.len) {
            sum += @as(u32, inv[j]);
            j += 1;
        }
        i += 1;
    }

    return sum;
}

export fn KBenchmarkDecodeSumStats(size: u32) u32 {
    _ = KarmemReader.SetSize(size);

    var monsters = km.NewMonstersViewer(&KarmemReader, 0);
    var monsterList = monsters.Monsters(&KarmemReader);

    var i: usize = 0;
    var sum: km.WeaponData = km.NewWeaponData();
    while (i < monsterList.len) {
        var weapons = monsterList[i].Data(&KarmemReader).Weapons();

        var j: usize = 0;
        while (j < weapons.len) {
            var data = weapons[j].Data(&KarmemReader);
            sum.Ammo += data.Ammo();
            sum.Damage += data.Damage();
            sum.ClipSize += data.ClipSize();
            sum.ReloadTime += data.ReloadTime();
            sum.Range += data.Range();
            j += 1;
        }
        i += 1;
    }

    KarmemWriter.Reset();
    _ = sum.WriteAsRoot(&KarmemWriter) catch unreachable;

    return @as(u32, KarmemWriter.Bytes().len);
}

export fn KNOOP() u32 {
    return 42;
}
