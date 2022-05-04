export class Writer {
    private Memory: Array<u8>;
    private isFixed: boolean;

    constructor(array: Array<u8>, fixed: boolean) {
        this.Memory = array;
        this.isFixed = fixed;
    }

    @inline Pointer(): usize {
        return this.Memory.dataStart;
    }

    @inline Reset(): void {
        this.Memory.length = 0;
    }

    @inline Alloc(n: u32): u32 {
        let offset: u32 = this.Memory.length;
        let totalSize = offset + n;
        if (totalSize > u32(this.Memory.byteLength) && this.isFixed) {
            return 0xFFFFFFFF;
        }
        this.Memory.length = totalSize
        return offset;
    }

    @inline WriteAt<T>(offset: u32, data: T): void {
        store<T>(this.Memory.dataStart + offset, data)
    }

    @inline WriteArrayAt<T>(offset: u32, data: T): void {
        memory.copy(this.Memory.dataStart + offset, changetype<usize>(data), data.length * sizeof<valueof<T>>())
    }

    @inline WriteSliceAt<T>(offset: u32, data: T): void {
        memory.copy(this.Memory.dataStart + offset, data.dataStart, data.length * sizeof<valueof<T>>())
    }

    @inline Bytes(): Array<u8> {
        return this.Memory
    }
}

export function NewWriter(capacity: u32): Writer {
    let array = new Array<u8>(capacity)
    array.length = 0
    return new Writer(array, false);
}

export function NewFixedWriter(array: Array<u8>): Writer {
    return new Writer(array, true);
}

export class Reader {
    Memory: StaticArray<u8>;
    Pointer: usize;
    Size: u64;
    Min: usize;
    Max: usize;

    constructor(array: StaticArray<u8>) {
        this.Memory = array;
        this.Pointer = changetype<usize>(array);
        this.Size = u64(array.length);
        this.Min = this.Pointer;
        this.Max = usize(u64(this.Pointer) + this.Size);
    }

    IsValidOffset(ptr: u32, size: u32): boolean {
        return this.Size >= u64(ptr) + u64(size)
    }
}

export function NewReader(array: StaticArray<u8>): Reader {
    return new Reader(array)
}

export class Slice<T> {
    [key: number]: T;

    readonly ptr: usize;
    readonly length: i32;
    private readonly size: i32;

    constructor(ptr: usize, length: u32, size: u32) {
        this.ptr = ptr;
        this.length = <i32>length;
        this.size = <i32>size;
    }

    @operator("[]") Get(index: i32): T {
        if (index >= this.length) {
            throw new RangeError("Out of range on index:" + index.toString() + "with limit:" + this.length.toString());
        }
        if (isFloat<T>() || isBoolean<T>() || isInteger<T>()) {
            return load<T>(this.ptr + (this.size * index));
        }
        return changetype<T>(this.ptr + usize(this.size * index));
    }
}