package tj

import (
	gconv "github.com/og/x/conv"
)

type Formatter interface {
	StringRequiredFailMessage(name string, value string) string
	StringMinRuneLen(name string, value string, length int) string
}
type CNFormat struct {
}
func (CNFormat) StringRequiredFailMessage(name string, value string) string {
	return name  + "必填"
}
func (CNFormat) StringMinRuneLen(name string, value string, length int) string {
	return name + "长度不能小于" + gconv.IntString(length)
}