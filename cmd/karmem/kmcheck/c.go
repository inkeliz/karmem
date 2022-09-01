package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewC()) }

type C struct {
	RestrictedWords      *RestrictedWords
	CollisionArraySuffix *CollisionArraySuffix
}

func (v *C) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
	v.CollisionArraySuffix.CheckStruct(mutex, parsed, target)
}

func (v *C) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	v.RestrictedWords.CheckEnum(mutex, parsed, target)
}

func NewC() *C {
	return &C{
		RestrictedWords: &RestrictedWords{
			Language: kmparser.LanguageC,
			Rules: []WordRule{
				NewMatchRule("auto"),
				NewMatchRule("else"),
				NewMatchRule("long"),
				NewMatchRule("switch"),
				NewMatchRule("break"),
				NewMatchRule("enum"),
				NewMatchRule("register"),
				NewMatchRule("typedef"),
				NewMatchRule("case"),
				NewMatchRule("extern"),
				NewMatchRule("return"),
				NewMatchRule("union"),
				NewMatchRule("char"),
				NewMatchRule("float"),
				NewMatchRule("short"),
				NewMatchRule("unsigned"),
				NewMatchRule("const"),
				NewMatchRule("for"),
				NewMatchRule("signed"),
				NewMatchRule("void"),
				NewMatchRule("continue"),
				NewMatchRule("goto"),
				NewMatchRule("sizeof"),
				NewMatchRule("volatile"),
				NewMatchRule("default"),
				NewMatchRule("if"),
				NewMatchRule("static"),
				NewMatchRule("while"),
				NewMatchRule("do"),
				NewMatchRule("int"),
				NewMatchRule("struct"),
				NewMatchRule("double"),
				NewMatchRegexRule("^int[0-9]+_t$"),
				NewMatchRegexRule("^uint[0-9]+_t$"),
				NewMatchRegexRule("^int_fast[0-9]+_t$"),
				NewMatchRegexRule("^uint_fast[0-9]+_t$"),
				NewMatchRegexRule("^int_least[0-9]+_t$"),
				NewMatchRegexRule("^uint_least[0-9]+_t$"),
				NewMatchRegexRule("true"),
				NewMatchRegexRule("false"),
				NewMatchRegexRule("null"),
				NewMatchRegexRule("bool"),
			},
		},
		CollisionArraySuffix: &CollisionArraySuffix{
			Language: kmparser.LanguageC,
			Rules: []WordRule{
				NewMatchSuffix("size"),
				NewMatchSuffix("pointer"),
				NewMatchSuffix("length"),
			},
		},
	}
}
