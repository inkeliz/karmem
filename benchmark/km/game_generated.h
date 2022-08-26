
#include "stdint.h"
#include "stdlib.h"
#include "string.h"
#include "stdbool.h"
#include "../../c/karmem.h"

uint8_t _Null[111];


typedef uint8_t EnumColor;
const EnumColor EnumColorRed = 0UL;
const EnumColor EnumColorGreen = 1UL;
const EnumColor EnumColorBlue = 2UL;

typedef uint8_t EnumTeam;
const EnumTeam EnumTeamHumans = 0UL;
const EnumTeam EnumTeamOrcs = 1UL;
const EnumTeam EnumTeamZombies = 2UL;
const EnumTeam EnumTeamRobots = 3UL;
const EnumTeam EnumTeamAliens = 4UL;


#pragma pack(1)
typedef struct __attribute__((packed)) {
    uint8_t _data[12];
} Vec3Viewer;
#pragma options align=reset

uint32_t Vec3ViewerSize(Vec3Viewer * x) {
    return 12;
}

Vec3Viewer * NewVec3Viewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 12) == false) {
        return (Vec3Viewer *) &_Null;
    }
    Vec3Viewer * v = (Vec3Viewer *) &reader->pointer[offset];
    return v;
}

float Vec3Viewer_X(Vec3Viewer * x) {
    float r;
    memcpy(&r, &x->_data[0], 4);
    return r;
}

float Vec3Viewer_Y(Vec3Viewer * x) {
    float r;
    memcpy(&r, &x->_data[4], 4);
    return r;
}

float Vec3Viewer_Z(Vec3Viewer * x) {
    float r;
    memcpy(&r, &x->_data[8], 4);
    return r;
}

#pragma pack(1)
typedef struct __attribute__((packed)) {
    uint8_t _data[12];
} WeaponDataViewer;
#pragma options align=reset

uint32_t WeaponDataViewerSize(WeaponDataViewer * x) {
    uint32_t r;
    memcpy(&r, x, 4);
    return r;
}

WeaponDataViewer * NewWeaponDataViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 4) == false) {
        return (WeaponDataViewer *) &_Null;
    }
    WeaponDataViewer * v = (WeaponDataViewer *) &reader->pointer[offset];
    if (KarmemReaderIsValidOffset(reader, offset, WeaponDataViewerSize(v)) == false) {
        return (WeaponDataViewer *) &_Null;
    }
    return v;
}

int32_t WeaponDataViewer_Damage(WeaponDataViewer * x) {
    if ((4 + 4) > WeaponDataViewerSize(x)) {
        return 0;
    }
    int32_t r;
    memcpy(&r, &x->_data[4], 4);
    return r;
}

int32_t WeaponDataViewer_Range(WeaponDataViewer * x) {
    if ((8 + 4) > WeaponDataViewerSize(x)) {
        return 0;
    }
    int32_t r;
    memcpy(&r, &x->_data[8], 4);
    return r;
}

#pragma pack(1)
typedef struct __attribute__((packed)) {
    uint8_t _data[4];
} WeaponViewer;
#pragma options align=reset

uint32_t WeaponViewerSize(WeaponViewer * x) {
    return 4;
}

WeaponViewer * NewWeaponViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 4) == false) {
        return (WeaponViewer *) &_Null;
    }
    WeaponViewer * v = (WeaponViewer *) &reader->pointer[offset];
    return v;
}

WeaponDataViewer * WeaponViewer_Data(WeaponViewer * x, KarmemReader * reader) {
    uint32_t offset;
    memcpy(&offset, &x->_data[0], 4);
    return NewWeaponDataViewer(reader, offset);
}

#pragma pack(1)
typedef struct __attribute__((packed)) {
    uint8_t _data[111];
} MonsterDataViewer;
#pragma options align=reset

uint32_t MonsterDataViewerSize(MonsterDataViewer * x) {
    uint32_t r;
    memcpy(&r, x, 4);
    return r;
}

MonsterDataViewer * NewMonsterDataViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 4) == false) {
        return (MonsterDataViewer *) &_Null;
    }
    MonsterDataViewer * v = (MonsterDataViewer *) &reader->pointer[offset];
    if (KarmemReaderIsValidOffset(reader, offset, MonsterDataViewerSize(v)) == false) {
        return (MonsterDataViewer *) &_Null;
    }
    return v;
}

Vec3Viewer * MonsterDataViewer_Pos(MonsterDataViewer * x) {
    if ((4 + 12) > MonsterDataViewerSize(x)) {
        return (Vec3Viewer *) &_Null;
    }
        return (Vec3Viewer *) &x->_data[4];
}

int16_t MonsterDataViewer_Mana(MonsterDataViewer * x) {
    if ((16 + 2) > MonsterDataViewerSize(x)) {
        return 0;
    }
    int16_t r;
    memcpy(&r, &x->_data[16], 2);
    return r;
}

int16_t MonsterDataViewer_Health(MonsterDataViewer * x) {
    if ((18 + 2) > MonsterDataViewerSize(x)) {
        return 0;
    }
    int16_t r;
    memcpy(&r, &x->_data[18], 2);
    return r;
}
uint32_t MonsterDataViewer_NameLength(MonsterDataViewer * x, KarmemReader * reader) {
    if ((20 + 8) > MonsterDataViewerSize(x)) {
        return 0;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[20], 4);
    uint32_t size;
    memcpy(&size, &x->_data[20 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return 0;
    }
    uint32_t length = size / 1;
    if (length > 512) {
        length = 512;
    }
    return length;
}

uint8_t * MonsterDataViewer_Name(MonsterDataViewer * x, KarmemReader * reader) {
    if ((20 + 8) > MonsterDataViewerSize(x)) {
        return (uint8_t *) &_Null;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[20], 4);
    uint32_t size;
    memcpy(&size, &x->_data[20 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return (uint8_t *) &_Null;
    }
    uint32_t length = size / 1;
    return (uint8_t *) &reader->pointer[offset];
}

EnumTeam MonsterDataViewer_Team(MonsterDataViewer * x) {
    if ((28 + 1) > MonsterDataViewerSize(x)) {
        return 0;
    }
        return * (EnumTeam * ) &x->_data[28];
}
uint32_t MonsterDataViewer_InventoryLength(MonsterDataViewer * x, KarmemReader * reader) {
    if ((29 + 8) > MonsterDataViewerSize(x)) {
        return 0;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[29], 4);
    uint32_t size;
    memcpy(&size, &x->_data[29 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return 0;
    }
    uint32_t length = size / 1;
    if (length > 128) {
        length = 128;
    }
    return length;
}

uint8_t * MonsterDataViewer_Inventory(MonsterDataViewer * x, KarmemReader * reader) {
    if ((29 + 8) > MonsterDataViewerSize(x)) {
        return (uint8_t *) &_Null;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[29], 4);
    uint32_t size;
    memcpy(&size, &x->_data[29 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return (uint8_t *) &_Null;
    }
    uint32_t length = size / 1;
    return (uint8_t *) &reader->pointer[offset];
}

EnumColor MonsterDataViewer_Color(MonsterDataViewer * x) {
    if ((37 + 1) > MonsterDataViewerSize(x)) {
        return 0;
    }
        return * (EnumColor * ) &x->_data[37];
}
uint32_t MonsterDataViewer_HitboxLength(MonsterDataViewer * x) {
    if ((38 + 40) > MonsterDataViewerSize(x)) {
        return 0;
    }
    return 5;
}

double * MonsterDataViewer_Hitbox(MonsterDataViewer * x) {
    if ((38 + 40) > MonsterDataViewerSize(x)) {
        return (double *) &_Null;
    }
    return (double *) &x->_data[38];
}
uint32_t MonsterDataViewer_StatusLength(MonsterDataViewer * x, KarmemReader * reader) {
    if ((78 + 8) > MonsterDataViewerSize(x)) {
        return 0;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[78], 4);
    uint32_t size;
    memcpy(&size, &x->_data[78 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return 0;
    }
    uint32_t length = size / 4;
    if (length > 10) {
        length = 10;
    }
    return length;
}

int32_t * MonsterDataViewer_Status(MonsterDataViewer * x, KarmemReader * reader) {
    if ((78 + 8) > MonsterDataViewerSize(x)) {
        return (int32_t *) &_Null;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[78], 4);
    uint32_t size;
    memcpy(&size, &x->_data[78 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return (int32_t *) &_Null;
    }
    uint32_t length = size / 4;
    return (int32_t *) &reader->pointer[offset];
}
uint32_t MonsterDataViewer_WeaponsLength(MonsterDataViewer * x) {
    if ((86 + 16) > MonsterDataViewerSize(x)) {
        return 0;
    }
    return 4;
}

WeaponViewer * MonsterDataViewer_Weapons(MonsterDataViewer * x) {
    if ((86 + 16) > MonsterDataViewerSize(x)) {
        return (WeaponViewer *) &_Null;
    }
    return (WeaponViewer *) &x->_data[86];
}
uint32_t MonsterDataViewer_PathLength(MonsterDataViewer * x, KarmemReader * reader) {
    if ((102 + 8) > MonsterDataViewerSize(x)) {
        return 0;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[102], 4);
    uint32_t size;
    memcpy(&size, &x->_data[102 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return 0;
    }
    uint32_t length = size / 12;
    if (length > 2000) {
        length = 2000;
    }
    return length;
}

Vec3Viewer * MonsterDataViewer_Path(MonsterDataViewer * x, KarmemReader * reader) {
    if ((102 + 8) > MonsterDataViewerSize(x)) {
        return (Vec3Viewer *) &_Null;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[102], 4);
    uint32_t size;
    memcpy(&size, &x->_data[102 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return (Vec3Viewer *) &_Null;
    }
    uint32_t length = size / 12;
    return (Vec3Viewer *) &reader->pointer[offset];
}

bool MonsterDataViewer_IsAlive(MonsterDataViewer * x) {
    if ((110 + 1) > MonsterDataViewerSize(x)) {
        return false;
    }
    bool r;
    memcpy(&r, &x->_data[110], 1);
    return r;
}

#pragma pack(1)
typedef struct __attribute__((packed)) {
    uint8_t _data[4];
} MonsterViewer;
#pragma options align=reset

uint32_t MonsterViewerSize(MonsterViewer * x) {
    return 4;
}

MonsterViewer * NewMonsterViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 4) == false) {
        return (MonsterViewer *) &_Null;
    }
    MonsterViewer * v = (MonsterViewer *) &reader->pointer[offset];
    return v;
}

MonsterDataViewer * MonsterViewer_Data(MonsterViewer * x, KarmemReader * reader) {
    uint32_t offset;
    memcpy(&offset, &x->_data[0], 4);
    return NewMonsterDataViewer(reader, offset);
}

#pragma pack(1)
typedef struct __attribute__((packed)) {
    uint8_t _data[12];
} MonstersViewer;
#pragma options align=reset

uint32_t MonstersViewerSize(MonstersViewer * x) {
    uint32_t r;
    memcpy(&r, x, 4);
    return r;
}

MonstersViewer * NewMonstersViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 4) == false) {
        return (MonstersViewer *) &_Null;
    }
    MonstersViewer * v = (MonstersViewer *) &reader->pointer[offset];
    if (KarmemReaderIsValidOffset(reader, offset, MonstersViewerSize(v)) == false) {
        return (MonstersViewer *) &_Null;
    }
    return v;
}
uint32_t MonstersViewer_MonstersLength(MonstersViewer * x, KarmemReader * reader) {
    if ((4 + 8) > MonstersViewerSize(x)) {
        return 0;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[4], 4);
    uint32_t size;
    memcpy(&size, &x->_data[4 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return 0;
    }
    uint32_t length = size / 4;
    if (length > 2000) {
        length = 2000;
    }
    return length;
}

MonsterViewer * MonstersViewer_Monsters(MonstersViewer * x, KarmemReader * reader) {
    if ((4 + 8) > MonstersViewerSize(x)) {
        return (MonsterViewer *) &_Null;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[4], 4);
    uint32_t size;
    memcpy(&size, &x->_data[4 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return (MonsterViewer *) &_Null;
    }
    uint32_t length = size / 4;
    return (MonsterViewer *) &reader->pointer[offset];
}
typedef uint64_t EnumPacketIdentifier;

const EnumPacketIdentifier EnumPacketIdentifierVec3 = 10268726485798425099UL;
const EnumPacketIdentifier EnumPacketIdentifierWeaponData = 15342010214468761012UL;
const EnumPacketIdentifier EnumPacketIdentifierWeapon = 8029074423243608167UL;
const EnumPacketIdentifier EnumPacketIdentifierMonsterData = 12254962724431809041UL;
const EnumPacketIdentifier EnumPacketIdentifierMonster = 5593793986513565154UL;
const EnumPacketIdentifier EnumPacketIdentifierMonsters = 14096677544474027661UL;

    
typedef struct {
    float X;
    float Y;
    float Z;
} Vec3;

EnumPacketIdentifier Vec3PacketIdentifier(Vec3 * x) {
    return EnumPacketIdentifierVec3;
}

uint32_t Vec3Write(Vec3 * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 12;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t __XOffset = offset + 0;
    KarmemWriterWriteAt(writer, __XOffset, (void *) &x->X, 4);
    uint32_t __YOffset = offset + 4;
    KarmemWriterWriteAt(writer, __YOffset, (void *) &x->Y, 4);
    uint32_t __ZOffset = offset + 8;
    KarmemWriterWriteAt(writer, __ZOffset, (void *) &x->Z, 4);

    return offset;
}

uint32_t Vec3WriteAsRoot(Vec3 * x, KarmemWriter * writer) {
    return Vec3Write(x, writer, 0);
}

void Vec3Read(Vec3 * x, Vec3Viewer * viewer, KarmemReader * reader) {
    x->X = Vec3Viewer_X(viewer);
    x->Y = Vec3Viewer_Y(viewer);
    x->Z = Vec3Viewer_Z(viewer);
}

void Vec3ReadAsRoot(Vec3 * x, KarmemReader * reader) {
    return Vec3Read(x, NewVec3Viewer(reader, 0), reader);
}

void Vec3Reset(Vec3 * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 111);
    Vec3Read(x, (Vec3Viewer *) &_Null, &reader);
}

Vec3 NewVec3() {
    Vec3 r;
    memset(&r, 0, sizeof(Vec3));
    return r;
}
typedef struct {
    int32_t Damage;
    int32_t Range;
} WeaponData;

EnumPacketIdentifier WeaponDataPacketIdentifier(WeaponData * x) {
    return EnumPacketIdentifierWeaponData;
}

uint32_t WeaponDataWrite(WeaponData * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 12;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t sizeData = 12;
    KarmemWriterWriteAt(writer, offset, (void *)&sizeData, 4);
    uint32_t __DamageOffset = offset + 4;
    KarmemWriterWriteAt(writer, __DamageOffset, (void *) &x->Damage, 4);
    uint32_t __RangeOffset = offset + 8;
    KarmemWriterWriteAt(writer, __RangeOffset, (void *) &x->Range, 4);

    return offset;
}

uint32_t WeaponDataWriteAsRoot(WeaponData * x, KarmemWriter * writer) {
    return WeaponDataWrite(x, writer, 0);
}

void WeaponDataRead(WeaponData * x, WeaponDataViewer * viewer, KarmemReader * reader) {
    x->Damage = WeaponDataViewer_Damage(viewer);
    x->Range = WeaponDataViewer_Range(viewer);
}

void WeaponDataReadAsRoot(WeaponData * x, KarmemReader * reader) {
    return WeaponDataRead(x, NewWeaponDataViewer(reader, 0), reader);
}

void WeaponDataReset(WeaponData * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 111);
    WeaponDataRead(x, (WeaponDataViewer *) &_Null, &reader);
}

WeaponData NewWeaponData() {
    WeaponData r;
    memset(&r, 0, sizeof(WeaponData));
    return r;
}
typedef struct {
    WeaponData Data;
} Weapon;

EnumPacketIdentifier WeaponPacketIdentifier(Weapon * x) {
    return EnumPacketIdentifierWeapon;
}

uint32_t WeaponWrite(Weapon * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 4;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t __DataSize = 12;
    uint32_t __DataOffset = KarmemWriterAlloc(writer, __DataSize);

    KarmemWriterWriteAt(writer, offset+0, (void *) &__DataOffset, 4);
    if (WeaponDataWrite(&x->Data, writer, __DataOffset) == 0) {
        return 0;
    }

    return offset;
}

uint32_t WeaponWriteAsRoot(Weapon * x, KarmemWriter * writer) {
    return WeaponWrite(x, writer, 0);
}

void WeaponRead(Weapon * x, WeaponViewer * viewer, KarmemReader * reader) {
    WeaponDataRead(&x->Data, WeaponViewer_Data(viewer, reader), reader);
}

void WeaponReadAsRoot(Weapon * x, KarmemReader * reader) {
    return WeaponRead(x, NewWeaponViewer(reader, 0), reader);
}

void WeaponReset(Weapon * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 111);
    WeaponRead(x, (WeaponViewer *) &_Null, &reader);
}

Weapon NewWeapon() {
    Weapon r;
    memset(&r, 0, sizeof(Weapon));
    return r;
}
typedef struct {
    Vec3 Pos;
    int16_t Mana;
    int16_t Health;
    uint8_t * Name;
    uint32_t _Name_len;
    uint32_t _Name_cap;
    EnumTeam Team;
    uint8_t * Inventory;
    uint32_t _Inventory_len;
    uint32_t _Inventory_cap;
    EnumColor Color;
    double Hitbox[5];
    int32_t * Status;
    uint32_t _Status_len;
    uint32_t _Status_cap;
    Weapon Weapons[4];
    Vec3 * Path;
    uint32_t _Path_len;
    uint32_t _Path_cap;
    bool IsAlive;
} MonsterData;

EnumPacketIdentifier MonsterDataPacketIdentifier(MonsterData * x) {
    return EnumPacketIdentifierMonsterData;
}

uint32_t MonsterDataWrite(MonsterData * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 111;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t sizeData = 111;
    KarmemWriterWriteAt(writer, offset, (void *)&sizeData, 4);
    uint32_t __PosOffset = offset + 4;
    if (Vec3Write(&x->Pos, writer, __PosOffset) == 0) {
        return 0;
    }
    uint32_t __ManaOffset = offset + 16;
    KarmemWriterWriteAt(writer, __ManaOffset, (void *) &x->Mana, 2);
    uint32_t __HealthOffset = offset + 18;
    KarmemWriterWriteAt(writer, __HealthOffset, (void *) &x->Health, 2);
    uint32_t __NameSize = 1 * x->_Name_len;
    uint32_t __NameOffset = KarmemWriterAlloc(writer, __NameSize);

    KarmemWriterWriteAt(writer, offset+20, (void *) &__NameOffset, 4);
    KarmemWriterWriteAt(writer, offset+20+4, (void *) &__NameSize, 4);
    KarmemWriterWriteAt(writer, __NameOffset, (void *) x->Name, __NameSize);
    uint32_t __TeamOffset = offset + 28;
    KarmemWriterWriteAt(writer, __TeamOffset, (void *) &x->Team, 1);
    uint32_t __InventorySize = 1 * x->_Inventory_len;
    uint32_t __InventoryOffset = KarmemWriterAlloc(writer, __InventorySize);

    KarmemWriterWriteAt(writer, offset+29, (void *) &__InventoryOffset, 4);
    KarmemWriterWriteAt(writer, offset+29+4, (void *) &__InventorySize, 4);
    KarmemWriterWriteAt(writer, __InventoryOffset, (void *) x->Inventory, __InventorySize);
    uint32_t __ColorOffset = offset + 37;
    KarmemWriterWriteAt(writer, __ColorOffset, (void *) &x->Color, 1);
    uint32_t __HitboxSize = 40;
    uint32_t __HitboxOffset = offset + 38;
    KarmemWriterWriteAt(writer, __HitboxOffset,(void *) x->Hitbox, __HitboxSize);
    uint32_t __StatusSize = 4 * x->_Status_len;
    uint32_t __StatusOffset = KarmemWriterAlloc(writer, __StatusSize);

    KarmemWriterWriteAt(writer, offset+78, (void *) &__StatusOffset, 4);
    KarmemWriterWriteAt(writer, offset+78+4, (void *) &__StatusSize, 4);
    KarmemWriterWriteAt(writer, __StatusOffset, (void *) x->Status, __StatusSize);
    uint32_t __WeaponsSize = 16;
    uint32_t __WeaponsOffset = offset + 86;
    uint32_t __WeaponsIndex = 0;
    uint32_t __WeaponsEnd = __WeaponsOffset + __WeaponsSize;
    while (__WeaponsOffset < __WeaponsEnd) {
        if (WeaponWrite(&x->Weapons[__WeaponsIndex], writer, __WeaponsOffset) == 0) {
            return 0;
        }
        __WeaponsOffset = __WeaponsOffset + 4;
        __WeaponsIndex = __WeaponsIndex + 1;
    }
    uint32_t __PathSize = 12 * x->_Path_len;
    uint32_t __PathOffset = KarmemWriterAlloc(writer, __PathSize);

    KarmemWriterWriteAt(writer, offset+102, (void *) &__PathOffset, 4);
    KarmemWriterWriteAt(writer, offset+102+4, (void *) &__PathSize, 4);
    uint32_t __PathIndex = 0;
    uint32_t __PathEnd = __PathOffset + __PathSize;
    while (__PathOffset < __PathEnd) {
        if (Vec3Write(&x->Path[__PathIndex], writer, __PathOffset) == 0) {
            return 0;
        }
        __PathOffset = __PathOffset + 12;
        __PathIndex = __PathIndex + 1;
    }
    uint32_t __IsAliveOffset = offset + 110;
    KarmemWriterWriteAt(writer, __IsAliveOffset, (void *) &x->IsAlive, 1);

    return offset;
}

uint32_t MonsterDataWriteAsRoot(MonsterData * x, KarmemWriter * writer) {
    return MonsterDataWrite(x, writer, 0);
}

void MonsterDataRead(MonsterData * x, MonsterDataViewer * viewer, KarmemReader * reader) {
    Vec3Read(&x->Pos, MonsterDataViewer_Pos(viewer), reader);
    x->Mana = MonsterDataViewer_Mana(viewer);
    x->Health = MonsterDataViewer_Health(viewer);
    uint8_t * __NameSlice = MonsterDataViewer_Name(viewer, reader);
    uint32_t __NameLen = MonsterDataViewer_NameLength(viewer, reader);
    if (__NameLen > x->_Name_cap) {
        uint32_t __NameCapacityTarget = __NameLen;
        x->Name = (uint8_t *) realloc(x->Name, __NameCapacityTarget);
        uint32_t __NameNewIndex = x->_Name_cap;
        while (__NameNewIndex < __NameCapacityTarget) {
            x->Name[__NameNewIndex] = 0;
            __NameNewIndex = __NameNewIndex + 1;
        }
        x->_Name_cap = __NameCapacityTarget;
    }
    if (__NameLen > x->_Name_len) {
        x->_Name_len = __NameLen;
    }
    uint32_t __NameIndex = 0;
    while (__NameIndex < __NameLen) {
        x->Name[__NameIndex] = __NameSlice[__NameIndex];
        __NameIndex = __NameIndex + 1;
    }
    x->_Name_len = __NameLen;
    x->Team = MonsterDataViewer_Team(viewer);
    uint8_t * __InventorySlice = MonsterDataViewer_Inventory(viewer, reader);
    uint32_t __InventoryLen = MonsterDataViewer_InventoryLength(viewer, reader);
    if (__InventoryLen > x->_Inventory_cap) {
        uint32_t __InventoryCapacityTarget = __InventoryLen;
        x->Inventory = (uint8_t *) realloc(x->Inventory, __InventoryCapacityTarget * sizeof(uint8_t));
        uint32_t __InventoryNewIndex = x->_Inventory_cap;
        while (__InventoryNewIndex < __InventoryCapacityTarget) {
            x->Inventory[__InventoryNewIndex] = 0;
            __InventoryNewIndex = __InventoryNewIndex + 1;
        }
        x->_Inventory_cap = __InventoryCapacityTarget;
    }
    if (__InventoryLen > x->_Inventory_len) {
        x->_Inventory_len = __InventoryLen;
    }
    uint32_t __InventoryIndex = 0;
    while (__InventoryIndex < __InventoryLen) {
        x->Inventory[__InventoryIndex] = __InventorySlice[__InventoryIndex];
        __InventoryIndex = __InventoryIndex + 1;
    }
    x->_Inventory_len = __InventoryLen;
    x->Color = MonsterDataViewer_Color(viewer);
    double * __HitboxSlice = MonsterDataViewer_Hitbox(viewer);
    uint32_t __HitboxLen = MonsterDataViewer_HitboxLength(viewer);
    if (__HitboxLen > 5) {
        __HitboxLen = 5;
    }
    uint32_t __HitboxIndex = 0;
    while (__HitboxIndex < __HitboxLen) {
        x->Hitbox[__HitboxIndex] = __HitboxSlice[__HitboxIndex];
        __HitboxIndex = __HitboxIndex + 1;
    }
    while (__HitboxIndex < 5) {
        x->Hitbox[__HitboxIndex] = 0;
        __HitboxIndex = __HitboxIndex + 1;
    }
    int32_t * __StatusSlice = MonsterDataViewer_Status(viewer, reader);
    uint32_t __StatusLen = MonsterDataViewer_StatusLength(viewer, reader);
    if (__StatusLen > x->_Status_cap) {
        uint32_t __StatusCapacityTarget = __StatusLen;
        x->Status = (int32_t *) realloc(x->Status, __StatusCapacityTarget * sizeof(int32_t));
        uint32_t __StatusNewIndex = x->_Status_cap;
        while (__StatusNewIndex < __StatusCapacityTarget) {
            x->Status[__StatusNewIndex] = 0;
            __StatusNewIndex = __StatusNewIndex + 1;
        }
        x->_Status_cap = __StatusCapacityTarget;
    }
    if (__StatusLen > x->_Status_len) {
        x->_Status_len = __StatusLen;
    }
    uint32_t __StatusIndex = 0;
    while (__StatusIndex < __StatusLen) {
        x->Status[__StatusIndex] = __StatusSlice[__StatusIndex];
        __StatusIndex = __StatusIndex + 1;
    }
    x->_Status_len = __StatusLen;
    WeaponViewer * __WeaponsSlice = MonsterDataViewer_Weapons(viewer);
    uint32_t __WeaponsLen = MonsterDataViewer_WeaponsLength(viewer);
    if (__WeaponsLen > 4) {
        __WeaponsLen = 4;
    }
    uint32_t __WeaponsIndex = 0;
    while (__WeaponsIndex < __WeaponsLen) {
        WeaponRead(&x->Weapons[__WeaponsIndex], &__WeaponsSlice[__WeaponsIndex], reader);
         __WeaponsIndex = __WeaponsIndex + 1;
    }
    while (__WeaponsIndex < 4) {
        WeaponReset(&x->Weapons[__WeaponsIndex]);
        __WeaponsIndex = __WeaponsIndex + 1;
    }
    Vec3Viewer * __PathSlice = MonsterDataViewer_Path(viewer, reader);
    uint32_t __PathLen = MonsterDataViewer_PathLength(viewer, reader);
    if (__PathLen > x->_Path_cap) {
        uint32_t __PathCapacityTarget = __PathLen;
        x->Path = (Vec3 *) realloc(x->Path, __PathCapacityTarget * sizeof(Vec3));
        uint32_t __PathNewIndex = x->_Path_cap;
        while (__PathNewIndex < __PathCapacityTarget) {
            x->Path[__PathNewIndex] = NewVec3();
            __PathNewIndex = __PathNewIndex + 1;
        }
        x->_Path_cap = __PathCapacityTarget;
    }
    if (__PathLen > x->_Path_len) {
        x->_Path_len = __PathLen;
    }
    uint32_t __PathIndex = 0;
    while (__PathIndex < __PathLen) {
        Vec3Read(&x->Path[__PathIndex], &__PathSlice[__PathIndex], reader);
         __PathIndex = __PathIndex + 1;
    }
    x->_Path_len = __PathLen;
    x->IsAlive = MonsterDataViewer_IsAlive(viewer);
}

void MonsterDataReadAsRoot(MonsterData * x, KarmemReader * reader) {
    return MonsterDataRead(x, NewMonsterDataViewer(reader, 0), reader);
}

void MonsterDataReset(MonsterData * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 111);
    MonsterDataRead(x, (MonsterDataViewer *) &_Null, &reader);
}

MonsterData NewMonsterData() {
    MonsterData r;
    memset(&r, 0, sizeof(MonsterData));
    return r;
}
typedef struct {
    MonsterData Data;
} Monster;

EnumPacketIdentifier MonsterPacketIdentifier(Monster * x) {
    return EnumPacketIdentifierMonster;
}

uint32_t MonsterWrite(Monster * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 4;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t __DataSize = 111;
    uint32_t __DataOffset = KarmemWriterAlloc(writer, __DataSize);

    KarmemWriterWriteAt(writer, offset+0, (void *) &__DataOffset, 4);
    if (MonsterDataWrite(&x->Data, writer, __DataOffset) == 0) {
        return 0;
    }

    return offset;
}

uint32_t MonsterWriteAsRoot(Monster * x, KarmemWriter * writer) {
    return MonsterWrite(x, writer, 0);
}

void MonsterRead(Monster * x, MonsterViewer * viewer, KarmemReader * reader) {
    MonsterDataRead(&x->Data, MonsterViewer_Data(viewer, reader), reader);
}

void MonsterReadAsRoot(Monster * x, KarmemReader * reader) {
    return MonsterRead(x, NewMonsterViewer(reader, 0), reader);
}

void MonsterReset(Monster * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 111);
    MonsterRead(x, (MonsterViewer *) &_Null, &reader);
}

Monster NewMonster() {
    Monster r;
    memset(&r, 0, sizeof(Monster));
    return r;
}
typedef struct {
    Monster * Monsters;
    uint32_t _Monsters_len;
    uint32_t _Monsters_cap;
} Monsters;

EnumPacketIdentifier MonstersPacketIdentifier(Monsters * x) {
    return EnumPacketIdentifierMonsters;
}

uint32_t MonstersWrite(Monsters * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 12;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t sizeData = 12;
    KarmemWriterWriteAt(writer, offset, (void *)&sizeData, 4);
    uint32_t __MonstersSize = 4 * x->_Monsters_len;
    uint32_t __MonstersOffset = KarmemWriterAlloc(writer, __MonstersSize);

    KarmemWriterWriteAt(writer, offset+4, (void *) &__MonstersOffset, 4);
    KarmemWriterWriteAt(writer, offset+4+4, (void *) &__MonstersSize, 4);
    uint32_t __MonstersIndex = 0;
    uint32_t __MonstersEnd = __MonstersOffset + __MonstersSize;
    while (__MonstersOffset < __MonstersEnd) {
        if (MonsterWrite(&x->Monsters[__MonstersIndex], writer, __MonstersOffset) == 0) {
            return 0;
        }
        __MonstersOffset = __MonstersOffset + 4;
        __MonstersIndex = __MonstersIndex + 1;
    }

    return offset;
}

uint32_t MonstersWriteAsRoot(Monsters * x, KarmemWriter * writer) {
    return MonstersWrite(x, writer, 0);
}

void MonstersRead(Monsters * x, MonstersViewer * viewer, KarmemReader * reader) {
    MonsterViewer * __MonstersSlice = MonstersViewer_Monsters(viewer, reader);
    uint32_t __MonstersLen = MonstersViewer_MonstersLength(viewer, reader);
    if (__MonstersLen > x->_Monsters_cap) {
        uint32_t __MonstersCapacityTarget = __MonstersLen;
        x->Monsters = (Monster *) realloc(x->Monsters, __MonstersCapacityTarget * sizeof(Monster));
        uint32_t __MonstersNewIndex = x->_Monsters_cap;
        while (__MonstersNewIndex < __MonstersCapacityTarget) {
            x->Monsters[__MonstersNewIndex] = NewMonster();
            __MonstersNewIndex = __MonstersNewIndex + 1;
        }
        x->_Monsters_cap = __MonstersCapacityTarget;
    }
    if (__MonstersLen > x->_Monsters_len) {
        x->_Monsters_len = __MonstersLen;
    }
    uint32_t __MonstersIndex = 0;
    while (__MonstersIndex < __MonstersLen) {
        MonsterRead(&x->Monsters[__MonstersIndex], &__MonstersSlice[__MonstersIndex], reader);
         __MonstersIndex = __MonstersIndex + 1;
    }
    x->_Monsters_len = __MonstersLen;
}

void MonstersReadAsRoot(Monsters * x, KarmemReader * reader) {
    return MonstersRead(x, NewMonstersViewer(reader, 0), reader);
}

void MonstersReset(Monsters * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 111);
    MonstersRead(x, (MonstersViewer *) &_Null, &reader);
}

Monsters NewMonsters() {
    Monsters r;
    memset(&r, 0, sizeof(Monsters));
    return r;
}
