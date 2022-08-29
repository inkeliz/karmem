# KARMEM

Karmem is a fast binary serialization format. The priority of Karmem is to be
easy to use while been fast as possible. It's optimized to take Golang and
TinyGo's maximum performance and is efficient for repeatable reads, reading
different content of the same type. Karmem demonstrates to be ten times faster
than Google Flatbuffers, with the additional overhead of bounds-checking
included.

> ⚠️ Karmem still under development, the API is not stable. However, serialization-format itself is unlike to change and
> should remain backward compatible with older versions.

# Contents

- [🧐 Motivation](#Motivation)
- [🧠 Usage](#Usage)
- [🏃 Benchmark](#Benchmark)
- [🌎 Languages](#Languages)
- [📙 Schema](#Schema)
    - Example
    - Types
    - Structs
    - Enum
- [🛠️ Generator](#Generator)
- [🔒 Security](#Security)

# Motivation

Karmem was create to solve one single issue: make easy to transfer data between WebAssembly host and guest. While still
portable for non-WebAssembly languages. We are experimenting with an "event-command pattern" between wasm-host and
wasm-guest in one project, but sharing data is very expensive, and FFI calls are not cheap either. Karmem encodes once
and shares the same content with multiple guests, regardless of the language, making it very efficient. Also, even using
Object-API to decode, it's fast enough, and Karmem was designed to take advantage of that pattern, avoid allocations,
and re-use the same struct for multiple data.

Why not use [Witx](https://github.com/jedisct1/witx-codegen)? It is good project and aimed to WASM, however it seems
more complex and defines not just data-structure, but functions, which I'm trying to avoid. Also, it is not intended to
be portable to non-wasm. Why not use [Flatbuffers](https://google.github.io/flatbuffers/)? We tried, but it's not fast
enough and also causes panics due to the lack of bound-checking. Why not use [Cap'n'Proto](https://capnproto.org/)? It's
a good alternative but lacks implementation for Zig and AssemblyScript, which is top-priority, it also has more
allocations and the generated API is harder to use, compared than Karmem.

# Usage

That is a small example of how use Karmem.

### Schema

```go
karmem app @packed(true) @golang.package(`app`);  
  
enum SocialNetwork uint8 { Unknown; Facebook; Instagram; Twitter; TikTok; }  
  
struct ProfileData table {  
    Network  SocialNetwork;  
    Username []char;  
    ID       uint64;  
}  
  
struct Profile inline {  
    Data ProfileData;  
}  
  
struct AccountData table {  
    ID       uint64;  
    Email    []char;  
    Profiles []Profile;  
}
```

Generate the code using `go run karmem.org/cmd/karmem build --golang -o "km" app.km`.

### Encoding

In order to encode, use should create an native struct and then encode it.

```go
var writerPool = sync.Pool{New: func() any { return karmem.NewWriter(1024) }}

func main() {
	writer := writerPool.Get().(*karmem.Writer)

	content := app.AccountData{
		ID:    42,
		Email: "example@email.com",
		Profiles: []app.Profile{
			{Data: app.ProfileData{
				Network:  app.SocialNetworkFacebook,
				Username: "inkeliz",
				ID:       123,
			}},
			{Data: app.ProfileData{
				Network:  app.SocialNetworkFacebook,
				Username: "karmem",
				ID:       231,
			}},
			{Data: app.ProfileData{
				Network:  app.SocialNetworkInstagram,
				Username: "inkeliz",
				ID:       312,
			}},
		},
	}

	if _, err := content.WriteAsRoot(writer); err != nil {
		panic(err)
	}

	encoded := writer.Bytes()
	_ = encoded // Do something with encoded data

	writer.Reset()
	writerPool.Put(writer)
}
```

### Reading

Instead of decoding it to another struct, you can read some fields directly, without any additional decoding. In this
example, we only need the username of each profile.

```go
func decodes(encoded []byte) {
	reader := karmem.NewReader(encoded)
	account := app.NewAccountDataViewer(reader, 0)

	profiles := account.Profiles(reader)
	for i := range profiles {
		fmt.Println(profiles[i].Data(reader).Username(reader))
	}
}
```

Notice: we use `NewAccountDataViewer`, any `Viewer` is just a Viewer, and doesn't copy the backend data. Some 
languages (C#, AssemblyScript) uses UTF-16, while Karmem uses UTF-8, in those cases you have some performance 
penalty.

### Decoding

You can also decode it to an existent struct. In some cases, it's better if you re-use the same struct for multiples
reads.

```go
var accountPool = sync.Pool{New: func() any { return new(app.AccountData) }}

func decodes(encoded []byte) {
	account := accountPool.Get().(*app.AccountData)
	account.ReadAsRoot(karmem.NewReader(encoded))

	profiles := account.Profiles
	for i := range profiles {
		fmt.Println(profiles[i].Data.Username)
	}

	accountPool.Put(account)
}
```

# Benchmark

### Flatbuffers vs Karmem

Using similar schema with Flatbuffers and Karmem. Karmem is almost 10 times faster than Google Flatbuffers.

**Native (MacOS/ARM64 - M1):**

```
name               old time/op    new time/op    delta
EncodeObjectAPI-8    2.54ms ± 0%    0.51ms ± 0%   -79.85%  (p=0.008 n=5+5)
DecodeObjectAPI-8    3.57ms ± 0%    0.20ms ± 0%   -94.30%  (p=0.008 n=5+5)
DecodeSumVec3-8      1.44ms ± 0%    0.16ms ± 0%   -88.86%  (p=0.008 n=5+5)

name               old alloc/op   new alloc/op   delta
EncodeObjectAPI-8    12.1kB ± 0%     0.0kB       -100.00%  (p=0.008 n=5+5)
DecodeObjectAPI-8    2.87MB ± 0%    0.00MB       -100.00%  (p=0.008 n=5+5)
DecodeSumVec3-8       0.00B          0.00B           ~     (all equal)

name               old allocs/op  new allocs/op  delta
EncodeObjectAPI-8     1.00k ± 0%     0.00k       -100.00%  (p=0.008 n=5+5)
DecodeObjectAPI-8      110k ± 0%        0k       -100.00%  (p=0.008 n=5+5)
DecodeSumVec3-8        0.00           0.00           ~     (all equal)
```

**WebAssembly on Wazero (MacOS/ARM64 - M1):**

```
name               old time/op    new time/op    delta
EncodeObjectAPI-8    17.2ms ± 0%     4.0ms ± 0%  -76.51%  (p=0.008 n=5+5)
DecodeObjectAPI-8    50.7ms ± 2%     1.9ms ± 0%  -96.18%  (p=0.008 n=5+5)
DecodeSumVec3-8      5.74ms ± 0%    0.75ms ± 0%  -86.87%  (p=0.008 n=5+5)

name               old alloc/op   new alloc/op   delta
EncodeObjectAPI-8    3.28kB ± 0%    3.02kB ± 0%   -7.80%  (p=0.008 n=5+5)
DecodeObjectAPI-8    3.47MB ± 2%    0.02MB ± 0%  -99.56%  (p=0.008 n=5+5)
DecodeSumVec3-8      1.25kB ± 0%    1.25kB ± 0%     ~     (all equal)

name               old allocs/op  new allocs/op  delta
EncodeObjectAPI-8      4.00 ± 0%      4.00 ± 0%     ~     (all equal)
DecodeObjectAPI-8      5.00 ± 0%      4.00 ± 0%  -20.00%  (p=0.008 n=5+5)
DecodeSumVec3-8        5.00 ± 0%      5.00 ± 0%     ~     (all equal)
```

### Raw-Struct vs Karmem

The performance is nearly the same when comparing reading non-serialized data from a native struct and reading it from a
karmem-serialized data.

**Native (MacOS/ARM64 - M1):**

```
name             old time/op    new time/op    delta
DecodeSumVec3-8     154µs ± 0%     160µs ± 0%  +4.36%  (p=0.008 n=5+5)

name             old alloc/op   new alloc/op   delta
DecodeSumVec3-8     0.00B          0.00B         ~     (all equal)

name             old allocs/op  new allocs/op  delta
DecodeSumVec3-8      0.00           0.00         ~     (all equal)
```

### Karmem vs Karmem

That is an comparison with all supported languages.

**WebAssembly on Wazero (MacOS/ARM64 - M1):**

```
name \ time/op     result/wasi-go-km.out  result/wasi-as-km.out  result/wasi-zig-km.out  result/wasi-swift-km.out  result/wasi-c-km.out  result/wasi-odin-km.out  result/wasi-dotnet-km.out
DecodeSumVec3-8               757µs ± 0%            1651µs ± 0%              369µs ± 0%               9145µs ± 6%            368µs ± 0%              1330µs ± 0%               75671µs ± 0% 
DecodeObjectAPI-8            1.59ms ± 0%            6.13ms ± 0%             1.04ms ± 0%              30.59ms ±34%           0.90ms ± 1%              4.06ms ± 0%              231.72ms ± 0% 
EncodeObjectAPI-8            3.96ms ± 0%            4.51ms ± 1%             1.20ms ± 0%               8.26ms ± 0%           1.03ms ± 0%              5.19ms ± 0%              237.99ms ± 0% 

name \ alloc/op    result/wasi-go-km.out  result/wasi-as-km.out  result/wasi-zig-km.out  result/wasi-swift-km.out  result/wasi-c-km.out  result/wasi-odin-km.out  result/wasi-dotnet-km.out 
DecodeSumVec3-8              1.25kB ± 0%           21.75kB ± 0%             1.25kB ± 0%               1.82kB ± 0%           1.25kB ± 0%              5.34kB ± 0%              321.65kB ± 0% 
DecodeObjectAPI-8            15.0kB ± 0%           122.3kB ± 1%            280.8kB ± 1%              108.6kB ± 3%            1.2kB ± 0%              23.8kB ± 0%               386.5kB ± 0% 
EncodeObjectAPI-8            3.02kB ± 0%           58.00kB ± 1%             1.23kB ± 0%               1.82kB ± 0%           1.23kB ± 0%              8.91kB ± 0%              375.82kB ± 0% 

name \ allocs/op   result/wasi-go-km.out  result/wasi-as-km.out  result/wasi-zig-km.out  result/wasi-swift-km.out  result/wasi-c-km.out  result/wasi-odin-km.out  result/wasi-dotnet-km.out 
DecodeSumVec3-8                5.00 ± 0%              5.00 ± 0%               5.00 ± 0%                32.00 ± 0%             5.00 ± 0%                6.00 ± 0%                 11.00 ± 0% 
DecodeObjectAPI-8              5.00 ± 0%              4.00 ± 0%               4.00 ± 0%                32.00 ± 0%             4.00 ± 0%                6.00 ± 0%                340.00 ± 0% 
EncodeObjectAPI-8              4.00 ± 0%              3.00 ± 0%               3.00 ± 0%                30.00 ± 0%             3.00 ± 0%                5.00 ± 0%                 40.00 ± 0% 
```

# Languages

Currently, we have focus on WebAssembly, and because of that those are the languages supported:

- AssemblyScript 0.20.16
- C/Emscripten
- C#/.NET 7
- Golang 1.19/TinyGo 0.25.0
- Odin
- Swift 5.7/SwiftWasm 5.7
- Zig 0.10

Some languages still under development, and doesn't have any backward compatibility promise. We will try 
to keep up with the latest version. Currently, the API generated and the libraries should not consider stable.

### Features

| Features | Go/TinyGo | Zig | AssemblyScript | Swift | C | C#/.NET | Odin |
| -- | -- | -- | -- | -- | -- | -- | -- |
| Performance | Good | Excellent | Good | Poor | Excellent | Horrible | Good |
| Priority | High | High | High | Low | High | Medium | Low |
| **Encoding** | | | | | | | |
| Object Encoding | ✔️ |✔️ |✔️ | ✔️ | ✔️| ✔️ | ✔️ |
| Raw Encoding | ❌ |❌ | ❌| ❌ | ❌ | ❌ | ❌ |
| Zero-Copy |❌ | ❌ |❌ | ❌ | ❌ | ❌ | ❌ |
| **Decoding** | | | | | | | |
| Object Decoding |✔️ |✔️ |✔️ | ✔️ |✔️ | ✔️ | ✔️|
| Object Re-Use |✔️ |✔️ |✔️ | ❌ |✔️ |  ✔️|  ✔️ |
| Random-Access |✔️ |✔️ |✔️ | ✔️ |✔️ | ✔️ |  ✔️|
| Zero-Copy |✔️ | ✔️ |✔️ | ❌ |✔️ |  ✔️| ✔️|
| Zero-Copy-String |✔️ | ✔️ |❌ | ✔️ |✔️ |  ❌| ✔️|
| Native Array | ✔️ |✔️ |❌ | ❌ |✔️ |  ❌️| ✔️|

# Schema

Karmem uses a custom schema language, which defines structs, enums and types.

### Example

The schema is very simple to understand and define:

```go
karmem game @packed(true) @golang.package(`km`) @assemblyscript.import(`../../assemblyscript/karmem`);

enum Team uint8 {Humans;Orcs;Zombies;Robots;Aliens;}

struct Vec3 inline {
    X float32;
    Y float32;
    Z float32;
}

struct MonsterData table {
    Pos       Vec3;
    Mana      int16;
    Health    int16;
    Name      []char;
    Team      Team;
    Inventory [<128]byte;
    Hitbox    [4]float64;
    Status    []int32;
    Path      [<128]Vec3;
}

struct Monster inline {
    Data MonsterData;
}

struct State table {
    Monsters [<2000]Monster;
}
```

### Header:

Every file must begin with: `karmem {name} [@tag()];`. Other optional tags can be defined, as shown above, it's
recommended to use the `@packed(true)` option.

### Types:

**Primitives**:

- Unsigned Integers:
  `uint8`, `uint16`, `uint32`, `uint64`
- Signed Integers:
  `int8`, `int16`, `int32,` `int64`
- Floats:
  `float32`, `float64`
- Boolean:
  `bool`
- Byte:
  `byte`, `char`

It's not possible to define optional or nullable types.

**Arrays**:

- Fixed:
  `[{Length}]{Type}` (example: `[123]uint16`, `[3]float32`)
- Dynamic:
  `[]{Type}` (example: `[]char`, `[]uint64`)
- Limited:
  `[<{Length}]{Type}` (example: `[<512]float64`, `[<42]byte`)

It's not possible to have slice of tables or slices of enums or slice of slices. However, it's possible to wrap those
types inside one inline-struct.

### Struct:

Currently, Karmem has two structs types: inline and table.

**Inline:**
Inline structs, as the name suggests, are inlined when used. That reduces the size and may improve the performance.
However, it can't have their definition changed. In order words: you can't edit the field of one inline struct
without breaking compatibility.

```go
struct Vec3 inline {
    X float32;
    Y float32;
    Z float32;
}
```

*That struct is exactly the same of `[3]float32` and will have the same serialization result. Because of that, any
change of this struct (for instance, change it to `float64` or adding new fields) will break the compatibility.*

**Tables:**
Tables can be used when backward compatibility matters. For example, tables can have new fields append at the bottom
without breaking compatibility.

```go
struct User table {
    Name     []char;
    Email    []char;
    Password []char;
}
```

Let's consider that you need another field... For tables, it's not an issue:

```go
struct User table {
    Name      []char;
    Email     []char;
    Password  []char;
    Telephone []char;
}
```

Since it's a table, you can add new fields at the bottom of the struct, and both versions are compatible between them. If
the message is sent to a client that doesn't understand the new field, it will be ignored. If one outdated client sends a
message to a newer client, the new field will have the default value (0, false, empty string, etc).

### Enums:

Enums can be used as an alias to Integers type, such as `uint8`.

```go
enum Team uint8 {
    Unknown;
    Humans;
    Orcs;
    Zombies = 255;
}
```

Enums must start with a zero value, the default value in all cases. If the value of any enum is omitted, it will use the
order of enum as value.

# Generator

Once you have a schema defined, you can generate the code. First, you need to `karmem` installed, get it from the
releases page or run it with go.

```
karmem build --assemblyscript -o "output-folder" your-schema.km
```

*If you already have Golang installed, you can use `go karmem.org/cmd/karmem build --zig -o "output-folder" your-schema.km`
instead.*

**Commands:**

**`build`**

- `--zig`: Enable generation for Zig
- `--swift`: Enable generation for Swift/SwiftWasm
- `--odin`: Enable generation for Odin
- `--golang`: Enable generation for Golang/TinyGo
- `--dotnet`: Enable generation for .NET
- `--c`: Enable generation for C
- `--assemblyscript`: Enable generation for AssemblyScript
- `-o <dir>`: Defines the output folder
- `<input-file>`: Defines the input schema

# Security

Karmem is fast and is also aimed to be secure and stable for general usage.

**Out Of Bounds**

Karmem includes bounds-checking to prevent out-of-bounds reading and avoid crashes and panics. That is something that
Google Protobuf doesn't have, and malformed content will cause panic. However, it doesn't fix all possible
vulnerabilities.

**Resource Exhaustion**

Karmem allows one pointer/offset can be re-used multiple times in the same message. Unfortunately, that behaviour makes
it possible for a short message to generate more extensive arrays than the message size. Currently, the only mitigation
for that issue is using Limited-Arrays instead of Arrays and avoiding Object-API decode.

**Data Leak**

Karmem doesn't clear the memory before encoding, that may leak information from the previous message or from system
memory itself. That can be solved using `@packed(true)` tag, as previously described. The `packed` tag will remove
padding from the message, which will prevent the leak. Alternatively, you can clear the memory before encoding,
manually.

# Limitations

Karmem have some limitations compared to other serialization libraries, such as:

**Maximum Size**

Similar to Google Protobuf and Google Flatbuffers, Karmem has a maximum size of 2GB. That is the maximum size of the entire
message, not the maximum size of each array. This limitation is due to the fact that WASM is designed to be 32-bit,
and the maximum size of 2GB seems adequate for the current needs. The current Writer doesn't enforce this limitation,
but reading a message that is bigger than 2GB will cause undefined behaviour.

**Arrays of Arrays/Tables**

Karmem doesn't support arrays of arrays or arrays of tables. However, it's possible to wrap those types inside one
inline-struct, as mentioned above. That limitation was imposed to take advantage of native arrays/slice from the
language. Most languages encapsulates the pointer and the size of the array inside a struct-like, that requires 
the size of each element to be known, consequently preventing arrays of items with variable size/strides.

**UTF-8**

Karmem only supports UTF-8 and doesn't support other encodings.

