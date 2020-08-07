package tj

import (
	"github.com/hoisie/mustache"
	"log"
	"reflect"
)

type ArraySpec struct {
	Name string
	Path string
	MinLen OptionInt
	MinLenMessage string
	MaxLen OptionInt
	MaxLenMessage string
}
func (r *Rule) Array(v interface{}, spec ArraySpec){
	rValue := reflect.ValueOf(v)
	if rValue.Kind() != reflect.Slice {
		log.Print("r.Array(v []Type) v must be slice")
		return
	}
	if spec.CheckMinLen(rValue.Len(), r) {return}
	if spec.CheckMaxLen(rValue.Len(), r) {return}
}
type arraySpecRender struct {
	Value interface{}
	ArraySpec
}
func (spec ArraySpec) render (message string, value interface{}) string {
	context := arraySpecRender{
		Value: value,
		ArraySpec: spec,
	}
	return mustache.Render(message, context)
}
func (spec ArraySpec) CheckMinLen(v int, r *Rule) (fail bool) {
	if !spec.MinLen.Valid() {
		return
	}
	min := spec.MinLen.Unwrap()
	pass := v >= min
	if !pass {
		message := r.CreateMessage(spec.MinLenMessage, func() string {
			return r.Format.IntMinLen(spec.Name, v, min)
		})
		r.Break(spec.render(message, v))
	}
	return
}
func (spec ArraySpec) CheckMaxLen(v int, r *Rule) (fail bool) {
	if !spec.MaxLen.Valid() {
		return
	}
	max := spec.MaxLen.Unwrap()
	pass := v <= max
	if !pass {
		message := r.CreateMessage(spec.MaxLenMessage, func() string {
			return r.Format.IntMaxLen(spec.Name, v, max)
		})
		r.Break(spec.render(message, v))
	}
	return
}