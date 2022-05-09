//go:build wazero
// +build wazero

package main

import (
	"context"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi"
	"os"
	"path/filepath"
)

func initWasm(b interface {
	Error(...any)
	Fatal(...any)
}, fn ...string) Wasm {
	w := &WasmWazero{}
	var err error
	runtime := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigJIT().WithWasmCore2())
	w.modules[0], err = wasi.InstantiateSnapshotPreview1(context.Background(), runtime)
	if err != nil {
		b.Fatal(err)
	}

	config := wazero.NewModuleConfig().WithStdout(os.Stdout).WithStderr(os.Stdout)

	w.modules[1], err = runtime.NewModuleBuilder("env").ExportFunction("abort", func(_, _, _, _ int32) { panic("xxx") }).Instantiate(context.Background())
	if err != nil {
		b.Fatal(err)
	}

	wasifile, err := os.ReadFile(filepath.Join("testdata", "wasi", FileWasm))
	if err != nil {
		b.Fatal(err)
	}

	w.mainModule, err = runtime.InstantiateModuleFromCodeWithConfig(context.Background(), wasifile, config)
	if err != nil {
		b.Fatal(err)
	}

	input, err := w.mainModule.ExportedFunction("InputMemoryPointer").Call(context.Background())
	if err != nil {
		b.Fatal(err)
	}

	w.mainModule.Memory().Write(context.Background(), uint32(input[0]), make([]byte, len(InputMemory)))

	output, err := w.mainModule.ExportedFunction("OutputMemoryPointer").Call(context.Background())
	if err != nil {
		b.Fatal(err)
	}

	w.mainModule.Memory().Write(context.Background(), uint32(output[0]), make([]byte, len(InputMemory)))

	if len(input) == 0 || len(output) == 0 || input[0] == 0 || output[0] == 0 {
		b.Fatal("invalid ptr")
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
	modules    [5]api.Module
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
	return w.functions[s].Call(context.Background(), v...)
}

func (w *WasmWazero) Close() error {
	for _, m := range w.modules {
		if m != nil {
			m.Close(context.Background())
		}
	}
	return w.mainModule.Close(context.Background())
}
