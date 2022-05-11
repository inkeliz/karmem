package kmgen

import (
	"bytes"
	"io"
	"karmem.org/cmd/karmem/kmparser"
	"text/template"
)

type Swift struct {
	buf             *bytes.Buffer
	template        *template.Template
	translateTypes  ValueTypeTranslator
	translateViewer ValueTypeTranslator
	*kmparser.File
}

func SwiftGenerator() Generator {
	t, err := template.ParseFS(templateFiles, "swift_template.*")
	if err != nil {
		panic(err)
	}

	return &Swift{
		buf:      bytes.NewBuffer(nil),
		template: t,
		translateTypes: ValueTypeTranslator{
			Byte:    "UInt8",
			Bool:    "Bool",
			Char:    "[UInt8]",
			Uint8:   "UInt8",
			Uint16:  "UInt16",
			Uint32:  "UInt32",
			Uint64:  "UInt64",
			Int8:    "Int8",
			Int16:   "Int16",
			Int32:   "Int32",
			Int64:   "Int64",
			Float32: "Float",
			Float64: "Double",
			Enum:    NewValueTypeFormatter(`Enum{{.}}`),
			Array:   NewValueTypeFormatter(`[{{.PlainType}}]`),
			Slice:   NewValueTypeFormatter(`[{{.PlainType}}]`),
		},
		translateViewer: ValueTypeTranslator{
			Byte:    "UInt8",
			Bool:    "Bool",
			Char:    "karmem.Slice<UInt8>",
			Uint8:   "UInt8",
			Uint16:  "UInt16",
			Uint32:  "UInt32",
			Uint64:  "UInt64",
			Int8:    "Int8",
			Int16:   "Int16",
			Int32:   "Int32",
			Int64:   "Int64",
			Float32: "Float",
			Float64: "Double",
			Enum:    NewValueTypeFormatter(`Enum{{.}}`),
			Struct:  NewValueTypeFormatter(`{{.PlainType}}`),
			Array:   NewValueTypeFormatter(`karmem.Slice<{{.PlainType}}>`),
			Slice:   NewValueTypeFormatter(`karmem.Slice<{{.PlainType}}>`),
		},
	}
}

func (gen *Swift) Start(file *kmparser.File) error {
	gen.translateTypes.TranslateTypes(file)
	gen.translateViewer.TranslateViewerTypes(file)
	gen.File = file

	for _, f := range []func() error{gen.start, gen.enums, gen.structs, gen.enumsBuilder, gen.structBuilder} {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func (gen *Swift) start() error {
	pkg := gen.Header.Name
	if p := gen.Header.GetTag("swift.package"); p != nil {
		pkg = p.Value
	}

	imports := "karmem/swift/karmem"
	if p := gen.Header.GetTag("swift.import"); p != nil {
		imports = p.Value
	}

	return gen.template.ExecuteTemplate(gen.buf, `header`, struct {
		Package, Import string
	}{
		Package: pkg,
		Import:  imports,
	})
}

func (gen *Swift) enums() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums`, gen.File.Enum)
}

func (gen *Swift) enumsBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums_builder`, gen.File.Enum)
}

func (gen *Swift) structs() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct`, gen.File.Struct)
}

func (gen *Swift) structBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct_builder`, gen.File.Struct)
}

func (gen *Swift) Save(output io.Writer) error {
	_, err := io.Copy(output, gen.buf)
	return err
}
