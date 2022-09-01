package kmcheck

import (
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

// Validators is a list of all Validator available and registered by RegisterValidator
var Validators []Validator

// RegisterValidator register the given Validator.
// You should use it on `init` function.
func RegisterValidator(v Validator) {
	Validators = append(Validators, v)
}

type Validator interface {
	CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData)
	CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData)
}

// Check checks potential naming conflicts and collisions,
// that function will set warnings into kmparser.Content.
func Check(parsed *kmparser.Content) {
	var (
		mutex sync.Mutex
		group sync.WaitGroup
	)

	for i := range Validators {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			check(&mutex, Validators[i], parsed)
		}(i)
	}

	group.Wait()
}

func check(mutex *sync.Mutex, v Validator, parsed *kmparser.Content) {
	for i := range parsed.Structs {
		v.CheckStruct(mutex, parsed, &parsed.Structs[i].Data)
	}
	for i := range parsed.Enums {
		v.CheckEnum(mutex, parsed, &parsed.Enums[i].Data)
	}
}
