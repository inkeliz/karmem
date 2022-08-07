using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;
using System.Text;

namespace Karmem;

/// <summary>
///     IViewer is one marker interface implemented by all
///     viewers living in the generated code.
///
///     The implementer MUST be on STRUCT and have one ulong field
///     at index 0.
/// </summary>
public interface IViewer { }

/// <summary>
///     The Slice is similar to Span<T>, but supports more than 2GB, assuming that the Reader is unmanaged.
///     The Slice doesn't copy the values, instead it's just an encapsulation over the existent Reader data,
///     beware that is unsafe to use if Reader.Dispose or Reader.Reset is called.
///
///     If T is one IViewer, the given T is expected to be one struct and MUST have one ulong at the
///     beginning of the struct (0 index).
/// </summary>
public unsafe ref struct Slice<T>
{
    private readonly ulong _ptr;
    private readonly ulong _length;
    private readonly ulong _stride;
    private readonly bool _isViewer;

    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public Slice(ulong ptr, uint length, uint stride)
    {
        _ptr = ptr;
        _length = (ulong)length;
        _stride = (ulong)stride;
        _isViewer = typeof(IViewer).IsAssignableFrom(typeof(T));
    }

    public ulong Length => this._length;

    public ulong Count => Length;

    /// <summary>
    ///     Return the element at the given index.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public T Get(ulong index)
    {
        ulong d = 0;
        if (this._isViewer)
        {
            // The IViewer MUST have one ulong at the beginning of the struct.
            // In order words: ulong have the same memory layout of struct { [FieldOffset(0)]private ulong _ptr = 0; },
            // as expected. So, the d is the same as such struct, populated with the _ptr field.
            d = (_ptr + (ulong)(index * _stride));
        }
        else
        {
            // ulong is the largest primitive type, so everything can be casted
            // from ulong to T.
            d = *(ulong*)(_ptr + (ulong)(index * _stride));
        }

        return Unsafe.As<ulong, T>(ref d);
    }
    
    public T this[ulong index]
    {
        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        get => Get(index);
    }
    
    public T this[int index]
    {
        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        get => Get((ulong)index);
    }
    
    public T this[uint index]
    {
        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        get => Get((ulong)index);
    }
    
    public T this[long index]
    {
        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        get => Get((ulong)index);
    }
}

/// <summary>
///     A struct to write KARMEM, that is consumed by generated code. Most
///     functions are not designed to be used by hand.
///     That is not concurrently safe, and should not write multiple data
///     at the same time.
/// </summary>
public unsafe struct Writer
{
    public IntPtr Memory;
    public long Size = 0;
    public long Capacity;
    private readonly bool _isFixed = false;
    private readonly GCHandle? _handle = null;

    public Writer(int capacity)
    {
        Memory = Marshal.AllocHGlobal(capacity);
        Capacity = capacity;
    }

    public Writer(IntPtr memory, int capacity, GCHandle? gcHandle)
    {
        Memory = memory;
        Capacity = capacity;
        _isFixed = true;
        _handle = gcHandle;
    }

    /// <summary>
    ///     Creates a new Writer with the specified capacity. It will grow the buffer if needed.
    /// </summary>
    public static Writer NewWriter(int capacity)
    {
        return new Writer(capacity);
    }

    /// <summary>
    ///     Creates a new Writer with the given memory as buffer. It will not grow the buffer, and
    ///     you must make sure to not deallocate the memory before the Writer is disposed.
    /// </summary>
    public static Writer NewFixedWriter(IntPtr memory, int capacity)
    {
        return new Writer(memory, capacity, null);
    }

    /// <summary>
    ///     Creates a new Writer with the given memory from an Managed Memory as buffer. It will not grow
    ///     the buffer, and will not free the memory until the Writer is disposed.
    /// </summary>
    public static Writer NewFixedManagedWriter(byte[] memory)
    {
        var h = GCHandle.Alloc(memory, GCHandleType.Pinned);
        return new Writer(GCHandle.ToIntPtr(h), memory.Length, h);
    }

    /// <summary>
    ///     Allocates n bytes inside the buffer.
    ///     It returns the offset and may return uint.MinValue if it's not possible to allocate.
    ///     It must call Write() to write the data after allocating.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public uint Alloc(uint n)
    {
        var ptr = Size;
        var total = ptr + n;
        if (total > Capacity)
        {
            if (_isFixed) return uint.MaxValue;

            var target = Capacity * 2;
            if (target < total) target = total * 2;

            Memory = Marshal.ReAllocHGlobal(Memory, (IntPtr)target);
            Unsafe.InitBlockUnaligned((void*)(Memory.ToInt64() + ptr), 0, (uint)(target - ptr));
            Capacity = target;
            Size = total;
        }
        else
        {
            Size = total;
        }
        return (uint)ptr;
    }

    /// <summary>
    ///     Writes the given src data into the writer buffer, starting at the given offset.
    ///     That function is type-agnostic, as such it can be used to write any type of data,
    ///     the size should be in bytes.
    ///     Notice that strings must be UTF-8, or use the WriteAt(long offset, string src) function.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, IntPtr src, long size)
    {
        Buffer.MemoryCopy(src.ToPointer(), (void*)(Memory.ToInt64() + offset), size, size);
    }

    /// <summary>
    ///     Writes the given src UTF-16 string into the writer buffer, starting at the given offset.
    ///     That function returns the size of the UTF-8 string in bytes. That functions assumes that the
    ///     previous allocation is the size of (char count) * 4.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public int WriteAt(long offset, string src)
    {
        if (src.Length == 0) return 0;

        fixed (char* p = src) // Equivalent of GCHandle.Alloc(src, GCHandleType.Pinned)
        {
            var size = Encoding.UTF8.GetBytes(p, src.Length, (byte*)(Memory.ToInt64() + offset), src.Length * 4);
            return size;
        }
    }

    /// <summary>
    ///     Writes the single-byte value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, bool src)
    {
        *(bool*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the single-byte value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, byte src)
    {
        *(byte*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the single-byte value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, sbyte src)
    {
        *(sbyte*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the two-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, short src)
    {
        *(short*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the two-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, ushort src)
    {
        *(ushort*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the four-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, int src)
    {
        *(int*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the four-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, uint src)
    {
        *(uint*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the eight-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, long src)
    {
        *(long*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the eight-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, ulong src)
    {
        *(ulong*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the four-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, float src)
    {
        *(float*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Writes the eight-bytes value into the writer buffer, starting at the given offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void WriteAt(long offset, double src)
    {
        *(double*)(Memory.ToInt64() + offset) = src;
    }

    /// <summary>
    ///     Resets the writer buffer, setting the size to 0.
    ///     It doesn't deallocate the memory.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Reset()
    {
        Size = 0;
    }

    /// <summary>
    ///     Releases the writer buffer, deallocating the memory.
    ///     Only call this if you created the writer with NewWriter or NewFixedManagedWriter.
    /// </summary>
    public void Dispose()
    {
        if (_isFixed)
        {
            if (_handle is null) return;
            _handle?.Free();
        }
        else
        {
            Marshal.FreeHGlobal(Memory);
        }
    }

    /// <summary>
    ///     Bytes returns the bytes of the writer buffer.
    ///     Beware that the returned array is not a copy,
    ///     the content will be lost or corrupted when
    ///     call Reset() or Dispose().
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public Span<byte> Bytes()
    {
        return new Span<byte>(Memory.ToPointer(), (int)Size);
    }
}

/// <summary>
///     A struct to read KARMEM encoded data. That is consumed by generated code.
///     The backed buffer must not be modified when any data-viewer is using it.
/// </summary>
public struct Reader
{
    public IntPtr Memory;
    public long Size;
    public readonly long Capacity;
    public GCHandle? Handle;

    public Reader(IntPtr memory, long size, long capacity, GCHandle? gcHandle)
    {
        Memory = memory;
        Size = size;
        Capacity = capacity;
        Handle = gcHandle;
    }

    /// <summary>
    ///     Creates a new reader from unmanaged memory, with the given size and capacity.
    /// </summary>
    public static Reader NewReader(IntPtr memory, long size, long capacity)
    {
        return new Reader(memory, size, capacity, null);
    }

    /// <summary>
    ///     Creates a new Reader from the given memory, which is allocated
    ///     on the Managed memory heap. The size will be the capacity.
    ///     You must call Dispose() on the returned Reader when you are done with it,
    ///     otherwise the memory will be leaked.
    /// </summary>
    public static Reader NewManagedReader(byte[] memory)
    {
        return NewManagedReader(memory, memory.Length);
    }

    /// <summary>
    ///     Creates a new Reader from the given memory, which is allocated
    ///     on the Managed memory heap. The size and capacity can be different, but
    ///     the capacity should be equal or greater than the size, and undefined behavior
    ///     will occur if it's not.
    ///     You must call Dispose() on the returned Reader when you are done with it,
    ///     otherwise the memory will be leaked.
    /// </summary>
    public static Reader NewManagedReader(byte[] memory, long size)
    {
        var h = GCHandle.Alloc(memory, GCHandleType.Pinned);
        return new Reader(h.AddrOfPinnedObject(), size, memory.Length, h);
    }

    /// <summary>
    ///     Checks the given offset is valid for the given size and offset.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool IsValidOffset(uint offset, uint size)
    {
        return Size >= offset + (long)size;
    }

    /// <summary>
    ///     Re-defines the bounds of the memory, useful when the
    ///     backend slice is being re-used for multiples contents.
    ///     It returns false if the given size is greater than the capacity, or
    ///     invalid.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public bool SetSize(uint size)
    {
        if ((size == 0) | (size > Capacity)) return false;
        Size = size;
        return true;
    }

    /// <summary>
    ///     Release the memory associated with the reader, when using the NewManagedReader.
    /// </summary>
    [MethodImpl(MethodImplOptions.AggressiveInlining)]
    public void Dispose()
    {
        if (Handle is null) return;
        Handle?.Free();
    }
}