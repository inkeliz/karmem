//go:build wazero
// +build wazero

package main

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/tetratelabs/wazero/imports/assemblyscript"
	"github.com/tetratelabs/wazero/imports/emscripten"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func initBridge(b interface {
	Error(...any)
	Fatal(...any)
}, fn ...string,
) Bridge {
	w := &WasmWazero{}
	var err error
	w.runtime = wazero.NewRuntime(context.Background())

	_, err = wasi_snapshot_preview1.Instantiate(context.Background(), w.runtime)
	if err != nil {
		b.Fatal(err)
	}

	config := wazero.NewModuleConfig().WithStdout(os.Stdout).WithStderr(os.Stdout)

	envBuilder := w.runtime.NewHostModuleBuilder("env")
	emscripten.NewFunctionExporter().ExportFunctions(envBuilder)
	assemblyscript.NewFunctionExporter().ExportFunctions(envBuilder)
	_, err = envBuilder.Instantiate(context.Background())
	if err != nil {
		b.Fatal(err)
	}

	wasifile, err := os.ReadFile(filepath.Join("testdata", "wasi", FileWasm))
	if err != nil {
		b.Fatal(err)
	}

	compiledWasi, err := w.runtime.CompileModule(context.Background(), wasifile)
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

	if _, ok := w.mainModule.Memory().Read(uint32(input[0]), uint32(len(InputMemory))); !ok {
		b.Fatal("invalid ptr", input)
	}

	w.mainModule.Memory().Write(uint32(input[0]), make([]byte, len(InputMemory)))

	output, err := w.mainModule.ExportedFunction("OutputMemoryPointer").Call(context.Background())
	if err != nil {
		b.Fatal(err)
	}

	if _, ok := w.mainModule.Memory().Read(uint32(output[0]), uint32(len(InputMemory))); !ok {
		b.Fatal("invalid ptr", output)
	}

	w.mainModule.Memory().Write(uint32(output[0]), make([]byte, len(InputMemory)))

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
	return w.mainModule.Memory().Write(uint32(w.input), b)
}

func (w *WasmWazero) Reader(l uint32) []byte {
	out, ok := w.mainModule.Memory().Read(uint32(w.output), l)
	if !ok {
		return nil
	}
	return out
}

func (w *WasmWazero) ReaderReset(b []byte) {
	w.mainModule.Memory().Write(uint32(w.output), b)
}

func (w *WasmWazero) Run(s Functions, v ...uint64) ([]uint64, error) {
	f, ok := w.functions[s.String()]
	if !ok || f == nil {
		return nil, errors.New("invalid function of " + s.String())
	}
	return f.Call(context.Background(), v...)
}

func (w *WasmWazero) Close() error {
	err := w.runtime.Close(context.Background())
	if err != nil {
		return err
	}
	*w = WasmWazero{}
	return nil
}
