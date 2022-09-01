package kmcheck

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"karmem.org/cmd/karmem/kmparser"
)

type WordRule func(needle string) bool

func NewMatchRule(match string) WordRule {
	return func(needle string) bool {
		return strings.EqualFold(match, needle)
	}
}

func NewMatchRegexRule(reg string) WordRule {
	r := regexp.MustCompile(reg)
	return func(needle string) bool {
		return r.MatchString(strings.ToLower(needle))
	}
}

func NewMatchPrefix(match string) WordRule {
	return func(needle string) bool {
		return strings.HasPrefix(strings.ToLower(needle), strings.ToLower(match))
	}
}

func NewMatchSuffix(match string) WordRule {
	return func(needle string) bool {
		return strings.HasSuffix(strings.ToLower(needle), strings.ToLower(match))
	}
}

type RestrictedWords struct {
	Language kmparser.Language
	Rules    []WordRule
}

func (r *RestrictedWords) newWarning(mutex *sync.Mutex, k *[]kmparser.Warning, msg string) {
	mutex.Lock()
	defer mutex.Unlock()

	*k = append(*k, kmparser.Warning{Data: kmparser.WarningData{
		Message:   msg,
		Rule:      "RestrictedWords",
		Type:      kmparser.RuleTypeReservedName,
		Languages: r.Language,
	}})
}

func (r *RestrictedWords) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	for _, word := range r.Rules {
		if word(target.Name) {
			r.newWarning(mutex, &target.Warnings, fmt.Sprintf("%s is a restricted name", target.Name))
		}
		for i := range target.Fields {
			field := &target.Fields[i]
			if word(field.Data.Name) {
				r.newWarning(mutex, &field.Data.Warnings, fmt.Sprintf("%s is a restricted name", field.Data.Name))
			}
		}
	}
}

func (r *RestrictedWords) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	for _, word := range r.Rules {
		if word(target.Name) {
			r.newWarning(mutex, &target.Warnings, fmt.Sprintf("%s is a restricted name", target.Name))
		}
	}
}

type CollisionParentChildField struct {
	Language kmparser.Language
}

func (r *CollisionParentChildField) newWarning(mutex *sync.Mutex, k *[]kmparser.Warning, msg string) {
	mutex.Lock()
	defer mutex.Unlock()

	*k = append(*k, kmparser.Warning{Data: kmparser.WarningData{
		Message:   msg,
		Rule:      "FieldNameMatchesStructName",
		Type:      kmparser.RuleTypeDuplicateName,
		Languages: r.Language,
	}})
}

func (r *CollisionParentChildField) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	for i := range target.Fields {
		x := &target.Fields[i]
		if target.Name == x.Data.Name {
			r.newWarning(mutex, &x.Data.Warnings, fmt.Sprintf("%s is the same name of parent struct", target.Name))
		}
	}
}

func (r *CollisionParentChildField) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
	for i := range target.Fields {
		x := &target.Fields[i]
		if target.Name == x.Data.Name {
			r.newWarning(mutex, &target.Warnings, fmt.Sprintf("%s is the same name of parent enum", target.Name))
		}
	}
}

type CollisionArraySuffix struct {
	Language kmparser.Language
	Rules    []WordRule
}

func (r *CollisionArraySuffix) newWarning(mutex *sync.Mutex, k *[]kmparser.Warning, msg string) {
	mutex.Lock()
	defer mutex.Unlock()
	*k = append(*k, kmparser.Warning{Data: kmparser.WarningData{
		Message:   msg,
		Rule:      "",
		Type:      kmparser.RuleTypeDuplicateName,
		Languages: kmparser.LanguageC,
	}})
}

func (r *CollisionArraySuffix) CheckStruct(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.StructData) {
	for i := range target.Fields {
		x := &target.Fields[i]
		for _, v := range r.Rules {
			if !v(x.Data.Name) {
				continue
			}
			for j := range target.Fields {
				if i == j {
					continue
				}
				y := &target.Fields[j]
				if strings.HasPrefix(strings.ToLower(x.Data.Name), strings.ToLower(y.Data.Name)) {
					r.newWarning(mutex, &x.Data.Warnings, fmt.Sprintf("%s may collides with %s, the generated code may generate functions with such suffix", x.Data.Name, y.Data.Name))
				}
			}
		}
	}
}

func (r *CollisionArraySuffix) CheckEnum(mutex *sync.Mutex, parsed *kmparser.Content, target *kmparser.EnumData) {
}
