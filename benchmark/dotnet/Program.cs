using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;
using BitConverter = System.BitConverter;
#if IS_KARMEM
    using km;
#endif
#if IS_WASM
#else
using System.Net;
using System.Net.Sockets;
#endif

var global = new dotnet.Benchmark(); // Init global benchmark object
dotnet._Global.global = global; // Suppress GC

unsafe {
#if IS_WASM // Compile with: dotnet build -c Release -p:IS_WASM=true -p:BYTEBUFFER_NO_BOUNDS_CHECK=true -p:UNSAFE_BYTEBUFFE=true

    // Give the C code the pointer to the Benchmark class.
    dotnet.Benchmark.Ready(global); 

#else // Compile with: dotnet build -c Release -p:IS_WASM=false -p:BYTEBUFFER_NO_BOUNDS_CHECK=true -p:UNSAFE_BYTEBUFFE=true
    try
    {
        var server = new TcpListener(IPAddress.Parse("127.0.0.1"), 13000);
        server.Start();

        var buf = new Span<byte>(global.InputMemory.ToPointer(), 8_000_000);
        var len = new byte[4];
        var fn = new byte[4];

        while (true)
        {
            TcpClient client = server.AcceptTcpClient();
            client.ReceiveTimeout = 3_600_000;
            client.SendTimeout = 3_600_000;
            client.SendBufferSize = 9_000_000;
            client.ReceiveBufferSize = 9_000_000;
            
            NetworkStream stream = client.GetStream();
            
            while (stream.Read(len, 0, 4) != 0)
            {
                stream.Read(fn, 0, 4);
                if (BitConverter.ToInt32(len, 0) > 0)
                {
                    stream.ReadAtLeast(buf, BitConverter.ToInt32(len, 0));
                }

                switch (BitConverter.ToUInt32(fn, 0))
                {
                    case 1:
                        global.KBenchmarkDecodeObjectAPI(BitConverter.ToUInt32(len, 0));
                        stream.Write(BitConverter.GetBytes(0));
                        continue;
                    case 2:
                        global.KBenchmarkEncodeObjectAPI();
                        var bufout = new ReadOnlySpan<byte>(global.OutputMemory.ToPointer(), 8_000_000);
                        stream.Write(BitConverter.GetBytes((int)8_000_000));
                        stream.Write(bufout);
                        continue;
                    case 3:
                        var r = global.KBenchmarkDecodeSumVec3(BitConverter.ToUInt32(len, 0));
                        stream.Write(BitConverter.GetBytes(r));
                        continue;
                }
            }
        }
    
    }
    catch (Exception exception)
    {
        Console.WriteLine(exception.ToString());
    }
    
#endif

}

namespace dotnet
{

    public static class _Global
    {
        public static dotnet.Benchmark? global = null;
    }

    public unsafe class Benchmark
    {
    #if IS_KARMEM

        // The InputMemory/OutputMemory will leak, but it's ok, since that class is alive as long as the program is running.
        public IntPtr InputMemory = Marshal.AllocHGlobal(8_000_000);
        public IntPtr OutputMemory = Marshal.AllocHGlobal(8_000_000);

        public km.Monsters Structure = new km.Monsters();
        public Karmem.Reader Reader;
        public Karmem.Writer Writer;

    #endif
    #if IS_FLATBUFFERS

        public IntPtr InputMemory;
        public IntPtr OutputMemory;

        public game.MonstersT Structure = new game.MonstersT();
        public FlatBuffers.ByteBuffer Reader;
        public FlatBuffers.FlatBufferBuilder Writer;

    #endif

        public Benchmark()
        {
        #if IS_KARMEM
            Unsafe.InitBlockUnaligned(InputMemory.ToPointer(), 0, 8_000_000);
            Unsafe.InitBlockUnaligned(OutputMemory.ToPointer(), 0, 8_000_000);
            Reader = Karmem.Reader.NewReader(InputMemory, 8_000_000, 8_000_000);
            Writer = Karmem.Writer.NewFixedWriter(OutputMemory, 8_000_000);
        #endif
        #if IS_FLATBUFFERS
            var i = new byte[8_000_000];
            var handle = GCHandle.Alloc(i, GCHandleType.Pinned);
            fixed (void* f= &i[0])
            {
                InputMemory = new IntPtr(f);
            }
            Reader = new FlatBuffers.ByteBuffer(i);

            var o = new byte[8_000_000];
            var handle2 = GCHandle.Alloc(o, GCHandleType.Pinned);
            fixed (void* f= &o[0])
            {
                OutputMemory = new IntPtr(f);
            }
            Writer = new FlatBuffers.FlatBufferBuilder(new FlatBuffers.ByteBuffer(o));
        #endif
        }

        [MethodImpl(MethodImplOptions.InternalCall)]
        public static extern void Ready(Benchmark owner);

        public uint InputMemoryPointer()
        {
            return (uint)InputMemory.ToInt64();
        }

        public uint OutputMemoryPointer()
        {
            return (uint)OutputMemory.ToInt64();
        }

        // Must be exported to WASM.
        public void KBenchmarkEncodeObjectAPI()
        {
        #if IS_KARMEM
            this.Writer.Reset();
            if (!Structure.WriteAsRoot(this.Writer))
            {
                throw new System.Exception("Failed to write object");
            }
        #endif
        #if IS_FLATBUFFERS
            Writer.Clear();
            var r = game.Monsters.Pack(Writer, Structure);
            Writer.Finish(r.Value);
        #endif
        }

        // Must be exported to WASM.
        public void KBenchmarkDecodeObjectAPI(uint size)
        {
        #if IS_KARMEM
            this.Reader.SetSize(size);
            Structure.ReadAsRoot(this.Reader);
        #endif
        #if IS_FLATBUFFERS
            var r = game.Monsters.GetRootAsMonsters(Reader);
            r.UnPackTo(Structure);
        #endif
        }

        // Must be exported to WASM.
        public float KBenchmarkDecodeSumVec3(uint size)
        {
        #if IS_KARMEM
            this.Reader.SetSize(size);

            var sum = new Vec3();
            var monsters = MonstersViewer.NewMonstersViewer(this.Reader, 0);
            var monstersList = monsters.Monsters(this.Reader);

            for (var i = 0; i < monstersList.Length; i++)
            {
                var paths = monstersList[i].Data(this.Reader).Path(this.Reader);
                for (var p = 0; p < paths.Length; p++)
                {
                    var path = paths[p];
                    sum._X += path.X();
                    sum._Y += path.Y();
                    sum._Z += path.Z();
                }
            }
            return sum._X + sum._Y + sum._Z;
        #endif
        #if IS_FLATBUFFERS
            var sum = new game.Vec3T();
            var monsters = game.Monsters.GetRootAsMonsters(Reader);

            for (var i = 0; i < monsters.MonstersLength; i++)
            {
                var monster = monsters._Monsters(i).Value;
                for (var p = 0; p < monster.PathLength; p++)
                {
                    var path = monster.Path(p).Value;
                    sum.X += path.X;
                    sum.Y += path.Y;
                    sum.Z += path.Z;
                }
            }
            return sum.X + sum.Y + sum.Z;
        #endif
        }

    }
}