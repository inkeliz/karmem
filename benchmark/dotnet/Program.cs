// See https://aka.ms/new-console-template for more information

using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;
using km;

// Run test manually:
unsafe
{
    var g = new Global();
    g.KBenchmarkDecodeObjectAPIFrom(g.InputMemory);

    float expected = 50150224.000000F;
    if (g.KBenchmarkDecodeSumVec3((uint)g.InputMemory.Length) != expected)
    {
        Console.WriteLine("KBenchmarkDecodeSumVec3 failed");
        Console.WriteLine((uint)g.KBenchmarkDecodeSumVec3((uint)g.InputMemory.Length));
    }
    g.KBenchmarkDecodeObjectAPI((uint)g.InputMemory.Length);
    g.KBenchmarkEncodeObjectAPI();
}

public class Global
{
    public byte[] InputMemory = File.ReadAllBytes("../testdata/kmbin/km.bin");
    public IntPtr OutputMemory = Marshal.AllocHGlobal(8_000_000);
    public km.Monsters Structure = new km.Monsters(); 
    public Karmem.Reader Reader;
    public Karmem.Writer Writer;

    public Global()
    {
        Reader = Karmem.Reader.NewManagedReader(InputMemory);
        Writer = Karmem.Writer.NewFixedWriter(OutputMemory, 8_000_000);
    }
    
    // Must be exported to WASM.
    public void KBenchmarkEncodeObjectAPI()
    {
        this.Writer.Reset();
        if (!Structure.WriteAsRoot(this.Writer)) {
            throw new System.Exception("Failed to write object");
        }
    }

    // Must be exported to WASM.
    public void KBenchmarkDecodeObjectAPI(uint size)
    {
        this.Reader.SetSize(size);
        Structure.ReadAsRoot(this.Reader);
    }

    // Must be exported to WASM.
    public void KBenchmarkDecodeObjectAPIFrom(byte[] b)
    {
        var reader = Karmem.Reader.NewManagedReader(b);
        Structure.ReadAsRoot(reader);
        reader.Dispose();
    }

    // Must be exported to WASM.
    public float KBenchmarkDecodeSumVec3(uint size)
    {
        this.Reader.SetSize(size);
        
        var monsters = MonstersViewer.NewMonstersViewer(this.Reader, 0);
        var monstersList = monsters.Monsters(this.Reader);
        
        var sum = new Vec3();
        for (var i = 0; i < monstersList.Count; i++)
        {
            var path = monstersList[i].Data(this.Reader).Path(this.Reader);
            for (var p = 0; p < path.Count; p++)
            {
         
                sum._X += path[p].X();
                sum._Y += path[p].Y();
                sum._Z += path[p].Z();
            }
        }

        return sum._X + sum._Y + sum._Z;
    }
    
}