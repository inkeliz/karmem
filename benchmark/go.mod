module benchmark.karmem.org

go 1.18

replace karmem.org => ../

require (
	github.com/google/flatbuffers v2.0.6+incompatible
	github.com/google/gofuzz v1.2.0
	github.com/r3labs/diff/v3 v3.0.0
	github.com/tetratelabs/wazero v0.0.0-20220506015640-0561190cb9af
	golang.org/x/crypto v0.0.0-20220507011949-2cf3adece122
	karmem.org v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/perf v0.0.0-20220411212318-84e58bfe0a7e // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.6 // indirect
)
