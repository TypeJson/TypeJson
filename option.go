package tj

import (
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
