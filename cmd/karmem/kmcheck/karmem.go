package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

func init() { RegisterValidator(NewKarmem()) }

type Karmem struct {
	RestrictedWords *RestrictedWords
}

func (v *Karmem) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	v.RestrictedWords.CheckStruct(mutex, parsed, target)
}

func (v *Karmem) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {}

func NewKarmem() *Karmem {
	return &Karmem{
		RestrictedWords: &RestrictedWords{
			Rules: []WordRule{
				NewMatchRule("KarmemPointer"),
				NewMatchRule("KarmemSize"),
			},
		},
	}
}
