package kmgen

import (
	"fmt"
	"go/format"
	"io"
	"karmem.org/cmd/karmem/kmparser"
)

type Golang struct {
	content *kmparser.Content
}

func init() { RegisterGenerator(GolangGenerator()) }

func GolangGenerator() Generator {
	return &Golang{}
}

func (gen *Golang) Start(file *kmparser.Content) (compiler Compiler, err error) {
	gen.content = file
	return NewTemplate(
		[]string{`header`, `enums`, `enums_builder`, `struct`, `struct_builder`},
		gen.functions(),
		"golang_template.*",
	), nil
}

func (gen *Golang) functions() (f TemplateFunctions) {
	f = NewTemplateFunctions(gen, gen.content)

	f.ToDefault = func(typ kmparser.Type) string {
		t := typ.PlainSchema
		switch t {
		case "bool":
			return "false"
		case "char":
			return `""`
		default:
			return "0"
		}
	}

	f.ToPlainType = func(typ kmparser.Type) string {
		t := typ.PlainSchema
		switch t {
		case "byte":
			return "byte"
		case "bool":
			return "bool"
		case "char":
			return "string"
		case "uint8":
			return "uint8"
		case "uint16":
			return "uint16"
		case "uint32":
			return "uint32"
		case "uint64":
			return "uint64"
		case "int8":
			return "int8"
		case "int16":
			return "int16"
		case "int32":
			return "int32"
		case "int64":
			return "int64"
		case "float32":
			return "float32"
		case "float64":
			return "float64"
		default:
			return t
		}
	}

	f.ToType = func(typ kmparser.Type) string {
		p := f.ToPlainType(typ)
		switch {
		case typ.PlainSchema == "char":
			return p
		case typ.Model == kmparser.TypeModelSlice, typ.Model == kmparser.TypeModelSliceLimited:
			return fmt.Sprintf(`[]%s`, p)
		case typ.Model == kmparser.TypeModelArray:
			return fmt.Sprintf(`[%d]%s`, typ.Length, p)
		default:
			return p
		}
	}

	f.ToPlainTypeView = func(typ kmparser.Type) string {
		p := f.ToPlainType(typ)
		switch typ.Format {
		case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
			return fmt.Sprintf(`*%sViewer`, p)
		default:
			return p
		}
	}

	f.ToTypeView = func(typ kmparser.Type) string {
		p := f.ToPlainTypeView(typ)
		if typ.PlainSchema == "char" {
			return p
		}
		switch typ.Model {
		case kmparser.TypeModelArray, kmparser.TypeModelSlice, kmparser.TypeModelSliceLimited:
			switch typ.Format {
			case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
				return fmt.Sprintf(`[]%sViewer`, f.ToPlainType(typ))
			default:
				return fmt.Sprintf(`[]%s`, p)
			}
		default:
			return p
		}
	}

	return f
}

func (gen *Golang) Language() string { return "golang" }

func (gen *Golang) Options() map[string]string {
	return map[string]string{"package": "", "import": "karmem.org/golang"}
}

func (gen *Golang) Extensions() []string { return []string{`.go`} }

func (gen *Golang) Finish(output io.Writer, buffer io.Reader) error {
	buf, err := io.ReadAll(buffer)
	if err != nil {
		return err
	}
	b, err := format.Source(buf)
	if err != nil {
		return err
	}
	_, err = output.Write(b)
	return err
}
