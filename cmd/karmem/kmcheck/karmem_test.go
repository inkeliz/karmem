package kmcheck

import (
	"testing"
)

func TestKarmem(t *testing.T) {
	testValidation(t, NewKarmem(), `karmem test @packed(true);

	enum TestEnum uint32 {
		None;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	enum KarmemSize uint32 {
		None;
		One;
	}

	enum TestEnumValid uint32 {
		None;
		KarmemPointer;
	}

	struct TestStructErr table {
		KarmemSize []char @error();
	}

	struct KarmemPointer table @error() {
		Foo []char;
	}
`)
}
