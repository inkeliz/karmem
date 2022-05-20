package kmgen

import (
	"fmt"
	"karmem.org/cmd/karmem/kmparser"
)

type C struct {
	content *kmparser.Content
	generatorFinishCopy
}

func init() { RegisterGenerator(CGenerator()) }

func CGenerator() Generator {
	return &C{}
}

func (gen *C) Start(file *kmparser.Content) (compiler Compiler, err error) {
	gen.content = file
	return NewTemplate(
		[]string{`header`, `enums`, `enums_builder`, `struct_builder`, `struct`},
		gen.functions(),
		"c_template.*",
	), nil
}

func (gen *C) functions() (f TemplateFunctions) {
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
		if typ.Format == kmparser.TypeFormatEnum {
			return fmt.Sprintf(`Enum%s`, t)
		}
		switch t {
		case "byte":
			return "uint8_t"
		case "bool":
			return "bool"
		case "char":
			return "uint8_t *"
		case "uint8":
			return "uint8_t"
		case "uint16":
			return "uint16_t"
		case "uint32":
			return "uint32_t"
		case "uint64":
			return "uint64_t"
		case "int8":
			return "int8_t"
		case "int16":
			return "int16_t"
		case "int32":
			return "int32_t"
		case "int64":
			return "int64_t"
		case "float32":
			return "float"
		case "float64":
			return "double"
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
			return fmt.Sprintf(`%s *`, p)
		case typ.Model == kmparser.TypeModelArray:
			return fmt.Sprintf(`%s`, p)
		default:
			return p
		}
	}

	f.ToPlainTypeView = func(typ kmparser.Type) string {
		p := f.ToPlainType(typ)
		switch typ.Format {
		case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
			return fmt.Sprintf(`%sViewer *`, p)
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
				return fmt.Sprintf(`%sViewer *`, f.ToPlainType(typ))
			default:
				return fmt.Sprintf(`%s *`, p)
			}
		default:
			return p
		}
	}

	return f
}

func (gen *C) Language() string { return "c" }

func (gen *C) Options() map[string]string {
	return map[string]string{"prefix": "", "import": "karmem.h"}
}

func (gen *C) Extensions() []string { return []string{`.h`} }
