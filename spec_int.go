package tj

import (
	"github.com/hoisie/mustache"
	gconv "github.com/og/x/conv"
)
// option int simulate int? (int or nil)
// tj.Int(18) equal OptionInt(valid: true, int: 18)
type OptionInt struct {
	valid bool
	int int
}
func (o OptionInt) Valid() bool {
	return o.valid
}
func (o OptionInt) String() string {
	if !o.valid {return ""}
	return gconv.IntString(o.int)
}
func (o OptionInt) Unwrap() int {
	if o.valid {return o.int}
	panic("OptionInt: valid is false, can not unwrap")
}
func Int(i int) OptionInt {
	return OptionInt{true, i}
}
type IntSpec struct {
	Name string
	AllowZero bool
	Unsigned bool
	Min OptionInt
	MinMessage string
	Max OptionInt
	MaxMessage string
	Pattern []string
	BanPattern []string
	PatternMessage string
}
type intSpecRender struct {
	Value interface{}
	IntSpec
}
func (spec IntSpec) render (message string, value interface{}) string {
	context := intSpecRender{
		Value: value,
		IntSpec: spec,
	}
	return mustache.Render(message, context)
}
func (r *Rule) Int(v int, spec IntSpec) {
	if r.Fail {return}
	if v == 0 && !spec.AllowZero {
		r.Break(r.Format.IntNotAllowEmpty(spec.Name))
		return
	}
	if spec.CheckMin(v, r) { return }
	if spec.CheckMax(v ,r) { return }
	if spec.CheckPattern(v, r) {return}
	if spec.CheckBanPattern(v, r) {return}
	return
}
func (spec IntSpec) CheckMin(v int, r *Rule) (fail bool) {
	var min int
	if spec.Min.Valid() {
		min = spec.Min.Unwrap()
	} else {
		return
	}
	pass := v >= min
	if !pass {
		message := r.CreateMessage(spec.MinMessage, func() string {
			return r.Format.IntMin(spec.Name, v, min)
		})
		r.Break(spec.render(message, v))
	}
	return
}
func (spec IntSpec) CheckMax(v int, r *Rule) (fail bool) {
	var max int
	if spec.Max.Valid() {
		max = spec.Max.Unwrap()
	} else {
		return
	}
	pass := v <= max
	if !pass {
		message := r.CreateMessage(spec.MaxMessage, func() string {
			return r.Format.IntMax(spec.Name, v, max)
		})
		r.Break(spec.render(message, v))
	}
	return
}
func (spec IntSpec) CheckPattern(v int, r *Rule) (fail bool) {
	return checkPattern(patternData{
		Pattern:        spec.Pattern,
		PatternMessage: spec.PatternMessage,
		Name:           spec.Name,
	}, spec.render, gconv.IntString(v), r)
}

func (spec IntSpec) CheckBanPattern(v int, r *Rule) (fail bool) {
	return checkBanPattern(banPatternData{
		BanPattern:        spec.BanPattern,
		PatternMessage: spec.PatternMessage,
		Name:           spec.Name,
	}, spec.render, gconv.IntString(v), r)
}