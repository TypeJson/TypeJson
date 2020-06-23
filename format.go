package tj

import (
	gconv "github.com/og/x/conv"
)

type Formatter interface {
	StringRequiredFailMessage(field string) string
	StringMinRuneLen(name string, value string, length int) string
	StringMaxRuneLen(name string, value string, length int) string
	StringPattern(name string, value string, pattern string) string
}
type CNFormat struct {
}
func (CNFormat) StringRequiredFailMessage(field string) string {
	return field  + "必填"
}
func (CNFormat) StringMinRuneLen(name string, value string, length int) string {
	return name + "长度不能小于" + gconv.IntString(length)
}
func (CNFormat) StringMaxRuneLen(name string, value string, length int) string {
	return name + "长度不能大于" + gconv.IntString(length)
}
func (CNFormat) StringPattern(name string, value string, pattern string) string {
	return name + "格式错误"
}