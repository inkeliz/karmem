package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewAssemblyScript()) }

type AssemblyScript struct {
	RestrictedWords *RestrictedWords
}

func (v *AssemblyScript) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
}

func (v *AssemblyScript) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	v.RestrictedWords.CheckEnum(mutex, parsed, target)
}

func NewAssemblyScript() *AssemblyScript {
	return &AssemblyScript{
		RestrictedWords: &RestrictedWords{
			Language: kmparser.LanguageAssemblyScript,
			Rules: []WordRule{
				NewMatchRule("await"),
				NewMatchRule("break"),
				NewMatchRule("case"),
				NewMatchRule("catch"),
				NewMatchRule("class"),
				NewMatchRule("const"),
				NewMatchRule("continue"),
				NewMatchRule("debugger"),
				NewMatchRule("default"),
				NewMatchRule("delete"),
				NewMatchRule("do"),
				NewMatchRule("else"),
				NewMatchRule("enum"),
				NewMatchRule("export"),
				NewMatchRule("extends"),
				NewMatchRule("false"),
				NewMatchRule("finally"),
				NewMatchRule("for"),
				NewMatchRule("function"),
				NewMatchRule("if"),
				NewMatchRule("import"),
				NewMatchRule("in"),
				NewMatchRule("of"),
				NewMatchRule("instanceof"),
				NewMatchRule("new"),
				NewMatchRule("null"),
				NewMatchRule("return"),
				NewMatchRule("super"),
				NewMatchRule("switch"),
				NewMatchRule("this"),
				NewMatchRule("throw"),
				NewMatchRule("true"),
				NewMatchRule("try"),
				NewMatchRule("typeof"),
				NewMatchRule("type"),
				NewMatchRule("var"),
				NewMatchRule("void"),
				NewMatchRule("while"),
				NewMatchRule("with"),
				NewMatchRule("yield"),
				NewMatchRule("let"),
				NewMatchRule("static"),
				NewMatchRule("as"),
				NewMatchRule("any"),
				NewMatchRule("set"),
				NewMatchRule("from"),
				NewMatchRule("constructor"),
				NewMatchRule("module"),
				NewMatchRule("require"),
				NewMatchRule("implements"),
				NewMatchRule("interface"),
				NewMatchRule("package"),
				NewMatchRule("private"),
				NewMatchRule("protected"),
				NewMatchRule("and"),
				NewMatchRule("public"),
				NewMatchRule("i8"),
				NewMatchRule("i16"),
				NewMatchRule("i32"),
				NewMatchRule("i64"),
				NewMatchRule("u8"),
				NewMatchRule("u16"),
				NewMatchRule("u32"),
				NewMatchRule("u64"),
				NewMatchRule("f32"),
				NewMatchRule("f64"),
				NewMatchRule("bool"),
				NewMatchRule("boolean"),
				NewMatchRule("isize"),
				NewMatchRule("usize"),
				NewMatchRule("v128"),
				NewMatchRule("externref"),
				NewMatchRule("funcref"),
				NewMatchRule("string"),
				NewMatchRule("number"),
				NewMatchRule("symbol"),
				NewMatchRule("undefined"),
			},
		},
	}
}
