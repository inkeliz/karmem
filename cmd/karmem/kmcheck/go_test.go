package kmcheck

import (
	"testing"
)

func TestGolang(t *testing.T) {
	testValidation(t, NewGolang(), `karmem test @packed(true);

	enum TestEnum uint32 {
		None;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum Make uint32 @error() {
		None;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		Append;
	}

	struct TestStructErr table {
		Append []char @error();
	}

	struct Append table @error() {
		Foo []char;
	}
`)
}
