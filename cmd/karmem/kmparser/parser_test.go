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

	for _, v := range k.Structs {
		if v.Data.ID == 0 {
			t.Errorf("invalid id")
		}
		for _, v := range v.Data.Fields {
			switch {
			case strings.Contains(v.Data.Name, "BasicPtr"):
				if v.Data.Type.IsInline() {
					t.Errorf("expect inline at %s %v", v.Data.Name, v.Data.Type.IsInline())
				}
				if !v.Data.Type.IsBasic() {
					t.Errorf("expect array at %s %v", v.Data.Name, v.Data.Type.IsInline())
				}
			case strings.Contains(v.Data.Name, "Basic"):
				if !v.Data.Type.IsInline() {
					t.Errorf("expect inline at %s %v", v.Data.Name, v.Data.Type.IsInline())
				}
				if !v.Data.Type.IsBasic() {
					t.Errorf("expect array at %s %v", v.Data.Name, v.Data.Type.IsInline())
				}
			case strings.Contains(v.Data.Name, "Limited"):
				if !v.Data.Type.IsLimited() {
					t.Errorf("expect limited at %s", v.Data.Name)
				}
				if v.Data.Type.Length == 0 {
					t.Errorf("unexpected zero lenght at %s", v.Data.Name)
				}
				fallthrough
			case strings.Contains(v.Data.Name, "String"):
				fallthrough
			case strings.Contains(v.Data.Name, "Slice"):
				if v.Data.Type.IsInline() {
					t.Errorf("expect inline at %s %v", v.Data.Name, v.Data.Type.IsInline())
				}
				if !v.Data.Type.IsSlice() {
					t.Errorf("expect slice at %s %v", v.Data.Name, v.Data.Type.IsInline())
				}
				if v.Data.Size.Field != 12 {
					t.Errorf("wrong size at %s with size %d", v.Data.Name, v.Data.Size.Field)
				}
			default:
				if v.Data.Type.IsLimited() {
					t.Errorf("unexpected limited at %s %v", v.Data.Name, v)
				}
				if !v.Data.Type.IsInline() {
					t.Errorf("unexpected limited at %s", v.Data.Name)
				}
				if !v.Data.Type.IsArray() {
					t.Errorf("expect array at %s %v", v.Data.Name, v.Data.Type.IsArray())
				}
				if v.Data.Type.Length == 0 {
					t.Errorf("unexpected zero lenght at %s", v.Data.Name)
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

	if parsed.Structs[0].Data.Name != "Point" || parsed.Structs[1].Data.Name != "User" {
		t.Error("invalid struct name")
	}

	h, _ := blake2b.New(8, nil)
	entropy, ok := Tags(parsed.Tags).Get("entropy")
	if !ok {
		t.Error("not found entropy")
	}
	h.Write([]byte(entropy))
	h.Write([]byte(parsed.Structs[0].Data.Name))

	if *(*uint64)(unsafe.Pointer(&h.Sum(nil)[0])) != parsed.Structs[0].Data.ID {
		t.Error("invalid id", *(*uint64)(unsafe.Pointer(&h.Sum(nil)[0])), parsed.Structs[0].Data.ID)
	}

	h.Reset()
	entropy, ok = Tags(parsed.Tags).Get("entropy")
	if !ok {
		t.Error("not found entropy")
	}
	h.Write([]byte(entropy))
	h.Write([]byte(parsed.Structs[1].Data.Name))
	if *(*uint64)(unsafe.Pointer(&h.Sum(nil)[0])) != parsed.Structs[1].Data.ID {
		t.Error("invalid id")
	}

}
