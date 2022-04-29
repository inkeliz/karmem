package kmgen

import (
	"bytes"
	"go/format"
	"io"
	"karmem.org/cmd/karmem/kmparser"
	"text/template"
)

type Golang struct {
	buf             *bytes.Buffer
	template        *template.Template
	translateTypes  ValueTypeTranslator
	translateViewer ValueTypeTranslator
	*kmparser.File
}

func GolangGenerator() Generator {
	t, err := template.ParseFS(templateFiles, "golang_template.*")
	if err != nil {
		panic(err)
	}

	return &Golang{
		buf:      bytes.NewBuffer(nil),
		template: t,
		translateTypes: ValueTypeTranslator{
			Byte:    "byte",
			Bool:    "bool",
			Char:    "string",
			Uint8:   "uint8",
			Uint16:  "uint16",
			Uint32:  "uint32",
			Uint64:  "uint64",
			Int8:    "int8",
			Int16:   "int16",
			Int32:   "int32",
			Int64:   "int64",
			Float32: "float32",
			Float64: "float64",
			Array:   NewValueTypeFormatter(`[{{.Length}}]{{.PlainType}}`),
			Slice:   NewValueTypeFormatter(`[]{{.PlainType}}`),
		},
		translateViewer: ValueTypeTranslator{
			Byte:    "byte",
			Bool:    "bool",
			Char:    "[]byte",
			Uint8:   "uint8",
			Uint16:  "uint16",
			Uint32:  "uint32",
			Uint64:  "uint64",
			Int8:    "int8",
			Int16:   "int16",
			Int32:   "int32",
			Int64:   "int64",
			Float32: "float32",
			Float64: "float64",
			Enum:    NewValueTypeFormatter(`{{.}}`),
			Struct:  NewValueTypeFormatter(`*{{.PlainType}}`),
			Array:   NewValueTypeFormatter(`[]{{.PlainType}}`),
			Slice:   NewValueTypeFormatter(`[]{{.PlainType}}`),
		},
	}
}

func (gen *Golang) Start(file *kmparser.File) error {
	gen.translateTypes.TranslateTypes(file)
	gen.translateViewer.TranslateViewerTypes(file)
	gen.File = file

	for _, f := range []func() error{gen.start, gen.enums, gen.enumsBuilder, gen.structs, gen.structBuilder} {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func (gen *Golang) start() error {
	pkg := gen.Header.Name
	if p := gen.Header.GetTag("golang.package"); p != nil {
		pkg = p.Value
	}

	return gen.template.ExecuteTemplate(gen.buf, `header`, pkg)
}

func (gen *Golang) enums() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums`, gen.File.Enum)
}

func (gen *Golang) enumsBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums_builder`, gen.File.Enum)
}

func (gen *Golang) structs() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct`, gen.File.Struct)
}

func (gen *Golang) structBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct_builder`, gen.File.Struct)
}

func (gen *Golang) Save(output io.Writer) error {
	b, err := format.Source(gen.buf.Bytes())
	if err != nil {
		return err
	}
	_, err = output.Write(b)
	return err
}
