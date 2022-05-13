package kmparser

import (
	"fmt"
	"strconv"
	"strings"
)

// File contains all contents from the karmem file.
type File struct {
	Header Header
	Enum   []Enum
	Struct []Struct
}

// Header contains the header of the file (package names).
type Header struct {
	Name string
	Tags
}

// Enum contains the decoded information of enum-type.
type Enum struct {
	Name string
	Type string
	Size uint32
	ValueType
	Fields []EnumField
}

func (s *Enum) Save(p *File) bool {
	for _, v := range s.Fields {
		size, ok := valueSize[ValueType(v.Type.PlainType())]
		if !ok {
			return false
		}
		s.Size = size
	}
	return true
}

// EnumField contains the decoded information of one enum-field.
type EnumField struct {
	Name  string
	Type  ValueType
	Value string
}

// Struct contains the decoded information of struct-type.
type Struct struct {
	Name      string
	Size      uint32
	SizeGroup []struct{} // Each item is one u64 that matches the Size.
	MinSize   uint32
	Tags
	Fields      []StructField
	Padding     uint32
	ContentSize uint32
	Class       StructClass
}

func (s *Struct) IsTable() bool {
	return s.Class.IsTable()
}

func (s *Struct) Save(p *File) bool {
	if s.IsTable() {
		s.Size = 4
	}
	s.MinSize = 8
	s.ContentSize = s.Size
	for i, v := range s.Fields {
		s.Fields[i].offset = s.Size
		s.Size += v.size
		s.ContentSize += v.size
	}

	s.Padding = 8 - (s.Size % 8)
	s.Size += s.Padding
	s.SizeGroup = make([]struct{}, s.Size/8)

	if s.Class.IsInline() {
		s.MinSize = s.Size
	}

	return true
}

// StructClass is the type of Struct (inline or table)
type StructClass string

func (x StructClass) IsValid() bool {
	return x == "inline" || x == "table"
}

func (x StructClass) IsInline() bool {
	return x == "inline"
}

func (x StructClass) IsTable() bool {
	return x == "table"
}

// StructField contains decoded information of one struct-field
type StructField struct {
	Name            string
	Type            string
	PlainViewerType string
	ViewerType      string
	PlainType       string
	BackgroundType  string
	ValueType

	isEnum    bool
	offset    uint32
	minSize   uint32
	allocSize uint32
	size      uint32
	inline    bool
	Tags
}

func (x *StructField) Offset() uint32 {
	return x.offset
}

func (x *StructField) Size() uint32 {
	return x.size
}

func (x *StructField) AllocSize() uint32 {
	return x.allocSize
}

func (x *StructField) MinSize() uint32 {
	return x.minSize
}

func (x *StructField) IsPadding() bool {
	return false
}

func (x *StructField) IsInline() bool {
	return x.inline
}

func (x *StructField) IsEnum() bool {
	return x.isEnum
}

func (x *StructField) Save(p *File) error {
	x.inline = true
	size, ok := valueSize[ValueType(x.ValueType.PlainType())]
	var minSize uint32
	if !ok {
		tn := x.ValueType.PlainType()
		for _, e := range p.Enum {
			if tn == e.Name {
				size, ok = e.Size, true
				x.isEnum = true
				x.BackgroundType = string(e.ValueType)
				break
			}
		}
		for _, e := range p.Struct {
			if tn == e.Name {
				size, ok = e.Size, true
				x.inline = e.Class.IsInline()
				minSize = x.minSize
				break
			}
		}
	}

	if minSize == 0 {
		minSize = size
	}

	if !ok {
		return fmt.Errorf("invalid type of %s. That type is unkown and must be declared before use", x.ValueType.PlainType())
	}

	if !x.ValueType.IsNative() && !x.ValueType.IsBasic() && !x.inline {
		return fmt.Errorf("invalid slice of %s. Table structs can't be used inside slices or arrays, warp it inside one inline struct", x.ValueType.PlainType())
	}

	x.allocSize = size
	if !x.inline {
		size = valueSize["*"]
	}

	if x.ValueType.IsArray() {
		x.inline = true
		size *= x.ValueType.Length()
		minSize *= x.ValueType.Length()
	}

	if x.ValueType.IsSlice() {
		x.inline = false
		size = valueSize["[]"]
	}

	x.minSize = minSize
	x.size = size

	return nil
}

// ValueType is the type of the field/enum.
type ValueType string

var valueSize = map[ValueType]uint32{
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

	"*":  4,     // Pointer
	"[]": 4 * 3, // Slice
}

// IsValidNumeric returns true if it's integers.
func (v ValueType) IsValidNumeric() bool {
	switch v {
	case "uint8", "uint16", "uint32", "uint64":
		return true
	case "int8", "int16", "int32", "int64":
		return true
	default:
		return false
	}
}

// IsValidNumericFor returns true if the given string is in range with the type.
func (v ValueType) IsValidNumericFor(s string) bool {
	if !v.IsValidNumeric() {
		return false
	}

	m := 8
	for _, x := range []string{"uint8", "uint16", "uint32", "uint64"} {
		if x != string(v) {
			continue
		}
		if _, err := strconv.ParseUint(s, 10, m); err == nil {
			return true
		}
		m *= 2
	}

	m = 8
	for _, x := range []string{"int8", "int16", "int32", "int64"} {
		if x != string(v) {
			continue
		}
		if _, err := strconv.ParseInt(s, 10, m); err == nil {
			return true
		}
		m *= 2
	}

	return false
}

// Save updates the ValueType and returns true if it's a valid type.
func (v *ValueType) Save(p *File) bool {
	if _, err := v.length(); err != nil {
		return false
	}

	tn := v.PlainType()
	if tn == "char" && !v.IsSlice() {
		return false
	}

	if _, ok := valueSize[ValueType(tn)]; ok {
		return true
	}

	for _, e := range p.Enum {
		if tn == e.Name {
			return true
		}
	}
	for _, e := range p.Struct {
		if tn == e.Name {
			return true
		}
	}

	return false
}

func (v ValueType) PlainType() string {
	_, t, found := strings.Cut(string(v), "]")
	if !found {
		t = string(v)
	}
	return t
}

func (v ValueType) IsNative() bool {
	_, ok := valueSize[ValueType(v.PlainType())]
	return ok
}

func (v ValueType) Default() string {
	switch {
	case v.IsString():
		return `""`
	case v.IsBool():
		return "false"
	}
	return "0"
}

func (v ValueType) IsBasic() bool {
	return !v.IsSlice() && !v.IsArray()
}

func (v ValueType) IsArray() bool {
	if strings.HasPrefix(string(v), "[]") {
		return false
	}
	if v.IsLimited() {
		return false
	}
	_, _, found := strings.Cut(string(v), "]")
	return found
}

func (v ValueType) IsSlice() bool {
	if v.IsArray() {
		return false
	}
	if v.IsLimited() {
		return true
	}
	_, _, found := strings.Cut(string(v), "]")
	return found
}

func (v ValueType) IsLimited() bool {
	return strings.Contains(string(v), "<")
}

func (v ValueType) Length() uint32 {
	n, _ := v.length()
	return uint32(n)
}

func (v ValueType) IsBool() bool {
	return v == "bool"
}

func (v ValueType) IsBytes() bool {
	return v.PlainType() == "byte"
}

func (v ValueType) IsInteger() bool {
	switch v {
	case "uint8", "uint16", "uint32", "uint64":
		return true
	case "int8", "int16", "int32", "int64":
		return true
	}
	return false
}

func (v ValueType) IsFloat() bool {
	return v.PlainType() == "float32" || v.PlainType() == "float64"
}

func (v ValueType) IsString() bool {
	return v.PlainType() == "char" && v.IsSlice()
}

func (v ValueType) length() (uint64, error) {
	if !v.IsSlice() && !v.IsArray() {
		return 1, nil
	}

	a, _, found := strings.Cut(string(v), "]")
	if !found {
		return 1, nil
	}
	a = strings.TrimPrefix(a, "[")
	a = strings.TrimPrefix(a, "<")

	if a == "" {
		return 1, nil
	}
	return strconv.ParseUint(a, 10, 32)
}

// Tags contains all custom defined tags.
type Tags map[string]TagValue

// TagValue contains all decoded information of one tag.
type TagValue struct {
	Name   string
	Value  string
	Line   int
	Column int
}

// AddTag includes the tag into the Tags, it returns error
// if the tag is duplicated.
func (t Tags) AddTag(v TagValue) error {
	if _, ok := t[v.Name]; ok {
		return fmt.Errorf(`duplicated "%s" tag`, v.Name)
	}
	t[v.Name] = v
	return nil
}

// GetTag gets the tag or nil if not exists
func (t Tags) GetTag(s string) *TagValue {
	x, ok := t[s]
	if !ok {
		return nil
	}
	return &x
}
