package kmparser

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"io"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/crypto/blake2b"
)

//go:generate go run ../main.go build -golang kmgen.km
//go:generate go run ../main.go fmt -w kmgen.km

// Reader reads and decodes Karmem files.
type Reader struct {
	Parsed           Content
	WarningsDisabled bool

	hasher     func(s string) uint64
	hasherHash hash.Hash
	hasherKey  []byte

	path  string
	buf   *bufio.Reader
	error error

	line   int
	column int
}

// NewReader accepts any karmem file as io.Reader.
// In order to give errors, the path is used.
func NewReader(path string, r io.Reader) *Reader {
	v := &Reader{path: path, buf: bufio.NewReader(r)}
	v.Parsed = Content{}
	return v
}

type parserFunc func() parserFunc

// Parser will try to parse the given karmem file.
func (r *Reader) Parser() (d *Content, err error) {
	var x parserFunc
	x = r.headerInit()
	for r.error == nil {
		if x = x(); x == nil {
			break
		}
	}
	if r.error == io.EOF {
		return &r.Parsed, nil
	}
	return nil, fmt.Errorf("error on %s:%d:%d: %s", r.path, r.line, r.column, r.error)
}

func (r *Reader) headerInit() parserFunc {
	if r.Parsed.Name == "" && r.peekEqual(string([]byte{0xEF, 0xBB, 0xBF})) {
		r.skip(3)
	}
	if !r.peekEqual("karmem") {
		r.error = errors.New(`invalid header, expecting "karmem [package name] [@tag(value)];"`)
		return nil
	}
	r.skip(len("karmem "))
	return r.headerPackage
}

func (r *Reader) headerPackage() parserFunc {
	b := r.nextRune()
	switch {
	case b == '@':
		r.error = fmt.Errorf(`invalid header, expecting package name got %q`, string(b))
		return nil

	case b == ';':
		r.prevRune()
		return r.headerTagsOrEnd

	case unicodeSpaceTab(b):
		return r.headerTagsOrEnd

	case unicode.IsControl(b):
		r.error = fmt.Errorf(`invalid header, expecting ";" got %q`, string(b))
		return nil

	case unicode.IsLetter(b):
		r.Parsed.Name += string(b)

	default:
		r.error = fmt.Errorf(`invalid header, expecting ";" got %q`, string(b))
		return nil
	}

	return r.headerPackage
}

func (r *Reader) headerTagsOrEnd() parserFunc {
	b := r.nextRune()
	switch {
	case b == '@':
		r.Parsed.Tags = append(r.Parsed.Tags, Tag{})
		return func() parserFunc { return r.tagsInit(&r.Parsed.Tags[len(r.Parsed.Tags)-1], r.headerTagsOrEnd) }

	case b == ';':
		v, err := Tags(r.Parsed.Tags).GetBoolean("packed")
		if err != nil {
			r.error = err
			return nil
		}
		r.Parsed.Packed = v == 1

		key, ok := Tags(r.Parsed.Tags).Get("key")
		if ok {
			k := blake2b.Sum256([]byte(key))
			r.hasherKey = k[:]
		}
		r.hasherHash, _ = blake2b.New(8, r.hasherKey)
		r.hasher = func(s string) uint64 {
			r.hasherHash.Reset()
			r.hasherHash.Write([]byte(s))
			return binary.LittleEndian.Uint64(r.hasherHash.Sum(nil))
		}

		return r.bodyInit
	}
	return r.headerTagsOrEnd
}

func (r *Reader) tagsInit(t *Tag, f parserFunc) parserFunc {
	b := r.nextRune()
	switch {
	case r.error != nil:
		return nil

	case b == '(':
		return r.tagsValue(t, f, false)

	case unicode.IsLetter(b) || b == '.':
		t.Name += string(b)

	default:
		r.error = fmt.Errorf(`invalid tag, expecting tag formatted as "@tag(value)" and receives %q`, string(b))
		return nil
	}

	return r.tagsInit(t, f)
}

func (r *Reader) tagsValue(t *Tag, f parserFunc, ignore bool) parserFunc {
	b := r.nextRune()
	switch {
	case r.error != nil:
		return nil

	case b == ';':
		r.prevRune()
		fallthrough
	case b == ')':
		return f

	case b == '`':
		ignore = !ignore

	case ignore:
		if unicode.IsSpace(b) && !unicodeSpaceTab(b) {
			r.error = fmt.Errorf(`invalid tag, expecting tag formatted as "@tag(value)" and receives %q`, string(b))
			return nil
		}
		fallthrough
	case unicode.IsLetter(b) || unicode.IsNumber(b):
		t.Value += string(b)

	default:
		r.error = fmt.Errorf(`invalid tag, expecting tag formatted as "@tag(value)" and receives %q`, string(b))
		return nil
	}

	return r.tagsValue(t, f, ignore)
}

func (r *Reader) bodyInit() parserFunc {
	b := r.nextRune()
	if unicode.IsSpace(b) {
		return r.bodyInit
	}
	if r.error != nil {
		return nil
	}
	r.prevRune()

	switch {
	case r.peekEqual("enum"):
		return r.enumInit

	case r.peekEqual("struct"):
		return r.structInit

	default:
		r.error = fmt.Errorf(`invalid type, expecting "enum" or "struct" and got %q`, string(b))
		return nil
	}
}

func (r *Reader) enumInit() parserFunc {
	if !r.peekEqual("enum ") {
		r.error = errors.New(`invalid enum, expecting "enum Name [byte|uint8|unit16|uint32|uint64|int8|int16|int32|int64] {}"`)
	}
	r.skip(len("enum "))
	r.Parsed.Enums = append(r.Parsed.Enums, Enumeration{Data: EnumData{IsSequential: true}})
	return r.skipSpace(r.enumName)
}

func (r *Reader) skipSpace(f parserFunc) parserFunc {
	if r.error == nil && unicodeSpaceTab(r.nextRune()) {
		return r.skipSpace(f)
	}
	r.prevRune()
	return f
}

func unicodeSpaceTab(r rune) bool {
	return r == ' ' || r == '\t'
}

func (r *Reader) enumName() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Enums[len(r.Parsed.Enums)-1].Data
	t.Position.File = r.path
	t.Position.Line = uint32(r.line)
	t.Position.Column = uint32(r.column)

	switch {
	case unicodeSpaceTab(b) && len(t.Name) > 0:
		return r.skipSpace(r.enumType)

	case unicode.IsLetter(b):
		if len(t.Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got %q`, string(b))
			return nil
		}
		t.Name += string(b)
		return r.enumName

	default:
		r.error = fmt.Errorf(`invalid enum name, expecting "enum [name] [type] {}" got %q`, string(b))
		return nil
	}
}

func (r *Reader) enumType() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Enums[len(r.Parsed.Enums)-1].Data
	switch {
	case unicodeSpaceTab(b):
		return r.enumType

	case b == '@':
		if t.Type.Schema == "" {
			r.error = fmt.Errorf(`invalid enum, expecting "[type] [@tag()] {" got tag before the type: %q`, string(b))
			return nil
		}

		t.Tags = append(t.Tags, Tag{})
		return func() parserFunc { return r.tagsInit(&t.Tags[len(t.Tags)-1], r.enumType) }

	case b == '{':
		t.Fields = append(t.Fields, EnumField{})
		return r.enumFieldName

	case unicode.IsLetter(b) || unicode.IsNumber(b):
		t.Type.Schema += string(b)
		return r.enumType

	default:
		r.error = fmt.Errorf(`invalid enum, expecting "[type] {" got %q`, string(b))
		return nil
	}
}

func (r *Reader) enumFieldName() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Enums[len(r.Parsed.Enums)-1].Data
	f := &t.Fields[len(t.Fields)-1].Data
	f.Position.File = r.path
	f.Position.Line = uint32(r.line)
	f.Position.Column = uint32(r.column)

	switch {
	case b == ';':
		r.prevRune()
		return r.enumFieldValue

	case unicode.IsSpace(b) && len(f.Name) == 0:
		return r.enumFieldName

	case unicodeSpaceTab(b):
		if len(f.Name) > 0 {
			return r.skipSpace(r.enumFieldValue)
		}
		return r.enumFieldName

	case unicode.IsNumber(b):
		if len(f.Name) == 0 {
			r.error = fmt.Errorf(`invalid name, enum names can't use number as first char, got %q`, string(b))
			return nil
		}
		fallthrough
	case unicode.IsLetter(b):
		if len(f.Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got %q`, string(b))
			return nil
		}
		f.Name += string(b)
		return r.enumFieldName

	default:
		r.error = fmt.Errorf(`invalid enum name, got %q`, string(b))
		return nil
	}
}

func _valueInRangeOfType(v string, t string) (err error) {
	bits := 0
	switch t[len(t)-1] {
	case '8':
		bits = 8
	case '6':
		bits = 16
	case '2':
		bits = 32
	case '4':
		bits = 64
	}

	switch t[0] {
	case 'u':
		_, err = strconv.ParseUint(v, 10, bits)
	case 'i':
		_, err = strconv.ParseInt(v, 10, bits)
	}
	return err
}

func (r *Reader) enumFieldValue() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Enums[len(r.Parsed.Enums)-1].Data
	f := &t.Fields[len(t.Fields)-1].Data
	switch {
	case b == ';':
		if len(f.Value) == 0 {
			f.Value = strconv.Itoa(len(t.Fields) - 1)
		} else {
			if f.Value != strconv.Itoa(len(t.Fields)-1) {
				t.IsSequential = false
			}
		}

		switch t.Type.Schema {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
			if err := _valueInRangeOfType(f.Value, t.Type.Schema); err != nil {
				r.error = err
				return nil
			}
		default:
			r.error = fmt.Errorf(`invalid enum type`)
		}

		for i := range t.Fields {
			if t.Fields[i].Data.Name == f.Name && &t.Fields[i].Data != f {
				r.error = fmt.Errorf(`duplicated enum field name of %q`, f.Name)
				return nil
			}
		}

		if t.Type.PlainSchema == "" {
			t.Type.PlainSchema = t.Type.Schema
		}

		t.Type.Model = TypeModelSingle
		t.Type.Format = TypeFormatPrimitive

		return r.enumEnd
	case b == '=':
		return r.enumFieldValue

	case len(t.Type.Schema) > 0 && unicode.IsNumber(b):
		f.Value += string(b)
		return r.enumFieldValue

	case unicodeSpaceTab(b):
		return r.enumFieldValue

	default:
		r.error = fmt.Errorf(`invalid enum value, expecting "name [= value];" got %q`, string(b))
		return nil
	}
}

func (r *Reader) _uniqueName(s string) bool {
	if len(r.Parsed.Enums) > 0 {
		for i := range r.Parsed.Enums[:len(r.Parsed.Enums)-1] {
			if r.Parsed.Enums[i].Data.Name == s {
				return false
			}
		}
	}
	if len(r.Parsed.Structs) > 0 {
		for i := range r.Parsed.Structs[:len(r.Parsed.Structs)-1] {
			if r.Parsed.Structs[i].Data.Name == s {
				return false
			}
		}
	}
	return true
}

func (r *Reader) enumEnd() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Enums[len(r.Parsed.Enums)-1].Data
	f := &t.Fields[len(t.Fields)-1].Data
	switch {
	case b == '}':
		if !r._uniqueName(f.Name) {
			r.error = fmt.Errorf(`duplicated enum name of %q`, f.Name)
			return nil
		}
		var hasDefault bool
		for _, v := range t.Fields {
			if v.Data.Value == "0" {
				hasDefault = true
				break
			}
		}
		if !hasDefault {
			r.error = fmt.Errorf(`enum %q doesn't have default value. It must have one enum with value 0`, t.Name)
			return nil
		}
		return r.bodyInit
	case unicode.IsSpace(b):
		return r.enumEnd
	case unicode.IsLetter(b):
		r.prevRune()
		t.Fields = append(t.Fields, EnumField{})
		return r.enumFieldName
	default:
		r.error = fmt.Errorf(`invalid enum value, expecting "name [= value];" got %q`, string(b))
		return nil
	}
}

func (r *Reader) structInit() parserFunc {
	if !r.peekEqual("struct ") {
		r.error = errors.New(`invalid struct, expecting "struct Name [inline | table] {}"`)
	}
	r.skip(len("struct "))
	r.Parsed.Structs = append(r.Parsed.Structs, Structure{})
	return r.skipSpace(r.structName)
}

func (r *Reader) structName() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Structs[len(r.Parsed.Structs)-1].Data
	t.Position.File = r.path
	t.Position.Line = uint32(r.line)
	t.Position.Column = uint32(r.column)

	switch {
	case unicodeSpaceTab(b) && len(t.Name) > 0:
		return r.skipSpace(r.structType)

	case unicode.IsNumber(b):
		if len(t.Name) == 0 {
			r.error = fmt.Errorf(`invalid struct name, can't start with number, got %q`, string(b))
			return nil
		}
		t.Name += string(b)
		return r.structName
	case unicode.IsLetter(b):
		if len(t.Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got %q`, string(b))
		}
		t.Name += string(b)
		return r.structName

	default:
		r.error = fmt.Errorf(`invalid struct name, expecting "struct [name] [inline | table] [@tag()] {}" got %q`, string(b))
		return nil
	}
}

func (r *Reader) structTypeEnd() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Structs[len(r.Parsed.Structs)-1].Data

	switch {
	case b == '@':
		t.Tags = append(t.Tags, Tag{})
		return func() parserFunc { return r.tagsInit(&t.Tags[len(t.Tags)-1], r.structTypeEnd) }
	case b == '{':
		if t.Class == StructClassNone {
			r.error = errors.New(`invalid struct, missing type (inline/table)"`)
			return nil
		}

		v, err := Tags(t.Tags).GetBoolean("packed")
		if err != nil {
			r.error = err
			return nil
		}
		if v > -1 {
			t.Packed = v == 1
		} else {
			t.Packed = r.Parsed.Packed
		}

		t.Fields = append(t.Fields, StructField{})
		return r.structFieldName

	case unicodeSpaceTab(b):
		return r.skipSpace(r.structTypeEnd)

	default:
		r.error = errors.New(`invalid struct, missing open "{" after type`)
		return nil
	}
}

func (r *Reader) structType() parserFunc {
	t := &r.Parsed.Structs[len(r.Parsed.Structs)-1].Data

	switch {
	case r.peekEqual("inline"):
		t.Class = StructClassInline
		r.skip(len("inline"))
		return r.structTypeEnd

	case r.peekEqual("table"):
		t.Class = StructClassTable
		t.Size.Content += 4
		r.skip(len("table"))
		return r.structTypeEnd
	}

	b := r.nextRune()
	switch {
	case unicodeSpaceTab(b):
		return r.skipSpace(r.structType)
	case b == '{':
		r.prevRune()
		return r.structTypeEnd
	default:
		r.error = fmt.Errorf(`invalid struct format got %q`, string(b))
		return nil
	}
}

func (r *Reader) structFieldName() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Structs[len(r.Parsed.Structs)-1].Data
	f := &t.Fields[len(t.Fields)-1].Data
	f.Position.File = r.path
	f.Position.Line = uint32(r.line)
	f.Position.Column = uint32(r.column)

	switch {
	case unicode.IsSpace(b) && len(f.Name) == 0:
		return r.structFieldName

	case unicodeSpaceTab(b):
		if len(f.Name) > 0 {
			return r.skipSpace(r.structFieldType)
		}
		return r.structFieldName

	case unicode.IsNumber(b):
		if len(f.Name) == 0 {
			r.error = fmt.Errorf(`invalid name, names can't start with numbers, got %q`, string(b))
			return nil
		}
		fallthrough
	case unicode.IsLetter(b):
		if len(f.Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got %q`, string(b))
		}
		f.Name += string(b)
		return r.structFieldName

	default:
		r.error = fmt.Errorf(`invalid struct field name got %q`, string(b))
		return nil
	}
}

var _paddings = map[uint32]PaddingType{
	1: {Data: Type{Schema: "uint8", PlainSchema: "uint8", Length: 1, Format: TypeFormatPrimitive, Model: TypeModelSingle}},
	2: {Data: Type{Schema: "uint16", PlainSchema: "uint16", Length: 2, Format: TypeFormatPrimitive, Model: TypeModelSingle}},
	4: {Data: Type{Schema: "uint32", PlainSchema: "uint32", Length: 4, Format: TypeFormatPrimitive, Model: TypeModelSingle}},
	8: {Data: Type{Schema: "uint64", PlainSchema: "uint64", Length: 8, Format: TypeFormatPrimitive, Model: TypeModelSingle}},
}

var _primitives = map[string]uint32{
	"byte": 1,
	"bool": 1,

	"char": 1,

	"uint8":  1,
	"uint16": 2,
	"uint32": 4,
	"uint64": 8,

	"int8":  1,
	"int16": 2,
	"int32": 4,
	"int64": 8,

	"float32": 4,
	"float64": 8,

	"*":         4,     // Pointer
	"[PACKED]":  4 * 2, // Dynamic Array
	"[PADDING]": 4 * 3, // Padding Dynamic Array
}

func (r *Reader) _typeFormatOf(t *StructData, field *StructFieldData) (StructFieldSize, TypeFormat, error) {
	size := StructFieldSize{}
	if s, ok := _primitives[field.Type.PlainSchema]; ok {
		switch field.Type.Model {
		case TypeModelSlice, TypeModelSliceLimited:
			size.Minimum = s
			size.Field = _primitives["[PACKED]"]
			if !t.Packed {
				size.Field = _primitives["[PADDING]"]
			}
			size.Allocation = s
		case TypeModelArray:
			size.Minimum = s * field.Type.Length
			size.Field = size.Minimum
			size.Allocation = s
		default:
			size.Minimum = s
			size.Field = s
			size.Allocation = s
		}

		return size, TypeFormatPrimitive, nil
	}

	if len(r.Parsed.Structs) > 0 {
		for i := range r.Parsed.Structs[:len(r.Parsed.Structs)-1] {
			if r.Parsed.Structs[i].Data.Name != field.Type.PlainSchema {
				continue
			}
			if r.Parsed.Structs[i].Data.Class == StructClassInline {
				size.Minimum = r.Parsed.Structs[i].Data.Size.Total * field.Type.Length
				if field.Type.Model == TypeModelSlice || field.Type.Model == TypeModelSliceLimited {
					size.Field = _primitives["[PACKED]"]
					if !t.Packed {
						size.Field = _primitives["[PADDING]"]
					}
					size.Allocation = r.Parsed.Structs[i].Data.Size.Total
				} else {
					size.Field = size.Minimum
					size.Allocation = r.Parsed.Structs[i].Data.Size.Total
				}
				return size, TypeFormatStruct, nil
			} else {
				size.Minimum = r.Parsed.Structs[i].Data.Size.Minimum
				size.Field = _primitives["*"]
				size.Allocation = r.Parsed.Structs[i].Data.Size.Total
				return size, TypeFormatTable, nil
			}
		}
	}

	if len(r.Parsed.Enums) > 0 {
		for i := range r.Parsed.Enums {
			if r.Parsed.Enums[i].Data.Name != field.Type.PlainSchema {
				continue
			}
			size.Minimum = _primitives[r.Parsed.Enums[i].Data.Type.Schema]
			size.Field = size.Minimum
			size.Allocation = size.Minimum
			return size, TypeFormatEnum, nil
		}
	}

	return size, TypeFormatNone, fmt.Errorf("invalid type of %q", field.Type.PlainSchema)
}

func (r *Reader) structFieldType() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Structs[len(r.Parsed.Structs)-1].Data
	f := &t.Fields[len(t.Fields)-1].Data

	switch {
	case b == ';':
		for i := range t.Fields[:len(t.Fields)-1] {
			if t.Fields[i].Data.Name == f.Name {
				r.error = fmt.Errorf(`duplicated struct field name of %s`, f.Name)
				return nil
			}
		}

		if f.Type.Model == TypeModelNone {
			f.Type.Model = TypeModelSingle
			f.Type.Length = 1
		}

		size, format, err := r._typeFormatOf(t, f)
		if err != nil {
			r.error = err
			return nil
		}
		f.Type.Format = format
		f.Size = size
		f.Offset = t.Size.Content
		t.Size.Content += f.Size.Field

		if f.Type.Format == TypeFormatEnum || f.Type.Format == TypeFormatTable {
			if f.Type.Model != TypeModelSingle {
				r.error = fmt.Errorf("invalid usafe of array/slice. enum and tables can't be used in array/slices, wrap it first")
				return nil
			}
		}

		return r.structEnd

	case b == '@':
		f.Tags = append(f.Tags, Tag{})
		return func() parserFunc { return r.tagsInit(&f.Tags[len(f.Tags)-1], r.structFieldType) }

	case unicodeSpaceTab(b) && len(f.Type.Schema) > 0:
		return r.skipSpace(r.structFieldType)

	case b == '[':
		if f.Type.Schema != "" {
			r.error = fmt.Errorf(`invalid array type name, types must be [<n] or [n] or [] got %q`, string(b))
			return nil
		}
		f.Type.Schema += string(b)
		return r.structFieldType

	case b == '<':
		if !strings.HasPrefix(f.Type.Schema, "[") {
			r.error = fmt.Errorf(`invalid array type name, types must be [<n] or [n] or [] got %q`, string(b))
			return nil
		}
		f.Type.Schema += string(b)
		return r.structFieldType

	case b == ']':
		if !strings.Contains(f.Type.Schema, "[") || strings.Contains(f.Type.Schema, "]") {
			r.error = fmt.Errorf(`invalid array type name, types must be [n] or [] got %q`, string(b))
			return nil
		}

		if f.Type.Schema[len(f.Type.Schema)-1] == '[' {
			f.Type.Model = TypeModelSlice
		}
		if i := strings.Index(f.Type.Schema, "<"); i >= 1 {
			f.Type.Model = TypeModelSliceLimited
			length, err := strconv.ParseUint(f.Type.Schema[i+1:], 10, 32)
			if err != nil {
				r.error = fmt.Errorf(`invalid length of %q`, f.Type.Schema)
				return nil
			}
			f.Type.Length = uint32(length)
		}
		if f.Type.Model == TypeModelNone {
			f.Type.Model = TypeModelArray
			length, err := strconv.ParseUint(f.Type.Schema[1:], 10, 32)
			if err != nil {
				r.error = fmt.Errorf(`invalid length of %q`, f.Type.Schema[1:])
				return nil
			}
			f.Type.Length = uint32(length)
		}

		f.Type.Schema += string(b)

		return r.structFieldType

	case unicode.IsNumber(b):
		if len(f.Type.Schema) == 0 || strings.HasSuffix(f.Type.Schema, "]") {
			r.error = fmt.Errorf(`invalid type name, types must start with letter got %q`, string(b))
			return nil
		}
		f.Type.Schema += string(b)
		if strings.Contains(f.Type.Schema, "]") || !strings.Contains(f.Type.Schema, "[") {
			f.Type.PlainSchema += string(b)
		}
		return r.structFieldType

	case unicode.IsLetter(b):
		if strings.HasPrefix(f.Type.Schema, "[") && !strings.Contains(f.Type.Schema, "]") {
			r.error = fmt.Errorf(`invalid array length, got %q`, string(b))
			return nil
		}
		f.Type.Schema += string(b)
		f.Type.PlainSchema += string(b)
		return r.structFieldType

	default:
		r.error = fmt.Errorf(`invalid struct field type, expecting "Name type [@tag(value)];" got %q`, string(b))
		return nil
	}
}

func (r *Reader) structEnd() parserFunc {
	b := r.nextRune()
	t := &r.Parsed.Structs[len(r.Parsed.Structs)-1].Data

	switch {
	case b == '}':
		if !r._uniqueName(t.Name) {
			r.error = fmt.Errorf(`duplicated struct name of %q`, t.Name)
			return nil
		}

		t.Size.Padding = 0
		if !t.Packed {
			switch t.Class {
			case StructClassTable:
				// StructClassTable is always packed.
				// That is backward compatible with old versions of Karmem.
			case StructClassInline:
				// StructClassInline uses padding if Packed isn't enabled.
				// That is a bug in old versions of Karmem, but we have to support it.
				t.Size.Padding = 8 - (t.Size.Content % 8)
			}
		}

		t.Size.Total = t.Size.Padding + t.Size.Content
		t.Size.TotalGroup = make([]PaddingType, 0, (t.Size.Total/8)+1)
		for total := t.Size.Total; total != 0; {
			for _, pt := range []uint32{8, 4, 2, 1} {
				if total >= pt {
					t.Size.TotalGroup = append(t.Size.TotalGroup, _paddings[pt])
					total -= pt
					break
				}
			}
		}

		if t.Class == StructClassTable {
			if !t.Packed {
				t.Size.Minimum = 8
			} else {
				t.Size.Minimum = 4
			}
		} else {
			t.Size.Minimum = t.Size.Content
		}

		if t.Size.Total > r.Parsed.Size.Largest {
			r.Parsed.Size.Largest = t.Size.Total
		}

		if v, ok := Tags(t.Tags).Get("id"); ok {
			id, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				r.error = fmt.Errorf(`invalid id of "%s"`, v)
				return nil
			}
			t.ID = id
		} else {
			if r.hasher != nil {
				t.ID = r.hasher(t.Name)
				t.Tags = append(t.Tags, Tag{"id", fmt.Sprintf("%d", t.ID)})
			}
		}

		return r.bodyInit

	case unicode.IsLetter(b):
		r.prevRune()

		t.Fields = append(t.Fields, StructField{})
		return r.structFieldName

	case unicode.IsSpace(b):
		return r.structEnd

	default:
		r.error = fmt.Errorf(`invalid struct, expecting "}" got %q`, string(b))
		return nil
	}
}

func (r *Reader) nextRune() rune {
	b, _, err := r.buf.ReadRune()
	if err != nil {
		r.error = err
		return 0
	}
	if b == 0 {
		r.error = io.EOF
		return 0
	}
	r.column++
	if b == '\n' {
		r.line++
		r.column = 0
	}
	return b
}

func (r *Reader) prevRune() rune {
	err := r.buf.UnreadRune()
	if err != nil {
		r.error = err
	}
	r.column--
	return 0
}

func (r *Reader) skip(l int) {
	l, err := r.buf.Discard(l)
	if err != nil {
		r.error = err
	}
	r.column += l
}

func (r *Reader) peekEqual(expected string) bool {
	v, _ := r.buf.Peek(len(expected))
	return string(v) == expected
}
