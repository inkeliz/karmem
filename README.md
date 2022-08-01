# KARMEM

Karmem is a fast binary serialization format. The priority of Karmem is to be
easy to use while been fast as possible. It's optimized to take Golang and
TinyGo's maximum performance and is efficient for repeatable reads, reading
different content of the same type. Karmem has proven to be ten times faster
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
a good alternative but lacks implementation for Zig and AssemblyScript, which is top-priority, it also have more
allocations and the generated API is harder to use, compared than Karmem.

# Usage

That is a small example of how use Karmem.

### Schema

```go
karmem app @golang.package(`app`);  
  
enum SocialNetwork uint8 { Unknown; Facebook; Instagram; Twitter; TikTok; }  
  
struct ProfileData table {  
    Network SocialNetwork;  
    Username []char;  
    ID uint64;  
}  
  
struct Profile inline {  
    Data ProfileData;  
}  
  
struct AccountData table {  
    ID uint64;  
    Email []char;  
    Profiles []Profile;  
}
```

Generate the code using `go run karmem.org/cmd/karmem build --golang -o "km" app.km`.

### Encoding

In order to encode, use should create an native struct and then encode it.

```go
var writerPool = sync.Pool{New: func() any { return karmem.NewWriter(1024) }}

func main() {
	var writer = writerPool.Get().(*karmem.Writer)

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
		fmt.Println(string(profiles[i].Data(reader).Username(reader)))
	}
}
```

Notice, we use `NewAccountDataViewer`, any `Viewer` is just a Viewer, and doesn't copy the backend data.

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
name               flatbuffers/op    karmem/op    delta
EncodeObjectAPI-8    1.46ms ± 0%    0.32ms ± 0%   -78.22%  (p=0.008 n=5+5)
DecodeObjectAPI-8    2.16ms ± 0%    0.15ms ± 0%   -93.14%  (p=0.008 n=5+5)
DecodeSumVec3-8       887µs ± 1%      99µs ± 1%   -88.86%  (p=0.008 n=5+5)

name               flatbuffers/op   karmem/op   delta
EncodeObjectAPI-8    12.1kB ± 0%     0.0kB       -100.00%  (p=0.008 n=5+5)
DecodeObjectAPI-8    2.74MB ± 0%    0.03MB ± 0%   -98.83%  (p=0.008 n=5+5)
DecodeSumVec3-8       0.00B          0.00B           ~     (all equal)    

name               flatbuffers/op  karmem/op  delta
EncodeObjectAPI-8     1.00k ± 0%     0.00k       -100.00%  (p=0.008 n=5+5)
DecodeObjectAPI-8      108k ± 0%        1k ± 0%   -99.07%  (p=0.008 n=5+5)
DecodeSumVec3-8        0.00           0.00           ~     (all equal)
```

**WebAssembly on Wazero (MacOS/ARM64 - M1):**

```
name               flatbuffers/op    karmem/op    delta
EncodeObjectAPI-8    10.1ms ± 0%     2.5ms ± 0%  -75.27%  (p=0.016 n=4+5)
DecodeObjectAPI-8    31.1ms ± 0%     1.2ms ± 0%  -96.18%  (p=0.008 n=5+5)
DecodeSumVec3-8      4.44ms ± 0%    0.47ms ± 0%  -89.41%  (p=0.008 n=5+5)

name               flatbuffers/op   karmem/op   delta
EncodeObjectAPI-8    3.02kB ± 0%    3.02kB ± 0%     ~     (all equal)
DecodeObjectAPI-8    2.16MB ± 0%    0.01MB ± 0%  -99.45%  (p=0.008 n=5+5)
DecodeSumVec3-8      1.25kB ± 0%    1.25kB ± 0%     ~     (all equal)

name               flatbuffers/op  karmem/op  delta
EncodeObjectAPI-8      4.00 ± 0%      4.00 ± 0%     ~     (all equal)
DecodeObjectAPI-8      5.00 ± 0%      5.00 ± 0%     ~     (all equal)
DecodeSumVec3-8        5.00 ± 0%      5.00 ± 0%     ~     (all equal)
```

### Raw-Struct vs Karmem

The performance is nearly the same when comparing reading non-serialized data from a native struct and reading it from a
karmem-serialized data.

**Native (MacOS/ARM64 - M1):**

```
name             old time/op    new time/op    delta
DecodeSumVec3-8    93.7µs ± 0%    98.8µs ± 1%  +5.38%  (p=0.008 n=5+5)

name             old alloc/op   new alloc/op   delta
DecodeSumVec3-8     0.00B          0.00B         ~     (all equal)

name             old allocs/op  new allocs/op  delta
DecodeSumVec3-8      0.00           0.00         ~     (all equal)
```

### Karmem vs Karmem

That is an comparison with all supported languages.

**WebAssembly on Wazero (MacOS/ARM64 - M1):**

```
name \ time/op     result/wasi-go-km.out  result/wasi-as-km.out  result/wasi-zig-km.out  result/wasi-c-km.out  result/wasi-swift-km.out
DecodeSumVec3-8               470µs ± 0%             932µs ± 0%              231µs ± 0%            230µs ± 0%              97822µs ± 5%
DecodeObjectAPI-8            1.19ms ± 0%            3.70ms ± 0%             0.62ms ± 0%           0.56ms ± 0%              74.72ms ± 4%
EncodeObjectAPI-8            2.52ms ± 0%            2.98ms ± 2%             0.71ms ± 0%           0.67ms ± 0%              42.45ms ± 7%

name \ alloc/op    result/wasi-go-km.out  result/wasi-as-km.out  result/wasi-zig-km.out  result/wasi-c-km.out  result/wasi-swift-km.out
DecodeSumVec3-8              1.25kB ± 0%           12.72kB ± 0%             1.25kB ± 0%           1.25kB ± 0%               2.99kB ± 0%
DecodeObjectAPI-8            11.9kB ± 1%            74.2kB ± 0%            164.3kB ± 0%            1.2kB ± 0%              291.7kB ± 3%
EncodeObjectAPI-8            3.02kB ± 0%           38.38kB ± 0%             1.23kB ± 0%           1.23kB ± 0%               2.98kB ± 0%

name \ allocs/op   result/wasi-go-km.out  result/wasi-as-km.out  result/wasi-zig-km.out  result/wasi-c-km.out  result/wasi-swift-km.out
DecodeSumVec3-8                5.00 ± 0%              5.00 ± 0%               5.00 ± 0%             5.00 ± 0%                35.00 ± 0%
DecodeObjectAPI-8              5.00 ± 0%              4.00 ± 0%               4.00 ± 0%             4.00 ± 0%                35.00 ± 0%
EncodeObjectAPI-8              4.00 ± 0%              3.00 ± 0%               3.00 ± 0%             3.00 ± 0%                33.00 ± 0%
```

# Languages

Currently, we have focus on WebAssembly, and because of that those are the languages supported:

- AssemblyScript
- Golang/TinyGo
- ~~Swift/SwiftWasm~~
- Zig
- C

### Features

| Features | Golang | Zig | AssemblyScript | Swift | C |
|--|-- | -- | --| -- | -- |
| Performance | Good | Excellent | Good | Horrible | Excellent |
| Priority | High | High | High | Low | High |
| **Encoding** | | | | | |
| Object Encoding | ✔️ |✔️ |✔️ | ✔️ | ✔️|
| Raw Encoding | ❌ |❌ | ❌| ❌ | ❌ |
| Zero-Copy |❌ | ❌ |❌ | ❌ | ❌ |
| **Decoding** | | | | | |
| Object Decoding |✔️ |✔️ |✔️ | ✔️ |✔️ |
| Object Re-Use |✔️ |✔️ |✔️ | ❌ |✔️ |
| Random-Access |✔️ |✔️ |✔️ | ✔️ |✔️ |
| Zero-Copy |✔️ | ✔️ |✔️ | ❌ |✔️ |
| Native Array | ✔️ |✔️ |❌ | ❌ |✔️ |

# Schema

Karmem uses a custom schema language, which defines structs, enums and types.

### Example

The schema is very simple to understand and define:

```go
karmem game @golang.package(`km`) @assemblyscript.import(`../../assemblyscript/karmem`);

enum Team uint8 {Humans;Orcs;Zombies;Robots;Aliens;}

struct Vec3 inline {
    X float32;
    Y float32;
    Z float32;
}

struct MonsterData table {
    Pos Vec3;
    Mana int16;
    Health int16;
    Name []char;
    Team Team;
    Inventory [<128]byte;
    Hitbox [4]float64;
    Status []int32;
    Path [<128]Vec3;
}

struct Monster inline {
    Data MonsterData;
}

struct State table {
    Monsters [<2000]Monster;
}
```

### Header:

Every file must begin with: `karmem {name};`, other optional options can be defined, as shown above.

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

It's not possible to defined optional or nullable types.

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
However, it can't have their definition changed. In order words: you can't edit the description of one inline struct
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
Name []char;
Email []char;
Password []char;
}
```

Let's consider that you need another field... For tables, it's not an issue:

```go
struct User table {
Name []char;
Email []char;
Password []char;
Telephone []char;
}
```

Since it's a table, you can add new fields at the bottom of the struct, and both versions are compatible between them.

### Enums:

Enums can be used as an alias to Integers type, such as `uint8`.

```
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

*If you already have Golang installed, you can use `go karmem.org/cmd build --zig -o "output-folder" your-schema.km`
instead.*

**Commands:**

**`build`**

- `--zig`: Enable generation for Zig
- `--golang`: Enable generation for Golang
- `--assemblyscript`: Enable generation for AssemblyScript
- `--swift`: Enable generation for Swift
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




