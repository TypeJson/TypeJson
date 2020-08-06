package tj

import (
	gconv "github.com/og/x/conv"
	"strings"
)

type Formatter interface {
	StringNotAllowEmpty(name string) string
	StringMinRuneLen(name string, value string, length int) string
	StringMaxRuneLen(name string, value string, length int) string
	StringPattern   (name string, value string, pattern []string, failPattern string) string
	StringBanPattern   (name string, value string, banPattern []string, failBanPattern string) string
	StringEnum (name string, value string, enum []string) string
	IntNotAllowEmpty(name string) string
	IntMin(name string, v int, min int) string
	IntMax(name string, v int, max int) string
}
type CNFormat struct {
}
func (CNFormat) StringNotAllowEmpty(name string) string {
	return name  + "必填"
}
func (CNFormat) StringMinRuneLen(name string, value string, length int) string {
	return name + "长度不能小于" + gconv.IntString(length)
}
func (CNFormat) StringMaxRuneLen(name string, value string, length int) string {
	return name + "长度不能大于" + gconv.IntString(length)
}
func (CNFormat) StringPattern(name string, value string, pattern []string, failPattern string) string {
	return name + "格式错误"
}
func (CNFormat) StringBanPattern(name string, value string, banPattern []string, failBanPattern string) string {
	return name + "格式错误"
}
func (CNFormat) StringEnum(name string, value string, enum []string) string {
	return name + "参数错误，只允许("+ strings.Join(enum, " ") + ")"
}
func (CNFormat) IntNotAllowEmpty(name string) string {
	return name + "不允许为0"
}
func (CNFormat) IntMin(name string, value int, min int) string {
	return name + "不能小于" + gconv.IntString(min)
}
func (CNFormat) IntMax(name string, value int, max int) string {
	return name + "不能大于" + gconv.IntString(max)
}