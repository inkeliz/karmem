//go:build wazero
// +build wazero

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
	"golang.org/x/text/encoding/unicode"
	"os"
	"path/filepath"
)

func initWasm(b interface {
	Error(...any)
	Fatal(...any)
}, fn ...string) Wasm {
	w := &WasmWazero{}
	var err error
	w.runtime = wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigCompiler().WithWasmCore2())

	_, err = wasi_snapshot_preview1.Instantiate(context.Background(), w.runtime)
	if err != nil {
		b.Fatal(err)
	}

	config := wazero.NewModuleConfig().WithStdout(os.Stdout).WithStderr(os.Stdout)

	getStringImpl := func(ptr int32) string {
		len, _ := w.mainModule.Memory().ReadUint32Le(nil, uint32(ptr+(-4)))
		wtf16, _ := w.mainModule.Memory().Read(nil, uint32(ptr), len)

		b, err := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Bytes(wtf16)
		if err != nil {
			panic(err)
		}
		return string(b)
	}

	__getString := func(ptr int32) string {
		return getStringImpl(ptr)
	}

	_, err = w.runtime.NewModuleBuilder("env").
		ExportFunction("abort", func(msg, file, line, column int32) {
			var m string
			m += fmt.Sprint(`msg:`, __getString(msg))
			m += fmt.Sprint(`file:`, __getString(file))
			m += fmt.Sprint(`line:`, __getString(line))
			m += fmt.Sprint(`column:`, __getString(column))
			panic(m)
		}).
		ExportFunction("emscripten_notify_memory_growth", func(_ int32) {
		}).
		Instantiate(context.Background(), w.runtime)
	if err != nil {
		b.Fatal(err)
	}

	wasifile, err := os.ReadFile(filepath.Join("testdata", "wasi", FileWasm))
	if err != nil {
		b.Fatal(err)
	}

	compiledWasi, err := w.runtime.CompileModule(context.Background(), wasifile, wazero.NewCompileConfig())
	if err != nil {
		b.Fatal(err)
	}

	w.mainModule, err = w.runtime.InstantiateModule(context.Background(), compiledWasi, config)
	if err != nil {
		b.Fatal(err)
	}

	input, err := w.mainModule.ExportedFunction("InputMemoryPointer").Call(context.Background())
	if err != nil {
		b.Fatal(err)
	}

	if _, ok := w.mainModule.Memory().Read(nil, uint32(input[0]), uint32(len(InputMemory))); !ok {
		b.Fatal("invalid ptr", input)
	}

	w.mainModule.Memory().Write(context.Background(), uint32(input[0]), make([]byte, len(InputMemory)))

	output, err := w.mainModule.ExportedFunction("OutputMemoryPointer").Call(context.Background())
	if err != nil {
		b.Fatal(err)
	}

	if _, ok := w.mainModule.Memory().Read(nil, uint32(output[0]), uint32(len(InputMemory))); !ok {
		b.Fatal("invalid ptr", output)
	}

	w.mainModule.Memory().Write(context.Background(), uint32(output[0]), make([]byte, len(InputMemory)))

	if len(input) == 0 || len(output) == 0 || input[0] == 0 || output[0] == 0 {
		b.Fatal("invalid ptr", input, output)
	}

	functions := make(map[string]api.Function, len(fn))
	for _, fn := range fn {
		functions[fn] = w.mainModule.ExportedFunction(fn)
	}

	w.input = input[0]
	w.output = output[0]
	w.functions = functions

	return w
}

type WasmWazero struct {
	runtime    wazero.Runtime
	mainModule api.Module
	input      uint64
	output     uint64
	functions  map[string]api.Function
}

func (w *WasmWazero) Write(b []byte) bool {
	return w.mainModule.Memory().Write(context.Background(), uint32(w.input), b)
}

func (w *WasmWazero) Reader(l uint32) []byte {
	out, _ := w.mainModule.Memory().Read(context.Background(), uint32(w.output), l)
	return out
}

func (w *WasmWazero) ReaderReset(b []byte) {
	w.mainModule.Memory().Write(context.Background(), uint32(w.output), b)
}

func (w *WasmWazero) Run(s string, v ...uint64) ([]uint64, error) {
	f, ok := w.functions[s]
	if !ok || f == nil {
		return nil, errors.New("invalid function of " + s)
	}
	return f.Call(context.Background(), v...)
}

func (w *WasmWazero) Close() error {
	return w.runtime.Close(context.Background())
}
