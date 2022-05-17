#flatc -g --gen-object-api --gen-onefile --no-includes --go-namespace "fbs" -o "fbs" testdata/game.fbs
go run karmem.org/cmd/karmem build --golang -o "km" testdata/game.km
go run karmem.org/cmd/karmem build --assemblyscript -o "km" testdata/game.km
go run karmem.org/cmd/karmem build --zig -o "km" testdata/game.km
go run karmem.org/cmd/karmem build --swift -o "km" testdata/game.km
go run karmem.org/cmd/karmem build --c -o "km" testdata/game.km

go test -tags fbs -v -bench=. -benchmem -benchtime=5s -count 5 . > result/fbs.out
go test -tags km -v -bench=. -benchmem -benchtime=5s -count 5 . > result/km.out

tinygo build -target "wasi" -tags "km" -o "testdata/wasi/go.wasi" -gc conservative --no-debug -opt 2 -scheduler none -wasm-abi generic .
go test -tags wasi,wazero,km,golang -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-go-km.out

tinygo build -target "wasi" -tags "km" -o "testdata/wasi/go.wasi" -gc leaking --no-debug -opt 2 -scheduler none -wasm-abi generic .
go test -tags wasi,wazero,km,golang -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-go-leaking-km.out

tinygo build -target "wasi" -tags "fbs" -o "testdata/wasi/go.wasi" -gc conservative --no-debug -opt 2 -scheduler none -wasm-abi generic .
go test -tags wasi,wazero,fbs,golang -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-go-fbs.out

tinygo build -target "wasi" -tags "fbs" -o "testdata/wasi/go.wasi" -gc leaking --no-debug -opt 2 -scheduler none -wasm-abi generic .
go test -tags wasi,wazero,fbs,golang -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-go-leaking-fbs.out

../node_modules/.bin/asc main.ts -Ospeed --optimizeLevel 3 --target release --disable bulk-memory --zeroFilledMemory --exportStart "_start" --outFile "./testdata/wasi/as.wasi"
go test -tags wasi,wazero,km,assemblyscript -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-as-km.out

zig build install -Dtarget="wasm32-wasi" -Drelease-fast
go test -tags wasi,wazero,km,zig -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-zig-km.out

xcrun --toolchain swiftwasm swift build --triple wasm32-unknown-wasi -v -c release
xcrun --toolchain swiftwasm swiftc -L ".build/wasm32-unknown-wasi/release" -o "testdata/wasi/swift.wasi" -module-name benchmark -emit-executable @.build/wasm32-unknown-wasi/release/benchmark.product/Objects.LinkFileList -target wasm32-unknown-wasi -sdk /Library/Developer/Toolchains/swift-wasm-5.6.0-RELEASE.xctoolchain/usr/share/wasi-sysroot -L /Library/Developer/Toolchains/swift-wasm-5.6.0-RELEASE.xctoolchain/usr/lib -O -Xlinker "--export=_start" -Xlinker "--export=InputMemoryPointer" -Xlinker "--export=OutputMemoryPointer" -Xlinker "--export=KBenchmarkEncodeObjectAPI"  -Xlinker "--export=KBenchmarkDecodeObjectAPI" -Xlinker "--export=KBenchmarkDecodeSumVec3"
go test -tags wasi,wazero,km,swift -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-swift-km.out

emcc c/main.c -o testdata/wasi/c.wasm -s "EXPORTED_FUNCTIONS=_InputMemoryPointer,_OutputMemoryPointer,__start,_KBenchmarkDecodeSumVec3,_KBenchmarkDecodeObjectAPI,_KBenchmarkEncodeObjectAPI, _KBenchmarkDecodeSumVec3" --no-entry -sALLOW_MEMORY_GROWTH -O3 -flto
go test -tags wasi,wazero,km,c -v -bench=. -benchmem -benchtime=5s -count 5 . > result/wasi-c-km.out

benchstat result/fbs.out result/km.out
benchstat result/wasi-go-km.out result/wasi-go-fbs.out
benchstat result/wasi-go-leaking-km.out result/wasi-go-leaking-fbs.out
benchstat result/wasi-go-km.out result/wasi-as-km.out result/wasi-zig-km.out result/wasi-swift-km.out result/wasi-c-km.out
