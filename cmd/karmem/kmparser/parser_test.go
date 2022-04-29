package kmparser

import (
	"os"
	"strings"
	"testing"
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
