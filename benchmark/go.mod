module benchmark.karmem.org

go 1.18

replace karmem.org => ../

require (
	github.com/google/flatbuffers v2.0.6+incompatible
	github.com/google/gofuzz v1.2.0
	github.com/r3labs/diff/v3 v3.0.0
	github.com/tetratelabs/wazero v1.0.0-pre.1.0.20220929053752-9a623c4f88f3
	golang.org/x/crypto v0.0.0-20220513210258-46612604a0f9
	golang.org/x/text v0.3.7
	karmem.org v1.0.0
)

require (
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	google.golang.org/appengine v1.6.6 // indirect
)
