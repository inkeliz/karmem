package kmparser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// Reader reads and decodes Karmem files.
type Reader struct {
	Parsed File

	path  string
	buf   *bufio.Reader
	error error

	line   int
	column int

	lastTag    TagValue
	lastEnum   Enum
	lastStruct Struct
}

// NewReader accepts any karmem file as io.Reader.
// In order to give errors, the path is used.
func NewReader(path string, r io.Reader) *Reader {
	v := &Reader{path: path, buf: bufio.NewReader(r)}
	v.Parsed = File{Header: Header{Tags: Tags{}}}
	return v
}

type parserFunc func() parserFunc

// Parser will try to parse the given karmem file.
func (r *Reader) Parser() (d *File, err error) {
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
	if !r.peekEqual("karmem") {
		r.error = errors.New(`invalid header, expecting "karmem [package name] [@tag(value)];"`)
	}
	r.skip(len("karmem "))
	return r.headerPackage
}

func (r *Reader) headerPackage() parserFunc {
	b := r.nextRune()
	switch {
	case b == '@':
		r.error = fmt.Errorf(`invalid header, expecting package name got "%s"`, string(b))
		return nil

	case b == ';':
		return r.bodyInit

	case unicodeSpaceTab(b):
		return r.headerTagsOrEnd

	case unicode.IsControl(b):
		r.error = fmt.Errorf(`invalid header, expecting ";" got "%s"`, string(b))
		return nil

	case unicode.IsLetter(b):
		r.Parsed.Header.Name += string(b)

	default:
		r.error = fmt.Errorf(`invalid header, expecting ";" got "%s"`, string(b))
		return nil
	}

	return r.headerPackage
}

func (r *Reader) headerTagsOrEnd() parserFunc {
	b := r.nextRune()
	switch {
	case b == '@':
		return func() parserFunc { return r.tagsInit(r.Parsed.Header.Tags, r.headerTagsOrEnd) }

	case b == ';':
		return r.bodyInit
	}
	return r.headerTagsOrEnd
}

func (r *Reader) tagsInit(t Tags, f parserFunc) parserFunc {
	if r.lastTag.Column+r.lastTag.Line == 0 {
		r.lastTag.Column = r.column
		r.lastTag.Line = r.line
	}

	b := r.nextRune()
	switch {
	case r.error != nil:
		return nil

	case b == '(':
		return r.tagsValue(t, f, false)

	case unicode.IsLetter(b) || b == '.':
		r.lastTag.Name += string(b)

	default:
		r.error = fmt.Errorf(`invalid tag, expecting tag formatted as "@tag(value)" and receives "%s"`, string(b))
		return nil
	}

	return r.tagsInit(t, f)
}

func (r *Reader) tagsValue(t Tags, f parserFunc, ignore bool) parserFunc {
	b := r.nextRune()
	switch {
	case r.error != nil:
		return nil

	case b == ';':
		r.prevRune()
		fallthrough
	case b == ')':
		if err := t.AddTag(r.lastTag); err != nil {
			r.error = err
			return nil
		}
		r.lastTag = TagValue{}
		return f

	case b == '`':
		ignore = !ignore

	case ignore:
		if unicode.IsSpace(b) && !unicodeSpaceTab(b) {
			r.error = fmt.Errorf(`invalid tag, expecting tag formatted as "@tag(value)" and receives "%s"`, string(b))
			return nil
		}
		fallthrough
	case unicode.IsLetter(b) || unicode.IsNumber(b):
		r.lastTag.Value += string(b)

	default:
		r.error = fmt.Errorf(`invalid tag, expecting tag formatted as "@tag(value)" and receives "%s"`, string(b))
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
		r.error = fmt.Errorf(`invalid type, expecting "enum" or "struct" and got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) enumInit() parserFunc {
	if !r.peekEqual("enum ") {
		r.error = errors.New(`invalid enum, expecting "enum Name [byte|uint8|unit16|uint32|uint64|int8|int16|int32|int64] {}"`)
	}
	r.skip(len("enum "))
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
	switch {
	case unicodeSpaceTab(b) && len(r.lastEnum.Name) > 0:
		return r.skipSpace(r.enumType)

	case unicode.IsLetter(b):
		if len(r.lastEnum.Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got "%s"`, string(b))
			return nil
		}
		r.lastEnum.Name += string(b)
		return r.enumName

	default:
		r.error = fmt.Errorf(`invalid enum name, expecting "enum [name] [type] {}" got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) enumType() parserFunc {
	b := r.nextRune()
	switch {
	case unicodeSpaceTab(b):
		if len(r.lastEnum.ValueType) == 0 {
			return r.enumType
		}
		return r.skipSpace(r.enumType)

	case b == '{':
		r.lastEnum.Fields = append(r.lastEnum.Fields, EnumField{})
		return r.enumFieldName

	case unicode.IsLetter(b) || unicode.IsNumber(b):
		r.lastEnum.ValueType += ValueType(b)
		return r.enumType

	default:
		r.error = fmt.Errorf(`invalid enum, expecting "[type] {" got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) enumFieldName() parserFunc {
	b := r.nextRune()
	l := len(r.lastEnum.Fields) - 1
	switch {
	case b == ';':
		r.prevRune()
		return r.enumFieldValue

	case unicode.IsSpace(b) && len(r.lastEnum.Fields[l].Name) == 0:
		return r.enumFieldName

	case unicodeSpaceTab(b):
		if len(r.lastEnum.Fields[l].Name) > 0 {
			return r.skipSpace(r.enumFieldValue)
		}
		return r.enumFieldName

	case unicode.IsNumber(b):
		if len(r.lastEnum.Name) == 0 {
			r.error = fmt.Errorf(`invalid name, enum names can't use number as first char, got "%s"`, string(b))
			return nil
		}
		fallthrough
	case unicode.IsLetter(b):
		if len(r.lastEnum.Fields[l].Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got "%s"`, string(b))
			return nil
		}
		r.lastEnum.Fields[l].Name += string(b)
		return r.enumFieldName

	default:
		r.error = fmt.Errorf(`invalid enum name, got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) enumFieldValue() parserFunc {
	b := r.nextRune()
	l := len(r.lastEnum.Fields) - 1
	switch {
	case b == ';':
		if len(r.lastEnum.Fields[l].Type) == 0 {
			r.lastEnum.Fields[l].Type = r.lastEnum.ValueType
		}
		if len(r.lastEnum.Fields[l].Value) == 0 {
			r.lastEnum.Fields[l].Value = strconv.Itoa(len(r.lastEnum.Fields) - 1)
		}

		if !r.lastEnum.Fields[l].Type.IsValidNumericFor(r.lastEnum.Fields[l].Value) {
			r.error = fmt.Errorf(`invalid enum value ("%s") for "%s" type`, string(b), r.lastEnum.Fields[l].Type)
			return nil
		}

		return r.enumEnd
	case b == '=':
		if len(r.lastEnum.Fields[l].Type) > 0 {
			r.error = fmt.Errorf(`invalid enum value, expecting a number and got "%s"`, string(b))
			return nil
		}
		r.lastEnum.Fields[l].Type = r.lastEnum.ValueType
		return r.enumFieldValue

	case len(r.lastEnum.Fields[l].Type) > 0 && unicode.IsNumber(b):
		r.lastEnum.Fields[l].Value += string(b)
		return r.enumFieldValue

	case unicodeSpaceTab(b):
		return r.enumFieldValue

	default:
		r.error = fmt.Errorf(`invalid enum value, expecting "name [= value];" got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) enumEnd() parserFunc {
	b := r.nextRune()
	switch {
	case b == '}':

		if ok := r.lastEnum.Save(&r.Parsed); !ok {
			r.error = fmt.Errorf(`invalid enum`)
			return nil
		}

		r.Parsed.Enum = append(r.Parsed.Enum, r.lastEnum)
		r.lastEnum = Enum{}

		return r.bodyInit
	case unicode.IsSpace(b):
		return r.enumEnd
	case unicode.IsLetter(b):
		r.prevRune()
		r.lastEnum.Fields = append(r.lastEnum.Fields, EnumField{})
		return r.enumFieldName
	default:
		r.error = fmt.Errorf(`invalid enum value, expecting "name [= value];" got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) structInit() parserFunc {
	if !r.peekEqual("struct ") {
		r.error = errors.New(`invalid enum, expecting "struct Name [inline | table] {}"`)
	}
	r.skip(len("struct "))
	return r.skipSpace(r.structName)
}

func (r *Reader) structName() parserFunc {
	b := r.nextRune()
	switch {
	case unicodeSpaceTab(b) && len(r.lastStruct.Name) > 0:
		return r.skipSpace(r.structType)

	case unicode.IsNumber(b):
		if len(r.lastStruct.Name) == 0 {
			r.error = fmt.Errorf(`invalid struct name, can't start with number, got "%s"`, string(b))
			return nil
		}
		r.lastStruct.Name += string(b)
		return r.structName
	case unicode.IsLetter(b):
		if len(r.lastStruct.Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got "%s"`, string(b))
		}
		r.lastStruct.Name += string(b)
		return r.structName

	default:
		r.error = fmt.Errorf(`invalid struct name, expecting "struct Name [inline | table] {}" got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) structType() parserFunc {
	b := r.nextRune()
	switch {
	case b == '{' && len(r.lastStruct.Name) > 0:
		if !r.lastStruct.Class.IsValid() {
			r.error = fmt.Errorf(`invalid struct type, expecting "[inline | table]" got "%s"`, r.lastStruct.Class)
			return nil
		}

		r.lastStruct.Fields = append(r.lastStruct.Fields, StructField{Tags: Tags{}})
		return r.structFieldName

	case unicodeSpaceTab(b) && len(r.lastStruct.Name) > 0:
		return r.skipSpace(r.structType)

	case unicode.IsLower(b) && unicode.IsLetter(b):
		r.lastStruct.Class += StructClass(b)
		return r.structType

	default:
		r.error = fmt.Errorf(`invalid struct name, expecting "struct Name {}" got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) structFieldName() parserFunc {
	b := r.nextRune()
	l := len(r.lastStruct.Fields) - 1
	switch {
	case unicode.IsSpace(b) && len(r.lastStruct.Fields[l].Name) == 0:
		return r.structFieldName

	case unicodeSpaceTab(b):
		if len(r.lastStruct.Fields[l].Name) > 0 {
			return r.skipSpace(r.structFieldType)
		}
		return r.structFieldName

	case unicode.IsNumber(b):
		if len(r.lastStruct.Fields[l].Name) == 0 {
			r.error = fmt.Errorf(`invalid name, names can't start with numbers, got "%s"`, string(b))
			return nil
		}
		fallthrough
	case unicode.IsLetter(b):
		if len(r.lastStruct.Fields[l].Name) == 0 && unicode.IsLower(b) {
			r.error = fmt.Errorf(`invalid field name, names can't start with lowercase, got "%s"`, string(b))
		}
		r.lastStruct.Fields[l].Name += string(b)
		return r.structFieldName

	default:
		r.error = fmt.Errorf(`invalid struct field name got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) structFieldType() parserFunc {
	b := r.nextRune()
	l := len(r.lastStruct.Fields) - 1
	switch {
	case b == ';':
		if !r.lastStruct.Fields[l].ValueType.Save(&r.Parsed) {
			r.error = fmt.Errorf(`invalid type of "%s"`, string(r.lastStruct.Fields[l].ValueType))
			return nil
		}

		r.prevRune()
		return r.structEnd

	case b == '@':
		return func() parserFunc { return r.tagsInit(r.lastStruct.Fields[l].Tags, r.structFieldTagsOrEnd) }

	case unicodeSpaceTab(b) && len(r.lastStruct.Fields[l].ValueType) > 0:
		return r.skipSpace(r.structFieldType)

	case b == '[':
		if strings.Contains(string(r.lastStruct.Fields[l].ValueType), "[") || strings.Contains(string(r.lastStruct.Fields[l].ValueType), "]") {
			if strings.HasSuffix(string(r.lastStruct.Fields[l].ValueType), "]") {
				r.error = fmt.Errorf(`invalid array type name, types must be [n] or [] got "%s"`, string(b))
				return nil
			}
		}
		r.lastStruct.Fields[l].ValueType += ValueType(b)
		return r.structFieldType
	case b == '<':
		if !strings.HasPrefix(string(r.lastStruct.Fields[l].ValueType), "[") {
			r.error = fmt.Errorf(`invalid array type name, types must be [<n] or [n] or [] got "%s"`, string(b))
			return nil
		}
		r.lastStruct.Fields[l].ValueType += ValueType(b)
		return r.structFieldType
	case b == ']':
		if !strings.Contains(string(r.lastStruct.Fields[l].ValueType), "[") || strings.Contains(string(r.lastStruct.Fields[l].ValueType), "]") {
			if strings.HasSuffix(string(r.lastStruct.Fields[l].ValueType), "]") {
				r.error = fmt.Errorf(`invalid array type name, types must be [n] or [] got "%s"`, string(b))
				return nil
			}
		}
		r.lastStruct.Fields[l].ValueType += ValueType(b)
		return r.structFieldType

	case unicode.IsNumber(b):
		if len(r.lastStruct.Fields[l].ValueType) == 0 || strings.HasSuffix(string(r.lastStruct.Fields[l].ValueType), "]") {
			r.error = fmt.Errorf(`invalid type name, types must start with letter got "%s"`, string(b))
			return nil
		}
		r.lastStruct.Fields[l].ValueType += ValueType(b)
		return r.structFieldType

	case unicode.IsLetter(b):
		if strings.HasPrefix(string(r.lastStruct.Fields[l].ValueType), "[") && !strings.Contains(string(r.lastStruct.Fields[l].ValueType), "]") {
			r.error = fmt.Errorf(`invalid array lenght, got "%s"`, string(b))
			return nil
		}
		r.lastStruct.Fields[l].ValueType += ValueType(b)
		return r.structFieldType

	default:
		r.error = fmt.Errorf(`invalid struct field type, expecting "Name type [@tag(value)];" got "%s"`, string(b))
		return nil
	}
}

func (r *Reader) structFieldTagsOrEnd() parserFunc {
	b := r.nextRune()
	l := len(r.lastStruct.Fields) - 1
	switch {
	case b == '@':
		return func() parserFunc { return r.tagsInit(r.lastStruct.Fields[l].Tags, r.structFieldTagsOrEnd) }

	case b == ';':
		r.prevRune()
		return r.structEnd
	}
	return r.structFieldTagsOrEnd
}

func (r *Reader) structEnd() parserFunc {
	b := r.nextRune()
	switch {
	case b == '}':

		if err := r.lastStruct.Fields[len(r.lastStruct.Fields)-1].Save(&r.Parsed); err != nil {
			r.error = err
			return nil
		}

		if ok := r.lastStruct.Save(&r.Parsed); !ok {
			r.error = fmt.Errorf(`invalid enum`)
			return nil
		}

		r.Parsed.Struct = append(r.Parsed.Struct, r.lastStruct)
		r.lastStruct = Struct{}
		return r.bodyInit

	case b == ';':
		return r.structEnd

	case unicode.IsLetter(b):
		r.prevRune()

		if err := r.lastStruct.Fields[len(r.lastStruct.Fields)-1].Save(&r.Parsed); err != nil {
			r.error = err
			return nil
		}

		r.lastStruct.Fields = append(r.lastStruct.Fields, StructField{Tags: Tags{}})
		return r.structFieldName

	case unicode.IsSpace(b):
		return r.structEnd

	default:
		r.error = fmt.Errorf(`invalid struct, expecting "}" got "%s"`, string(b))
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
