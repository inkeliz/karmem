package kmgen

import (
	"bytes"
	"io"
	"karmem.org/cmd/karmem/kmparser"
	"text/template"
)

type AssemblyScript struct {
	buf             *bytes.Buffer
	template        *template.Template
	translateTypes  ValueTypeTranslator
	translateViewer ValueTypeTranslator
	*kmparser.File
}

func AssemblyScriptGenerator() Generator {
	t, err := template.ParseFS(templateFiles, "assemblyscript_template.*")
	if err != nil {
		panic(err)
	}

	return &AssemblyScript{
		buf:      bytes.NewBuffer(nil),
		template: t,
		translateTypes: ValueTypeTranslator{
			Byte:    "u8",
			Bool:    "bool",
			Char:    "string",
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
			Array:   NewValueTypeFormatter(`StaticArray<{{.PlainType}}>`),
			Slice:   NewValueTypeFormatter(`Array<{{.PlainType}}>`),
		},
		translateViewer: ValueTypeTranslator{
			Byte:    "u8",
			Bool:    "bool",
			Char:    "string",
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
			Enum:    NewValueTypeFormatter(`{{.}}`),
			Struct:  NewValueTypeFormatter(`{{.PlainType}}`),
			Array:   NewValueTypeFormatter(`karmem.Slice<{{.PlainType}}>`),
			Slice:   NewValueTypeFormatter(`karmem.Slice<{{.PlainType}}>`),
		},
	}
}

func (gen *AssemblyScript) Start(file *kmparser.File) error {
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

func (gen *AssemblyScript) start() error {
	pkg := gen.Header.Name
	if p := gen.Header.GetTag("assemblyscript.package"); p != nil {
		pkg = p.Value
	}

	imports := "karmem/assemblyscript/karmem"
	if p := gen.Header.GetTag("assemblyscript.import"); p != nil {
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
		Package: pkg,
		Import:  imports,
		Largest: largest,
	})
}

func (gen *AssemblyScript) enums() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums`, gen.File.Enum)
}

func (gen *AssemblyScript) enumsBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `enums_builder`, gen.File.Enum)
}

func (gen *AssemblyScript) structs() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct`, gen.File.Struct)
}

func (gen *AssemblyScript) structBuilder() error {
	return gen.template.ExecuteTemplate(gen.buf, `struct_builder`, gen.File.Struct)
}

func (gen *AssemblyScript) Save(output io.Writer) error {
	_, err := io.Copy(output, gen.buf)
	return err
}
