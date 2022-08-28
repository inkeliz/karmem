
using global::System.Reflection;
using global::System.Runtime.CompilerServices;
using global::System.Runtime.InteropServices;
using global::Karmem;

namespace km;

internal static unsafe class _Globals
{
    private static long _largest = 111;
    private static void* _null = null;
    private static Karmem.Reader? _nullReader = null;

    public static void* Null()
    {
        if (_null == null)
        {
            var n = Marshal.AllocHGlobal((int)_largest);
            Unsafe.InitBlockUnaligned(n.ToPointer(), 0, (uint)_largest);
            _null = n.ToPointer();
        }
        return _null;
    }
    public static Karmem.Reader NullReader()
    {
        _nullReader ??= Karmem.Reader.NewReader(new IntPtr(Null()), _largest, _largest);
        return _nullReader.Value;
    }
}

public enum Color : byte {
    Red = 0,
    Green = 1,
    Blue = 2,
}
    
public enum Team : byte {
    Humans = 0,
    Orcs = 1,
    Zombies = 2,
    Robots = 3,
    Aliens = 4,
}
    
public enum PacketIdentifier : ulong {
    Vec3 = 10268726485798425099,
    WeaponData = 15342010214468761012,
    Weapon = 8029074423243608167,
    MonsterData = 12254962724431809041,
    Monster = 5593793986513565154,
    Monsters = 14096677544474027661,
}
    
public unsafe struct Vec3 {
    public float _X = 0;
    public float _Y = 0;
    public float _Z = 0;

    public Vec3() {}

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static Vec3 NewVec3() {
        return new Vec3();
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static PacketIdentifier GetPacketIdentifier() {
        return PacketIdentifier.Vec3;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Reset() {
        this.ReadAsRoot(_Globals.NullReader());
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool WriteAsRoot(Karmem.Writer writer) {
        return this.Write(writer, 0);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool Write(Karmem.Writer writer, uint start) {
        var offset = start;
        var size = (uint)12;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == uint.MaxValue) {
                return false;
            }
        }
        var __XOffset = offset+0;
        writer.WriteAt(__XOffset, this._X);
        var __YOffset = offset+4;
        writer.WriteAt(__YOffset, this._Y);
        var __ZOffset = offset+8;
        writer.WriteAt(__ZOffset, this._Z);

        return true;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void ReadAsRoot(Karmem.Reader reader) {
        this.Read(Vec3Viewer.NewVec3Viewer(reader, 0), reader);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Read(Vec3Viewer viewer, Karmem.Reader reader) {
        this._X = viewer.X();
        this._Y = viewer.Y();
        this._Z = viewer.Z();
    }
}
public unsafe struct WeaponData {
    public int _Damage = 0;
    public ushort _Ammo = 0;
    public byte _ClipSize = 0;
    public float _ReloadTime = 0;
    public int _Range = 0;

    public WeaponData() {}

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static WeaponData NewWeaponData() {
        return new WeaponData();
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static PacketIdentifier GetPacketIdentifier() {
        return PacketIdentifier.WeaponData;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Reset() {
        this.ReadAsRoot(_Globals.NullReader());
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool WriteAsRoot(Karmem.Writer writer) {
        return this.Write(writer, 0);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool Write(Karmem.Writer writer, uint start) {
        var offset = start;
        var size = (uint)19;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == uint.MaxValue) {
                return false;
            }
        }
        writer.WriteAt(offset, (uint)19);
        var __DamageOffset = offset+4;
        writer.WriteAt(__DamageOffset, this._Damage);
        var __AmmoOffset = offset+8;
        writer.WriteAt(__AmmoOffset, this._Ammo);
        var __ClipSizeOffset = offset+10;
        writer.WriteAt(__ClipSizeOffset, this._ClipSize);
        var __ReloadTimeOffset = offset+11;
        writer.WriteAt(__ReloadTimeOffset, this._ReloadTime);
        var __RangeOffset = offset+15;
        writer.WriteAt(__RangeOffset, this._Range);

        return true;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void ReadAsRoot(Karmem.Reader reader) {
        this.Read(WeaponDataViewer.NewWeaponDataViewer(reader, 0), reader);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Read(WeaponDataViewer viewer, Karmem.Reader reader) {
        this._Damage = viewer.Damage();
        this._Ammo = viewer.Ammo();
        this._ClipSize = viewer.ClipSize();
        this._ReloadTime = viewer.ReloadTime();
        this._Range = viewer.Range();
    }
}
public unsafe struct Weapon {
    public WeaponData _Data = new WeaponData();

    public Weapon() {}

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static Weapon NewWeapon() {
        return new Weapon();
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static PacketIdentifier GetPacketIdentifier() {
        return PacketIdentifier.Weapon;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Reset() {
        this.ReadAsRoot(_Globals.NullReader());
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool WriteAsRoot(Karmem.Writer writer) {
        return this.Write(writer, 0);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool Write(Karmem.Writer writer, uint start) {
        var offset = start;
        var size = (uint)4;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == uint.MaxValue) {
                return false;
            }
        }
        var __DataSize = (uint)19;
        var __DataOffset = writer.Alloc(__DataSize);
        if (offset == uint.MaxValue) {
            return false;
        }
        writer.WriteAt(offset+0, (uint)__DataOffset);
            if (!this._Data.Write(writer, __DataOffset)) {
                return false;
            }

        return true;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void ReadAsRoot(Karmem.Reader reader) {
        this.Read(WeaponViewer.NewWeaponViewer(reader, 0), reader);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Read(WeaponViewer viewer, Karmem.Reader reader) {
        this._Data.Read(viewer.Data(reader), reader);
    }
}
public unsafe struct MonsterData {
    public Vec3 _Pos = new Vec3();
    public short _Mana = 0;
    public short _Health = 0;
    public string _Name = "";
    public Team _Team = 0;
    public List<byte> _Inventory = new List<byte>();
    public Color _Color = 0;
    public List<double> _Hitbox = new List<double>();
    public List<int> _Status = new List<int>();
    public List<Weapon> _Weapons = new List<Weapon>();
    public List<Vec3> _Path = new List<Vec3>();
    public bool _IsAlive = false;

    public MonsterData() {}

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static MonsterData NewMonsterData() {
        return new MonsterData();
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static PacketIdentifier GetPacketIdentifier() {
        return PacketIdentifier.MonsterData;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Reset() {
        this.ReadAsRoot(_Globals.NullReader());
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool WriteAsRoot(Karmem.Writer writer) {
        return this.Write(writer, 0);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool Write(Karmem.Writer writer, uint start) {
        var offset = start;
        var size = (uint)111;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == uint.MaxValue) {
                return false;
            }
        }
        writer.WriteAt(offset, (uint)111);
        var __PosOffset = offset+4;
            if (!this._Pos.Write(writer, __PosOffset)) {
                return false;
            }
        var __ManaOffset = offset+16;
        writer.WriteAt(__ManaOffset, this._Mana);
        var __HealthOffset = offset+18;
        writer.WriteAt(__HealthOffset, this._Health);
        var __NameSize = (uint)(4 * this._Name.Length);
        var __NameOffset = writer.Alloc(__NameSize);
        if (offset == uint.MaxValue) {
            return false;
        }
        writer.WriteAt(offset+20, (uint)__NameOffset);
        var __NameStringSize = writer.WriteAt(__NameOffset, this._Name);
        writer.WriteAt(offset+20 + 4, (uint)__NameStringSize);
        var __TeamOffset = offset+28;
        writer.WriteAt(__TeamOffset, (long)this._Team);
        var __InventorySize = (uint)(1 * this._Inventory.Count);
        var __InventoryOffset = writer.Alloc(__InventorySize);
        if (offset == uint.MaxValue) {
            return false;
        }
        writer.WriteAt(offset+29, (uint)__InventoryOffset);
        writer.WriteAt(offset+29 + 4, (uint)__InventorySize);
        for (var i = 0; i < this._Inventory.Count; i++) {
            writer.WriteAt(__InventoryOffset, this._Inventory[i]);
            __InventoryOffset += 1;
        }
        var __ColorOffset = offset+37;
        writer.WriteAt(__ColorOffset, (long)this._Color);
        var __HitboxOffset = offset+38;
        for (var i = 0; i < 5; i++) {
            if (i < this._Hitbox.Count) {
                writer.WriteAt(__HitboxOffset, this._Hitbox[i]);
            } else {
                writer.WriteAt(__HitboxOffset, 0);
            }
            __HitboxOffset += 8;
        }
        var __StatusSize = (uint)(4 * this._Status.Count);
        var __StatusOffset = writer.Alloc(__StatusSize);
        if (offset == uint.MaxValue) {
            return false;
        }
        writer.WriteAt(offset+78, (uint)__StatusOffset);
        writer.WriteAt(offset+78 + 4, (uint)__StatusSize);
        for (var i = 0; i < this._Status.Count; i++) {
            writer.WriteAt(__StatusOffset, this._Status[i]);
            __StatusOffset += 4;
        }
        var __WeaponsOffset = offset+86;
            for (var i = 0; i < this._Weapons.Count; i++) {
                if (!this._Weapons[i].Write(writer, __WeaponsOffset)) {
                    return false;
                }
                __WeaponsOffset += 4;
            }
        var __PathSize = (uint)(12 * this._Path.Count);
        var __PathOffset = writer.Alloc(__PathSize);
        if (offset == uint.MaxValue) {
            return false;
        }
        writer.WriteAt(offset+102, (uint)__PathOffset);
        writer.WriteAt(offset+102 + 4, (uint)__PathSize);
            for (var i = 0; i < this._Path.Count; i++) {
                if (!this._Path[i].Write(writer, __PathOffset)) {
                    return false;
                }
                __PathOffset += 12;
            }
        var __IsAliveOffset = offset+110;
        writer.WriteAt(__IsAliveOffset, this._IsAlive);

        return true;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void ReadAsRoot(Karmem.Reader reader) {
        this.Read(MonsterDataViewer.NewMonsterDataViewer(reader, 0), reader);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Read(MonsterDataViewer viewer, Karmem.Reader reader) {
        this._Pos.Read(viewer.Pos(), reader);
        this._Mana = viewer.Mana();
        this._Health = viewer.Health();
        this._Name = viewer.Name(reader);
        this._Team = (Team)(viewer.Team());
        var __InventorySlice = viewer.Inventory(reader);
        var __InventoryLen = __InventorySlice.Length;
        if (this._Inventory.Count != __InventoryLen) {
            if (__InventoryLen > this._Inventory.Capacity) {
                this._Inventory.EnsureCapacity(__InventoryLen);
                for (var i = this._Inventory.Count; i < __InventoryLen; i++) {
                    this._Inventory.Add(0);
                }
            }
            this._Inventory.GetType().GetField("_size", BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.GetField).SetValue(this._Inventory, __InventoryLen);
        }
        for (var i = 0; i < __InventoryLen; i++) {
            this._Inventory[i] = __InventorySlice[i];
        }
        this._Color = (Color)(viewer.Color());
        var __HitboxSlice = viewer.Hitbox();
        var __HitboxLen = __HitboxSlice.Length;
        if (__HitboxLen > 5) {
            __HitboxLen = 5;
        }
        if (this._Hitbox.Count != __HitboxLen) {
            if (__HitboxLen > this._Hitbox.Capacity) {
                this._Hitbox.EnsureCapacity(__HitboxLen);
                for (var i = this._Hitbox.Count; i < __HitboxLen; i++) {
                    this._Hitbox.Add(0);
                }
            }
            this._Hitbox.GetType().GetField("_size", BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.GetField).SetValue(this._Hitbox, __HitboxLen);
        }
        for (var i = 0; i < __HitboxLen; i++) {
            this._Hitbox[i] = __HitboxSlice[i];
        }
        for (var i = __HitboxLen; i < this._Hitbox.Count; i++) {
            this._Hitbox[i] = 0;
        }
        var __StatusSlice = viewer.Status(reader);
        var __StatusLen = __StatusSlice.Length;
        if (this._Status.Count != __StatusLen) {
            if (__StatusLen > this._Status.Capacity) {
                this._Status.EnsureCapacity(__StatusLen);
                for (var i = this._Status.Count; i < __StatusLen; i++) {
                    this._Status.Add(0);
                }
            }
            this._Status.GetType().GetField("_size", BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.GetField).SetValue(this._Status, __StatusLen);
        }
        for (var i = 0; i < __StatusLen; i++) {
            this._Status[i] = __StatusSlice[i];
        }
        var __WeaponsSlice = viewer.Weapons();
        var __WeaponsLen = __WeaponsSlice.Length;
        if (__WeaponsLen > 4) {
            __WeaponsLen = 4;
        }
        if (this._Weapons.Count != __WeaponsLen) {
            if (__WeaponsLen > this._Weapons.Capacity) {
                this._Weapons.EnsureCapacity(__WeaponsLen);
                for (var i = this._Weapons.Count; i < __WeaponsLen; i++) {
                    this._Weapons.Add(new Weapon());
                }
            }
            this._Weapons.GetType().GetField("_size", BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.GetField).SetValue(this._Weapons, __WeaponsLen);
        }
        var __WeaponsSpan = CollectionsMarshal.AsSpan(this._Weapons);
        for (var i = 0; i < __WeaponsLen; i++) {
            ref var __WeaponsItem = ref __WeaponsSpan[i];
            __WeaponsItem.Read(__WeaponsSlice[i], reader);
        }
        for (var i = __WeaponsLen; i < this._Weapons.Count; i++) {
            ref var __WeaponsItem = ref __WeaponsSpan[i];
            __WeaponsItem.Reset();
        }
        var __PathSlice = viewer.Path(reader);
        var __PathLen = __PathSlice.Length;
        if (this._Path.Count != __PathLen) {
            if (__PathLen > this._Path.Capacity) {
                this._Path.EnsureCapacity(__PathLen);
                for (var i = this._Path.Count; i < __PathLen; i++) {
                    this._Path.Add(new Vec3());
                }
            }
            this._Path.GetType().GetField("_size", BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.GetField).SetValue(this._Path, __PathLen);
        }
        var __PathSpan = CollectionsMarshal.AsSpan(this._Path);
        for (var i = 0; i < __PathLen; i++) {
            ref var __PathItem = ref __PathSpan[i];
            __PathItem.Read(__PathSlice[i], reader);
        }
        this._IsAlive = viewer.IsAlive();
    }
}
public unsafe struct Monster {
    public MonsterData _Data = new MonsterData();

    public Monster() {}

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static Monster NewMonster() {
        return new Monster();
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static PacketIdentifier GetPacketIdentifier() {
        return PacketIdentifier.Monster;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Reset() {
        this.ReadAsRoot(_Globals.NullReader());
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool WriteAsRoot(Karmem.Writer writer) {
        return this.Write(writer, 0);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool Write(Karmem.Writer writer, uint start) {
        var offset = start;
        var size = (uint)4;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == uint.MaxValue) {
                return false;
            }
        }
        var __DataSize = (uint)111;
        var __DataOffset = writer.Alloc(__DataSize);
        if (offset == uint.MaxValue) {
            return false;
        }
        writer.WriteAt(offset+0, (uint)__DataOffset);
            if (!this._Data.Write(writer, __DataOffset)) {
                return false;
            }

        return true;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void ReadAsRoot(Karmem.Reader reader) {
        this.Read(MonsterViewer.NewMonsterViewer(reader, 0), reader);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Read(MonsterViewer viewer, Karmem.Reader reader) {
        this._Data.Read(viewer.Data(reader), reader);
    }
}
public unsafe struct Monsters {
    public List<Monster> _Monsters = new List<Monster>();

    public Monsters() {}

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static Monsters NewMonsters() {
        return new Monsters();
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static PacketIdentifier GetPacketIdentifier() {
        return PacketIdentifier.Monsters;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Reset() {
        this.ReadAsRoot(_Globals.NullReader());
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool WriteAsRoot(Karmem.Writer writer) {
        return this.Write(writer, 0);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool Write(Karmem.Writer writer, uint start) {
        var offset = start;
        var size = (uint)12;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == uint.MaxValue) {
                return false;
            }
        }
        writer.WriteAt(offset, (uint)12);
        var __MonstersSize = (uint)(4 * this._Monsters.Count);
        var __MonstersOffset = writer.Alloc(__MonstersSize);
        if (offset == uint.MaxValue) {
            return false;
        }
        writer.WriteAt(offset+4, (uint)__MonstersOffset);
        writer.WriteAt(offset+4 + 4, (uint)__MonstersSize);
            for (var i = 0; i < this._Monsters.Count; i++) {
                if (!this._Monsters[i].Write(writer, __MonstersOffset)) {
                    return false;
                }
                __MonstersOffset += 4;
            }

        return true;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void ReadAsRoot(Karmem.Reader reader) {
        this.Read(MonstersViewer.NewMonstersViewer(reader, 0), reader);
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Read(MonstersViewer viewer, Karmem.Reader reader) {
        var __MonstersSlice = viewer.Monsters(reader);
        var __MonstersLen = __MonstersSlice.Length;
        if (this._Monsters.Count != __MonstersLen) {
            if (__MonstersLen > this._Monsters.Capacity) {
                this._Monsters.EnsureCapacity(__MonstersLen);
                for (var i = this._Monsters.Count; i < __MonstersLen; i++) {
                    this._Monsters.Add(new Monster());
                }
            }
            this._Monsters.GetType().GetField("_size", BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.GetField).SetValue(this._Monsters, __MonstersLen);
        }
        var __MonstersSpan = CollectionsMarshal.AsSpan(this._Monsters);
        for (var i = 0; i < __MonstersLen; i++) {
            ref var __MonstersItem = ref __MonstersSpan[i];
            __MonstersItem.Read(__MonstersSlice[i], reader);
        }
    }
}

[StructLayout(LayoutKind.Sequential, Pack=1, Size=12)]
public unsafe struct Vec3Viewer {
    private readonly ulong _0;
    private readonly uint _1;

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static ref Vec3Viewer NewVec3Viewer(Karmem.Reader reader, uint offset) {
        if (!reader.IsValidOffset(offset, 12)) {
            return ref *(Vec3Viewer*)(nuint)_Globals.Null();
        }
        ref Vec3Viewer v = ref *(Vec3Viewer*)(reader.MemoryPointer + offset);
        return ref v;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    private uint KarmemSizeOf() {
        return 12;
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public float X() {
        return *(float*)((nuint)Unsafe.AsPointer(ref this) + 0);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public float Y() {
        return *(float*)((nuint)Unsafe.AsPointer(ref this) + 4);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public float Z() {
        return *(float*)((nuint)Unsafe.AsPointer(ref this) + 8);
    }
}
    
[StructLayout(LayoutKind.Sequential, Pack=1, Size=19)]
public unsafe struct WeaponDataViewer {
    private readonly ulong _0;
    private readonly ulong _1;
    private readonly ushort _2;
    private readonly byte _3;

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static ref WeaponDataViewer NewWeaponDataViewer(Karmem.Reader reader, uint offset) {
        if (!reader.IsValidOffset(offset, 4)) {
            return ref *(WeaponDataViewer*)(nuint)_Globals.Null();
        }
        ref WeaponDataViewer v = ref *(WeaponDataViewer*)(reader.MemoryPointer + offset);
        if (!reader.IsValidOffset(offset, v.KarmemSizeOf())) {
            return ref *(WeaponDataViewer*)(nuint)_Globals.Null();
        }
        return ref v;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    private uint KarmemSizeOf() {
        return *(uint*)Unsafe.AsPointer(ref this);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public int Damage() {
        if (4 + 4 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(int*)((nuint)Unsafe.AsPointer(ref this) + 4);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ushort Ammo() {
        if (8 + 2 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(ushort*)((nuint)Unsafe.AsPointer(ref this) + 8);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public byte ClipSize() {
        if (10 + 1 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(byte*)((nuint)Unsafe.AsPointer(ref this) + 10);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public float ReloadTime() {
        if (11 + 4 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(float*)((nuint)Unsafe.AsPointer(ref this) + 11);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public int Range() {
        if (15 + 4 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(int*)((nuint)Unsafe.AsPointer(ref this) + 15);
    }
}
    
[StructLayout(LayoutKind.Sequential, Pack=1, Size=4)]
public unsafe struct WeaponViewer {
    private readonly uint _0;

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static ref WeaponViewer NewWeaponViewer(Karmem.Reader reader, uint offset) {
        if (!reader.IsValidOffset(offset, 4)) {
            return ref *(WeaponViewer*)(nuint)_Globals.Null();
        }
        ref WeaponViewer v = ref *(WeaponViewer*)(reader.MemoryPointer + offset);
        return ref v;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    private uint KarmemSizeOf() {
        return 4;
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ref WeaponDataViewer Data(Karmem.Reader reader) {
        var offset = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 0);
        return ref WeaponDataViewer.NewWeaponDataViewer(reader, offset);
    }
}
    
[StructLayout(LayoutKind.Sequential, Pack=1, Size=111)]
public unsafe struct MonsterDataViewer {
    private readonly ulong _0;
    private readonly ulong _1;
    private readonly ulong _2;
    private readonly ulong _3;
    private readonly ulong _4;
    private readonly ulong _5;
    private readonly ulong _6;
    private readonly ulong _7;
    private readonly ulong _8;
    private readonly ulong _9;
    private readonly ulong _10;
    private readonly ulong _11;
    private readonly ulong _12;
    private readonly uint _13;
    private readonly ushort _14;
    private readonly byte _15;

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static ref MonsterDataViewer NewMonsterDataViewer(Karmem.Reader reader, uint offset) {
        if (!reader.IsValidOffset(offset, 4)) {
            return ref *(MonsterDataViewer*)(nuint)_Globals.Null();
        }
        ref MonsterDataViewer v = ref *(MonsterDataViewer*)(reader.MemoryPointer + offset);
        if (!reader.IsValidOffset(offset, v.KarmemSizeOf())) {
            return ref *(MonsterDataViewer*)(nuint)_Globals.Null();
        }
        return ref v;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    private uint KarmemSizeOf() {
        return *(uint*)Unsafe.AsPointer(ref this);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ref Vec3Viewer Pos() {
        if (4 + 12 > this.KarmemSizeOf()) {
            return ref *(Vec3Viewer*)((nuint)_Globals.Null());
        }
        return ref *(Vec3Viewer*)((nuint)Unsafe.AsPointer(ref this) + 4);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public short Mana() {
        if (16 + 2 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(short*)((nuint)Unsafe.AsPointer(ref this) + 16);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public short Health() {
        if (18 + 2 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(short*)((nuint)Unsafe.AsPointer(ref this) + 18);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public string Name(Karmem.Reader reader) {
        if (20 + 8 > this.KarmemSizeOf()) {
            return "";
        }
        var offset = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 20);
        var size = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 20 + 4);
        if (!reader.IsValidOffset(offset, size)) {
            return "";
        }
        var length = size / 1;
        if (length > 512) {
            length = 512;
        }
        return Marshal.PtrToStringUTF8((IntPtr)(reader.Memory.ToInt64() + offset), (int)length);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public Team Team() {
        if (28 + 1 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(Team*)((nuint)Unsafe.AsPointer(ref this) + 28);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ReadOnlySpan<byte> Inventory(Karmem.Reader reader) {
        if (29 + 8 > this.KarmemSizeOf()) {
            return new ReadOnlySpan<byte>();
        }
        var offset = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 29);
        var size = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 29 + 4);
        if (!reader.IsValidOffset(offset, size)) {
            return new ReadOnlySpan<byte>();
        }
        var length = size / 1;
        if (length > 128) {
            length = 128;
        }
        return new ReadOnlySpan<byte>((void*)(reader.MemoryPointer + offset), (int)length);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public Color Color() {
        if (37 + 1 > this.KarmemSizeOf()) {
            return 0;
        }
        return *(Color*)((nuint)Unsafe.AsPointer(ref this) + 37);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ReadOnlySpan<double> Hitbox() {
        if (38 + 40 > this.KarmemSizeOf()) {
            return new ReadOnlySpan<double>();
        }
        return new ReadOnlySpan<double>((void*)((nuint)Unsafe.AsPointer(ref this) + 38), 5);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ReadOnlySpan<int> Status(Karmem.Reader reader) {
        if (78 + 8 > this.KarmemSizeOf()) {
            return new ReadOnlySpan<int>();
        }
        var offset = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 78);
        var size = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 78 + 4);
        if (!reader.IsValidOffset(offset, size)) {
            return new ReadOnlySpan<int>();
        }
        var length = size / 4;
        if (length > 10) {
            length = 10;
        }
        return new ReadOnlySpan<int>((void*)(reader.MemoryPointer + offset), (int)length);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ReadOnlySpan<WeaponViewer> Weapons() {
        if (86 + 16 > this.KarmemSizeOf()) {
            return new ReadOnlySpan<WeaponViewer>();
        }
        return new ReadOnlySpan<WeaponViewer>((void*)((nuint)Unsafe.AsPointer(ref this) + 86), 4);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ReadOnlySpan<Vec3Viewer> Path(Karmem.Reader reader) {
        if (102 + 8 > this.KarmemSizeOf()) {
            return new ReadOnlySpan<Vec3Viewer>();
        }
        var offset = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 102);
        var size = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 102 + 4);
        if (!reader.IsValidOffset(offset, size)) {
            return new ReadOnlySpan<Vec3Viewer>();
        }
        var length = size / 12;
        if (length > 2000) {
            length = 2000;
        }
        return new ReadOnlySpan<Vec3Viewer>((void*)(reader.MemoryPointer + offset), (int)length);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool IsAlive() {
        if (110 + 1 > this.KarmemSizeOf()) {
            return false;
        }
        return *(bool*)((nuint)Unsafe.AsPointer(ref this) + 110);
    }
}
    
[StructLayout(LayoutKind.Sequential, Pack=1, Size=4)]
public unsafe struct MonsterViewer {
    private readonly uint _0;

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static ref MonsterViewer NewMonsterViewer(Karmem.Reader reader, uint offset) {
        if (!reader.IsValidOffset(offset, 4)) {
            return ref *(MonsterViewer*)(nuint)_Globals.Null();
        }
        ref MonsterViewer v = ref *(MonsterViewer*)(reader.MemoryPointer + offset);
        return ref v;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    private uint KarmemSizeOf() {
        return 4;
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ref MonsterDataViewer Data(Karmem.Reader reader) {
        var offset = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 0);
        return ref MonsterDataViewer.NewMonsterDataViewer(reader, offset);
    }
}
    
[StructLayout(LayoutKind.Sequential, Pack=1, Size=12)]
public unsafe struct MonstersViewer {
    private readonly ulong _0;
    private readonly uint _1;

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public static ref MonstersViewer NewMonstersViewer(Karmem.Reader reader, uint offset) {
        if (!reader.IsValidOffset(offset, 4)) {
            return ref *(MonstersViewer*)(nuint)_Globals.Null();
        }
        ref MonstersViewer v = ref *(MonstersViewer*)(reader.MemoryPointer + offset);
        if (!reader.IsValidOffset(offset, v.KarmemSizeOf())) {
            return ref *(MonstersViewer*)(nuint)_Globals.Null();
        }
        return ref v;
    }

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    private uint KarmemSizeOf() {
        return *(uint*)Unsafe.AsPointer(ref this);
    }
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public ReadOnlySpan<MonsterViewer> Monsters(Karmem.Reader reader) {
        if (4 + 8 > this.KarmemSizeOf()) {
            return new ReadOnlySpan<MonsterViewer>();
        }
        var offset = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 4);
        var size = *(uint*)((nuint)Unsafe.AsPointer(ref this) + 4 + 4);
        if (!reader.IsValidOffset(offset, size)) {
            return new ReadOnlySpan<MonsterViewer>();
        }
        var length = size / 4;
        if (length > 2000) {
            length = 2000;
        }
        return new ReadOnlySpan<MonsterViewer>((void*)(reader.MemoryPointer + offset), (int)length);
    }
}
    
