package kmcheck

import (
	"testing"
)

func TestAssemblyScript(t *testing.T) {
	testValidation(t, NewAssemblyScript(), `karmem test @packed(true);

	enum TestEnum uint32 {
		None;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum Package uint32 @error() {
		None;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		Private;
	}

	struct TestStructErr table {
		And []char @error();
	}

	struct From table @error() {
		Foo []char;
	}
`)
}
