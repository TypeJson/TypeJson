package tj

import (
	"strings"
)

type Rule struct {
	Fail bool
	Message string
	Format Formatter
}
func (r *Rule) Break(message string) {
	r.Fail = true
	r.Message = message
}
func (r Rule) MessageIsEmpty(message string) (empty bool) {
	message = strings.TrimPrefix(message, " ")
	message = strings.TrimSuffix(message, " ")
	if len(message) == 0 {
		return true
	}
	return false
}
func (r *Rule) Check(pass bool, message string) {
	if !pass {
		r.Fail = true
		r.Message = message
	}
}