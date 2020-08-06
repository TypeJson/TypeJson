package tj

import (
	"github.com/hoisie/mustache"
	ge "github.com/og/x/error"
	glist "github.com/og/x/list"
	"regexp"
)

type StringSpec struct {
	Name string
	AllowEmpty bool
	MinRuneLen int
	MinRuneLenMessage string
	MaxRuneLen int
	MaxRuneLenMessage string
	Pattern []string
	BanPattern []string
	PatternMessage string
	Enum []string
}
type stringSpecRender struct {
	Value string
	StringSpec
}
func (spec StringSpec) render (message string, value string) string {
	context := stringSpecRender{
		Value: value,
		StringSpec: spec,
	}
	return mustache.Render(message, context)
}
func (r *Rule) String(v string, spec StringSpec) {
	if r.Fail { return }
	if v == "" && !spec.AllowEmpty {
		r.Break(r.Format.StringNotAllowEmpty(spec.Name))
		return
	}
	if spec.CheckMinRuneLen(v, r) { return }
	if spec.CheckMaxRuneLen(v, r) { return }
	if spec.CheckPattern   (v, r) { return }
	if spec.CheckBanPattern(v, r) { return }
	if spec.CheckEnum(v, r) {return}
}

func (spec StringSpec) CheckMaxRuneLen(v string, r *Rule) (fail bool) {
	if spec.MaxRuneLen == 0 {
		return false
	}
	length := len([]rune(v))
	pass := length <= spec.MaxRuneLen
	if !pass {
		message := r.CreateMessage(spec.MaxRuneLenMessage, func() string {
			return r.Format.StringMaxRuneLen(spec.Name, v, spec.MaxRuneLen)
		})
		r.Break(spec.render(message, v))
	}
	return r.Fail
}

func (spec StringSpec) CheckMinRuneLen(v string, r *Rule) (fail bool) {
	length := len([]rune(v))
	pass := length >= spec.MinRuneLen
	if !pass {
		message := r.CreateMessage(spec.MinRuneLenMessage, func() string {
			return r.Format.StringMinRuneLen(spec.Name, v, spec.MinRuneLen)
		})
		r.Break(spec.render(message, v))
	}
	return r.Fail
}
func (spec StringSpec) CheckPattern(v string, r *Rule) (fail bool) {
	if len(spec.Pattern) == 0 {
		return false
	}
	for _, pattern := range spec.Pattern {
		matched, err := regexp.MatchString(pattern, v) ; ge.Check(err)
		pass := matched
		if !pass {
			message := r.CreateMessage(spec.PatternMessage, func() string {
				return r.Format.StringPattern(spec.Name, v, spec.Pattern, pattern)
			})
			r.Break(spec.render(message, v))
			break
		}
	}
	return r.Fail
}
func (spec StringSpec) CheckBanPattern(v string, r *Rule) (fail bool) {
	if len(spec.BanPattern) == 0 {
		return false
	}
	for _, pattern := range spec.BanPattern {
		matched, err := regexp.MatchString(pattern, v) ; ge.Check(err)
		pass := !matched
		if !pass {
			message := r.CreateMessage(spec.PatternMessage, func() string {
				return r.Format.StringBanPattern(spec.Name, v, spec.BanPattern, pattern)
			})
			r.Break(spec.render(message, v))
			break
		}
	}
	return r.Fail
}
func (spec StringSpec) CheckEnum(v string, r *Rule) (fail bool) {
	if len(spec.Enum) == 0 {
		return false
	}
	sList := glist.StringList{Value:spec.Enum}
	pass := sList.In(v)
	if !pass {
		message := r.Format.StringEnum(spec.Name, v, spec.Enum)
		r.Break(spec.render(message, v))
	}
	return r.Fail
}
