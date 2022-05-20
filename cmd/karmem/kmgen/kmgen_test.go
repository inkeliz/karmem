package kmgen

import (
	"bytes"
	"karmem.org/cmd/karmem/kmparser"
	"os"
	"strings"
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
			k := k.Clone()

			k.Headers["Package"] = k.Header.Name
			for o, d := range gen.Options() {
				if v := k.Header.GetTag(gen.Language() + "." + o); v != nil {
					k.Headers[strings.Title(o)] = strings.TrimSpace(v.Value)
				} else {
					if d != "" {
						k.Headers[strings.Title(o)] = strings.TrimSpace(d)
					}
				}
			}

			//fmt.Println(k.Headers["Package"], len(k.Headers["Package"]))

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
