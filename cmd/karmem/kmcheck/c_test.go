package kmcheck

import (
	"testing"
)

func TestC(t *testing.T) {
	testValidation(t, NewC(), `karmem test @packed(true);

	enum TestEnum uint32 {
		None;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum SizeOf uint32 @error() {
		None;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		If;
	}

	struct TestStructErr table {
		While []char @error();
	}

	struct Static table @error() {
		Foo []char;
	}

	struct SomethingErr table {
		Foo []char;
		FooLength int32 @error();
	}

	struct Something table {
		FooLength int32;
	}
`)
}
