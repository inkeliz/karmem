const std = @import("std");
const testing = std.testing;
const Allocator = std.mem.Allocator;
const mem = @import("std").mem;

pub const Writer = struct {
    memory: []u8,
    capacity: usize,
    allocator: Allocator,
    isFixed: bool,

    pub fn Alloc(self: *Writer, len: usize) Allocator.Error!u32 {
        var offset: usize = self.memory.len;
        var total: usize = offset + len;
        if (total > self.capacity) {
            var capacityTarget: usize = self.capacity * 2;
            if (total > capacityTarget) {
                capacityTarget = total;
            }
            try Writer.Grow(self, capacityTarget);
        }
        self.memory.len = total;
        return @intCast(u32, offset);
    }

    pub fn Grow(self: *Writer, cap: usize) Allocator.Error!void {
        if (self.isFixed) {
            return std.mem.Allocator.Error.OutOfMemory;
        }
        const new_memory = try self.allocator.reallocAtLeast(self.memory.ptr[0..self.capacity], cap);
        self.memory.ptr = new_memory.ptr;
        self.capacity = new_memory.len;
    }

    pub fn WriteAt(self: *Writer, offset: u32, src: [*]const u8, len: usize) void {
        @memcpy(self.memory.ptr[offset .. offset + len].ptr, src, len);
    }

    pub fn Bytes(self: *Writer) []u8 {
        return self.memory;
    }

    pub fn Reset(self: *Writer) void {
        self.memory.len = 0;
    }

    pub fn Free(self: *Writer) void {
        self.allocator.free(self.memory.ptr[0..self.capacity]);
    }
};

pub fn NewWriter(allocator: Allocator, cap: usize) !Writer {
    var r: Writer = .{
        .memory = &[_]u8{},
        .capacity = 0,
        .allocator = allocator,
        .isFixed = false,
    };

    try Writer.Grow(&r, cap);
    return r;
}

pub fn NewFixedWriter(slice: []u8) Writer {
    var r: Writer = .{
        .memory = slice,
        .capacity = slice.len,
        .allocator = std.heap.page_allocator,
        .isFixed = true,
    };

    Writer.Reset(&r);
    return r;
}

pub const Reader = struct {
    memory: []u8,
    length: u64,
    allocator: Allocator,

    pub fn IsValidOffset(x: *Reader, offset: u32, size: u32) bool {
        return x.length >= (@intCast(u64, offset) + @intCast(u64, size));
    }

    pub fn SetSize(x: *Reader, size: u32) bool {
        if (size > x.memory.len) {
            return false;
        }
        x.length = @intCast(u64, size);
        return true;
    }
};

pub fn NewReader(allocator: Allocator, memory: []u8) Reader {
    return Reader{
        .memory = memory,
        .length = @intCast(u64, memory.len),
        .allocator = allocator,
    };
}
