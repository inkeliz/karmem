import * as karmem from "../assemblyscript/karmem"
import * as km from "./km/game_generated"
import {Writer} from "../assemblyscript/karmem";

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
  KarmemReader.SetSize(size);
  km.Monsters.ReadAsRoot(KarmemStruct, KarmemReader);
}

export function KBenchmarkDecodeSumVec3(size: u32) : f32 {
  KarmemReader.SetSize(size);

  let monsters = km.NewMonstersViewer(KarmemReader, 0);
  let monsterList = monsters.Monsters(KarmemReader);

  let sum : f32 = 0;
  for (let i = 0; i < monsterList.length; i++) {
    let data = monsterList[i].Data(KarmemReader);
    let path = data.Path(KarmemReader);

    for (let p = 0; p < path.length; p++) {
      let pp = path[p];
      sum += pp.X() + pp.Y() + pp.Z();
    }
  }

  return sum;
}

export function KBenchmarkDecodeSumUint8(size: u32) : u32 {
  KarmemReader.SetSize(size);

  let monsters = km.NewMonstersViewer(KarmemReader, 0);
  let monstersList = monsters.Monsters(KarmemReader);

  let sum: u32 = 0
  for (let i = 0; i < monstersList.length; i++) {
    let inv = monstersList[i].Data(KarmemReader).Inventory(KarmemReader);

    for (let j = 0; j < inv.length; j++) {
      sum += u32(inv[j]);
    }
  }

  return sum;
}

//export KBenchmarkDecodeSumStats
export function KBenchmarkDecodeSumStats(size: u32) : u32 {
  KarmemReader.SetSize(size)

  let monsters = km.NewMonstersViewer(KarmemReader, 0)
  let monstersList = monsters.Monsters(KarmemReader)

  let sum: km.WeaponData = km.NewWeaponData()
  for (let i = 0; i < monstersList.length; i++) {
    let weapons = monstersList[i].Data(KarmemReader).Weapons()

    for (let j = 0; j < weapons.length; j++) {
      let data = weapons[j].Data(KarmemReader)
      sum.Ammo += data.Ammo()
      sum.Damage += data.Damage()
      sum.ClipSize += data.ClipSize()
      sum.ReloadTime += data.ReloadTime()
      sum.Range += data.Range()
    }
  }

  KarmemWriter.Reset()
  km.WeaponData.WriteAsRoot(sum, KarmemWriter);
  return KarmemWriter.Bytes().length
}

export function KNOOP() : u32 {
  return 42;
}