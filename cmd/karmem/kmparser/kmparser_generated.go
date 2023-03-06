package kmparser

import (
	"unsafe"

	karmem "karmem.org/golang"
)

var _ unsafe.Pointer

var _Null = [54]byte{}
var _NullReader = karmem.NewReader(_Null[:])

type (
	StructClass uint8
)

const (
	StructClassNone   StructClass = 0
	StructClassTable  StructClass = 1
	StructClassInline StructClass = 2
)

type (
	TypeModel uint8
)

const (
	TypeModelNone         TypeModel = 0
	TypeModelSingle       TypeModel = 1
	TypeModelArray        TypeModel = 2
	TypeModelSlice        TypeModel = 3
	TypeModelSliceLimited TypeModel = 4
)

type (
	TypeFormat uint8
)

const (
	TypeFormatNone      TypeFormat = 0
	TypeFormatPrimitive TypeFormat = 1
	TypeFormatEnum      TypeFormat = 2
	TypeFormatStruct    TypeFormat = 3
	TypeFormatTable     TypeFormat = 4
)

type (
	RuleType uint16
)

const (
	RuleTypeNone          RuleType = 0
	RuleTypeReservedName  RuleType = 1
	RuleTypeReservedType  RuleType = 2
	RuleTypeDuplicateName RuleType = 3
	RuleTypeDuplicateType RuleType = 4
	RuleTypeCollisionName RuleType = 5
	RuleTypeCollisionType RuleType = 6
)

type (
	Language uint64
)

const (
	LanguageNone           Language = 0
	LanguageAssemblyScript Language = 1
	LanguageZig            Language = 2
	LanguageC              Language = 4
	LanguageSwift          Language = 8
	LanguageDotnet         Language = 16
	LanguageOdin           Language = 32
	LanguageGolang         Language = 64
	LanguageRust           Language = 128
	LanguageNim            Language = 256
	LanguageKotlin         Language = 512
	LanguageD              Language = 1024
	LanguageHaskell        Language = 2048
	LanguageCrystal        Language = 4096
	LanguageTypescript     Language = 8192
	LanguagePHP            Language = 16384
	LanguagePython         Language = 32768
)

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierSchemaPosition  = 8944966152016866568
	PacketIdentifierType            = 2206764383142231373
	PacketIdentifierWarningData     = 9648278712160971894
	PacketIdentifierWarning         = 5319256325625522744
	PacketIdentifierPaddingType     = 6449815373135188035
	PacketIdentifierTag             = 9280816983786621498
	PacketIdentifierStructSize      = 2296279785726396957
	PacketIdentifierStructFieldSize = 3117293985139574571
	PacketIdentifierEnumFieldData   = 6917629752752470509
	PacketIdentifierEnumField       = 18350873289003309128
	PacketIdentifierEnumData        = 18057555498029063613
	PacketIdentifierEnumeration     = 1253319329451847685
	PacketIdentifierStructFieldData = 17962757807284521522
	PacketIdentifierStructField     = 12155838558451759529
	PacketIdentifierStructData      = 8290009745541165076
	PacketIdentifierStructure       = 18088017590773436939
	PacketIdentifierContentSize     = 8764462619562198222
	PacketIdentifierContent         = 6792576797909524956
)

type SchemaPosition struct {
	File   string
	Line   uint32
	Column uint32
}

func NewSchemaPosition() SchemaPosition {
	return SchemaPosition{}
}

func (x *SchemaPosition) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierSchemaPosition
}

func (x *SchemaPosition) Reset() {
	x.Read((*SchemaPositionViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *SchemaPosition) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *SchemaPosition) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(20)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(20))
	__FileSize := uint(1 * len(x.File))
	__FileOffset, err := writer.Alloc(__FileSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__FileOffset))
	writer.Write4At(offset+4+4, uint32(__FileSize))
	__FileSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.File)), __FileSize, __FileSize}
	writer.WriteAt(__FileOffset, *(*[]byte)(unsafe.Pointer(&__FileSlice)))
	__LineOffset := offset + 12
	writer.Write4At(__LineOffset, *(*uint32)(unsafe.Pointer(&x.Line)))
	__ColumnOffset := offset + 16
	writer.Write4At(__ColumnOffset, *(*uint32)(unsafe.Pointer(&x.Column)))

	return offset, nil
}

func (x *SchemaPosition) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewSchemaPositionViewer(reader, 0), reader)
}

func (x *SchemaPosition) Read(viewer *SchemaPositionViewer, reader *karmem.Reader) {
	__FileString := viewer.File(reader)
	if x.File != __FileString {
		__FileStringCopy := make([]byte, len(__FileString))
		copy(__FileStringCopy, __FileString)
		x.File = *(*string)(unsafe.Pointer(&__FileStringCopy))
	}
	x.Line = viewer.Line()
	x.Column = viewer.Column()
}

type Type struct {
	Schema      string
	PlainSchema string
	Length      uint32
	Format      TypeFormat
	Model       TypeModel
}

func NewType() Type {
	return Type{}
}

func (x *Type) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierType
}

func (x *Type) Reset() {
	x.Read((*TypeViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *Type) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Type) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(26)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(26))
	__SchemaSize := uint(1 * len(x.Schema))
	__SchemaOffset, err := writer.Alloc(__SchemaSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__SchemaOffset))
	writer.Write4At(offset+4+4, uint32(__SchemaSize))
	__SchemaSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Schema)), __SchemaSize, __SchemaSize}
	writer.WriteAt(__SchemaOffset, *(*[]byte)(unsafe.Pointer(&__SchemaSlice)))
	__PlainSchemaSize := uint(1 * len(x.PlainSchema))
	__PlainSchemaOffset, err := writer.Alloc(__PlainSchemaSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__PlainSchemaOffset))
	writer.Write4At(offset+12+4, uint32(__PlainSchemaSize))
	__PlainSchemaSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.PlainSchema)), __PlainSchemaSize, __PlainSchemaSize}
	writer.WriteAt(__PlainSchemaOffset, *(*[]byte)(unsafe.Pointer(&__PlainSchemaSlice)))
	__LengthOffset := offset + 20
	writer.Write4At(__LengthOffset, *(*uint32)(unsafe.Pointer(&x.Length)))
	__FormatOffset := offset + 24
	writer.Write1At(__FormatOffset, *(*uint8)(unsafe.Pointer(&x.Format)))
	__ModelOffset := offset + 25
	writer.Write1At(__ModelOffset, *(*uint8)(unsafe.Pointer(&x.Model)))

	return offset, nil
}

func (x *Type) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewTypeViewer(reader, 0), reader)
}

func (x *Type) Read(viewer *TypeViewer, reader *karmem.Reader) {
	__SchemaString := viewer.Schema(reader)
	if x.Schema != __SchemaString {
		__SchemaStringCopy := make([]byte, len(__SchemaString))
		copy(__SchemaStringCopy, __SchemaString)
		x.Schema = *(*string)(unsafe.Pointer(&__SchemaStringCopy))
	}
	__PlainSchemaString := viewer.PlainSchema(reader)
	if x.PlainSchema != __PlainSchemaString {
		__PlainSchemaStringCopy := make([]byte, len(__PlainSchemaString))
		copy(__PlainSchemaStringCopy, __PlainSchemaString)
		x.PlainSchema = *(*string)(unsafe.Pointer(&__PlainSchemaStringCopy))
	}
	x.Length = viewer.Length()
	x.Format = TypeFormat(viewer.Format())
	x.Model = TypeModel(viewer.Model())
}

type WarningData struct {
	Message   string
	Rule      string
	Type      RuleType
	Languages Language
}

func NewWarningData() WarningData {
	return WarningData{}
}

func (x *WarningData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierWarningData
}

func (x *WarningData) Reset() {
	x.Read((*WarningDataViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *WarningData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *WarningData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(30)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(30))
	__MessageSize := uint(1 * len(x.Message))
	__MessageOffset, err := writer.Alloc(__MessageSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__MessageOffset))
	writer.Write4At(offset+4+4, uint32(__MessageSize))
	__MessageSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Message)), __MessageSize, __MessageSize}
	writer.WriteAt(__MessageOffset, *(*[]byte)(unsafe.Pointer(&__MessageSlice)))
	__RuleSize := uint(1 * len(x.Rule))
	__RuleOffset, err := writer.Alloc(__RuleSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__RuleOffset))
	writer.Write4At(offset+12+4, uint32(__RuleSize))
	__RuleSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Rule)), __RuleSize, __RuleSize}
	writer.WriteAt(__RuleOffset, *(*[]byte)(unsafe.Pointer(&__RuleSlice)))
	__TypeOffset := offset + 20
	writer.Write2At(__TypeOffset, *(*uint16)(unsafe.Pointer(&x.Type)))
	__LanguagesOffset := offset + 22
	writer.Write8At(__LanguagesOffset, *(*uint64)(unsafe.Pointer(&x.Languages)))

	return offset, nil
}

func (x *WarningData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewWarningDataViewer(reader, 0), reader)
}

func (x *WarningData) Read(viewer *WarningDataViewer, reader *karmem.Reader) {
	__MessageString := viewer.Message(reader)
	if x.Message != __MessageString {
		__MessageStringCopy := make([]byte, len(__MessageString))
		copy(__MessageStringCopy, __MessageString)
		x.Message = *(*string)(unsafe.Pointer(&__MessageStringCopy))
	}
	__RuleString := viewer.Rule(reader)
	if x.Rule != __RuleString {
		__RuleStringCopy := make([]byte, len(__RuleString))
		copy(__RuleStringCopy, __RuleString)
		x.Rule = *(*string)(unsafe.Pointer(&__RuleStringCopy))
	}
	x.Type = RuleType(viewer.Type())
	x.Languages = Language(viewer.Languages())
}

type Warning struct {
	Data WarningData
}

func NewWarning() Warning {
	return Warning{}
}

func (x *Warning) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierWarning
}

func (x *Warning) Reset() {
	x.Read((*WarningViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *Warning) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Warning) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(30)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *Warning) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewWarningViewer(reader, 0), reader)
}

func (x *Warning) Read(viewer *WarningViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type PaddingType struct {
	Data Type
}

func NewPaddingType() PaddingType {
	return PaddingType{}
}

func (x *PaddingType) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierPaddingType
}

func (x *PaddingType) Reset() {
	x.Read((*PaddingTypeViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *PaddingType) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *PaddingType) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(26)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *PaddingType) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewPaddingTypeViewer(reader, 0), reader)
}

func (x *PaddingType) Read(viewer *PaddingTypeViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type Tag struct {
	Name  string
	Value string
}

func NewTag() Tag {
	return Tag{}
}

func (x *Tag) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierTag
}

func (x *Tag) Reset() {
	x.Read((*TagViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *Tag) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Tag) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(16)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__NameOffset))
	writer.Write4At(offset+0+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__ValueSize := uint(1 * len(x.Value))
	__ValueOffset, err := writer.Alloc(__ValueSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+8, uint32(__ValueOffset))
	writer.Write4At(offset+8+4, uint32(__ValueSize))
	__ValueSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Value)), __ValueSize, __ValueSize}
	writer.WriteAt(__ValueOffset, *(*[]byte)(unsafe.Pointer(&__ValueSlice)))

	return offset, nil
}

func (x *Tag) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewTagViewer(reader, 0), reader)
}

func (x *Tag) Read(viewer *TagViewer, reader *karmem.Reader) {
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	__ValueString := viewer.Value(reader)
	if x.Value != __ValueString {
		__ValueStringCopy := make([]byte, len(__ValueString))
		copy(__ValueStringCopy, __ValueString)
		x.Value = *(*string)(unsafe.Pointer(&__ValueStringCopy))
	}
}

type StructSize struct {
	Minimum    uint32
	Content    uint32
	Padding    uint32
	Total      uint32
	TotalGroup []PaddingType
}

func NewStructSize() StructSize {
	return StructSize{}
}

func (x *StructSize) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierStructSize
}

func (x *StructSize) Reset() {
	x.Read((*StructSizeViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *StructSize) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructSize) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(28)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(28))
	__MinimumOffset := offset + 4
	writer.Write4At(__MinimumOffset, *(*uint32)(unsafe.Pointer(&x.Minimum)))
	__ContentOffset := offset + 8
	writer.Write4At(__ContentOffset, *(*uint32)(unsafe.Pointer(&x.Content)))
	__PaddingOffset := offset + 12
	writer.Write4At(__PaddingOffset, *(*uint32)(unsafe.Pointer(&x.Padding)))
	__TotalOffset := offset + 16
	writer.Write4At(__TotalOffset, *(*uint32)(unsafe.Pointer(&x.Total)))
	__TotalGroupSize := uint(4 * len(x.TotalGroup))
	__TotalGroupOffset, err := writer.Alloc(__TotalGroupSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__TotalGroupOffset))
	writer.Write4At(offset+20+4, uint32(__TotalGroupSize))
	for i := range x.TotalGroup {
		if _, err := x.TotalGroup[i].Write(writer, __TotalGroupOffset); err != nil {
			return offset, err
		}
		__TotalGroupOffset += 4
	}

	return offset, nil
}

func (x *StructSize) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructSizeViewer(reader, 0), reader)
}

func (x *StructSize) Read(viewer *StructSizeViewer, reader *karmem.Reader) {
	x.Minimum = viewer.Minimum()
	x.Content = viewer.Content()
	x.Padding = viewer.Padding()
	x.Total = viewer.Total()
	__TotalGroupSlice := viewer.TotalGroup(reader)
	__TotalGroupLen := len(__TotalGroupSlice)
	if __TotalGroupLen > cap(x.TotalGroup) {
		x.TotalGroup = append(x.TotalGroup, make([]PaddingType, __TotalGroupLen-len(x.TotalGroup))...)
	}
	if __TotalGroupLen > len(x.TotalGroup) {
		x.TotalGroup = x.TotalGroup[:__TotalGroupLen]
	}
	for i := 0; i < __TotalGroupLen; i++ {
		x.TotalGroup[i].Read(&__TotalGroupSlice[i], reader)
	}
	x.TotalGroup = x.TotalGroup[:__TotalGroupLen]
}

type StructFieldSize struct {
	Minimum    uint32
	Allocation uint32
	Field      uint32
}

func NewStructFieldSize() StructFieldSize {
	return StructFieldSize{}
}

func (x *StructFieldSize) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierStructFieldSize
}

func (x *StructFieldSize) Reset() {
	x.Read((*StructFieldSizeViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *StructFieldSize) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructFieldSize) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(16)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(16))
	__MinimumOffset := offset + 4
	writer.Write4At(__MinimumOffset, *(*uint32)(unsafe.Pointer(&x.Minimum)))
	__AllocationOffset := offset + 8
	writer.Write4At(__AllocationOffset, *(*uint32)(unsafe.Pointer(&x.Allocation)))
	__FieldOffset := offset + 12
	writer.Write4At(__FieldOffset, *(*uint32)(unsafe.Pointer(&x.Field)))

	return offset, nil
}

func (x *StructFieldSize) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructFieldSizeViewer(reader, 0), reader)
}

func (x *StructFieldSize) Read(viewer *StructFieldSizeViewer, reader *karmem.Reader) {
	x.Minimum = viewer.Minimum()
	x.Allocation = viewer.Allocation()
	x.Field = viewer.Field()
}

type EnumFieldData struct {
	Name     string
	Value    string
	Tags     []Tag
	Position SchemaPosition
	Warnings []Warning
}

func NewEnumFieldData() EnumFieldData {
	return EnumFieldData{}
}

func (x *EnumFieldData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierEnumFieldData
}

func (x *EnumFieldData) Reset() {
	x.Read((*EnumFieldDataViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *EnumFieldData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *EnumFieldData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(40)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(40))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__ValueSize := uint(1 * len(x.Value))
	__ValueOffset, err := writer.Alloc(__ValueSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__ValueOffset))
	writer.Write4At(offset+12+4, uint32(__ValueSize))
	__ValueSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Value)), __ValueSize, __ValueSize}
	writer.WriteAt(__ValueOffset, *(*[]byte)(unsafe.Pointer(&__ValueSlice)))
	__TagsSize := uint(16 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__TagsOffset))
	writer.Write4At(offset+20+4, uint32(__TagsSize))
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 16
	}
	__PositionSize := uint(20)
	__PositionOffset, err := writer.Alloc(__PositionSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+28, uint32(__PositionOffset))
	if _, err := x.Position.Write(writer, __PositionOffset); err != nil {
		return offset, err
	}
	__WarningsSize := uint(4 * len(x.Warnings))
	__WarningsOffset, err := writer.Alloc(__WarningsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+32, uint32(__WarningsOffset))
	writer.Write4At(offset+32+4, uint32(__WarningsSize))
	for i := range x.Warnings {
		if _, err := x.Warnings[i].Write(writer, __WarningsOffset); err != nil {
			return offset, err
		}
		__WarningsOffset += 4
	}

	return offset, nil
}

func (x *EnumFieldData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewEnumFieldDataViewer(reader, 0), reader)
}

func (x *EnumFieldData) Read(viewer *EnumFieldDataViewer, reader *karmem.Reader) {
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	__ValueString := viewer.Value(reader)
	if x.Value != __ValueString {
		__ValueStringCopy := make([]byte, len(__ValueString))
		copy(__ValueStringCopy, __ValueString)
		x.Value = *(*string)(unsafe.Pointer(&__ValueStringCopy))
	}
	__TagsSlice := viewer.Tags(reader)
	__TagsLen := len(__TagsSlice)
	if __TagsLen > cap(x.Tags) {
		x.Tags = append(x.Tags, make([]Tag, __TagsLen-len(x.Tags))...)
	}
	if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := 0; i < __TagsLen; i++ {
		x.Tags[i].Read(&__TagsSlice[i], reader)
	}
	x.Tags = x.Tags[:__TagsLen]
	x.Position.Read(viewer.Position(reader), reader)
	__WarningsSlice := viewer.Warnings(reader)
	__WarningsLen := len(__WarningsSlice)
	if __WarningsLen > cap(x.Warnings) {
		x.Warnings = append(x.Warnings, make([]Warning, __WarningsLen-len(x.Warnings))...)
	}
	if __WarningsLen > len(x.Warnings) {
		x.Warnings = x.Warnings[:__WarningsLen]
	}
	for i := 0; i < __WarningsLen; i++ {
		x.Warnings[i].Read(&__WarningsSlice[i], reader)
	}
	x.Warnings = x.Warnings[:__WarningsLen]
}

type EnumField struct {
	Data EnumFieldData
}

func NewEnumField() EnumField {
	return EnumField{}
}

func (x *EnumField) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierEnumField
}

func (x *EnumField) Reset() {
	x.Read((*EnumFieldViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *EnumField) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *EnumField) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(40)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *EnumField) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewEnumFieldViewer(reader, 0), reader)
}

func (x *EnumField) Read(viewer *EnumFieldViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type EnumData struct {
	Name         string
	Type         Type
	Fields       []EnumField
	Tags         []Tag
	IsSequential bool
	Position     SchemaPosition
	Warnings     []Warning
}

func NewEnumData() EnumData {
	return EnumData{}
}

func (x *EnumData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierEnumData
}

func (x *EnumData) Reset() {
	x.Read((*EnumDataViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *EnumData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *EnumData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(45)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(45))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__TypeSize := uint(26)
	__TypeOffset, err := writer.Alloc(__TypeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__TypeOffset))
	if _, err := x.Type.Write(writer, __TypeOffset); err != nil {
		return offset, err
	}
	__FieldsSize := uint(4 * len(x.Fields))
	__FieldsOffset, err := writer.Alloc(__FieldsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__FieldsOffset))
	writer.Write4At(offset+16+4, uint32(__FieldsSize))
	for i := range x.Fields {
		if _, err := x.Fields[i].Write(writer, __FieldsOffset); err != nil {
			return offset, err
		}
		__FieldsOffset += 4
	}
	__TagsSize := uint(16 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+24, uint32(__TagsOffset))
	writer.Write4At(offset+24+4, uint32(__TagsSize))
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 16
	}
	__IsSequentialOffset := offset + 32
	writer.Write1At(__IsSequentialOffset, *(*uint8)(unsafe.Pointer(&x.IsSequential)))
	__PositionSize := uint(20)
	__PositionOffset, err := writer.Alloc(__PositionSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+33, uint32(__PositionOffset))
	if _, err := x.Position.Write(writer, __PositionOffset); err != nil {
		return offset, err
	}
	__WarningsSize := uint(4 * len(x.Warnings))
	__WarningsOffset, err := writer.Alloc(__WarningsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+37, uint32(__WarningsOffset))
	writer.Write4At(offset+37+4, uint32(__WarningsSize))
	for i := range x.Warnings {
		if _, err := x.Warnings[i].Write(writer, __WarningsOffset); err != nil {
			return offset, err
		}
		__WarningsOffset += 4
	}

	return offset, nil
}

func (x *EnumData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewEnumDataViewer(reader, 0), reader)
}

func (x *EnumData) Read(viewer *EnumDataViewer, reader *karmem.Reader) {
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	x.Type.Read(viewer.Type(reader), reader)
	__FieldsSlice := viewer.Fields(reader)
	__FieldsLen := len(__FieldsSlice)
	if __FieldsLen > cap(x.Fields) {
		x.Fields = append(x.Fields, make([]EnumField, __FieldsLen-len(x.Fields))...)
	}
	if __FieldsLen > len(x.Fields) {
		x.Fields = x.Fields[:__FieldsLen]
	}
	for i := 0; i < __FieldsLen; i++ {
		x.Fields[i].Read(&__FieldsSlice[i], reader)
	}
	x.Fields = x.Fields[:__FieldsLen]
	__TagsSlice := viewer.Tags(reader)
	__TagsLen := len(__TagsSlice)
	if __TagsLen > cap(x.Tags) {
		x.Tags = append(x.Tags, make([]Tag, __TagsLen-len(x.Tags))...)
	}
	if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := 0; i < __TagsLen; i++ {
		x.Tags[i].Read(&__TagsSlice[i], reader)
	}
	x.Tags = x.Tags[:__TagsLen]
	x.IsSequential = viewer.IsSequential()
	x.Position.Read(viewer.Position(reader), reader)
	__WarningsSlice := viewer.Warnings(reader)
	__WarningsLen := len(__WarningsSlice)
	if __WarningsLen > cap(x.Warnings) {
		x.Warnings = append(x.Warnings, make([]Warning, __WarningsLen-len(x.Warnings))...)
	}
	if __WarningsLen > len(x.Warnings) {
		x.Warnings = x.Warnings[:__WarningsLen]
	}
	for i := 0; i < __WarningsLen; i++ {
		x.Warnings[i].Read(&__WarningsSlice[i], reader)
	}
	x.Warnings = x.Warnings[:__WarningsLen]
}

type Enumeration struct {
	Data EnumData
}

func NewEnumeration() Enumeration {
	return Enumeration{}
}

func (x *Enumeration) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierEnumeration
}

func (x *Enumeration) Reset() {
	x.Read((*EnumerationViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *Enumeration) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Enumeration) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(45)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *Enumeration) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewEnumerationViewer(reader, 0), reader)
}

func (x *Enumeration) Read(viewer *EnumerationViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type StructFieldData struct {
	Name     string
	Type     Type
	Offset   uint32
	Tags     []Tag
	Size     StructFieldSize
	Position SchemaPosition
	Warnings []Warning
}

func NewStructFieldData() StructFieldData {
	return StructFieldData{}
}

func (x *StructFieldData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierStructFieldData
}

func (x *StructFieldData) Reset() {
	x.Read((*StructFieldDataViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *StructFieldData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructFieldData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(44)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(44))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__TypeSize := uint(26)
	__TypeOffset, err := writer.Alloc(__TypeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__TypeOffset))
	if _, err := x.Type.Write(writer, __TypeOffset); err != nil {
		return offset, err
	}
	__OffsetOffset := offset + 16
	writer.Write4At(__OffsetOffset, *(*uint32)(unsafe.Pointer(&x.Offset)))
	__TagsSize := uint(16 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__TagsOffset))
	writer.Write4At(offset+20+4, uint32(__TagsSize))
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 16
	}
	__SizeSize := uint(16)
	__SizeOffset, err := writer.Alloc(__SizeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+28, uint32(__SizeOffset))
	if _, err := x.Size.Write(writer, __SizeOffset); err != nil {
		return offset, err
	}
	__PositionSize := uint(20)
	__PositionOffset, err := writer.Alloc(__PositionSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+32, uint32(__PositionOffset))
	if _, err := x.Position.Write(writer, __PositionOffset); err != nil {
		return offset, err
	}
	__WarningsSize := uint(4 * len(x.Warnings))
	__WarningsOffset, err := writer.Alloc(__WarningsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+36, uint32(__WarningsOffset))
	writer.Write4At(offset+36+4, uint32(__WarningsSize))
	for i := range x.Warnings {
		if _, err := x.Warnings[i].Write(writer, __WarningsOffset); err != nil {
			return offset, err
		}
		__WarningsOffset += 4
	}

	return offset, nil
}

func (x *StructFieldData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructFieldDataViewer(reader, 0), reader)
}

func (x *StructFieldData) Read(viewer *StructFieldDataViewer, reader *karmem.Reader) {
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	x.Type.Read(viewer.Type(reader), reader)
	x.Offset = viewer.Offset()
	__TagsSlice := viewer.Tags(reader)
	__TagsLen := len(__TagsSlice)
	if __TagsLen > cap(x.Tags) {
		x.Tags = append(x.Tags, make([]Tag, __TagsLen-len(x.Tags))...)
	}
	if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := 0; i < __TagsLen; i++ {
		x.Tags[i].Read(&__TagsSlice[i], reader)
	}
	x.Tags = x.Tags[:__TagsLen]
	x.Size.Read(viewer.Size(reader), reader)
	x.Position.Read(viewer.Position(reader), reader)
	__WarningsSlice := viewer.Warnings(reader)
	__WarningsLen := len(__WarningsSlice)
	if __WarningsLen > cap(x.Warnings) {
		x.Warnings = append(x.Warnings, make([]Warning, __WarningsLen-len(x.Warnings))...)
	}
	if __WarningsLen > len(x.Warnings) {
		x.Warnings = x.Warnings[:__WarningsLen]
	}
	for i := 0; i < __WarningsLen; i++ {
		x.Warnings[i].Read(&__WarningsSlice[i], reader)
	}
	x.Warnings = x.Warnings[:__WarningsLen]
}

type StructField struct {
	Data StructFieldData
}

func NewStructField() StructField {
	return StructField{}
}

func (x *StructField) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierStructField
}

func (x *StructField) Reset() {
	x.Read((*StructFieldViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *StructField) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructField) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(44)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *StructField) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructFieldViewer(reader, 0), reader)
}

func (x *StructField) Read(viewer *StructFieldViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type StructData struct {
	ID       uint64
	Name     string
	Size     StructSize
	Fields   []StructField
	Class    StructClass
	Tags     []Tag
	Packed   bool
	Position SchemaPosition
	Warnings []Warning
}

func NewStructData() StructData {
	return StructData{}
}

func (x *StructData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierStructData
}

func (x *StructData) Reset() {
	x.Read((*StructDataViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *StructData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(54)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(54))
	__IDOffset := offset + 4
	writer.Write8At(__IDOffset, *(*uint64)(unsafe.Pointer(&x.ID)))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__NameOffset))
	writer.Write4At(offset+12+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__SizeSize := uint(28)
	__SizeOffset, err := writer.Alloc(__SizeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__SizeOffset))
	if _, err := x.Size.Write(writer, __SizeOffset); err != nil {
		return offset, err
	}
	__FieldsSize := uint(4 * len(x.Fields))
	__FieldsOffset, err := writer.Alloc(__FieldsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+24, uint32(__FieldsOffset))
	writer.Write4At(offset+24+4, uint32(__FieldsSize))
	for i := range x.Fields {
		if _, err := x.Fields[i].Write(writer, __FieldsOffset); err != nil {
			return offset, err
		}
		__FieldsOffset += 4
	}
	__ClassOffset := offset + 32
	writer.Write1At(__ClassOffset, *(*uint8)(unsafe.Pointer(&x.Class)))
	__TagsSize := uint(16 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+33, uint32(__TagsOffset))
	writer.Write4At(offset+33+4, uint32(__TagsSize))
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 16
	}
	__PackedOffset := offset + 41
	writer.Write1At(__PackedOffset, *(*uint8)(unsafe.Pointer(&x.Packed)))
	__PositionSize := uint(20)
	__PositionOffset, err := writer.Alloc(__PositionSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+42, uint32(__PositionOffset))
	if _, err := x.Position.Write(writer, __PositionOffset); err != nil {
		return offset, err
	}
	__WarningsSize := uint(4 * len(x.Warnings))
	__WarningsOffset, err := writer.Alloc(__WarningsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+46, uint32(__WarningsOffset))
	writer.Write4At(offset+46+4, uint32(__WarningsSize))
	for i := range x.Warnings {
		if _, err := x.Warnings[i].Write(writer, __WarningsOffset); err != nil {
			return offset, err
		}
		__WarningsOffset += 4
	}

	return offset, nil
}

func (x *StructData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructDataViewer(reader, 0), reader)
}

func (x *StructData) Read(viewer *StructDataViewer, reader *karmem.Reader) {
	x.ID = viewer.ID()
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	x.Size.Read(viewer.Size(reader), reader)
	__FieldsSlice := viewer.Fields(reader)
	__FieldsLen := len(__FieldsSlice)
	if __FieldsLen > cap(x.Fields) {
		x.Fields = append(x.Fields, make([]StructField, __FieldsLen-len(x.Fields))...)
	}
	if __FieldsLen > len(x.Fields) {
		x.Fields = x.Fields[:__FieldsLen]
	}
	for i := 0; i < __FieldsLen; i++ {
		x.Fields[i].Read(&__FieldsSlice[i], reader)
	}
	x.Fields = x.Fields[:__FieldsLen]
	x.Class = StructClass(viewer.Class())
	__TagsSlice := viewer.Tags(reader)
	__TagsLen := len(__TagsSlice)
	if __TagsLen > cap(x.Tags) {
		x.Tags = append(x.Tags, make([]Tag, __TagsLen-len(x.Tags))...)
	}
	if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := 0; i < __TagsLen; i++ {
		x.Tags[i].Read(&__TagsSlice[i], reader)
	}
	x.Tags = x.Tags[:__TagsLen]
	x.Packed = viewer.Packed()
	x.Position.Read(viewer.Position(reader), reader)
	__WarningsSlice := viewer.Warnings(reader)
	__WarningsLen := len(__WarningsSlice)
	if __WarningsLen > cap(x.Warnings) {
		x.Warnings = append(x.Warnings, make([]Warning, __WarningsLen-len(x.Warnings))...)
	}
	if __WarningsLen > len(x.Warnings) {
		x.Warnings = x.Warnings[:__WarningsLen]
	}
	for i := 0; i < __WarningsLen; i++ {
		x.Warnings[i].Read(&__WarningsSlice[i], reader)
	}
	x.Warnings = x.Warnings[:__WarningsLen]
}

type Structure struct {
	Data StructData
}

func NewStructure() Structure {
	return Structure{}
}

func (x *Structure) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierStructure
}

func (x *Structure) Reset() {
	x.Read((*StructureViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *Structure) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Structure) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(4)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(54)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *Structure) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructureViewer(reader, 0), reader)
}

func (x *Structure) Read(viewer *StructureViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type ContentSize struct {
	Largest uint32
}

func NewContentSize() ContentSize {
	return ContentSize{}
}

func (x *ContentSize) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierContentSize
}

func (x *ContentSize) Reset() {
	x.Read((*ContentSizeViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *ContentSize) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *ContentSize) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(8))
	__LargestOffset := offset + 4
	writer.Write4At(__LargestOffset, *(*uint32)(unsafe.Pointer(&x.Largest)))

	return offset, nil
}

func (x *ContentSize) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewContentSizeViewer(reader, 0), reader)
}

func (x *ContentSize) Read(viewer *ContentSizeViewer, reader *karmem.Reader) {
	x.Largest = viewer.Largest()
}

type Content struct {
	Tags    []Tag
	Structs []Structure
	Enums   []Enumeration
	Size    ContentSize
	Name    string
	Packed  bool
}

func NewContent() Content {
	return Content{}
}

func (x *Content) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierContent
}

func (x *Content) Reset() {
	x.Read((*ContentViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *Content) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Content) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(41)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(41))
	__TagsSize := uint(16 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__TagsOffset))
	writer.Write4At(offset+4+4, uint32(__TagsSize))
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 16
	}
	__StructsSize := uint(4 * len(x.Structs))
	__StructsOffset, err := writer.Alloc(__StructsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__StructsOffset))
	writer.Write4At(offset+12+4, uint32(__StructsSize))
	for i := range x.Structs {
		if _, err := x.Structs[i].Write(writer, __StructsOffset); err != nil {
			return offset, err
		}
		__StructsOffset += 4
	}
	__EnumsSize := uint(4 * len(x.Enums))
	__EnumsOffset, err := writer.Alloc(__EnumsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__EnumsOffset))
	writer.Write4At(offset+20+4, uint32(__EnumsSize))
	for i := range x.Enums {
		if _, err := x.Enums[i].Write(writer, __EnumsOffset); err != nil {
			return offset, err
		}
		__EnumsOffset += 4
	}
	__SizeSize := uint(8)
	__SizeOffset, err := writer.Alloc(__SizeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+28, uint32(__SizeOffset))
	if _, err := x.Size.Write(writer, __SizeOffset); err != nil {
		return offset, err
	}
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+32, uint32(__NameOffset))
	writer.Write4At(offset+32+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__PackedOffset := offset + 40
	writer.Write1At(__PackedOffset, *(*uint8)(unsafe.Pointer(&x.Packed)))

	return offset, nil
}

func (x *Content) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewContentViewer(reader, 0), reader)
}

func (x *Content) Read(viewer *ContentViewer, reader *karmem.Reader) {
	__TagsSlice := viewer.Tags(reader)
	__TagsLen := len(__TagsSlice)
	if __TagsLen > cap(x.Tags) {
		x.Tags = append(x.Tags, make([]Tag, __TagsLen-len(x.Tags))...)
	}
	if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := 0; i < __TagsLen; i++ {
		x.Tags[i].Read(&__TagsSlice[i], reader)
	}
	x.Tags = x.Tags[:__TagsLen]
	__StructsSlice := viewer.Structs(reader)
	__StructsLen := len(__StructsSlice)
	if __StructsLen > cap(x.Structs) {
		x.Structs = append(x.Structs, make([]Structure, __StructsLen-len(x.Structs))...)
	}
	if __StructsLen > len(x.Structs) {
		x.Structs = x.Structs[:__StructsLen]
	}
	for i := 0; i < __StructsLen; i++ {
		x.Structs[i].Read(&__StructsSlice[i], reader)
	}
	x.Structs = x.Structs[:__StructsLen]
	__EnumsSlice := viewer.Enums(reader)
	__EnumsLen := len(__EnumsSlice)
	if __EnumsLen > cap(x.Enums) {
		x.Enums = append(x.Enums, make([]Enumeration, __EnumsLen-len(x.Enums))...)
	}
	if __EnumsLen > len(x.Enums) {
		x.Enums = x.Enums[:__EnumsLen]
	}
	for i := 0; i < __EnumsLen; i++ {
		x.Enums[i].Read(&__EnumsSlice[i], reader)
	}
	x.Enums = x.Enums[:__EnumsLen]
	x.Size.Read(viewer.Size(reader), reader)
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	x.Packed = viewer.Packed()
}

type SchemaPositionViewer [20]byte

func NewSchemaPositionViewer(reader *karmem.Reader, offset uint32) (v *SchemaPositionViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*SchemaPositionViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*SchemaPositionViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*SchemaPositionViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *SchemaPositionViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *SchemaPositionViewer) File(reader *karmem.Reader) (v string) {
	if 4+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *SchemaPositionViewer) Line() (v uint32) {
	if 12+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
}
func (x *SchemaPositionViewer) Column() (v uint32) {
	if 16+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 16))
}

type TypeViewer [26]byte

func NewTypeViewer(reader *karmem.Reader, offset uint32) (v *TypeViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*TypeViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*TypeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*TypeViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *TypeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *TypeViewer) Schema(reader *karmem.Reader) (v string) {
	if 4+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *TypeViewer) PlainSchema(reader *karmem.Reader) (v string) {
	if 12+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *TypeViewer) Length() (v uint32) {
	if 20+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20))
}
func (x *TypeViewer) Format() (v TypeFormat) {
	if 24+1 > x.size() {
		return v
	}
	return *(*TypeFormat)(unsafe.Add(unsafe.Pointer(x), 24))
}
func (x *TypeViewer) Model() (v TypeModel) {
	if 25+1 > x.size() {
		return v
	}
	return *(*TypeModel)(unsafe.Add(unsafe.Pointer(x), 25))
}

type WarningDataViewer [30]byte

func NewWarningDataViewer(reader *karmem.Reader, offset uint32) (v *WarningDataViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*WarningDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*WarningDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*WarningDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *WarningDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *WarningDataViewer) Message(reader *karmem.Reader) (v string) {
	if 4+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *WarningDataViewer) Rule(reader *karmem.Reader) (v string) {
	if 12+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *WarningDataViewer) Type() (v RuleType) {
	if 20+2 > x.size() {
		return v
	}
	return *(*RuleType)(unsafe.Add(unsafe.Pointer(x), 20))
}
func (x *WarningDataViewer) Languages() (v Language) {
	if 22+8 > x.size() {
		return v
	}
	return *(*Language)(unsafe.Add(unsafe.Pointer(x), 22))
}

type WarningViewer [4]byte

func NewWarningViewer(reader *karmem.Reader, offset uint32) (v *WarningViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*WarningViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*WarningViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *WarningViewer) size() uint32 {
	return 4
}
func (x *WarningViewer) Data(reader *karmem.Reader) (v *WarningDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewWarningDataViewer(reader, offset)
}

type PaddingTypeViewer [4]byte

func NewPaddingTypeViewer(reader *karmem.Reader, offset uint32) (v *PaddingTypeViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*PaddingTypeViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*PaddingTypeViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *PaddingTypeViewer) size() uint32 {
	return 4
}
func (x *PaddingTypeViewer) Data(reader *karmem.Reader) (v *TypeViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewTypeViewer(reader, offset)
}

type TagViewer [16]byte

func NewTagViewer(reader *karmem.Reader, offset uint32) (v *TagViewer) {
	if !reader.IsValidOffset(offset, 16) {
		return (*TagViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*TagViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *TagViewer) size() uint32 {
	return 16
}
func (x *TagViewer) Name(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *TagViewer) Value(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 8))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 8+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}

type StructSizeViewer [28]byte

func NewStructSizeViewer(reader *karmem.Reader, offset uint32) (v *StructSizeViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*StructSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*StructSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*StructSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *StructSizeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *StructSizeViewer) Minimum() (v uint32) {
	if 4+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
}
func (x *StructSizeViewer) Content() (v uint32) {
	if 8+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 8))
}
func (x *StructSizeViewer) Padding() (v uint32) {
	if 12+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
}
func (x *StructSizeViewer) Total() (v uint32) {
	if 16+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 16))
}
func (x *StructSizeViewer) TotalGroup(reader *karmem.Reader) (v []PaddingTypeViewer) {
	if 20+8 > x.size() {
		return []PaddingTypeViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20+4))
	if !reader.IsValidOffset(offset, size) {
		return []PaddingTypeViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]PaddingTypeViewer)(unsafe.Pointer(&slice))
}

type StructFieldSizeViewer [16]byte

func NewStructFieldSizeViewer(reader *karmem.Reader, offset uint32) (v *StructFieldSizeViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*StructFieldSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*StructFieldSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*StructFieldSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *StructFieldSizeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *StructFieldSizeViewer) Minimum() (v uint32) {
	if 4+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
}
func (x *StructFieldSizeViewer) Allocation() (v uint32) {
	if 8+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 8))
}
func (x *StructFieldSizeViewer) Field() (v uint32) {
	if 12+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
}

type EnumFieldDataViewer [40]byte

func NewEnumFieldDataViewer(reader *karmem.Reader, offset uint32) (v *EnumFieldDataViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*EnumFieldDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*EnumFieldDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*EnumFieldDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *EnumFieldDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *EnumFieldDataViewer) Name(reader *karmem.Reader) (v string) {
	if 4+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *EnumFieldDataViewer) Value(reader *karmem.Reader) (v string) {
	if 12+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *EnumFieldDataViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 20+8 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 16)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}
func (x *EnumFieldDataViewer) Position(reader *karmem.Reader) (v *SchemaPositionViewer) {
	if 28+4 > x.size() {
		return (*SchemaPositionViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 28))
	return NewSchemaPositionViewer(reader, offset)
}
func (x *EnumFieldDataViewer) Warnings(reader *karmem.Reader) (v []WarningViewer) {
	if 32+8 > x.size() {
		return []WarningViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 32))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 32+4))
	if !reader.IsValidOffset(offset, size) {
		return []WarningViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]WarningViewer)(unsafe.Pointer(&slice))
}

type EnumFieldViewer [4]byte

func NewEnumFieldViewer(reader *karmem.Reader, offset uint32) (v *EnumFieldViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*EnumFieldViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*EnumFieldViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *EnumFieldViewer) size() uint32 {
	return 4
}
func (x *EnumFieldViewer) Data(reader *karmem.Reader) (v *EnumFieldDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewEnumFieldDataViewer(reader, offset)
}

type EnumDataViewer [45]byte

func NewEnumDataViewer(reader *karmem.Reader, offset uint32) (v *EnumDataViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*EnumDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*EnumDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*EnumDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *EnumDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *EnumDataViewer) Name(reader *karmem.Reader) (v string) {
	if 4+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *EnumDataViewer) Type(reader *karmem.Reader) (v *TypeViewer) {
	if 12+4 > x.size() {
		return (*TypeViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
	return NewTypeViewer(reader, offset)
}
func (x *EnumDataViewer) Fields(reader *karmem.Reader) (v []EnumFieldViewer) {
	if 16+8 > x.size() {
		return []EnumFieldViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return []EnumFieldViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]EnumFieldViewer)(unsafe.Pointer(&slice))
}
func (x *EnumDataViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 24+8 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 24))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 24+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 16)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}
func (x *EnumDataViewer) IsSequential() (v bool) {
	if 32+1 > x.size() {
		return v
	}
	return *(*bool)(unsafe.Add(unsafe.Pointer(x), 32))
}
func (x *EnumDataViewer) Position(reader *karmem.Reader) (v *SchemaPositionViewer) {
	if 33+4 > x.size() {
		return (*SchemaPositionViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 33))
	return NewSchemaPositionViewer(reader, offset)
}
func (x *EnumDataViewer) Warnings(reader *karmem.Reader) (v []WarningViewer) {
	if 37+8 > x.size() {
		return []WarningViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 37))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 37+4))
	if !reader.IsValidOffset(offset, size) {
		return []WarningViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]WarningViewer)(unsafe.Pointer(&slice))
}

type EnumerationViewer [4]byte

func NewEnumerationViewer(reader *karmem.Reader, offset uint32) (v *EnumerationViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*EnumerationViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*EnumerationViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *EnumerationViewer) size() uint32 {
	return 4
}
func (x *EnumerationViewer) Data(reader *karmem.Reader) (v *EnumDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewEnumDataViewer(reader, offset)
}

type StructFieldDataViewer [44]byte

func NewStructFieldDataViewer(reader *karmem.Reader, offset uint32) (v *StructFieldDataViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*StructFieldDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*StructFieldDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*StructFieldDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *StructFieldDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *StructFieldDataViewer) Name(reader *karmem.Reader) (v string) {
	if 4+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *StructFieldDataViewer) Type(reader *karmem.Reader) (v *TypeViewer) {
	if 12+4 > x.size() {
		return (*TypeViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
	return NewTypeViewer(reader, offset)
}
func (x *StructFieldDataViewer) Offset() (v uint32) {
	if 16+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 16))
}
func (x *StructFieldDataViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 20+8 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 16)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}
func (x *StructFieldDataViewer) Size(reader *karmem.Reader) (v *StructFieldSizeViewer) {
	if 28+4 > x.size() {
		return (*StructFieldSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 28))
	return NewStructFieldSizeViewer(reader, offset)
}
func (x *StructFieldDataViewer) Position(reader *karmem.Reader) (v *SchemaPositionViewer) {
	if 32+4 > x.size() {
		return (*SchemaPositionViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 32))
	return NewSchemaPositionViewer(reader, offset)
}
func (x *StructFieldDataViewer) Warnings(reader *karmem.Reader) (v []WarningViewer) {
	if 36+8 > x.size() {
		return []WarningViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 36))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 36+4))
	if !reader.IsValidOffset(offset, size) {
		return []WarningViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]WarningViewer)(unsafe.Pointer(&slice))
}

type StructFieldViewer [4]byte

func NewStructFieldViewer(reader *karmem.Reader, offset uint32) (v *StructFieldViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*StructFieldViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*StructFieldViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *StructFieldViewer) size() uint32 {
	return 4
}
func (x *StructFieldViewer) Data(reader *karmem.Reader) (v *StructFieldDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewStructFieldDataViewer(reader, offset)
}

type StructDataViewer [54]byte

func NewStructDataViewer(reader *karmem.Reader, offset uint32) (v *StructDataViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*StructDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*StructDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*StructDataViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *StructDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *StructDataViewer) ID() (v uint64) {
	if 4+8 > x.size() {
		return v
	}
	return *(*uint64)(unsafe.Add(unsafe.Pointer(x), 4))
}
func (x *StructDataViewer) Name(reader *karmem.Reader) (v string) {
	if 12+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *StructDataViewer) Size(reader *karmem.Reader) (v *StructSizeViewer) {
	if 20+4 > x.size() {
		return (*StructSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20))
	return NewStructSizeViewer(reader, offset)
}
func (x *StructDataViewer) Fields(reader *karmem.Reader) (v []StructFieldViewer) {
	if 24+8 > x.size() {
		return []StructFieldViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 24))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 24+4))
	if !reader.IsValidOffset(offset, size) {
		return []StructFieldViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]StructFieldViewer)(unsafe.Pointer(&slice))
}
func (x *StructDataViewer) Class() (v StructClass) {
	if 32+1 > x.size() {
		return v
	}
	return *(*StructClass)(unsafe.Add(unsafe.Pointer(x), 32))
}
func (x *StructDataViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 33+8 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 33))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 33+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 16)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}
func (x *StructDataViewer) Packed() (v bool) {
	if 41+1 > x.size() {
		return v
	}
	return *(*bool)(unsafe.Add(unsafe.Pointer(x), 41))
}
func (x *StructDataViewer) Position(reader *karmem.Reader) (v *SchemaPositionViewer) {
	if 42+4 > x.size() {
		return (*SchemaPositionViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 42))
	return NewSchemaPositionViewer(reader, offset)
}
func (x *StructDataViewer) Warnings(reader *karmem.Reader) (v []WarningViewer) {
	if 46+8 > x.size() {
		return []WarningViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 46))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 46+4))
	if !reader.IsValidOffset(offset, size) {
		return []WarningViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]WarningViewer)(unsafe.Pointer(&slice))
}

type StructureViewer [4]byte

func NewStructureViewer(reader *karmem.Reader, offset uint32) (v *StructureViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*StructureViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*StructureViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *StructureViewer) size() uint32 {
	return 4
}
func (x *StructureViewer) Data(reader *karmem.Reader) (v *StructDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	return NewStructDataViewer(reader, offset)
}

type ContentSizeViewer [8]byte

func NewContentSizeViewer(reader *karmem.Reader, offset uint32) (v *ContentSizeViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*ContentSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*ContentSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*ContentSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *ContentSizeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *ContentSizeViewer) Largest() (v uint32) {
	if 4+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
}

type ContentViewer [41]byte

func NewContentViewer(reader *karmem.Reader, offset uint32) (v *ContentViewer) {
	if !reader.IsValidOffset(offset, 4) {
		return (*ContentViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*ContentViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*ContentViewer)(unsafe.Pointer(&_Null[0]))
	}
	return v
}

func (x *ContentViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(x))
}
func (x *ContentViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 4+8 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 16)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}
func (x *ContentViewer) Structs(reader *karmem.Reader) (v []StructureViewer) {
	if 12+8 > x.size() {
		return []StructureViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return []StructureViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]StructureViewer)(unsafe.Pointer(&slice))
}
func (x *ContentViewer) Enums(reader *karmem.Reader) (v []EnumerationViewer) {
	if 20+8 > x.size() {
		return []EnumerationViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 20+4))
	if !reader.IsValidOffset(offset, size) {
		return []EnumerationViewer{}
	}
	length := uintptr(size / 4)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]EnumerationViewer)(unsafe.Pointer(&slice))
}
func (x *ContentViewer) Size(reader *karmem.Reader) (v *ContentSizeViewer) {
	if 28+4 > x.size() {
		return (*ContentSizeViewer)(unsafe.Pointer(&_Null[0]))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 28))
	return NewContentSizeViewer(reader, offset)
}
func (x *ContentViewer) Name(reader *karmem.Reader) (v string) {
	if 32+8 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 32))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 32+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *ContentViewer) Packed() (v bool) {
	if 40+1 > x.size() {
		return v
	}
	return *(*bool)(unsafe.Add(unsafe.Pointer(x), 40))
}
