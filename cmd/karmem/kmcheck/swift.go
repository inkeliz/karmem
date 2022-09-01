package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewSwift()) }

type Swift struct {
	RestrictedWords *RestrictedWords
}

func (v *Swift) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
}

func (v *Swift) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	v.RestrictedWords.CheckEnum(mutex, parsed, target)
}

func NewSwift() *Swift {
	return &Swift{
		RestrictedWords: &RestrictedWords{
			Language: kmparser.LanguageSwift,
			Rules: []WordRule{
				NewMatchRule("associatedtype"),
				NewMatchRule("class"),
				NewMatchRule("deinit"),
				NewMatchRule("enum"),
				NewMatchRule("extension"),
				NewMatchRule("fileprivate"),
				NewMatchRule("func"),
				NewMatchRule("import"),
				NewMatchRule("init"),
				NewMatchRule("inout"),
				NewMatchRule("internal"),
				NewMatchRule("let"),
				NewMatchRule("open"),
				NewMatchRule("operator"),
				NewMatchRule("private"),
				NewMatchRule("precedencegroup"),
				NewMatchRule("protocol"),
				NewMatchRule("public"),
				NewMatchRule("rethrows"),
				NewMatchRule("static"),
				NewMatchRule("struct"),
				NewMatchRule("subscript"),
				NewMatchRule("typealias"),
				NewMatchRule("var"),
				NewMatchRule("break"),
				NewMatchRule("case"),
				NewMatchRule("catch"),
				NewMatchRule("continue"),
				NewMatchRule("default"),
				NewMatchRule("defer"),
				NewMatchRule("do"),
				NewMatchRule("else"),
				NewMatchRule("fallthrough"),
				NewMatchRule("for"),
				NewMatchRule("guard"),
				NewMatchRule("if"),
				NewMatchRule("in"),
				NewMatchRule("repeat"),
				NewMatchRule("return"),
				NewMatchRule("throw"),
				NewMatchRule("switch"),
				NewMatchRule("where"),
				NewMatchRule("while"),
				NewMatchRule("Any"),
				NewMatchRule("as"),
				NewMatchRule("catch"),
				NewMatchRule("false"),
				NewMatchRule("is"),
				NewMatchRule("nil"),
				NewMatchRule("rethrows"),
				NewMatchRule("self"),
				NewMatchRule("Self"),
				NewMatchRule("super"),
				NewMatchRule("throw"),
				NewMatchRule("throws"),
				NewMatchRule("true"),
				NewMatchRule("try"),
				NewMatchRule("_"),
				NewMatchRule("associativity"),
				NewMatchRule("convenience"),
				NewMatchRule("didSet"),
				NewMatchRule("dynamic"),
				NewMatchRule("final"),
				NewMatchRule("get"),
				NewMatchRule("indirect"),
				NewMatchRule("infix"),
				NewMatchRule("lazy"),
				NewMatchRule("left"),
				NewMatchRule("mutating"),
				NewMatchRule("none"),
				NewMatchRule("nonmutating"),
				NewMatchRule("optional"),
				NewMatchRule("override"),
				NewMatchRule("postfix"),
				NewMatchRule("precedence"),
				NewMatchRule("prefix"),
				NewMatchRule("Protocol"),
				NewMatchRule("required"),
				NewMatchRule("right"),
				NewMatchRule("set"),
				NewMatchRule("some"),
				NewMatchRule("Type"),
				NewMatchRule("unowned"),
				NewMatchRule("weak"),
				NewMatchRule("willSet"),
				NewMatchRegexRule("^Int[0-9]+$"),
				NewMatchRegexRule("^Uint[0-9]+$"),
				NewMatchRegexRule("^Float[0-9]+$"),
				NewMatchRule("Float"),
				NewMatchRule("Double"),
				NewMatchRule("Bool"),
				NewMatchRule("Void"),
				NewMatchRule("String"),
			},
		},
	}
}
