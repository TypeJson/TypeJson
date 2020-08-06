package tj

import (
	"github.com/hoisie/mustache"
)

type IntSpec struct {
	Name string
	AllowZero bool
	Min int
	MinMessage string
	Max int
	MaxMessage string
}
type intSpecRender struct {
	Value int
	IntSpec
}
func (spec IntSpec) render (message string, value int) string {
	context := intSpecRender{
		Value: value,
		IntSpec: spec,
	}
	return mustache.Render(message, context)
}
func (spec IntSpec) CheckMin(v int, r *Rule) (fail bool) {
	if spec.Min == 0 {return}
	pass := v >= spec.Min
	if !pass {
		message := r.CreateMessage(spec.MinMessage, func() string {
			return r.Format.IntMin(spec.Name, v, spec.Min)
		})
		r.Break(spec.render(message, v))
	}
	return
}
func (spec IntSpec) CheckMax(v int, r *Rule) (fail bool) {
	if spec.Max == 0 {return}
	pass := v <= spec.Max
	if !pass {
		message := r.CreateMessage(spec.MaxMessage, func() string {
			return r.Format.IntMax(spec.Name, v, spec.Max)
		})
		r.Break(spec.render(message, v))
	}
	return
}
func (r *Rule)Int(v int, spec IntSpec) {
	if r.Fail {return}
	if v == 0 && !spec.AllowZero {
		r.Break(r.Format.IntNotAllowEmpty(spec.Name))
		return
	}
	if spec.CheckMin(v, r) { return }
	if spec.CheckMax(v ,r) { return }
	return
}