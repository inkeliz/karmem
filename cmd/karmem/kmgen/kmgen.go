package kmgen

import (
	"embed"
	"io"
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
	FromTags func(s string) string

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
				def = content.Module
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
			"FromTags": funcs.FromTags,

			"ToDefault":      funcs.ToDefault,
			"ToPlainDefault": funcs.ToPlainDefault,

			"ToPlainType":     funcs.ToPlainType,
			"ToType":          funcs.ToType,
			"ToPlainTypeView": funcs.ToPlainTypeView,
			"ToTypeView":      funcs.ToTypeView,
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
