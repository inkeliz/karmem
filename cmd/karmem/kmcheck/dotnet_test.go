package kmcheck

import (
	"testing"
)

func TestDotNet(t *testing.T) {
	testValidation(t, NewDotNet(), `karmem test @packed(true);

	enum TestEnum uint32 {
		None;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum Namespace uint32 @error() {
		None;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		Private;
	}

	struct TestStructErr table {
		Virtual []char @error();
	}

	struct Private table @error() {
		Foo []char;
	}

	struct Collision table {
		Collision []char @error();
	}

	enum Collision uint32 @error() {
		None;
		Collision;
	}
`)
}
