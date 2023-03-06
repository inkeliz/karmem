package kmcheck

import (
	"strings"
	"sync"
	"testing"

	"karmem.org/cmd/karmem/kmparser"
)

func TestCheck(t *testing.T) {
	parsed, err := kmparser.NewReader("", strings.NewReader(`karmem test @packed(true);

	enum TestEnum uint32 {
		Unknown;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum KarmemSize uint32 {
		Unknown;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		KarmemPointer;
	}

	struct TestStructErr table {
		KarmemSize []char @error();
	}

	struct Private table @error() {
		Foo []char;
	}

	struct KarmemPointer table @error() {
	Foo []char;
	}`)).Parser()

	if err != nil {
		t.Fatal(err)
	}

	Check(parsed)
	testValidate(t, parsed)
}

func testValidation(t *testing.T, v Validator, s string) {
	parsed, err := kmparser.NewReader("", strings.NewReader(s)).Parser()
	if err != nil {
		panic(err)
	}

	check(&sync.Mutex{}, v, parsed)
	testValidate(t, parsed)
}

func testValidate(t *testing.T, parsed *kmparser.Content) {
	for _, x := range parsed.Structs {
		result := len(x.Data.Warnings) > 0

		if _, ok := kmparser.Tags(x.Data.Tags).Get("error"); ok != result {
			t.Errorf("Invalid warnigs: %v", x.Data.Name)
		}

		for _, y := range x.Data.Fields {
			result := len(y.Data.Warnings) > 0
			if _, ok := kmparser.Tags(y.Data.Tags).Get("error"); ok != result {
				t.Errorf("Invalid warnigs: %v %v", y.Data.Name, y.Data.Warnings)
			}
		}
	}

	for _, x := range parsed.Enums {
		result := len(x.Data.Warnings) > 0

		if _, ok := kmparser.Tags(x.Data.Tags).Get("error"); ok != result {
			t.Errorf("Invalid warnigs: %v", x.Data.Name)
		}

		result = false
		for _, y := range x.Data.Fields {
			if len(y.Data.Warnings) > 0 {
				result = true
				break
			}
		}
		if _, ok := kmparser.Tags(x.Data.Tags).Get("hasError"); ok != result {
			t.Errorf("Invalid warnigs: %v %v", x.Data.Name, x.Data.Warnings)
		}
	}
}
