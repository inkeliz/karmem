package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewOdin()) }

type Odin struct {
	RestrictedWords *RestrictedWords
}

func (v *Odin) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
}

func (v *Odin) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	v.RestrictedWords.CheckEnum(mutex, parsed, target)
}

func NewOdin() *Odin {
	return &Odin{
		RestrictedWords: &RestrictedWords{
			Language: kmparser.LanguageOdin,
			Rules:    []WordRule{
				// TODO: Add rules for Odin.
			},
		},
	}
}
