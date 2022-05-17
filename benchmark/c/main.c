#include "stdint.h"
#include "stdlib.h"
#include "string.h"
#include "stdio.h"
#include "../km/game_generated.h"

uint8_t *InputMemory;
uint8_t *OutputMemory;

uint32_t InputMemoryPointer() {
    if (InputMemory == NULL) {
        return 0xFFFFFFFF;
    }
    return *(uint32_t *) &InputMemory;
}

uint32_t OutputMemoryPointer() {
    if (OutputMemory == NULL) {
        return 0xFFFFFFFF;
    }
    return *(uint32_t *) &OutputMemory;
}

Monsters _Struct;
KarmemWriter _Writer;
KarmemReader _Reader;

void _start() {
    OutputMemory = (uint8_t *) malloc(8000000);
    InputMemory = (uint8_t *) malloc(8000000);
    _Struct = NewMonsters();
    _Writer = KarmemNewFixedWriter(OutputMemory, 8000000);
    _Reader = KarmemNewReader(InputMemory, 8000000);
}

void KBenchmarkEncodeObjectAPI() {
    KarmemWriterReset(&_Writer);
    MonstersWriteAsRoot(&_Struct, &_Writer);
}

void KBenchmarkDecodeObjectAPI(uint32_t size) {
    KarmemReaderSetSize(&_Reader, size);
    MonstersReadAsRoot(&_Struct, &_Reader);
}

float KBenchmarkDecodeSumVec3(uint32_t size) {
    Vec3 sum = NewVec3();

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
                sum.X += Vec3Viewer_X(&path[j]);
                sum.Y += Vec3Viewer_Y(&path[j]);
                sum.Z += Vec3Viewer_Z(&path[j]);
                j++;
            }
        i++;
    }

    return sum.X + sum.Y + sum.Z;
}