package tj

import (
	"github.com/hoisie/mustache"
	ge "github.com/og/x/error"
	"regexp"
)

type StringSpec struct {
	Name string
	Path string
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
	if spec.CheckPattern(v, r)    { return }
	if spec.CheckBadPattern(v, r)    { return }
}

func (spec StringSpec) CheckMaxRuneLen(v string, r *Rule) (fail bool) {
	if spec.MaxRuneLen == 0 {
		return false
	}
	length := len([]rune(v))
	pass := length <= spec.MaxRuneLen
	if !pass {
		message := ""
		if r.MessageIsEmpty(spec.MaxRuneLenMessage) {
			message = r.Format.StringMaxRuneLen(spec.Name, v, spec.MaxRuneLen)
		} else {
			message = spec.MaxRuneLenMessage
		}
		r.Break(spec.render(message, v))
	}
	return r.Fail
}

func (spec StringSpec) CheckMinRuneLen(v string, r *Rule) (fail bool) {
	length := len([]rune(v))
	pass := length >= spec.MinRuneLen
	if !pass {
		message := ""
		if r.MessageIsEmpty(spec.MinRuneLenMessage) {
			message = r.Format.StringMinRuneLen(spec.Name, v, spec.MinRuneLen)
		} else {
			message = spec.MinRuneLenMessage
		}
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
			message := ""
			if r.MessageIsEmpty(spec.PatternMessage) {
				message = r.Format.StringPattern(spec.Name, v, spec.Pattern, pattern)
			} else {
				message = spec.PatternMessage
			}
			r.Break(spec.render(message, v))
			break
		}
	}
	return r.Fail
}
func (spec StringSpec) CheckBadPattern(v string, r *Rule) (fail bool) {
	if len(spec.BanPattern) == 0 {
		return false
	}
	for _, pattern := range spec.BanPattern {
		matched, err := regexp.MatchString(pattern, v) ; ge.Check(err)
		pass := !matched
		if !pass {
			message := ""
			if r.MessageIsEmpty(spec.PatternMessage) {
				message = r.Format.StringBadPattern(spec.Name, v, spec.BanPattern, pattern)
			} else {
				message = spec.PatternMessage
			}
			r.Break(spec.render(message, v))
			break
		}
	}
	return r.Fail
}