package tj

import (
	gconv "github.com/og/x/conv"
	"strconv"
)

// option int simulate int? (int or nil)
// tj.Int(18) equal OptionInt{valid: true, int: 18}
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

// option Float simulate Float? (Float or nil)
// tj.Float(18.1) equal OptionFloat{valid: true, float: 18.1}
type OptionFloat struct {
	valid bool
	float float64
}
func (o OptionFloat) Valid() bool {
	return o.valid
}
func (o OptionFloat) String() string {
	if !o.valid {return ""}
	return strconv.FormatFloat(o.float,'f', -1, 64)
}
func (o OptionFloat) Unwrap() float64 {
	if o.valid {return o.float}
	panic("OptionFloat: valid is false, can not unwrap")
}
func Float(f float64) OptionFloat {
	return OptionFloat{true, f}
}

