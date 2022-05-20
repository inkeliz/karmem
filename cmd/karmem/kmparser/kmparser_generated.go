package kmparser

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

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

type Tag struct {
	Name  string
	Value string
}

func (x *Tag) Reset() {
	x.Name = x.Name[:0]
	x.Value = x.Value[:0]
}

func (x *Tag) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Tag) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
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
	writer.Write4At(offset+0+4+4, 1)
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__ValueSize := uint(1 * len(x.Value))
	__ValueOffset, err := writer.Alloc(__ValueSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__ValueOffset))
	writer.Write4At(offset+12+4, uint32(__ValueSize))
	writer.Write4At(offset+12+4+4, 1)
	__ValueSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Value)), __ValueSize, __ValueSize}
	writer.WriteAt(__ValueOffset, *(*[]byte)(unsafe.Pointer(&__ValueSlice)))

	return offset, nil
}

func (x *Tag) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewTagViewer(reader, 0), reader)
}

func (x *Tag) Read(viewer *TagViewer, reader *karmem.Reader) {
	x.Name = string(viewer.Name(reader))
	x.Value = string(viewer.Value(reader))
}

type StructSize struct {
	Minimum    uint32
	Content    uint32
	Padding    uint32
	Total      uint32
	TotalGroup []uint8
}

func (x *StructSize) Reset() {
	x.Minimum = 0
	x.Content = 0
	x.Padding = 0
	x.Total = 0
	x.TotalGroup = x.TotalGroup[:0]
}

func (x *StructSize) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructSize) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(40)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__MinimumOffset := offset + 4
	writer.Write4At(__MinimumOffset, *(*uint32)(unsafe.Pointer(&x.Minimum)))
	__ContentOffset := offset + 8
	writer.Write4At(__ContentOffset, *(*uint32)(unsafe.Pointer(&x.Content)))
	__PaddingOffset := offset + 12
	writer.Write4At(__PaddingOffset, *(*uint32)(unsafe.Pointer(&x.Padding)))
	__TotalOffset := offset + 16
	writer.Write4At(__TotalOffset, *(*uint32)(unsafe.Pointer(&x.Total)))
	__TotalGroupSize := uint(1 * len(x.TotalGroup))
	__TotalGroupOffset, err := writer.Alloc(__TotalGroupSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__TotalGroupOffset))
	writer.Write4At(offset+20+4, uint32(__TotalGroupSize))
	writer.Write4At(offset+20+4+4, 1)
	__TotalGroupSlice := *(*[3]uint)(unsafe.Pointer(&x.TotalGroup))
	__TotalGroupSlice[1] = __TotalGroupSize
	__TotalGroupSlice[2] = __TotalGroupSize
	writer.WriteAt(__TotalGroupOffset, *(*[]byte)(unsafe.Pointer(&__TotalGroupSlice)))

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
		x.TotalGroup = append(x.TotalGroup, make([]uint8, __TotalGroupLen-len(x.TotalGroup))...)
	} else if __TotalGroupLen > len(x.TotalGroup) {
		x.TotalGroup = x.TotalGroup[:__TotalGroupLen]
	}
	copy(x.TotalGroup, __TotalGroupSlice)
	for i := __TotalGroupLen; i < len(x.TotalGroup); i++ {
		x.TotalGroup[i] = 0
	}
	x.TotalGroup = x.TotalGroup[:__TotalGroupLen]
}

type StructFieldSize struct {
	Minimum    uint32
	Allocation uint32
	Field      uint32
}

func (x *StructFieldSize) Reset() {
	x.Minimum = 0
	x.Allocation = 0
	x.Field = 0
}

func (x *StructFieldSize) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructFieldSize) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(24)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
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

type Type struct {
	Schema      string
	PlainSchema string
	Length      uint32
	Format      TypeFormat
	Model       TypeModel
}

func (x *Type) Reset() {
	x.Schema = x.Schema[:0]
	x.PlainSchema = x.PlainSchema[:0]
	x.Length = 0
	x.Format = 0
	x.Model = 0
}

func (x *Type) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Type) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(40)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__SchemaSize := uint(1 * len(x.Schema))
	__SchemaOffset, err := writer.Alloc(__SchemaSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__SchemaOffset))
	writer.Write4At(offset+4+4, uint32(__SchemaSize))
	writer.Write4At(offset+4+4+4, 1)
	__SchemaSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Schema)), __SchemaSize, __SchemaSize}
	writer.WriteAt(__SchemaOffset, *(*[]byte)(unsafe.Pointer(&__SchemaSlice)))
	__PlainSchemaSize := uint(1 * len(x.PlainSchema))
	__PlainSchemaOffset, err := writer.Alloc(__PlainSchemaSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__PlainSchemaOffset))
	writer.Write4At(offset+16+4, uint32(__PlainSchemaSize))
	writer.Write4At(offset+16+4+4, 1)
	__PlainSchemaSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.PlainSchema)), __PlainSchemaSize, __PlainSchemaSize}
	writer.WriteAt(__PlainSchemaOffset, *(*[]byte)(unsafe.Pointer(&__PlainSchemaSlice)))
	__LengthOffset := offset + 28
	writer.Write4At(__LengthOffset, *(*uint32)(unsafe.Pointer(&x.Length)))
	__FormatOffset := offset + 32
	writer.Write1At(__FormatOffset, *(*uint8)(unsafe.Pointer(&x.Format)))
	__ModelOffset := offset + 33
	writer.Write1At(__ModelOffset, *(*uint8)(unsafe.Pointer(&x.Model)))

	return offset, nil
}

func (x *Type) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewTypeViewer(reader, 0), reader)
}

func (x *Type) Read(viewer *TypeViewer, reader *karmem.Reader) {
	x.Schema = string(viewer.Schema(reader))
	x.PlainSchema = string(viewer.PlainSchema(reader))
	x.Length = viewer.Length()
	x.Format = TypeFormat(viewer.Format())
	x.Model = TypeModel(viewer.Model())
}

type EnumFieldData struct {
	Name  string
	Value string
	Tags  []Tag
}

func (x *EnumFieldData) Reset() {
	x.Name = x.Name[:0]
	x.Value = x.Value[:0]
	for i := range x.Tags {
		x.Tags[i].Reset()
	}
	x.Tags = x.Tags[:0]
}

func (x *EnumFieldData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *EnumFieldData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(48)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	writer.Write4At(offset+4+4+4, 1)
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__ValueSize := uint(1 * len(x.Value))
	__ValueOffset, err := writer.Alloc(__ValueSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__ValueOffset))
	writer.Write4At(offset+16+4, uint32(__ValueSize))
	writer.Write4At(offset+16+4+4, 1)
	__ValueSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Value)), __ValueSize, __ValueSize}
	writer.WriteAt(__ValueOffset, *(*[]byte)(unsafe.Pointer(&__ValueSlice)))
	__TagsSize := uint(32 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+28, uint32(__TagsOffset))
	writer.Write4At(offset+28+4, uint32(__TagsSize))
	writer.Write4At(offset+28+4+4, 32)
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 32
	}

	return offset, nil
}

func (x *EnumFieldData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewEnumFieldDataViewer(reader, 0), reader)
}

func (x *EnumFieldData) Read(viewer *EnumFieldDataViewer, reader *karmem.Reader) {
	x.Name = string(viewer.Name(reader))
	x.Value = string(viewer.Value(reader))
	__TagsSlice := viewer.Tags(reader)
	__TagsLen := len(__TagsSlice)
	if __TagsLen > cap(x.Tags) {
		x.Tags = append(x.Tags, make([]Tag, __TagsLen-len(x.Tags))...)
	} else if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := range x.Tags {
		if i >= __TagsLen {
			x.Tags[i].Reset()
		} else {
			x.Tags[i].Read(&__TagsSlice[i], reader)
		}
	}
	x.Tags = x.Tags[:__TagsLen]
}

type EnumField struct {
	Data EnumFieldData
}

func (x *EnumField) Reset() {
	x.Data.Reset()
}

func (x *EnumField) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *EnumField) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(48)
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
	Name   string
	Type   Type
	Fields []EnumField
}

func (x *EnumData) Reset() {
	x.Name = x.Name[:0]
	x.Type.Reset()
	for i := range x.Fields {
		x.Fields[i].Reset()
	}
	x.Fields = x.Fields[:0]
}

func (x *EnumData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *EnumData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(40)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	writer.Write4At(offset+4+4+4, 1)
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__TypeSize := uint(40)
	__TypeOffset, err := writer.Alloc(__TypeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__TypeOffset))
	if _, err := x.Type.Write(writer, __TypeOffset); err != nil {
		return offset, err
	}
	__FieldsSize := uint(8 * len(x.Fields))
	__FieldsOffset, err := writer.Alloc(__FieldsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+20, uint32(__FieldsOffset))
	writer.Write4At(offset+20+4, uint32(__FieldsSize))
	writer.Write4At(offset+20+4+4, 8)
	for i := range x.Fields {
		if _, err := x.Fields[i].Write(writer, __FieldsOffset); err != nil {
			return offset, err
		}
		__FieldsOffset += 8
	}

	return offset, nil
}

func (x *EnumData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewEnumDataViewer(reader, 0), reader)
}

func (x *EnumData) Read(viewer *EnumDataViewer, reader *karmem.Reader) {
	x.Name = string(viewer.Name(reader))
	x.Type.Read(viewer.Type(reader), reader)
	__FieldsSlice := viewer.Fields(reader)
	__FieldsLen := len(__FieldsSlice)
	if __FieldsLen > cap(x.Fields) {
		x.Fields = append(x.Fields, make([]EnumField, __FieldsLen-len(x.Fields))...)
	} else if __FieldsLen > len(x.Fields) {
		x.Fields = x.Fields[:__FieldsLen]
	}
	for i := range x.Fields {
		if i >= __FieldsLen {
			x.Fields[i].Reset()
		} else {
			x.Fields[i].Read(&__FieldsSlice[i], reader)
		}
	}
	x.Fields = x.Fields[:__FieldsLen]
}

type Enum struct {
	Data EnumData
}

func (x *Enum) Reset() {
	x.Data.Reset()
}

func (x *Enum) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Enum) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
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

func (x *Enum) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewEnumViewer(reader, 0), reader)
}

func (x *Enum) Read(viewer *EnumViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type StructFieldData struct {
	Name   string
	Type   Type
	Offset uint32
	Tags   []Tag
	Size   StructFieldSize
}

func (x *StructFieldData) Reset() {
	x.Name = x.Name[:0]
	x.Type.Reset()
	x.Offset = 0
	for i := range x.Tags {
		x.Tags[i].Reset()
	}
	x.Tags = x.Tags[:0]
	x.Size.Reset()
}

func (x *StructFieldData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructFieldData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(48)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	writer.Write4At(offset+4+4+4, 1)
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__TypeSize := uint(40)
	__TypeOffset, err := writer.Alloc(__TypeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__TypeOffset))
	if _, err := x.Type.Write(writer, __TypeOffset); err != nil {
		return offset, err
	}
	__OffsetOffset := offset + 20
	writer.Write4At(__OffsetOffset, *(*uint32)(unsafe.Pointer(&x.Offset)))
	__TagsSize := uint(32 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+24, uint32(__TagsOffset))
	writer.Write4At(offset+24+4, uint32(__TagsSize))
	writer.Write4At(offset+24+4+4, 32)
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 32
	}
	__SizeSize := uint(24)
	__SizeOffset, err := writer.Alloc(__SizeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+36, uint32(__SizeOffset))
	if _, err := x.Size.Write(writer, __SizeOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *StructFieldData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructFieldDataViewer(reader, 0), reader)
}

func (x *StructFieldData) Read(viewer *StructFieldDataViewer, reader *karmem.Reader) {
	x.Name = string(viewer.Name(reader))
	x.Type.Read(viewer.Type(reader), reader)
	x.Offset = viewer.Offset()
	__TagsSlice := viewer.Tags(reader)
	__TagsLen := len(__TagsSlice)
	if __TagsLen > cap(x.Tags) {
		x.Tags = append(x.Tags, make([]Tag, __TagsLen-len(x.Tags))...)
	} else if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := range x.Tags {
		if i >= __TagsLen {
			x.Tags[i].Reset()
		} else {
			x.Tags[i].Read(&__TagsSlice[i], reader)
		}
	}
	x.Tags = x.Tags[:__TagsLen]
	x.Size.Read(viewer.Size(reader), reader)
}

type StructField struct {
	Data StructFieldData
}

func (x *StructField) Reset() {
	x.Data.Reset()
}

func (x *StructField) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructField) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(48)
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
	ID     uint64
	Name   string
	Size   StructSize
	Fields []StructField
	Class  StructClass
}

func (x *StructData) Reset() {
	x.ID = 0
	x.Name = x.Name[:0]
	x.Size.Reset()
	for i := range x.Fields {
		x.Fields[i].Reset()
	}
	x.Fields = x.Fields[:0]
	x.Class = 0
}

func (x *StructData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *StructData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(48)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__IDOffset := offset + 4
	writer.Write8At(__IDOffset, *(*uint64)(unsafe.Pointer(&x.ID)))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__NameOffset))
	writer.Write4At(offset+12+4, uint32(__NameSize))
	writer.Write4At(offset+12+4+4, 1)
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__SizeSize := uint(40)
	__SizeOffset, err := writer.Alloc(__SizeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+24, uint32(__SizeOffset))
	if _, err := x.Size.Write(writer, __SizeOffset); err != nil {
		return offset, err
	}
	__FieldsSize := uint(8 * len(x.Fields))
	__FieldsOffset, err := writer.Alloc(__FieldsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+28, uint32(__FieldsOffset))
	writer.Write4At(offset+28+4, uint32(__FieldsSize))
	writer.Write4At(offset+28+4+4, 8)
	for i := range x.Fields {
		if _, err := x.Fields[i].Write(writer, __FieldsOffset); err != nil {
			return offset, err
		}
		__FieldsOffset += 8
	}
	__ClassOffset := offset + 40
	writer.Write1At(__ClassOffset, *(*uint8)(unsafe.Pointer(&x.Class)))

	return offset, nil
}

func (x *StructData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructDataViewer(reader, 0), reader)
}

func (x *StructData) Read(viewer *StructDataViewer, reader *karmem.Reader) {
	x.ID = viewer.ID()
	x.Name = string(viewer.Name(reader))
	x.Size.Read(viewer.Size(reader), reader)
	__FieldsSlice := viewer.Fields(reader)
	__FieldsLen := len(__FieldsSlice)
	if __FieldsLen > cap(x.Fields) {
		x.Fields = append(x.Fields, make([]StructField, __FieldsLen-len(x.Fields))...)
	} else if __FieldsLen > len(x.Fields) {
		x.Fields = x.Fields[:__FieldsLen]
	}
	for i := range x.Fields {
		if i >= __FieldsLen {
			x.Fields[i].Reset()
		} else {
			x.Fields[i].Read(&__FieldsSlice[i], reader)
		}
	}
	x.Fields = x.Fields[:__FieldsLen]
	x.Class = StructClass(viewer.Class())
}

type Struct struct {
	Data StructData
}

func (x *Struct) Reset() {
	x.Data.Reset()
}

func (x *Struct) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Struct) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(48)
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

func (x *Struct) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewStructViewer(reader, 0), reader)
}

func (x *Struct) Read(viewer *StructViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type ContentSize struct {
	Largest uint32
}

func (x *ContentSize) Reset() {
	x.Largest = 0
}

func (x *ContentSize) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *ContentSize) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(16)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
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

type ContentOptions struct {
	Module string
	Import string
	Prefix string
}

func (x *ContentOptions) Reset() {
	x.Module = x.Module[:0]
	x.Import = x.Import[:0]
	x.Prefix = x.Prefix[:0]
}

func (x *ContentOptions) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *ContentOptions) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(48)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__ModuleSize := uint(1 * len(x.Module))
	__ModuleOffset, err := writer.Alloc(__ModuleSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__ModuleOffset))
	writer.Write4At(offset+4+4, uint32(__ModuleSize))
	writer.Write4At(offset+4+4+4, 1)
	__ModuleSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Module)), __ModuleSize, __ModuleSize}
	writer.WriteAt(__ModuleOffset, *(*[]byte)(unsafe.Pointer(&__ModuleSlice)))
	__ImportSize := uint(1 * len(x.Import))
	__ImportOffset, err := writer.Alloc(__ImportSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__ImportOffset))
	writer.Write4At(offset+16+4, uint32(__ImportSize))
	writer.Write4At(offset+16+4+4, 1)
	__ImportSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Import)), __ImportSize, __ImportSize}
	writer.WriteAt(__ImportOffset, *(*[]byte)(unsafe.Pointer(&__ImportSlice)))
	__PrefixSize := uint(1 * len(x.Prefix))
	__PrefixOffset, err := writer.Alloc(__PrefixSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+28, uint32(__PrefixOffset))
	writer.Write4At(offset+28+4, uint32(__PrefixSize))
	writer.Write4At(offset+28+4+4, 1)
	__PrefixSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Prefix)), __PrefixSize, __PrefixSize}
	writer.WriteAt(__PrefixOffset, *(*[]byte)(unsafe.Pointer(&__PrefixSlice)))

	return offset, nil
}

func (x *ContentOptions) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewContentOptionsViewer(reader, 0), reader)
}

func (x *ContentOptions) Read(viewer *ContentOptionsViewer, reader *karmem.Reader) {
	x.Module = string(viewer.Module(reader))
	x.Import = string(viewer.Import(reader))
	x.Prefix = string(viewer.Prefix(reader))
}

type Content struct {
	Tags    []Tag
	Structs []Struct
	Enums   []Enum
	Size    ContentSize
	Options ContentOptions
	Module  string
}

func (x *Content) Reset() {
	for i := range x.Tags {
		x.Tags[i].Reset()
	}
	x.Tags = x.Tags[:0]
	for i := range x.Structs {
		x.Structs[i].Reset()
	}
	x.Structs = x.Structs[:0]
	for i := range x.Enums {
		x.Enums[i].Reset()
	}
	x.Enums = x.Enums[:0]
	x.Size.Reset()
	x.Options.Reset()
	x.Module = x.Module[:0]
}

func (x *Content) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Content) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(64)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(size))
	__TagsSize := uint(32 * len(x.Tags))
	__TagsOffset, err := writer.Alloc(__TagsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__TagsOffset))
	writer.Write4At(offset+4+4, uint32(__TagsSize))
	writer.Write4At(offset+4+4+4, 32)
	for i := range x.Tags {
		if _, err := x.Tags[i].Write(writer, __TagsOffset); err != nil {
			return offset, err
		}
		__TagsOffset += 32
	}
	__StructsSize := uint(8 * len(x.Structs))
	__StructsOffset, err := writer.Alloc(__StructsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__StructsOffset))
	writer.Write4At(offset+16+4, uint32(__StructsSize))
	writer.Write4At(offset+16+4+4, 8)
	for i := range x.Structs {
		if _, err := x.Structs[i].Write(writer, __StructsOffset); err != nil {
			return offset, err
		}
		__StructsOffset += 8
	}
	__EnumsSize := uint(8 * len(x.Enums))
	__EnumsOffset, err := writer.Alloc(__EnumsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+28, uint32(__EnumsOffset))
	writer.Write4At(offset+28+4, uint32(__EnumsSize))
	writer.Write4At(offset+28+4+4, 8)
	for i := range x.Enums {
		if _, err := x.Enums[i].Write(writer, __EnumsOffset); err != nil {
			return offset, err
		}
		__EnumsOffset += 8
	}
	__SizeSize := uint(16)
	__SizeOffset, err := writer.Alloc(__SizeSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+40, uint32(__SizeOffset))
	if _, err := x.Size.Write(writer, __SizeOffset); err != nil {
		return offset, err
	}
	__OptionsSize := uint(48)
	__OptionsOffset, err := writer.Alloc(__OptionsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+44, uint32(__OptionsOffset))
	if _, err := x.Options.Write(writer, __OptionsOffset); err != nil {
		return offset, err
	}
	__ModuleSize := uint(1 * len(x.Module))
	__ModuleOffset, err := writer.Alloc(__ModuleSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+48, uint32(__ModuleOffset))
	writer.Write4At(offset+48+4, uint32(__ModuleSize))
	writer.Write4At(offset+48+4+4, 1)
	__ModuleSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Module)), __ModuleSize, __ModuleSize}
	writer.WriteAt(__ModuleOffset, *(*[]byte)(unsafe.Pointer(&__ModuleSlice)))

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
	} else if __TagsLen > len(x.Tags) {
		x.Tags = x.Tags[:__TagsLen]
	}
	for i := range x.Tags {
		if i >= __TagsLen {
			x.Tags[i].Reset()
		} else {
			x.Tags[i].Read(&__TagsSlice[i], reader)
		}
	}
	x.Tags = x.Tags[:__TagsLen]
	__StructsSlice := viewer.Structs(reader)
	__StructsLen := len(__StructsSlice)
	if __StructsLen > cap(x.Structs) {
		x.Structs = append(x.Structs, make([]Struct, __StructsLen-len(x.Structs))...)
	} else if __StructsLen > len(x.Structs) {
		x.Structs = x.Structs[:__StructsLen]
	}
	for i := range x.Structs {
		if i >= __StructsLen {
			x.Structs[i].Reset()
		} else {
			x.Structs[i].Read(&__StructsSlice[i], reader)
		}
	}
	x.Structs = x.Structs[:__StructsLen]
	__EnumsSlice := viewer.Enums(reader)
	__EnumsLen := len(__EnumsSlice)
	if __EnumsLen > cap(x.Enums) {
		x.Enums = append(x.Enums, make([]Enum, __EnumsLen-len(x.Enums))...)
	} else if __EnumsLen > len(x.Enums) {
		x.Enums = x.Enums[:__EnumsLen]
	}
	for i := range x.Enums {
		if i >= __EnumsLen {
			x.Enums[i].Reset()
		} else {
			x.Enums[i].Read(&__EnumsSlice[i], reader)
		}
	}
	x.Enums = x.Enums[:__EnumsLen]
	x.Size.Read(viewer.Size(reader), reader)
	x.Options.Read(viewer.Options(reader), reader)
	x.Module = string(viewer.Module(reader))
}

type TagViewer struct {
	_data [32]byte
}

var _NullTagViewer = TagViewer{}

func NewTagViewer(reader *karmem.Reader, offset uint32) (v *TagViewer) {
	if !reader.IsValidOffset(offset, 32) {
		return &_NullTagViewer
	}
	v = (*TagViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *TagViewer) size() uint32 {
	return 32
}
func (x *TagViewer) Name(reader *karmem.Reader) (v []byte) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *TagViewer) Value(reader *karmem.Reader) (v []byte) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}

type StructSizeViewer struct {
	_data [40]byte
}

var _NullStructSizeViewer = StructSizeViewer{}

func NewStructSizeViewer(reader *karmem.Reader, offset uint32) (v *StructSizeViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullStructSizeViewer
	}
	v = (*StructSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructSizeViewer
	}
	return v
}

func (x *StructSizeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *StructSizeViewer) Minimum() (v uint32) {
	if 4+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
}
func (x *StructSizeViewer) Content() (v uint32) {
	if 8+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 8))
}
func (x *StructSizeViewer) Padding() (v uint32) {
	if 12+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12))
}
func (x *StructSizeViewer) Total() (v uint32) {
	if 16+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
}
func (x *StructSizeViewer) TotalGroup(reader *karmem.Reader) (v []uint8) {
	if 20+12 > x.size() {
		return []uint8{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 20))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 20+4))
	if !reader.IsValidOffset(offset, size) {
		return []uint8{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]uint8)(unsafe.Pointer(&slice))
}

type StructFieldSizeViewer struct {
	_data [24]byte
}

var _NullStructFieldSizeViewer = StructFieldSizeViewer{}

func NewStructFieldSizeViewer(reader *karmem.Reader, offset uint32) (v *StructFieldSizeViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullStructFieldSizeViewer
	}
	v = (*StructFieldSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructFieldSizeViewer
	}
	return v
}

func (x *StructFieldSizeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *StructFieldSizeViewer) Minimum() (v uint32) {
	if 4+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
}
func (x *StructFieldSizeViewer) Allocation() (v uint32) {
	if 8+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 8))
}
func (x *StructFieldSizeViewer) Field() (v uint32) {
	if 12+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12))
}

type TypeViewer struct {
	_data [40]byte
}

var _NullTypeViewer = TypeViewer{}

func NewTypeViewer(reader *karmem.Reader, offset uint32) (v *TypeViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullTypeViewer
	}
	v = (*TypeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullTypeViewer
	}
	return v
}

func (x *TypeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *TypeViewer) Schema(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *TypeViewer) PlainSchema(reader *karmem.Reader) (v []byte) {
	if 16+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *TypeViewer) Length() (v uint32) {
	if 28+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28))
}
func (x *TypeViewer) Format() (v TypeFormat) {
	if 32+1 > x.size() {
		return v
	}
	return *(*TypeFormat)(unsafe.Add(unsafe.Pointer(&x._data), 32))
}
func (x *TypeViewer) Model() (v TypeModel) {
	if 33+1 > x.size() {
		return v
	}
	return *(*TypeModel)(unsafe.Add(unsafe.Pointer(&x._data), 33))
}

type EnumFieldDataViewer struct {
	_data [48]byte
}

var _NullEnumFieldDataViewer = EnumFieldDataViewer{}

func NewEnumFieldDataViewer(reader *karmem.Reader, offset uint32) (v *EnumFieldDataViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullEnumFieldDataViewer
	}
	v = (*EnumFieldDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullEnumFieldDataViewer
	}
	return v
}

func (x *EnumFieldDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *EnumFieldDataViewer) Name(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *EnumFieldDataViewer) Value(reader *karmem.Reader) (v []byte) {
	if 16+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *EnumFieldDataViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 28+12 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 32)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}

type EnumFieldViewer struct {
	_data [8]byte
}

var _NullEnumFieldViewer = EnumFieldViewer{}

func NewEnumFieldViewer(reader *karmem.Reader, offset uint32) (v *EnumFieldViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullEnumFieldViewer
	}
	v = (*EnumFieldViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *EnumFieldViewer) size() uint32 {
	return 8
}
func (x *EnumFieldViewer) Data(reader *karmem.Reader) (v *EnumFieldDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	if !reader.IsValidOffset(offset, 48) {
		return &_NullEnumFieldDataViewer
	}
	v = (*EnumFieldDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullEnumFieldDataViewer
	}
	return v
}

type EnumDataViewer struct {
	_data [40]byte
}

var _NullEnumDataViewer = EnumDataViewer{}

func NewEnumDataViewer(reader *karmem.Reader, offset uint32) (v *EnumDataViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullEnumDataViewer
	}
	v = (*EnumDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullEnumDataViewer
	}
	return v
}

func (x *EnumDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *EnumDataViewer) Name(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *EnumDataViewer) Type(reader *karmem.Reader) (v *TypeViewer) {
	if 16+4 > x.size() {
		return &_NullTypeViewer
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	if !reader.IsValidOffset(offset, 40) {
		return &_NullTypeViewer
	}
	v = (*TypeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullTypeViewer
	}
	return v
}
func (x *EnumDataViewer) Fields(reader *karmem.Reader) (v []EnumFieldViewer) {
	if 20+12 > x.size() {
		return []EnumFieldViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 20))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 20+4))
	if !reader.IsValidOffset(offset, size) {
		return []EnumFieldViewer{}
	}
	length := uintptr(size / 8)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]EnumFieldViewer)(unsafe.Pointer(&slice))
}

type EnumViewer struct {
	_data [8]byte
}

var _NullEnumViewer = EnumViewer{}

func NewEnumViewer(reader *karmem.Reader, offset uint32) (v *EnumViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullEnumViewer
	}
	v = (*EnumViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *EnumViewer) size() uint32 {
	return 8
}
func (x *EnumViewer) Data(reader *karmem.Reader) (v *EnumDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	if !reader.IsValidOffset(offset, 40) {
		return &_NullEnumDataViewer
	}
	v = (*EnumDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullEnumDataViewer
	}
	return v
}

type StructFieldDataViewer struct {
	_data [48]byte
}

var _NullStructFieldDataViewer = StructFieldDataViewer{}

func NewStructFieldDataViewer(reader *karmem.Reader, offset uint32) (v *StructFieldDataViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullStructFieldDataViewer
	}
	v = (*StructFieldDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructFieldDataViewer
	}
	return v
}

func (x *StructFieldDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *StructFieldDataViewer) Name(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *StructFieldDataViewer) Type(reader *karmem.Reader) (v *TypeViewer) {
	if 16+4 > x.size() {
		return &_NullTypeViewer
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	if !reader.IsValidOffset(offset, 40) {
		return &_NullTypeViewer
	}
	v = (*TypeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullTypeViewer
	}
	return v
}
func (x *StructFieldDataViewer) Offset() (v uint32) {
	if 20+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 20))
}
func (x *StructFieldDataViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 24+12 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 24))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 24+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 32)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}
func (x *StructFieldDataViewer) Size(reader *karmem.Reader) (v *StructFieldSizeViewer) {
	if 36+4 > x.size() {
		return &_NullStructFieldSizeViewer
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 36))
	if !reader.IsValidOffset(offset, 24) {
		return &_NullStructFieldSizeViewer
	}
	v = (*StructFieldSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructFieldSizeViewer
	}
	return v
}

type StructFieldViewer struct {
	_data [8]byte
}

var _NullStructFieldViewer = StructFieldViewer{}

func NewStructFieldViewer(reader *karmem.Reader, offset uint32) (v *StructFieldViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullStructFieldViewer
	}
	v = (*StructFieldViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *StructFieldViewer) size() uint32 {
	return 8
}
func (x *StructFieldViewer) Data(reader *karmem.Reader) (v *StructFieldDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	if !reader.IsValidOffset(offset, 48) {
		return &_NullStructFieldDataViewer
	}
	v = (*StructFieldDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructFieldDataViewer
	}
	return v
}

type StructDataViewer struct {
	_data [48]byte
}

var _NullStructDataViewer = StructDataViewer{}

func NewStructDataViewer(reader *karmem.Reader, offset uint32) (v *StructDataViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullStructDataViewer
	}
	v = (*StructDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructDataViewer
	}
	return v
}

func (x *StructDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *StructDataViewer) ID() (v uint64) {
	if 4+8 > x.size() {
		return v
	}
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 4))
}
func (x *StructDataViewer) Name(reader *karmem.Reader) (v []byte) {
	if 12+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *StructDataViewer) Size(reader *karmem.Reader) (v *StructSizeViewer) {
	if 24+4 > x.size() {
		return &_NullStructSizeViewer
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 24))
	if !reader.IsValidOffset(offset, 40) {
		return &_NullStructSizeViewer
	}
	v = (*StructSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructSizeViewer
	}
	return v
}
func (x *StructDataViewer) Fields(reader *karmem.Reader) (v []StructFieldViewer) {
	if 28+12 > x.size() {
		return []StructFieldViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28+4))
	if !reader.IsValidOffset(offset, size) {
		return []StructFieldViewer{}
	}
	length := uintptr(size / 8)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]StructFieldViewer)(unsafe.Pointer(&slice))
}
func (x *StructDataViewer) Class() (v StructClass) {
	if 40+1 > x.size() {
		return v
	}
	return *(*StructClass)(unsafe.Add(unsafe.Pointer(&x._data), 40))
}

type StructViewer struct {
	_data [8]byte
}

var _NullStructViewer = StructViewer{}

func NewStructViewer(reader *karmem.Reader, offset uint32) (v *StructViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullStructViewer
	}
	v = (*StructViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *StructViewer) size() uint32 {
	return 8
}
func (x *StructViewer) Data(reader *karmem.Reader) (v *StructDataViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	if !reader.IsValidOffset(offset, 48) {
		return &_NullStructDataViewer
	}
	v = (*StructDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullStructDataViewer
	}
	return v
}

type ContentSizeViewer struct {
	_data [16]byte
}

var _NullContentSizeViewer = ContentSizeViewer{}

func NewContentSizeViewer(reader *karmem.Reader, offset uint32) (v *ContentSizeViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullContentSizeViewer
	}
	v = (*ContentSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullContentSizeViewer
	}
	return v
}

func (x *ContentSizeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *ContentSizeViewer) Largest() (v uint32) {
	if 4+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
}

type ContentOptionsViewer struct {
	_data [48]byte
}

var _NullContentOptionsViewer = ContentOptionsViewer{}

func NewContentOptionsViewer(reader *karmem.Reader, offset uint32) (v *ContentOptionsViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullContentOptionsViewer
	}
	v = (*ContentOptionsViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullContentOptionsViewer
	}
	return v
}

func (x *ContentOptionsViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *ContentOptionsViewer) Module(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *ContentOptionsViewer) Import(reader *karmem.Reader) (v []byte) {
	if 16+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *ContentOptionsViewer) Prefix(reader *karmem.Reader) (v []byte) {
	if 28+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}

type ContentViewer struct {
	_data [64]byte
}

var _NullContentViewer = ContentViewer{}

func NewContentViewer(reader *karmem.Reader, offset uint32) (v *ContentViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullContentViewer
	}
	v = (*ContentViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullContentViewer
	}
	return v
}

func (x *ContentViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *ContentViewer) Tags(reader *karmem.Reader) (v []TagViewer) {
	if 4+12 > x.size() {
		return []TagViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []TagViewer{}
	}
	length := uintptr(size / 32)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TagViewer)(unsafe.Pointer(&slice))
}
func (x *ContentViewer) Structs(reader *karmem.Reader) (v []StructViewer) {
	if 16+12 > x.size() {
		return []StructViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return []StructViewer{}
	}
	length := uintptr(size / 8)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]StructViewer)(unsafe.Pointer(&slice))
}
func (x *ContentViewer) Enums(reader *karmem.Reader) (v []EnumViewer) {
	if 28+12 > x.size() {
		return []EnumViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 28+4))
	if !reader.IsValidOffset(offset, size) {
		return []EnumViewer{}
	}
	length := uintptr(size / 8)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]EnumViewer)(unsafe.Pointer(&slice))
}
func (x *ContentViewer) Size(reader *karmem.Reader) (v *ContentSizeViewer) {
	if 40+4 > x.size() {
		return &_NullContentSizeViewer
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 40))
	if !reader.IsValidOffset(offset, 16) {
		return &_NullContentSizeViewer
	}
	v = (*ContentSizeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullContentSizeViewer
	}
	return v
}
func (x *ContentViewer) Options(reader *karmem.Reader) (v *ContentOptionsViewer) {
	if 44+4 > x.size() {
		return &_NullContentOptionsViewer
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 44))
	if !reader.IsValidOffset(offset, 48) {
		return &_NullContentOptionsViewer
	}
	v = (*ContentOptionsViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullContentOptionsViewer
	}
	return v
}
func (x *ContentViewer) Module(reader *karmem.Reader) (v []byte) {
	if 48+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 48))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 48+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
