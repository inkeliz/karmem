package kmparser

import (
	"fmt"
)

type Tags []Tag

func (x Tags) Get(s string) (string, bool) {
	for i := range x {
		if x[i].Name == s {
			return x[i].Value, true
		}
	}
	return "", false
}

func (x Tags) GetBoolean(s string) (r int, err error) {
	for i := range x {
		if x[i].Name == s {
			switch x[i].Value {
			case "":
				fallthrough
			case "true":
				return 1, nil
			case "false":
				return 0, nil
			default:
				return -1, fmt.Errorf(`invalid value of "%s" for %s`, x[i].Value, x[i].Name)
			}
		}
	}
	return -1, nil
}

func (x *Type) IsBasic() bool {
	return x.Model == TypeModelSingle
}

func (x *Type) IsNative() bool {
	return x.Format == TypeFormatPrimitive
}

func (x *Type) IsArray() bool {
	return x.Model == TypeModelArray
}

func (x *Type) IsSlice() bool {
	return x.Model == TypeModelSlice || x.Model == TypeModelSliceLimited
}

func (x *Type) IsLimited() bool {
	return x.Model == TypeModelSliceLimited
}

func (x *Type) IsInteger() bool {
	switch x.PlainSchema {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
		return true
	default:
		return false
	}
}

func (x *Type) IsString() bool {
	return x.PlainSchema == "char"
}

func (x *Type) IsBool() bool {
	return x.PlainSchema == "bool"
}

func (x *Type) IsInline() bool {
	switch x.Model {
	case TypeModelSlice, TypeModelSliceLimited:
		return false
	default:
		return x.Format != TypeFormatTable
	}
}

func (x *Type) IsEnum() bool {
	return x.Format == TypeFormatEnum
}

func (x StructData) IsTable() bool {
	return x.Class == StructClassTable
}
