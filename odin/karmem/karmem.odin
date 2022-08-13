package karmem

import "core:runtime"

Error :: enum {
    ERR_NONE,
	ERR_OUT_OF_MEMORY,
};

Writer :: struct {
    memory: [dynamic]u8,
    isFixed: bool,
}

NewWriter :: proc (capacity: int) -> Writer {
    return Writer{make([dynamic]u8, 0, capacity), false}
}

NewFixedWriter :: proc (mem: [dynamic]u8) -> Writer {
    return Writer{mem, true}
}

NewFixedWriterArray :: proc (mem: [^]u8, size: int) -> Writer {
    slice := [4]int{0, size, size, 0}
    (cast(^[4]rawptr)(&slice))[0] = rawptr(mem)
    return Writer{(cast(^[dynamic]u8)(&slice))^, true}
}

WriterAlloc :: proc(w: ^Writer, n: u32) -> (uint, Error) {
    ptr := len(w.memory)
    total := ptr + int(n)
    
    if total > cap(w.memory) {
        if w.isFixed {
            return 0, .ERR_OUT_OF_MEMORY
        }

        target := cap(w.memory) * 2
        if total > target {
            target = total
        }
        newmem := make([dynamic]u8, total, target)
        #no_bounds_check {
            for i := 0; i < len(w.memory); i += 1 {
                newmem[i] = w.memory[i]
            }
        }
        delete(w.memory)
        w.memory = newmem
    } else {
        (cast(^[3]int)(&w.memory))[1] = total
    }

    return uint(ptr), .ERR_NONE
}

WriterWriteAt :: #force_inline proc(w: ^Writer, offset: uint, i: rawptr, size: u32) {
    #no_bounds_check {
        runtime.mem_copy(rawptr(&w.memory[offset]), i, int(size))
    }
}

WriterWrite1At :: #force_inline proc(w: ^Writer, offset: uint, i: u8) {
    #no_bounds_check {
        (cast(^u8)(&w.memory[offset]))^ = i
    }
}

WriterWrite2At :: #force_inline proc(w: ^Writer, offset: uint, i: u16) {
    #no_bounds_check {
        (cast(^u16)(&w.memory[offset]))^ = i
    }
}

WriterWrite4At :: #force_inline proc(w: ^Writer, offset: uint, i: u32) {
    #no_bounds_check {
        (cast(^u32)(&w.memory[offset]))^ = i
    }
}

WriterWrite8At :: #force_inline proc(w: ^Writer, offset: uint, i: u64) {
    #no_bounds_check {
        (cast(^u64)(&w.memory[offset]))^ = i
    }
}

WriterBytes :: proc(w: ^Writer) -> [dynamic]u8 {
    return w.memory
}

WriterReset :: proc(w: ^Writer) {
    (cast(^[3]int)(&w.memory))[1] = 0
}

Reader :: struct {
    memory: [dynamic]u8,
    pointer: rawptr,
    size: u64,
}

NewReader :: proc(mem: [dynamic]u8) -> Reader {
    return Reader{mem, (rawptr(&mem[0])), u64(len(mem))}
}

NewReaderArray :: proc(mem: [^]u8, size: int) -> Reader {
    slice := [4]int{0, size, size, 0}
    (cast(^[4]rawptr)(&slice))[0] = rawptr(mem)
    return Reader{(cast(^[dynamic]u8)(&slice))^, (rawptr(&mem[0])), u64(size)}
}

ReaderIsValidOffset :: #force_inline proc(r: ^Reader, ptr: u32, size: u32) -> bool {
    return r.size >= (u64(ptr) + u64(size))
}

ReaderSetSize :: proc(r: ^Reader, size: u32) {
    (cast(^[3]int)(&r.memory))[1] = int(size)
    r.size = u64(len(r.memory))
    r.pointer = rawptr(&r.memory[0])
}