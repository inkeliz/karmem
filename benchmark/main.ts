import * as karmem from "../assemblyscript/karmem"
import * as km from "./km/game_generated"
import { Console } from "as-wasi/assembly"

let InputMemory = new StaticArray<u8>(8_000_000);
let OutputMemory = new Array<u8>(8_000_000);

export function InputMemoryPointer() : u32 {
  return <u32>changetype<usize>(InputMemory);
}

export function OutputMemoryPointer() : u32 {
  return <u32>OutputMemory.dataStart;
}

let KarmemStruct : km.Monsters = km.NewMonsters();
let KarmemWriter : karmem.Writer = karmem.NewFixedWriter(OutputMemory);
let KarmemReader : karmem.Reader = karmem.NewReader(InputMemory);

export function KBenchmarkEncodeObjectAPI() :void {
  KarmemWriter.Reset();
  km.Monsters.WriteAsRoot(KarmemStruct, KarmemWriter);
}

export function KBenchmarkDecodeObjectAPI(size: u32) : void {
  KarmemReader.SetSize(size)
  km.Monsters.ReadAsRoot(KarmemStruct, KarmemReader);
}

export function KBenchmarkDecodeSumVec3(size: u32) : f32 {
  KarmemReader.SetSize(size)

  let monsters = km.NewMonstersViewer(KarmemReader, 0);
  let monsterList = monsters.Monsters(KarmemReader);

  let sum : km.Vec3 = km.NewVec3();
  for (let i = 0; i < monsterList.length; i++) {
    let data = monsterList[i].Data(KarmemReader)
    let path = data.Path(KarmemReader);

    for (let p = 0; p < path.length; p++) {
      sum.X += path[p].X();
      sum.Y += path[p].Y();
      sum.Z += path[p].Z();
    }
  }

  return sum.X + sum.Y + sum.Z;
}