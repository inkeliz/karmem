package kmgen

import (
	"bytes"
	"io"
	"karmem.org/cmd/karmem/kmparser"
	"text/template"
)

type Zig struct {
	buf             *bytes.Buffer
	template        *template.Template
	translateTypes  ValueTypeTranslator
	translateViewer ValueTypeTranslator
	*kmparser.File
}

func ZigGenerator() Generator {
	t, err := template.ParseFS(templateFiles, "zig_template.*")
	if err != nil {
		panic(err)
	}

	return &Zig{
		buf:      bytes.NewBuffer(nil),
		template: t,
		translateTypes: ValueTypeTranslator{
			Byte:    "u8",
			Bool:    "bool",
			Char:    "[]u8",
			Uint8:   "u8",
			Uint16:  "u16",
			Uint32:  "u32",
			Uint64:  "u64",
			Int8:    "i8",
			Int16:   "i16",
			Int32:   "i32",
			Int64:   "i64",
			Float32: "f32",
			Float64: "f64",
			Enum:    NewValueTypeFormatter(`Enum{{.}}`),
			Array:   NewValueTypeFormatter(`[{{.Length}}]{{.PlainType}}`),
			Slice:   NewValueTypeFormatter(`[]{{.PlainType}}`),
		},
		translateViewer: ValueTypeTranslator{
			Byte:    "u8",
			Bool:    "bool",
			Char:    "[]u8",
			Uint8:   "u8",
			Uint16:  "u16",
			Uint32:  "u32",
			Uint64:  "u64",
			Int8:    "i8",
			Int16:   "i16",
			Int32:   "i32",
			Int64:   "i64",
			Float32: "f32",
			Float64: "f64",
			Enum:    NewValueTypeFormatter(`Enum{{.}}`),
			Struct:  NewValueTypeFormatter(`*const {{.PlainType}}`),
			Array:   NewValueTypeFormatter(`[]const {{.PlainType}}`),
			Slice:   NewValueTypeFormatter(`[]const {{.PlainType}}`),
		},
	}
}

func (gen *Zig) Start(file *kmparser.File) error {
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

func (gen *Zig) start() error {
	pkg := gen.Header.Name
	if p := gen.Header.GetTag("zig.package"); p != nil {
		pkg = p.Value
	}

	imports := "karmem"
	if p := gen.Header.GetTag("zig.import"); p != nil {
		imports = p.Value
	}

	return gen.template.ExecuteTemplate(gen.buf, `header`, struct {
		Package, Import string
	}{
		Package: pkg,
		Import:  imports,
	})
}

func (gen *Zig) enums() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums`, gen.File.Enum)
}

func (gen *Zig) enumsBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums_builder`, gen.File.Enum)
}

func (gen *Zig) structs() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct`, gen.File.Struct)
}

func (gen *Zig) structBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct_builder`, gen.File.Struct)
}

func (gen *Zig) Save(output io.Writer) error {
	_, err := io.Copy(output, gen.buf)
	return err
}
