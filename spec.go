package tj

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