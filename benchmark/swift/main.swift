import karmem
import km

var InputMemory: [UInt8] = Array(repeating: 0, count: 8_000_000)
var OutputMemory: [UInt8] = Array(repeating: 0, count: 8_000_001)

@_cdecl("InputMemoryPointer")
func InputMemoryPointer() -> UInt32 {
    return UInt32(Int(bitPattern: InputMemory.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) })))
}

@_cdecl("OutputMemoryPointer")
func OutputMemoryPointer() -> UInt32 {
    return UInt32(Int(bitPattern: OutputMemory.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) })))
}

var _KarmemStruct: km.Monsters = km.Monsters()
var _KarmemWriter = karmem.NewFixedWriter(OutputMemory)
var _KarmemReader = karmem.NewReader(InputMemory)

@_cdecl("KBenchmarkEncodeObjectAPI")
func KBenchmarkEncodeObjectAPI() {
    _KarmemWriter.Reset()
    _ = _KarmemStruct.WriteAsRoot(&_KarmemWriter)
}

@_cdecl("KBenchmarkDecodeObjectAPI")
func KBenchmarkDecodeObjectAPI(_ size: UInt32) {
    _ = _KarmemReader.SetSize(size)
    _KarmemStruct.ReadAsRoot(_KarmemReader)
}

@_cdecl("KBenchmarkDecodeSumVec3")
func KBenchmarkDecodeSumVec3(_ size: UInt32) -> Float {
    _ = _KarmemReader.SetSize(size)

    let monsters = km.NewMonstersViewer(_KarmemReader, 0)
    var monstersList = monsters.Monsters(_KarmemReader)

    var sum = km.Vec3()
    var i: Int = 0
    while (i < monstersList.count) {
        var path = monstersList[i].Data(_KarmemReader).Path(_KarmemReader)
        var p: Int = 0
        while (p < path.count) {
            sum.X += path[p].X()
            sum.Y += path[p].Y()
            sum.Z += path[p].Z()
            p += 1
        }
        i += 1
    }
    return sum.X + sum.Y + sum.Z
}

