package kmgen

import (
	"fmt"
	"karmem.org/cmd/karmem/kmparser"
)

type AssemblyScript struct {
	content *kmparser.Content
	generatorFinishCopy
}

func init() { RegisterGenerator(AssemblyScriptGenerator()) }

func AssemblyScriptGenerator() Generator {
	return &AssemblyScript{}
}

func (gen *AssemblyScript) Start(file *kmparser.Content) (compiler Compiler, err error) {
	gen.content = file
	return NewTemplate(
		[]string{`header`, `enums`, `enums_builder`, `struct`, `struct_builder`},
		gen.functions(),
		"assemblyscript_template.*",
	), nil
}

func (gen *AssemblyScript) functions() (f TemplateFunctions) {
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
			return "u8"
		case "bool":
			return "bool"
		case "char":
			return "string"
		case "uint8":
			return "u8"
		case "uint16":
			return "u16"
		case "uint32":
			return "u32"
		case "uint64":
			return "u64"
		case "int8":
			return "i8"
		case "int16":
			return "i16"
		case "int32":
			return "i32"
		case "int64":
			return "i64"
		case "float32":
			return "f32"
		case "float64":
			return "f64"
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
			return fmt.Sprintf(`Array<%s>`, p)
		case typ.Model == kmparser.TypeModelArray:
			return fmt.Sprintf(`StaticArray<%s>`, p)
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
			return p
		}
		switch typ.Model {
		case kmparser.TypeModelArray, kmparser.TypeModelSlice, kmparser.TypeModelSliceLimited:
			return fmt.Sprintf(`karmem.Slice<%s>`, p)
		default:
			return p
		}
	}

	return f
}

func (gen *AssemblyScript) Language() string { return "assemblyscript" }

func (gen *AssemblyScript) Options() map[string]string {
	return map[string]string{"import": "karmem/assemblyscript/karmem"}
}

func (gen *AssemblyScript) Extensions() []string { return []string{`.ts`} }
