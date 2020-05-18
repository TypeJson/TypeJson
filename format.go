package tj

import (
	gconv "github.com/og/x/conv"
	"reflect"
)

type Formatter interface {
	StringRequiredFailMessage(name string, rValue reflect.Value, fieldType reflect.StructField) string
	StringMinRuneLen(name string, length int) string
}
type CNFormat struct {
}
func (CNFormat) StringRequiredFailMessage(name string, rValue reflect.Value, fieldType reflect.StructField) string {
	return name  + "必填"
}
func (CNFormat) StringMinRuneLen(name string, length int) string {
	return name + "长度不能小于" + gconv.IntString(length)
}