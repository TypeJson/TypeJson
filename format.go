package tj

import (
	gconv "github.com/og/x/conv"
)

type Formatter interface {
	StringRequiredFailMessage(field string) string
	StringMinRuneLen(name string, value string, length int) string
}
type CNFormat struct {
}
func (CNFormat) StringRequiredFailMessage(field string) string {
	return field  + "必填"
}
func (CNFormat) StringMinRuneLen(name string, value string, length int) string {
	return name + "长度不能小于" + gconv.IntString(length)
}