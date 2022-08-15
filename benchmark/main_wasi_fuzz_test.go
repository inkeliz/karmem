//go:build !fbs && (wazero || wasmer || tcp) && !nofuzz
// +build !fbs
// +build wazero wasmer tcp
// +build !nofuzz

package main

import (
	"crypto/rand"
	"math"
	"math/big"
	"testing"

	"benchmark.karmem.org/km"
	fuzz "github.com/google/gofuzz"
	"github.com/r3labs/diff/v3"
	"golang.org/x/crypto/blake2b"
	karmem "karmem.org/golang"
)

func init() {
	initEncode()
}

type Entropy struct {
	b blake2b.XOF
}

func NewEntropy(b []byte) *Entropy {
	k := blake2b.Sum512(b)
	e, _ := blake2b.NewXOF(blake2b.OutputLengthUnknown, k[:])
	return &Entropy{b: e}
}

func (e *Entropy) Int63() int64 {
	i, _ := rand.Int(e.b, big.NewInt(math.MaxInt))
	return i.Int64()
}

func (e *Entropy) Seed(seed int64) { return }

func FuzzContent(f *testing.F) {
	m := initBridge(f, "KBenchmarkEncodeObjectAPI", "KBenchmarkDecodeObjectAPI")
	defer func() {
		if err := m.Close(); err != nil {
			f.Fatal(err)
		}
	}()

	fx := fuzz.New().NilChance(.5)

	writer := karmem.NewWriter(len(OutputMemory))

	var expected km.Monsters
	var result km.Monsters
	f.Fuzz(func(t *testing.T, b []byte) {
		writer.Reset()
		fx.RandSource(NewEntropy(b))
		fx.Fuzz(&expected)

		if _, err := expected.WriteAsRoot(writer); err != nil {
			t.Error(err)
		}

		encoded := writer.Bytes()
		if len(encoded) > len(OutputMemory) {
			// Avoid growing on WASM
			return
		}
		m.Write(encoded)

		_, err := m.Run(FunctionKBenchmarkDecodeObjectAPI, uint64(len(encoded)))
		if err != nil {
			t.Fatal(err)
		}

		_, err = m.Run(FunctionKBenchmarkEncodeObjectAPI)
		if err != nil {
			t.Fatal(err)
		}

		reader := karmem.NewReader(m.Reader(uint32(len(InputMemory))))
		result.ReadAsRoot(reader)

		di, err := diff.NewDiffer()
		if err != nil {
			t.Errorf("diff lib error: %v", err)
			return
		}

		changes, err := di.Diff(expected, result)
		if err != nil {
			t.Errorf("diff lib error: %v", err)
			return
		}

		if len(changes) > 0 {
			for _, x := range changes {
				t.Errorf("%s %d %s %d", expected.Monsters[0].Data.Name, len([]byte(expected.Monsters[0].Data.Name)), result.Monsters[0].Data.Name, len([]byte(result.Monsters[0].Data.Name)))
				t.Error(x.Type, x.Path)
			}
			t.Fatal("changes detected")
		}
	})
}

var x uint64 = 0

func FuzzRandom(f *testing.F) {
	m := initBridge(f, "KBenchmarkEncodeObjectAPI", "KBenchmarkDecodeObjectAPI")
	defer func() {
		if err := m.Close(); err != nil {
			f.Fatal(err)
		}
	}()
	f.Add(encode())

	clear := make([]byte, len(InputMemory))
	f.Fuzz(func(t *testing.T, b []byte) {
		if len(b) > len(InputMemory) {
			return
		}
		if !m.Write(b) {
			t.Fatal("impossible to write")
		}
		m.ReaderReset(clear)
		if _, err := m.Run(FunctionKBenchmarkDecodeObjectAPI, uint64(len(b))); err != nil {
			t.Fatal(err)
		}
	})
}
