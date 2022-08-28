package main

import "../odin/karmem"
import "./km"
import "core:runtime"
import "core:mem"

InputMemory := [8_000_000]u8{}
OutputMemory := [8_000_000]u8{}

HeapMemory := [64_000_000]u8{}
HeapArena : mem.Arena
HeapAllocator : mem.Allocator

KarmemReader : karmem.Reader
KarmemWriter : karmem.Writer
KarmemStruct : km.Monsters
 
main :: proc () {
    mem.init_arena(&HeapArena, HeapMemory[:])
    HeapAllocator = mem.arena_allocator(&HeapArena)

    context.allocator = HeapAllocator

    KarmemStruct = km.NewMonsters()    
    KarmemReader = karmem.NewReaderArray(InputMemory[:])
    KarmemWriter = karmem.NewFixedWriterArray(OutputMemory[:], 8_000_000)
}

@export
InputMemoryPointer :: proc "c" () -> u32 {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    ptr := rawptr(&InputMemory[0])
    return (cast(^u32)(&ptr))^
}

@export
OutputMemoryPointer :: proc "c" () -> u32 {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    ptr := rawptr(&OutputMemory[0])
    return (cast(^u32)(&ptr))^
}

@export
KBenchmarkEncodeObjectAPI :: proc "c" () {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    karmem.WriterReset(&KarmemWriter)
    km.MonstersWriteAsRoot(&KarmemStruct, &KarmemWriter)
}

@export
KBenchmarkDecodeObjectAPI :: proc "c" (size: u32) {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    karmem.ReaderSetSize(&KarmemReader, size)
    km.MonstersReadAsRoot(&KarmemStruct, &KarmemReader)
}

@export
KBenchmarkDecodeSumVec3 :: proc "c" (size: u32) -> f32 #no_bounds_check {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    karmem.ReaderSetSize(&KarmemReader, size)

    monsters := km.NewMonstersViewer(&KarmemReader, 0)
    monsterList := km.MonstersViewerMonsters(monsters, &KarmemReader)

    sum := f32(0)
    for i := 0; i < len(monsterList); i += 1 {
        path := km.MonsterDataViewerPath(km.MonsterViewerData(&monsterList[i], &KarmemReader), &KarmemReader)

        for p := 0; p < len(path); p += 1 {
            pp := &path[p]
            sum += km.Vec3ViewerX(pp) + km.Vec3ViewerY(pp) + km.Vec3ViewerZ(pp)
        }
    }

    return sum
}

@export
KBenchmarkDecodeSumUint8 :: proc "c" (size: u32) -> u32 #no_bounds_check {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    karmem.ReaderSetSize(&KarmemReader, size)

    monsters := km.NewMonstersViewer(&KarmemReader, 0)
    monsterList := km.MonstersViewerMonsters(monsters, &KarmemReader)

    sum := u32(0)
    for i := 0; i < len(monsterList); i += 1 {
        inv := km.MonsterDataViewerInventory(km.MonsterViewerData(&monsterList[i], &KarmemReader), &KarmemReader)

        for j := 0; j < len(inv); j += 1 {
            sum += u32(inv[j])
        }
    }

    return sum
}

@export
KBenchmarkDecodeSumStats :: proc "c" (size: u32) -> u32 #no_bounds_check {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    karmem.ReaderSetSize(&KarmemReader, size)

    monsters := km.NewMonstersViewer(&KarmemReader, 0)
    monsterList := km.MonstersViewerMonsters(monsters, &KarmemReader)

    sum := km.NewWeaponData()
    for i := 0; i < len(monsterList); i += 1 {
        weapons := km.MonsterDataViewerWeapons(km.MonsterViewerData(&monsterList[i], &KarmemReader))

        for j := 0; j < len(weapons); j += 1 {
			data := km.WeaponViewerData(&weapons[j], &KarmemReader)
			sum.Ammo += km.WeaponDataViewerAmmo(data)
			sum.Damage += km.WeaponDataViewerDamage(data)
			sum.ClipSize += km.WeaponDataViewerClipSize(data)
			sum.ReloadTime += km.WeaponDataViewerReloadTime(data)
			sum.Range += km.WeaponDataViewerRange(data)
        }
    }

    karmem.WriterReset(&KarmemWriter)
    km.WeaponDataWriteAsRoot(&sum, &KarmemWriter)
    return u32(len(karmem.WriterBytes(&KarmemWriter)))
}

@export
KNOOP :: proc "c" () -> u32 #no_bounds_check {
    context = runtime.default_context()
    context.allocator = HeapAllocator

    return 42
}
