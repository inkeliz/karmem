package kmgen

import (
	"embed"
	"fmt"
	"io"
	"karmem.org/cmd/karmem/kmparser"
	"strings"
	"text/template"
)

//go:embed *_template.*
var templateFiles embed.FS

type Generator interface {
	Start(file *kmparser.File) error
	Save(output io.Writer) error
}

type ValueTypeTranslator struct {
	Byte    string
	Bool    string
	Char    string
	Uint8   string
	Uint16  string
	Uint32  string
	Uint64  string
	Int8    string
	Int16   string
	Int32   string
	Int64   string
	Float32 string
	Float64 string

	Enum   ValueTypeFormatter
	Struct ValueTypeFormatter
	Array  ValueTypeFormatter
	Slice  ValueTypeFormatter
	Import string
}

type ValueTypeFormatter struct {
	*template.Template
}

func NewValueTypeFormatter(s string) ValueTypeFormatter {
	t, err := template.New("").Parse(s)
	if err != nil {
		panic(err)
	}
	return ValueTypeFormatter{Template: t}
}

func (x *ValueTypeTranslator) TranslateTypes(p *kmparser.File) {
	for pi, v := range p.Enum {
		p.Enum[pi].Type = x.translate(v.ValueType, false, "")
	}
	for pi, v := range p.Struct {
		for vi, v := range v.Fields {
			p.Struct[pi].Fields[vi].Type = x.translate(v.ValueType, v.IsEnum(), v.BackgroundType)
			p.Struct[pi].Fields[vi].BackgroundType = x.translate(kmparser.ValueType(v.BackgroundType), v.IsEnum(), v.BackgroundType)
			p.Struct[pi].Fields[vi].PlainType = x.translate(kmparser.ValueType(v.ValueType.PlainType()), v.IsEnum(), v.BackgroundType)
		}
	}
}

func (x *ValueTypeTranslator) TranslateViewerTypes(p *kmparser.File) {
	for pi, v := range p.Enum {
		p.Enum[pi].Type = x.translate(v.ValueType, false, "")
	}
	for pi, v := range p.Struct {
		for vi, v := range v.Fields {
			v.IsBool()
			viewer := v.ValueType
			if !v.IsNative() {
				viewer = v.ValueType + "Viewer"
			}
			p.Struct[pi].Fields[vi].ViewerType = x.translate(viewer, v.IsEnum(), v.BackgroundType)
			p.Struct[pi].Fields[vi].PlainViewerType = x.translate(kmparser.ValueType(viewer.PlainType()), v.IsEnum(), v.BackgroundType)
		}
	}
}

func (x *ValueTypeTranslator) translate(v kmparser.ValueType, isEnum bool, bg string) string {
	t := v.PlainType()
	switch t {
	case "byte":
		t = x.Byte
	case "bool":
		t = x.Bool
	case "char":
		t = x.Char
	case "uint8":
		t = x.Uint8
	case "uint16":
		t = x.Uint16
	case "uint32":
		t = x.Uint32
	case "uint64":
		t = x.Uint64
	case "int8":
		t = x.Int8
	case "int16":
		t = x.Int16
	case "int32":
		t = x.Int32
	case "int64":
		t = x.Int64
	case "float32":
		t = x.Float32
	case "float64":
		t = x.Float64
	}

	if v.IsBasic() && isEnum && x.Enum.Template != nil {
		var s strings.Builder
		if err := x.Enum.Execute(&s, kmparser.ValueType(fmt.Sprintf("%s", bg))); err != nil {
			panic("invalid enum format on translator")
		}
		return s.String()
	}
	if v.IsBasic() && !v.IsNative() && x.Struct.Template != nil {
		var s strings.Builder
		if err := x.Struct.Execute(&s, kmparser.ValueType(fmt.Sprintf("%s", t))); err != nil {
			panic("invalid struct format on translator")
		}
		return s.String()
	}
	if v.IsString() {
		return x.Char
	}
	if v.IsArray() {
		var s strings.Builder
		if err := x.Array.Execute(&s, kmparser.ValueType(fmt.Sprintf("[%d]%s", v.Length(), t))); err != nil {
			panic("invalid array format on translator")
		}
		return s.String()
	}
	if v.IsSlice() {
		var s strings.Builder
		if err := x.Slice.Execute(&s, kmparser.ValueType(fmt.Sprintf("[]%s", t))); err != nil {
			panic("invalid slice format on translator")
		}
		return s.String()
	}
	return t
}
