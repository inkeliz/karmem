package kmgen

import (
	"fmt"
	"karmem.org/cmd/karmem/kmparser"
)

type Swift struct {
	content *kmparser.Content
	generatorFinishCopy
}

func init() { RegisterGenerator(SwiftGenerator()) }

func SwiftGenerator() Generator {
	return &Swift{}
}

func (gen *Swift) Start(file *kmparser.Content) (compiler Compiler, err error) {
	gen.content = file
	return NewTemplate(
		[]string{`header`, `enums`, `enums_builder`, `struct`, `struct_builder`},
		gen.functions(),
		"swift_template.*",
	), nil
}

func (gen *Swift) functions() (f TemplateFunctions) {
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
			return "UInt8"
		case "bool":
			return "Bool"
		case "char":
			return "[UInt8]"
		case "uint8":
			return "UInt8"
		case "uint16":
			return "UInt16"
		case "uint32":
			return "UInt32"
		case "uint64":
			return "UInt64"
		case "int8":
			return "Int8"
		case "int16":
			return "Int16"
		case "int32":
			return "Int32"
		case "int64":
			return "Int64"
		case "float32":
			return "Float"
		case "float64":
			return "Double"
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
			fallthrough
		case typ.Model == kmparser.TypeModelArray:
			return fmt.Sprintf(`[%s]`, p)
		default:
			return p
		}
	}

	f.ToPlainTypeView = func(typ kmparser.Type) string {
		p := f.ToPlainType(typ)
		switch typ.Format {
		case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
			return fmt.Sprintf(`%sViewer`, p)
		default:
			return p
		}
	}

	f.ToTypeView = func(typ kmparser.Type) string {
		p := f.ToPlainTypeView(typ)
		if typ.PlainSchema == "char" {
			return fmt.Sprintf(`karmem.Slice<%s>`, "UInt8")
		}
		switch typ.Model {
		case kmparser.TypeModelArray, kmparser.TypeModelSlice, kmparser.TypeModelSliceLimited:
			switch typ.Format {
			case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
				return fmt.Sprintf(`karmem.Slice<%sViewer>`, f.ToPlainType(typ))
			default:
				return fmt.Sprintf(`karmem.Slice<%s>`, p)
			}
		default:
			return p
		}
	}

	return f
}

func (gen *Swift) Language() string { return "swift" }

func (gen *Swift) Options() map[string]string {
	return map[string]string{"import": "karmem"}
}

func (gen *Swift) Extensions() []string { return []string{`.swift`} }
