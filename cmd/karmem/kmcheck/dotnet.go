package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewDotNet()) }

type DotNet struct {
	RestrictedWords           *RestrictedWords
	CollisionParentChildField *CollisionParentChildField
}

func (v *DotNet) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
	v.CollisionParentChildField.CheckStruct(mutex, parsed, target)
}

func (v *DotNet) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	v.RestrictedWords.CheckEnum(mutex, parsed, target)
	v.CollisionParentChildField.CheckEnum(mutex, parsed, target)
}

func NewDotNet() *DotNet {
	return &DotNet{
		RestrictedWords: &RestrictedWords{
			Language: kmparser.LanguageDotnet,
			Rules: []WordRule{
				NewMatchRule("abstract"),
				NewMatchRule("as"),
				NewMatchRule("base"),
				NewMatchRule("bool"),
				NewMatchRule("break"),
				NewMatchRule("byte"),
				NewMatchRule("case"),
				NewMatchRule("catch"),
				NewMatchRule("char"),
				NewMatchRule("checked"),
				NewMatchRule("class"),
				NewMatchRule("const"),
				NewMatchRule("continue"),
				NewMatchRule("decimal"),
				NewMatchRule("default"),
				NewMatchRule("delegate"),
				NewMatchRule("do"),
				NewMatchRule("double"),
				NewMatchRule("else"),
				NewMatchRule("enum"),
				NewMatchRule("event"),
				NewMatchRule("explicit"),
				NewMatchRule("extern"),
				NewMatchRule("false"),
				NewMatchRule("finally"),
				NewMatchRule("fixed"),
				NewMatchRule("float"),
				NewMatchRule("for"),
				NewMatchRule("foreach"),
				NewMatchRule("goto"),
				NewMatchRule("if"),
				NewMatchRule("implicit"),
				NewMatchRule("in"),
				NewMatchRule("int"),
				NewMatchRule("interface"),
				NewMatchRule("internal"),
				NewMatchRule("is"),
				NewMatchRule("lock"),
				NewMatchRule("long"),
				NewMatchRule("namespace"),
				NewMatchRule("new"),
				NewMatchRule("null"),
				NewMatchRule("object"),
				NewMatchRule("operator"),
				NewMatchRule("out"),
				NewMatchRule("override"),
				NewMatchRule("params"),
				NewMatchRule("private"),
				NewMatchRule("protected"),
				NewMatchRule("public"),
				NewMatchRule("readonly"),
				NewMatchRule("ref"),
				NewMatchRule("return"),
				NewMatchRule("sbyte"),
				NewMatchRule("sealed"),
				NewMatchRule("short"),
				NewMatchRule("sizeof"),
				NewMatchRule("stackalloc"),
				NewMatchRule("static"),
				NewMatchRule("string"),
				NewMatchRule("struct"),
				NewMatchRule("switch"),
				NewMatchRule("this"),
				NewMatchRule("throw"),
				NewMatchRule("true"),
				NewMatchRule("try"),
				NewMatchRule("typeof"),
				NewMatchRule("uint"),
				NewMatchRule("ulong"),
				NewMatchRule("unchecked"),
				NewMatchRule("unsafe"),
				NewMatchRule("ushort"),
				NewMatchRule("using"),
				NewMatchRule("virtual"),
				NewMatchRule("void"),
				NewMatchRule("volatile"),
				NewMatchRule("while"),
				NewMatchRule("add"),
				NewMatchRule("and"),
				NewMatchRule("alias"),
				NewMatchRule("ascending"),
				NewMatchRule("args"),
				NewMatchRule("async"),
				NewMatchRule("await"),
				NewMatchRule("by"),
				NewMatchRule("descending"),
				NewMatchRule("dynamic"),
				NewMatchRule("equals"),
				NewMatchRule("from"),
				NewMatchRule("get"),
				NewMatchRule("global"),
				NewMatchRule("group"),
				NewMatchRule("init"),
				NewMatchRule("into"),
				NewMatchRule("join"),
				NewMatchRule("let"),
				NewMatchRule("managed"),
				NewMatchRule("nameof"),
				NewMatchRule("nint"),
				NewMatchRule("not"),
				NewMatchRule("notnull"),
				NewMatchRule("nuint"),
				NewMatchRule("on"),
				NewMatchRule("or"),
				NewMatchRule("orderby"),
				NewMatchRule("partial"),
				NewMatchRule("record"),
				NewMatchRule("remove"),
				NewMatchRule("required"),
				NewMatchRule("select"),
				NewMatchRule("set"),
				NewMatchRule("unmanage"),
				NewMatchRule("value"),
				NewMatchRule("var"),
				NewMatchRule("when "),
				NewMatchRule("where"),
				NewMatchRule("with"),
				NewMatchRule("yield"),
			},
		},
		CollisionParentChildField: &CollisionParentChildField{
			Language: kmparser.LanguageDotnet,
		},
	}
}
