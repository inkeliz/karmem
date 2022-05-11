const std = @import("std");
const km = @import("km/game_generated.zig");
const karmem = @import("karmem");
const print = @import("std").debug.print;

var InputMemory : []u8 = undefined;
var OutputMemory : []u8 = undefined;

var KarmemStruct : km.Monsters = km.NewMonsters();
var KarmemWriter : karmem.Writer = undefined;
var KarmemReader : karmem.Reader = undefined;

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

var x: f32 =0;

export fn KBenchmarkDecodeSumVec3(size: u32) f32 {
    _ = KarmemReader.SetSize(size);

    var monsters = km.NewMonstersViewer(&KarmemReader, 0);
    var monsterList = monsters.Monsters(&KarmemReader);

    var i : usize = 0;
    var sum : km.Vec3 = km.NewVec3();
    while (i < monsterList.len) {
        var path = monsterList[i].Data(&KarmemReader).Path(&KarmemReader);

        var p : usize = 0;
        while (p < path.len) {
            sum.X += path[p].X();
            sum.Y += path[p].Y();
            sum.Z += path[p].Z();
            p += 1;
        }
        i += 1;
    }

    return sum.X+sum.Y+sum.Z;
}