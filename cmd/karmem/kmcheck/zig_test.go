package kmcheck

import (
	"testing"
)

func TestZig(t *testing.T) {
	testValidation(t, NewZig(), `karmem test @packed(true);

	enum TestEnum uint32 {
		None;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum Asm uint32 @error() {
		None;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		Defer;
	}

	struct TestStructErr table {
		Align []char @error();
	}

	struct Async table @error() {
		Foo []char;
	}
`)
}
