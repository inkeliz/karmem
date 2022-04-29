package kmgen

import (
	"karmem.org/cmd/karmem/kmparser"
	"os"
	"testing"
)

func TestGolangGenerator(t *testing.T) {
	path := "testdata/basic.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := kmparser.NewReader(path, f)
	k, err := r.Parser()
	if err != nil {
		t.Error(err)
	}

	gen := GolangGenerator()
	if err := gen.Start(k); err != nil {
		t.Error(err)
	}
}

func TestGolangGeneratorPath(t *testing.T) {
	path := "testdata/paths.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
		return
	}

	r := kmparser.NewReader(path, f)
	k, err := r.Parser()
	if err != nil {
		t.Error(err)
	}

	gen := GolangGenerator()
	if err := gen.Start(k); err != nil {
		t.Error(err)
	}
}
