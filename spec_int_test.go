package tj

import (
	gtest "github.com/og/x/test"
	"testing"
)

type IntNotAllowZero struct {
	Age int
}
func (v IntNotAllowZero) TJ(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
	})
}
func TestIntNotAllowZero(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(IntNotAllowZero{Age:0}), Report{
		Fail:    true,
		Message: "年龄不允许为0",
	})
	as.Equal(checker.Scan(IntNotAllowZero{Age:1}), Report{
		Fail:    false,
		Message: "",
	})
}

type IntAllowZero struct {
	Age int
}
func (v IntAllowZero) TJ(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		AllowZero: true,
	})
}
func TestIntAllowZero(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(IntAllowZero{Age:0}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntAllowZero{Age:1}), Report{
		Fail:    false,
		Message: "",
	})
}

type IntMin struct {
	Age int
}
func (v IntMin) TJ(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: 18,
	})
}
func TestIntMin(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(IntMin{Age:17}), Report{
		Fail:    true,
		Message: "年龄不能小于18",
	})
	as.Equal(checker.Scan(IntMin{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntMin{Age:19}), Report{
		Fail:    false,
		Message: "",
	})
}

type IntMinMessage struct {
	Age int
}
func (v IntMinMessage) TJ(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Min: 18,
		MinMessage:"年龄不可以小于{{Min}}",
	})
}
func TestIntMinMessage(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(IntMinMessage{Age:17}), Report{
		Fail:    true,
		Message: "年龄不可以小于18",
	})
	as.Equal(checker.Scan(IntMinMessage{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntMinMessage{Age:19}), Report{
		Fail:    false,
		Message: "",
	})
}