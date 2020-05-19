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
	Pattern string
	PatternMessage string
	BanPattern string
	Enum []string
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
func (r *Rule) String(v string, spec StringSpec) {
	if spec.CheckMinRuneLen(v, r) {
		return }
}
type IntSpec struct {
	Name string
	Path string
	Min int
	Max int
}
func (r *Rule)Int(v int, spec IntSpec) {

}
type ArraySpec struct {
	Name string
	Path string
	MinLen int
	MaxLen int
}
func (r *Rule)Array(v interface{}, spec ArraySpec){

}
type BoolSpec struct {
	Name string
	Path string
	Equal bool
}
func (r *Rule) Bool(v bool, spec BoolSpec) {

}