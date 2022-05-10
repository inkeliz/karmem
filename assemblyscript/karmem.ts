export class Writer {
    private Memory: Array<u8>;
    private isFixed: boolean;

    @inline constructor(array: Array<u8>, fixed: boolean) {
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

@inline export function NewWriter(capacity: u32): Writer {
    let array = new Array<u8>(capacity)
    array.length = 0
    return new Writer(array, false);
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
        this.Memory = array;
        this.Pointer = changetype<usize>(array);
        this.Size = u64(array.length);
        this.Min = this.Pointer;
        this.Max = usize(u64(this.Pointer) + this.Size);
    }

    @inline IsValidOffset(ptr: u32, size: u32): boolean {
        return this.Size >= u64(ptr) + u64(size)
    }

    @inline SetSize(size: u32): boolean {
        if (size > <u32>this.Memory.length) {
            return false;
        }
        this.Pointer = changetype<usize>(this.Memory);
        this.Size = u64(size);
        this.Min = this.Pointer;
        this.Max = usize(u64(this.Pointer) + this.Size);
        return true;
    }
}

@inline export function NewReader(array: StaticArray<u8>): Reader {
    return new Reader(array)
}

export class Slice<T> {
    [key: number]: T;

    readonly ptr: i32;
    readonly length: i32;
    private readonly size: i32;

    @inline constructor(ptr: usize, length: u32, size: u32) {
        this.ptr = <i32>ptr;
        this.length = <i32>length;
        this.size = <i32>size;
    }

    @inline Length(): i32 {
        return this.length
    }

    @operator("[]") @inline Get(index: i32): T {
        if (index >= this.length) {
            throw new RangeError("Out of range on index:" + index.toString() + "with limit:" + this.length.toString());
        }
        if (isFloat<T>() || isBoolean<T>() || isInteger<T>()) {
            return load<T>(this.ptr + (this.size * index));
        }
        return changetype<T>(this.ptr + (this.size * index));
    }
}