package kmgen

import (
	"embed"
	"io"
	"strings"
	"text/template"

	"karmem.org/cmd/karmem/kmparser"
)

//go:embed *_template.*
var templateFiles embed.FS

// Generators is a list of all generators available and registered by RegisterGenerator
var Generators []Generator

// RegisterGenerator register the given Generator.
// You should use it on `init` function.
func RegisterGenerator(g Generator) {
	Generators = append(Generators, g)
}

type Generator interface {
	Start(file *kmparser.Content) (compiler Compiler, err error)
	Options() map[string]string
	Extensions() []string
	Language() string
	Finish(output io.Writer, buffer io.Reader) (err error)
}

type Compiler struct {
	Template []*template.Template
	Modules  []string
}

type TemplateFunctions struct {
	FromTags        func(s string) string
	FromStructClass func(cls kmparser.StructClass) string

	ToNamePadding func(val any, root any) string

	ToStructName   func(val any) string
	ToFieldName    func(val any) string
	ToEnumName     func(val any) string
	ToFunctionName func(val any) string

	ToDefault       func(typ kmparser.Type) string
	ToPlainDefault  func(typ kmparser.Type) string
	ToType          func(typ kmparser.Type) string
	ToPlainType     func(typ kmparser.Type) string
	ToTypeView      func(typ kmparser.Type) string
	ToPlainTypeView func(typ kmparser.Type) string
}

func NewTemplateFunctions(gen Generator, content *kmparser.Content) TemplateFunctions {
	return TemplateFunctions{
		FromTags: func(s string) string {
			def, ok := gen.Options()[s]
			if !ok {
				panic("invalid tag search")
			}

			if s == "package" {
				def = content.Name
			}

			tags := content.Tags
			name := gen.Language() + "." + s
			for i := range tags {
				if tags[i].Name == name {
					return tags[i].Value
				}
			}

			return def
		},
		FromStructClass: func(cls kmparser.StructClass) string {
			switch cls {
			case kmparser.StructClassTable:
				return "table"
			case kmparser.StructClassInline:
				return "inline"
			default:
				panic("invalid struct class")
			}
		},
		ToNamePadding: func(val any, root any) string {
			var largest int
			var name string
			switch val := val.(type) {
			case kmparser.StructField:
				name = strings.TrimSpace(val.Data.Name)
				root := root.(kmparser.Structure)
				for i := range root.Data.Fields {
					if l := len(root.Data.Fields[i].Data.Name); l > largest {
						largest = l
					}
				}
			case kmparser.EnumField:
				name = strings.TrimSpace(val.Data.Name)
				root := root.(kmparser.Enumeration)
				if root.Data.IsSequential {
					return name
				}
				for i := range root.Data.Fields {
					if l := len(root.Data.Fields[i].Data.Name); l > largest {
						largest = l
					}
				}
			default:
				panic("invalid type")
			}
			return name + strings.Repeat(" ", largest-len(name))
		},
	}
}

type TemplateData struct {
	*kmparser.Content
}

func NewTemplate(modules []string, funcs TemplateFunctions, pattern ...string) (compiler Compiler) {
	compiler.Modules = modules
	compiler.Template = make([]*template.Template, len(pattern))
	for i, v := range pattern {
		t := template.New("")
		t = t.Funcs(template.FuncMap{
			"FromTags":        funcs.FromTags,
			"FromStructClass": funcs.FromStructClass,

			"ToDefault":      funcs.ToDefault,
			"ToPlainDefault": funcs.ToPlainDefault,

			"ToPlainType":     funcs.ToPlainType,
			"ToType":          funcs.ToType,
			"ToPlainTypeView": funcs.ToPlainTypeView,
			"ToTypeView":      funcs.ToTypeView,

			"ToNamePadding": funcs.ToNamePadding,
		})
		t, err := t.ParseFS(templateFiles, v)
		if err != nil {
			panic(err)
		}
		compiler.Template[i] = t
	}
	return compiler
}

type generatorFinishCopy struct{}

func (*generatorFinishCopy) Finish(output io.Writer, buffer io.Reader) error {
	_, err := io.Copy(output, buffer)
	return err
}
