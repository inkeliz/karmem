package kmcheck

import (
	"testing"
)

func TestOdin(t *testing.T) {
	testValidation(t, NewOdin(), `karmem test @packed(true);

	enum TestEnum uint32 {
		None;
		One;
	}

	struct TestStruct table {
		Foo int32;
		Bar int16;
	}

	struct TestStructErr table {
		And []char;
	}
`)
}
