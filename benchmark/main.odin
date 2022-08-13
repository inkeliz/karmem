package main

import "../odin/karmem"
import "./km"
import "core:runtime"

InputMemory := [8_000_000]u8{}
OutputMemory := [8_000_000]u8{}

KarmemReader : karmem.Reader
KarmemWriter : karmem.Writer
KarmemStruct : km.Monsters
 
main :: proc () {
    KarmemStruct = km.NewMonsters()
    KarmemReader = karmem.NewReaderArray(cast([^]u8)(&InputMemory[0]), 8_000_000)
    KarmemWriter = karmem.NewFixedWriterArray(cast([^]u8)(&OutputMemory[0]), 8_000_000)
}

@export
InputMemoryPointer :: proc "c" () -> u32 {
    context = runtime.default_context()

    ptr := rawptr(&InputMemory)
    return (cast(^u32)(&ptr))^
}

@export
OutputMemoryPointer :: proc "c" () -> u32 {
    context = runtime.default_context()

    ptr := rawptr(&OutputMemory)
    return (cast(^u32)(&ptr))^
}

@export
KBenchmarkEncodeObjectAPI :: proc "c" () {
    context = runtime.default_context()

    karmem.WriterReset(&KarmemWriter)
    km.MonstersWriteAsRoot(&KarmemStruct, &KarmemWriter)
}

@export
KBenchmarkDecodeObjectAPI :: proc "c" (size: u32) {
    context = runtime.default_context()

    karmem.ReaderSetSize(&KarmemReader, size)
    km.MonstersReadAsRoot(&KarmemStruct, &KarmemReader)
}

@export
KBenchmarkDecodeSumVec3 :: proc "c" (size: u32) -> f32 #no_bounds_check {
    context = runtime.default_context()

    karmem.ReaderSetSize(&KarmemReader, size)

    monsters := km.NewMonstersViewer(&KarmemReader, 0)
    monsterList := km.MonstersViewerMonsters(monsters, &KarmemReader)

    sum := km.NewVec3();
    for i := 0; i < len(monsterList); i += 1 {
        path := km.MonsterDataViewerPath(km.MonsterViewerData(&monsterList[i], &KarmemReader), &KarmemReader)

        for p := 0; p < len(path); p += 1 {
            pp := &path[p]
            sum.X += km.Vec3ViewerX(pp)
            sum.Y += km.Vec3ViewerY(pp)
            sum.Z += km.Vec3ViewerZ(pp)
        }
    }

    return sum.X+sum.Y+sum.Z
}