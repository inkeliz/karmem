package kmgen

import (
	"bytes"
	"karmem.org/cmd/karmem/kmparser"
	"os"
	"testing"
)

func TestGolangGenerator(t *testing.T) {
	path := []string{"testdata/basic.km", "testdata/paths.km"}
	for _, path := range path {
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

		for _, gen := range Generators {
			compiler, err := gen.Start(k)
			if err != nil {
				t.Error(err)
				return
			}
			for _, c := range compiler.Template {
				var buffer bytes.Buffer
				var output bytes.Buffer

				for _, n := range compiler.Modules {
					if err := c.ExecuteTemplate(&buffer, n, k); err != nil {
						t.Error(err)
						return
					}
				}

				if err := gen.Finish(&output, &buffer); err != nil {
					t.Error(err)
					return
				}
			}
		}
	}
}
