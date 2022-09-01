package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewZig()) }

type Zig struct {
	RestrictedWords *RestrictedWords
}

func (v *Zig) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
}

func (v *Zig) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	v.RestrictedWords.CheckEnum(mutex, parsed, target)
}

func NewZig() *Zig {
	return &Zig{
		RestrictedWords: &RestrictedWords{
			Language: kmparser.LanguageZig,
			Rules: []WordRule{
				NewMatchRule("align"),
				NewMatchRule("allowzero"),
				NewMatchRule("and"),
				NewMatchRule("anyframe"),
				NewMatchRule("anytype"),
				NewMatchRule("asm"),
				NewMatchRule("async"),
				NewMatchRule("await"),
				NewMatchRule("break"),
				NewMatchRule("catch"),
				NewMatchRule("comptime"),
				NewMatchRule("const"),
				NewMatchRule("continue"),
				NewMatchRule("defer"),
				NewMatchRule("else"),
				NewMatchRule("enum"),
				NewMatchRule("errdefer"),
				NewMatchRule("error"),
				NewMatchRule("export"),
				NewMatchRule("extern"),
				NewMatchRule("false"),
				NewMatchRule("fn"),
				NewMatchRule("for"),
				NewMatchRule("if"),
				NewMatchRule("inline"),
				NewMatchRule("noalias"),
				NewMatchRule("nosuspend"),
				NewMatchRule("null"),
				NewMatchRule("or"),
				NewMatchRule("orelse"),
				NewMatchRule("packed"),
				NewMatchRule("pub"),
				NewMatchRule("resume"),
				NewMatchRule("return"),
				NewMatchRule("linksection"),
				NewMatchRule("struct"),
				NewMatchRule("suspend"),
				NewMatchRule("switch"),
				NewMatchRule("test"),
				NewMatchRule("threadlocal"),
				NewMatchRule("true"),
				NewMatchRule("try"),
				NewMatchRule("undefined"),
				NewMatchRule("union"),
				NewMatchRule("unreachable"),
				NewMatchRule("usingnamespace"),
				NewMatchRule("var"),
				NewMatchRule("volatile"),
				NewMatchRule("while"),
				NewMatchRule("i8"),
				NewMatchRule("u8"),
				NewMatchRule("i16"),
				NewMatchRule("u16"),
				NewMatchRule("i32"),
				NewMatchRule("u32"),
				NewMatchRule("i64"),
				NewMatchRule("u64"),
				NewMatchRule("i128"),
				NewMatchRule("u128"),
				NewMatchRule("isize"),
				NewMatchRule("usize"),
				NewMatchRule("c_short"),
				NewMatchRule("c_ushort"),
				NewMatchRule("c_int"),
				NewMatchRule("c_uint"),
				NewMatchRule("c_long"),
				NewMatchRule("c_ulong"),
				NewMatchRule("c_longlong"),
				NewMatchRule("c_ulonglong"),
				NewMatchRule("c_longdouble"),
				NewMatchRule("f16"),
				NewMatchRule("f32"),
				NewMatchRule("f64"),
				NewMatchRule("f128"),
				NewMatchRule("bool"),
				NewMatchRule("anyopaque"),
				NewMatchRule("void"),
				NewMatchRule("noreturn"),
				NewMatchRule("type"),
				NewMatchRule("anyerror"),
				NewMatchRule("comptime_int"),
				NewMatchRule("comptime_float"),
				NewMatchRegexRule("^[i|u][0-9]+$"),
			},
		},
	}
}
