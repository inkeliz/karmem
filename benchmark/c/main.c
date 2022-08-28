#include "stdint.h"
#include "stdlib.h"
#include "string.h"
#include "stdio.h"
#include "../km/game_generated.h"

uint8_t *InputMemory;
uint8_t *OutputMemory;

__attribute__((export_name("InputMemoryPointer")))
uint32_t InputMemoryPointer() {
    if (InputMemory == NULL) {
        return 0xFFFFFFFF;
    }
    return *(uint32_t *) &InputMemory;
}

__attribute__((export_name("OutputMemoryPointer")))
uint32_t OutputMemoryPointer() {
    if (OutputMemory == NULL) {
        return 0xFFFFFFFF;
    }
    return *(uint32_t *) &OutputMemory;
}

Monsters _Struct;
KarmemWriter _Writer;
KarmemReader _Reader;

__attribute__((export_name("_start")))
void _start() {
    OutputMemory = (uint8_t *) malloc(8000000);
    InputMemory = (uint8_t *) malloc(8000000);
    _Struct = NewMonsters();
    _Writer = KarmemNewFixedWriter(OutputMemory, 8000000);
    _Reader = KarmemNewReader(InputMemory, 8000000);
}

__attribute__((export_name("KBenchmarkEncodeObjectAPI")))
void KBenchmarkEncodeObjectAPI() {
    KarmemWriterReset(&_Writer);
    MonstersWriteAsRoot(&_Struct, &_Writer);
}

__attribute__((export_name("KBenchmarkDecodeObjectAPI")))
void KBenchmarkDecodeObjectAPI(uint32_t size) {
    KarmemReaderSetSize(&_Reader, size);
    MonstersReadAsRoot(&_Struct, &_Reader);
}

__attribute__((export_name("KBenchmarkDecodeSumVec3")))
float KBenchmarkDecodeSumVec3(uint32_t size) {
    float sum = 0;

    KarmemReaderSetSize(&_Reader, size);

    MonstersViewer * monsters = NewMonstersViewer(&_Reader, 0);

    uint32_t il = MonstersViewer_MonstersLength(monsters, &_Reader);
    uint32_t i = 0;

    MonsterViewer * monsterList = MonstersViewer_Monsters(monsters, &_Reader);
    while (i < il) {
            MonsterDataViewer * data = MonsterViewer_Data(&monsterList[i], &_Reader);

            uint32_t jl = MonsterDataViewer_PathLength(data, &_Reader);
            uint32_t j = 0;

            Vec3Viewer * path = MonsterDataViewer_Path(data, &_Reader);
            while (j < jl) {
                Vec3Viewer * pp = &path[j];
                sum += Vec3Viewer_X(pp) + Vec3Viewer_Y(pp) + Vec3Viewer_Z(pp);
                j++;
            }
        i++;
    }

    return sum;
}

__attribute__((export_name("KBenchmarkDecodeSumUint8")))
uint32_t KBenchmarkDecodeSumUint8(uint32_t size) {
    uint32_t sum = 0;

    KarmemReaderSetSize(&_Reader, size);

    MonstersViewer * monsters = NewMonstersViewer(&_Reader, 0);

    uint32_t il = MonstersViewer_MonstersLength(monsters, &_Reader);
    uint32_t i = 0;

    MonsterViewer * monsterList = MonstersViewer_Monsters(monsters, &_Reader);
    while (i < il) {
            MonsterDataViewer * data = MonsterViewer_Data(&monsterList[i], &_Reader);

            uint32_t jl = MonsterDataViewer_InventoryLength(data, &_Reader);
            uint32_t j = 0;

            uint8_t * inv = MonsterDataViewer_Inventory(data, &_Reader);
            while (j < jl) {
                sum += (uint32_t)(inv[j]);
                j++;
            }
        i++;
    }

    return sum;
}

__attribute__((export_name("KBenchmarkDecodeSumStats")))
uint32_t KBenchmarkDecodeSumStats(uint32_t size) {
    WeaponData sum = NewWeaponData();

    KarmemReaderSetSize(&_Reader, size);

    MonstersViewer * monsters = NewMonstersViewer(&_Reader, 0);

    uint32_t il = MonstersViewer_MonstersLength(monsters, &_Reader);
    uint32_t i = 0;

    MonsterViewer * monsterList = MonstersViewer_Monsters(monsters, &_Reader);
    while (i < il) {
            MonsterDataViewer * data = MonsterViewer_Data(&monsterList[i], &_Reader);

            uint32_t jl = MonsterDataViewer_WeaponsLength(data);
            uint32_t j = 0;

            WeaponViewer * weapons = MonsterDataViewer_Weapons(data);
            while (j < jl) {
                WeaponDataViewer * d = WeaponViewer_Data(&weapons[j], &_Reader);
                sum.Ammo += WeaponDataViewer_Ammo(d);
                sum.Damage += WeaponDataViewer_Damage(d);
                sum.ClipSize += WeaponDataViewer_ClipSize(d);
                sum.ReloadTime += WeaponDataViewer_ReloadTime(d);
                sum.Range += WeaponDataViewer_Range(d);
                j++;
            }
        i++;
    }


    KarmemWriterReset(&_Writer);
    WeaponDataWriteAsRoot(&sum, &_Writer);
    return (uint32_t)_Writer.length;
}

__attribute__((export_name("KNOOP")))
uint32_t KNOOP() {
    return 42;
}