import karmem
import km

var InputMemory: [UInt8] = Array(repeating: 0, count: 8_000_000)
var OutputMemory: [UInt8] = Array(repeating: 0, count: 8_000_000)

@_cdecl("InputMemoryPointer")
func InputMemoryPointer() -> UInt32 {
    return UInt32(Int(bitPattern: InputMemory.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) })))
}

@_cdecl("OutputMemoryPointer")
func OutputMemoryPointer() -> UInt32 {
    return UInt32(Int(bitPattern: OutputMemory.withUnsafeBufferPointer({ return UnsafePointer($0.baseAddress!) })))
}

var _Struct = km.Monsters()
var _Writer = karmem.NewFixedWriter(OutputMemory)
var _Reader = karmem.NewReader(InputMemory)

@_cdecl("KBenchmarkEncodeObjectAPI")
func KBenchmarkEncodeObjectAPI() {
    _Writer.Reset()
    _ = _Struct.WriteAsRoot(&_Writer)
}

@_cdecl("KBenchmarkDecodeObjectAPI")
func KBenchmarkDecodeObjectAPI(_ size: UInt32) {
    _ = _Reader.SetSize(size)
    _Struct.ReadAsRoot(_Reader)
}

@_cdecl("KBenchmarkDecodeSumVec3")
func KBenchmarkDecodeSumVec3(_ size: UInt32) -> Float32 {
    _ = _Reader.SetSize(size)

    let monsters = km.NewMonstersViewer(_Reader, 0)
    var monstersList = monsters.Monsters(_Reader)

    var sum: Float32 = 0.0
    var i: Int = 0
    while (i < monstersList.count) {
        var path = monstersList[i].Data(_Reader).Path(_Reader)
        var p: Int = 0
        while (p < path.count) {
            let pp = path[p]
            sum += (pp.X() + pp.Y() + pp.Z())
            p += 1
        }
        i += 1
    }
    return sum
}

@_cdecl("KBenchmarkDecodeSumUint8")
func KBenchmarkDecodeSumUint8(_ size: UInt32) -> UInt32 {
    _ = _Reader.SetSize(size)

    let monsters = km.NewMonstersViewer(_Reader, 0)
    var monstersList = monsters.Monsters(_Reader)

	var sum: UInt32 = 0
    var i: Int = 0
    while (i < monstersList.count) {
		var inv = monstersList[i].Data(_Reader).Inventory(_Reader)

        var j: Int = 0
        while (j < inv.count) {
			sum &+= UInt32(inv[j])
			j += 1
		}
		i += 1
	}

	return sum
}

@_cdecl("KBenchmarkDecodeSumStats")
func KBenchmarkDecodeSumStats(_ size: UInt32) -> UInt32 {
	_ = _Reader.SetSize(size)

	let monsters = km.NewMonstersViewer(_Reader, 0)
	var monstersList = monsters.Monsters(_Reader)

	var sum: km.WeaponData = km.NewWeaponData()
	var i: Int = 0
    while (i < monstersList.count) {
		var weapons = monstersList[i].Data(_Reader).Weapons()

        var j: Int = 0
        while (j < weapons.count) {
			let data = weapons[j].Data(_Reader)
			sum.Ammo &+= data.Ammo()
			sum.Damage &+= data.Damage()
			sum.ClipSize &+= data.ClipSize()
			sum.ReloadTime += data.ReloadTime()
			sum.Range &+= data.Range()
			j += 1
		}
		i += 1
	}

	_Writer.Reset()
	_ = sum.WriteAsRoot(&_Writer)
	return _Writer.Length()
}

@_cdecl("KNOOP")
func KNOOP() -> UInt32 {
	return 42
}