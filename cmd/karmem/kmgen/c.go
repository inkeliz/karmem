package kmgen

import (
	"bytes"
	"io"
	"karmem.org/cmd/karmem/kmparser"
	"text/template"
)

type C struct {
	buf             *bytes.Buffer
	template        *template.Template
	translateTypes  ValueTypeTranslator
	translateViewer ValueTypeTranslator
	*kmparser.File
}

func CGenerator() Generator {
	t, err := template.ParseFS(templateFiles, "c_template.*")
	if err != nil {
		panic(err)
	}

	return &C{
		buf:      bytes.NewBuffer(nil),
		template: t,
		translateTypes: ValueTypeTranslator{
			Byte:    "uint8_t",
			Bool:    "bool",
			Char:    "uint8_t *",
			Uint8:   "uint8_t",
			Uint16:  "uint16_t",
			Uint32:  "uint32_t",
			Uint64:  "uint64_t",
			Int8:    "int8_t",
			Int16:   "int16_t",
			Int32:   "int32_t",
			Int64:   "int64_t",
			Float32: "float",
			Float64: "double",
			Enum:    NewValueTypeFormatter(`Enum{{.}}`),
			Array:   NewValueTypeFormatter(`{{.PlainType}}`),
			Slice:   NewValueTypeFormatter(`{{.PlainType}} *`),
		},
		translateViewer: ValueTypeTranslator{
			Byte:    "uint8_t",
			Bool:    "bool",
			Char:    "uint8_t *",
			Uint8:   "uint8_t",
			Uint16:  "uint16_t",
			Uint32:  "uint32_t",
			Uint64:  "uint64_t",
			Int8:    "int8_t",
			Int16:   "int16_t",
			Int32:   "int32_t",
			Int64:   "int64_t",
			Float32: "float",
			Float64: "double",
			Enum:    NewValueTypeFormatter(`Enum{{.}}`),
			Struct:  NewValueTypeFormatter(`{{.PlainType}} *`),
			Array:   NewValueTypeFormatter(`{{.PlainType}} *`),
			Slice:   NewValueTypeFormatter(`{{.PlainType}} *`),
		},
	}
}

func (gen *C) Start(file *kmparser.File) error {
	gen.translateTypes.TranslateTypes(file)
	gen.translateViewer.TranslateViewerTypes(file)
	gen.File = file

	for _, f := range []func() error{gen.start, gen.enums, gen.enumsBuilder, gen.structBuilder, gen.structs} {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func (gen *C) start() error {
	imports := "karmem.h"
	if p := gen.Header.GetTag("c.import"); p != nil {
		imports = p.Value
	}

	largest := uint32(0)
	for _, s := range gen.File.Struct {
		if s.Size > largest {
			largest = s.Size
		}
	}

	return gen.template.ExecuteTemplate(gen.buf, `header`, struct {
		Package, Import string
		Largest         uint32
	}{
		Import:  imports,
		Largest: largest,
	})
}

func (gen *C) enums() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums`, gen.File.Enum)
}

func (gen *C) enumsBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums_builder`, gen.File.Enum)
}

func (gen *C) structs() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct`, gen.File.Struct)
}

func (gen *C) structBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct_builder`, gen.File.Struct)
}

func (gen *C) Save(output io.Writer) error {
	_, err := io.Copy(output, gen.buf)
	return err
}
