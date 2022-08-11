public struct Writer {
    public var memory: UnsafeMutableRawPointer
    var length: UInt32 = 0
    var capacity: UInt32 = 0
    var isFixed: Bool = false

    @usableFromInline
    init(cap: UInt32) {
        self.memory = UnsafeMutableRawPointer.allocate(byteCount: 0, alignment: 1)
        self.isFixed = false
        self.Grow(cap)
    }

    @usableFromInline
    init(memory: UnsafeMutableRawPointer, cap: UInt32) {
        self.memory = memory
        self.isFixed = true
        self.capacity = cap
    }

    @inline(__always)
    public mutating func Alloc(_ n: UInt32) -> UInt32 {
        let offset = self.length
        let total = offset + n
        if (total > self.capacity) {
            if (self.isFixed) {
                return 0xFFFFFFFF
            }
            var target = self.capacity * 2
            if (total > target) {
                target = total * 2
            }
            self.Grow(target)
        }
        var i = self.length
        while (i < total) {
            self.memory.storeBytes(of: UInt8(0), toByteOffset: Int(i), as: UInt8.self)
            i += 1
        }
        self.length = total
        return offset
    }

    @inline(__always)
    public mutating func Grow(_ n: UInt32) {
        let newMemory = UnsafeMutableRawPointer.allocate(byteCount: Int(n), alignment: 1)
        var i = 0
        let t = Int(self.length)
        while (i < Int(n)) {
            if (i > t) {
                newMemory.storeBytes(of: UInt8(0), toByteOffset: i, as: UInt8.self)
            } else {
                newMemory.storeBytes(of: self.memory.load(fromByteOffset: i, as: UInt8.self), toByteOffset: i, as: UInt8.self)
            }
            i += 1
        }
        self.memory.deallocate()
        self.memory = newMemory
        self.capacity = n
    }

    @inline(__always)
    public func Free() {
        if (self.isFixed) {
            return
        }
        self.memory.deallocate()
    }

    @available(*, deprecated, message: "New Karmem generated code doesn't use that function anymore")
    @inline(__always)
    public func WriteAt<T>(_ offset: UInt32, _ data: T) {
        withUnsafePointer(to: data) {
            self.memory.advanced(by: Int(offset)).copyMemory(from: UnsafeRawPointer($0), byteCount: MemoryLayout<T>.size)
        }
    }

    @available(*, deprecated, message: "New Karmem generated code doesn't use that function anymore")
    @inline(__always)
    public func WriteArrayAt<T: Any>(_ offset: UInt32, _ data: [T], _ size: UInt32) {
        var index = 0
        var off = offset
        while(index < data.count) {
            self.WriteAt(off, data[index])
            index = index + 1
            off = off + size
        }
    }

    @inline(__always)
    public func Bytes() -> UnsafeMutableRawPointer {
        return self.memory
    }

    @inline(__always)
    public mutating func Reset() {
        self.length = 0
    }
}

@inline(__always) public func NewWriter(_ cap: UInt32) -> Writer {
    return Writer.init(cap: cap)
}

@inline(__always) public func NewFixedWriter(_ array: [UInt8]) -> Writer {
    return Writer.init(memory: array.withUnsafeBufferPointer({return UnsafeMutableRawPointer(mutating: $0.baseAddress!)}), cap:UInt32(array.count))
}

public struct Reader {
    var memory: [UInt8]
    public var pointer: UnsafeRawPointer
    var size: UInt64

    @usableFromInline
    init(array: [UInt8]) {
        self.memory = array
        self.pointer = array.withUnsafeBufferPointer({return UnsafeRawPointer($0.baseAddress!)})
        self.size = UInt64(array.count)
    }

    @inline(__always)
    public func IsValidOffset(_ ptr: UInt32, _ size: UInt32) -> Bool {
        return self.size >= UInt64(ptr) + UInt64(size)
    }

    @inline(__always)
    public mutating func SetSize(_ size: UInt32) -> Bool {
        if (size > memory.count) {
            return false
        }
        self.size = UInt64(size)
        return true
    }

}

@inline(__always) public func NewReader(_ array: [UInt8]) -> Reader {
    return Reader.init(array: array)
}

public protocol StructureViewer {
    init(ptr: UnsafeRawPointer)
    var karmemPointer: UnsafeRawPointer { get }
}

public struct Slice<T: Any> {
    var pointer: UnsafeRawPointer
    public var count: Int
    var size: Int

    @usableFromInline
    init(ptr: UnsafeRawPointer, len: UInt32, size: UInt32, as: T.Type) {
        self.pointer = ptr
        self.count = Int(len)
        self.size = Int(size)
    }

    @inline(__always)
    public subscript(_ index: Int) -> T {
         mutating get {
            return self.Get(index)
        }
    }

    @inline(__always)
    public func Length() -> Int {
        return self.count
    }

    @inline(__always)
    public mutating func Get(_ index: Int) -> T {
        return (self.pointer + Int(index * self.size)).loadUnaligned(as: T.self)
    }
}

public struct SliceStructure<T: StructureViewer> {
    var pointer: UnsafeRawPointer
    public var count: Int
    var size: Int

    @usableFromInline
    init(ptr: UnsafeRawPointer, len: UInt32, size: UInt32, as: T.Type) {
        self.pointer = ptr
        self.count = Int(len)
        self.size = Int(size)
    }

    @inline(__always)
    public subscript(_ index: Int) -> T {
         mutating get {
            return self.Get(index)
        }
    }

    @inline(__always)
    public func Length() -> Int {
        return self.count
    }

    @inline(__always)
    public mutating func Get(_ index: Int) -> T {
        return T.self.init(ptr: self.pointer + Int(index * self.size))
    }
}

@available(*, deprecated, message: "New Karmem generated code doesn't use that function anymore")
@inline(__always)
public func NewSlice<T>(_ ptr: UnsafeRawPointer, _ len: UInt32, _ size: UInt32, _ stub: T) -> Slice<T> {
    return Slice<T>(ptr: ptr, len: len, size: size, as: T.self)
}

@inline(__always)
public func NewSliceUnaligned<T>(_ ptr: UnsafeRawPointer, _ len: UInt32, _ size: UInt32, as: T.Type) -> Slice<T> {
    return Slice<T>(ptr: ptr, len: len, size: size, as: T.self)
}

@inline(__always)
public func NewSliceStructure<T: StructureViewer>(_ ptr: UnsafeRawPointer, _ len: UInt32, _ size: UInt32, as: T.Type) -> SliceStructure<T> {
    return SliceStructure<T>(ptr: ptr, len: len, size: size, as: T.self)
}

@available(*, deprecated, message: "New Karmem generated code doesn't use that function anymore")
@inline(__always)
public func Load<T: Any>(_ memory: UnsafeRawPointer, _ offset: Int, _ type: inout T) -> T {
    return memory.loadUnaligned(fromByteOffset: offset, as: T.self)
}