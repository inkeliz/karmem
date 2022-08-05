package kmgen

import (
	"fmt"

	"karmem.org/cmd/karmem/kmparser"
)

type DotNet struct {
	content *kmparser.Content
	generatorFinishCopy
}

func init() { RegisterGenerator(DotNetGenerator()) }

func DotNetGenerator() Generator {
	return &DotNet{}
}

func (gen *DotNet) Start(file *kmparser.Content) (compiler Compiler, err error) {
	gen.content = file
	return NewTemplate(
		[]string{`header`, `enums`, `enums_builder`, `struct`, `struct_builder`},
		gen.functions(),
		"dotnet_template.*",
	), nil
}

func (gen *DotNet) functions() (f TemplateFunctions) {
	f = NewTemplateFunctions(gen, gen.content)

	f.ToDefault = func(typ kmparser.Type) string {
		if typ.IsString() {
			return f.ToPlainDefault(typ)
		}
		switch typ.Model {
		case kmparser.TypeModelArray, kmparser.TypeModelSlice, kmparser.TypeModelSliceLimited:
			return fmt.Sprintf("new List<%s>()", f.ToPlainType(typ))
		}
		switch typ.Format {
		case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
			return fmt.Sprintf("new %s()", f.ToPlainType(typ))
		}
		return f.ToPlainDefault(typ)
	}

	f.ToPlainDefault = func(typ kmparser.Type) string {
		t := typ.PlainSchema
		switch t {
		case "bool":
			return "false"
		case "char":
			return `""`
		default:
			switch typ.Format {
			case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
				return fmt.Sprintf("new %s()", f.ToPlainType(typ))
			}
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
			return "byte"
		case "uint16":
			return "ushort"
		case "uint32":
			return "uint"
		case "uint64":
			return "ulong"
		case "int8":
			return "sbyte"
		case "int16":
			return "short"
		case "int32":
			return "int"
		case "int64":
			return "long"
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
			fallthrough
		case typ.Model == kmparser.TypeModelArray:
			return fmt.Sprintf(`List<%s>`, p)
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
			switch typ.Format {
			case kmparser.TypeFormatStruct, kmparser.TypeFormatTable:
				return fmt.Sprintf(`Karmem.Slice<%sViewer>`, f.ToPlainType(typ))
			default:
				return fmt.Sprintf(`Karmem.Slice<%s>`, f.ToPlainType(typ))
			}
		default:
			return p
		}
	}

	return f
}

func (gen *DotNet) Language() string { return "dotnet" }

func (gen *DotNet) Options() map[string]string {
	return map[string]string{"package": "", "import": "karmem.org/dotnet"}
}

func (gen *DotNet) Extensions() []string { return []string{`.cs`} }
