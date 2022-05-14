package kmparser

import (
	"golang.org/x/crypto/blake2b"
	"os"
	"strings"
	"testing"
	"unsafe"
)

func TestNewReader(t *testing.T) {
	path := "testdata/basic.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := NewReader(path, f)
	if _, err := r.Parser(); err != nil {
		t.Error(err)
	}
}

func TestInlineOfTables(t *testing.T) {
	path := "testdata/tableinline.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := NewReader(path, f)
	if _, err := r.Parser(); err == nil {
		t.Error("inline of table is not valid")
	}
}

func TestInlineOfTables2(t *testing.T) {
	path := "testdata/tableinline2.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := NewReader(path, f)
	if _, err := r.Parser(); err == nil {
		t.Error("inline of table is not valid")
	}
}

func TestInlineOfEnums(t *testing.T) {
	path := "testdata/enuminline.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := NewReader(path, f)
	if _, err := r.Parser(); err == nil {
		t.Error("inline of enum is not valid")
	}
}

func TestInlineOfEnums2(t *testing.T) {
	path := "testdata/enuminline2.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := NewReader(path, f)
	if _, err := r.Parser(); err == nil {
		t.Error("inline of enum is not valid")
	}
}

func TestInlining(t *testing.T) {

	path := "testdata/inline.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := NewReader(path, f)
	k, err := r.Parser()
	if err != nil {
		t.Error(err)
	}

	for _, v := range k.Struct {
		for _, v := range v.Fields {
			switch {
			case strings.Contains(v.Name, "BasicPtr"):
				if v.inline {
					t.Errorf("expect inline at %s %v", v.Name, v.inline)
				}
				if !v.ValueType.IsBasic() {
					t.Errorf("expect array at %s %v", v.Name, v.inline)
				}
			case strings.Contains(v.Name, "Basic"):
				if !v.inline {
					t.Errorf("expect inline at %s %v", v.Name, v.inline)
				}
				if !v.ValueType.IsBasic() {
					t.Errorf("expect array at %s %v", v.Name, v.inline)
				}
			case strings.Contains(v.Name, "Limited"):
				if !v.ValueType.IsLimited() {
					t.Errorf("expect limited at %s", v.Name)
				}
				if v.ValueType.Length() == 0 {
					t.Errorf("unexpected zero lenght at %s", v.Name)
				}
				fallthrough
			case strings.Contains(v.Name, "String"):
				fallthrough
			case strings.Contains(v.Name, "Slice"):
				if v.inline {
					t.Errorf("expect inline at %s %v", v.Name, v.inline)
				}
				if !v.ValueType.IsSlice() {
					t.Errorf("expect slice at %s %v", v.Name, v.inline)
				}
				if v.Size() != 4 {
					t.Errorf("wrong size at %s with size %d", v.Name, v.Size())
				}
			default:
				if v.ValueType.IsLimited() {
					t.Errorf("unexpected limited at %s %v", v.Name, v)
				}
				if !v.inline {
					t.Errorf("unexpected limited at %s", v.Name)
				}
				if !v.ValueType.IsArray() {
					t.Errorf("expect array at %s %v", v.Name, v.inline)
				}
				if v.Size() == 4 {
					t.Errorf("wrong size at %s with %d", v.Name, v.Size())
				}
				if v.ValueType.Length() == 0 {
					t.Errorf("unexpected zero lenght at %s", v.Name)
				}
			}
		}
	}
}

func TestEntropyIdentifier(t *testing.T) {
	path := "testdata/key.km"
	f, err := os.Open(path)
	if err != nil {
		t.Error(f)
		return
	}

	r := NewReader(path, f)
	parsed, err := r.Parser()
	if err != nil {
		t.Error(err)
	}

	if parsed.Struct[0].Name != "Point" || parsed.Struct[1].Name != "User" {
		t.Error("invalid struct name")
	}

	h, _ := blake2b.New(8, nil)
	h.Write([]byte(parsed.Header.GetTag("entropy").Value))
	h.Write([]byte(parsed.Struct[0].Name))

	if *(*uint64)(unsafe.Pointer(&h.Sum(nil)[0])) != parsed.Struct[0].ID {
		t.Error("invalid id")
	}

	h.Reset()
	h.Write([]byte(parsed.Header.GetTag("entropy").Value))
	h.Write([]byte(parsed.Struct[1].Name))
	if *(*uint64)(unsafe.Pointer(&h.Sum(nil)[0])) != parsed.Struct[1].ID {
		t.Error("invalid id")
	}

}
