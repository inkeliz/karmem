package kmcheck

import (
	"testing"
)

func TestSwift(t *testing.T) {
	testValidation(t, NewSwift(), `karmem test @packed(true);

	enum TestEnum uint32 {
		Unknown;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum Protocol uint32 @error() {
		Unknown;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		Break;
	}

	struct TestStructErr table {
		While []char @error();
	}

	struct Lazy table @error() {
		Foo []char;
	}
`)
}
