package tj

import (
	"github.com/hoisie/mustache"
)

type StringSpec struct {
	Value interface{}
	Name string
	Path string
	MinRuneLen int
	MinRuneLenMessage string
	MaxRuneLen int
	MaxRuneLenMessage string
	Pattern string
	PatternMessage string
	BanPattern string
	Enum []string
}

func (r *Rule) String(v string, spec StringSpec) {
	if spec.CheckMinRuneLen(v, r) { return }
	if spec.CheckMaxRuneLen(v, r) { return }
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
		spec.Value = v
		r.Break(mustache.Render(message, spec))
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
		spec.Value = v
		r.Break(mustache.Render(message, spec))
	}
	return r.Fail
}