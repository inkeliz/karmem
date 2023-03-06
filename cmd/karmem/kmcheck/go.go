package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewGolang()) }

type Golang struct {
	RestrictedWords *RestrictedWords
}

func (v *Golang) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
}

func (v *Golang) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	v.RestrictedWords.CheckEnum(mutex, parsed, target)
}

func NewGolang() *Golang {
	return &Golang{
		RestrictedWords: &RestrictedWords{
			Language: kmparser.LanguageGolang,
			Rules: []WordRule{
				NewMatchRule("break"),
				NewMatchRule("default"),
				NewMatchRule("delete"),
				NewMatchRule("func"),
				NewMatchRule("interface"),
				NewMatchRule("select"),
				NewMatchRule("case"),
				NewMatchRule("defer"),
				NewMatchRule("go"),
				NewMatchRule("map"),
				NewMatchRule("struct"),
				NewMatchRule("chan"),
				NewMatchRule("else"),
				NewMatchRule("goto"),
				NewMatchRule("package"),
				NewMatchRule("switch"),
				NewMatchRule("const"),
				NewMatchRule("fallthrough"),
				NewMatchRule("if"),
				NewMatchRule("range"),
				NewMatchRule("type"),
				NewMatchRule("continue"),
				NewMatchRule("for"),
				NewMatchRule("import"),
				NewMatchRule("return"),
				NewMatchRule("var"),
				NewMatchRule("append"),
				NewMatchRule("bool"),
				NewMatchRule("byte"),
				NewMatchRule("cap"),
				NewMatchRule("close"),
				NewMatchRule("complex"),
				NewMatchRule("complex64"),
				NewMatchRule("complex128"),
				NewMatchRule("uint16"),
				NewMatchRule("copy"),
				NewMatchRule("false"),
				NewMatchRule("float32"),
				NewMatchRule("float64"),
				NewMatchRule("imag"),
				NewMatchRule("int"),
				NewMatchRule("int8"),
				NewMatchRule("int16"),
				NewMatchRule("uint32"),
				NewMatchRule("int32"),
				NewMatchRule("int64"),
				NewMatchRule("iota"),
				NewMatchRule("len"),
				NewMatchRule("make"),
				NewMatchRule("new"),
				NewMatchRule("nil"),
				NewMatchRule("panic"),
				NewMatchRule("uint64"),
				NewMatchRule("real"),
				NewMatchRule("recover"),
				NewMatchRule("string"),
				NewMatchRule("true"),
				NewMatchRule("uint"),
				NewMatchRule("uint8"),
				NewMatchRule("uintptr"),
			},
		},
	}
}
