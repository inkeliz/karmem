export class Writer {
    @inline constructor(
        private Memory: Array<u8>,
        private isFixed: bool = false
    ) {}

    @inline Pointer(): usize {
        return this.Memory.dataStart;
    }

    @inline Reset(): void {
        this.Memory.length = 0;
    }

    @inline Alloc(n: u32): u32 {
        let mem = this.Memory;
        let offset: u32 = mem.length;
        let totalSize = offset + n;
        if (this.isFixed && totalSize > <u32>mem.byteLength) {
            return 0xFFFFFFFF;
        }
        mem.length = totalSize;
        return offset;
    }

    @inline WriteAt<T>(offset: u32, data: T): void {
        store<T>(this.Memory.dataStart + offset, data);
    }

    @inline WriteArrayAt<T>(offset: u32, data: T): void {
        memory.copy(this.Memory.dataStart + offset, changetype<usize>(data), data.length * sizeof<valueof<T>>());
    }

    @inline WriteSliceAt<T>(offset: u32, data: T): void {
        memory.copy(this.Memory.dataStart + offset, data.dataStart, data.length * sizeof<valueof<T>>());
    }

    @inline Bytes(): Array<u8> {
        return this.Memory;
    }
}

@inline export function NewWriter(capacity: u32): Writer {
    let array = new Array<u8>(capacity);
    array.length = 0;
    return new Writer(array);
}

@inline export function NewFixedWriter(array: Array<u8>): Writer {
    return new Writer(array, true);
}

export class Reader {
    Memory: StaticArray<u8>;
    Pointer: usize;
    Size: u64;
    Min: usize;
    Max: usize;

    constructor(array: StaticArray<u8>) {
        let ptr  = changetype<usize>(array);
        let size = array.length;
        this.Memory = array;
        this.Pointer = ptr;
        this.Size = <u64>size;
        this.Min = ptr;
        this.Max = ptr + size;
    }

    @inline IsValidOffset(ptr: u32, size: u32): boolean {
        return this.Size >= u64(ptr) + u64(size);
    }

    @inline SetSize(size: u32): boolean {
        let mem = this.Memory;
        if (size > <u32>mem.length) {
            return false;
        }
        let ptr = changetype<usize>(mem);
        this.Pointer = ptr;
        this.Size = <u64>size;
        this.Min = ptr;
        this.Max = ptr + size;
        return true;
    }
}

@inline export function NewReader(array: StaticArray<u8>): Reader {
    return new Reader(array);
}

export class Slice<T> {
    [key: number]: T;

    readonly ptr: usize;
    readonly length: i32;
    private readonly size: i32;

    @inline constructor(ptr: usize, length: u32, size: u32) {
        this.ptr = ptr;
        this.length = <i32>length;
        this.size = <i32>size;
    }

    @inline Length(): i32 {
        return this.length;
    }

    @inline @operator("[]") Get(index: i32): T {
        if (index >= this.length) {
            throw new RangeError(`Out of range on index: ${index} with limit: ${this.length}`);
        }
        if (isFloat<T>() || isBoolean<T>() || isInteger<T>()) {
            return load<T>(this.ptr + (<usize>this.size * index));
        }
        return changetype<T>(this.ptr + (<usize>this.size * index));
    }
}
