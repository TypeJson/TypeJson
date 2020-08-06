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
		Min: Int(18),
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
		Min: Int(18),
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


type IntMax struct {
	Age int
}
func (v IntMax) TJ(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Max: Int(18),
	})
}
func TestIntMax(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(IntMax{Age:17}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntMax{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntMax{Age:19}), Report{
		Fail:    true,
		Message: "年龄不能大于18",
	})
}

type IntMaxMessage struct {
	Age int
}
func (v IntMaxMessage) TJ(r *Rule) {
	r.Int(v.Age, IntSpec{
		Name: "年龄",
		Max: Int(18),
		MaxMessage:"年龄不可以大于{{Max}}",
	})
}
func TestIntMaxMessage(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(IntMaxMessage{Age:17}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntMaxMessage{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntMaxMessage{Age:19}), Report{
		Fail:    true,
		Message: "年龄不可以大于18",
	})
}
type IntPattern struct {
	Number int
}
func (v IntPattern) TJ (r *Rule) {
	r.Int(v.Number, IntSpec{
		Name: "号码",
		Pattern: []string{`^138`},
		PatternMessage: "{{Name}}必须以138开头",
	})
}
func TestIntPattern(t *testing.T) {
	as := gtest.NewAS(t)
	_=as
	checker := NewCN()
	as.Equal(checker.Scan(IntPattern{Number: 11384}), Report{
		Fail:    true,
		Message: "号码必须以138开头",
	})
	as.Equal(checker.Scan(IntPattern{Number: 138}), Report{
		Fail:    false,
		Message: "",
	})
}

type IntBanPattern struct {
	Number int
}
func (v IntBanPattern) TJ (r *Rule) {
	r.Int(v.Number, IntSpec{
		Name: "号码",
		BanPattern: []string{`^138`, `^178`},
		PatternMessage: "{{Name}}不允许以138和178开头",
	})
}
func TestIntBanPattern(t *testing.T) {
	as := gtest.NewAS(t)
	_=as
	checker := NewCN()
	as.Equal(checker.Scan(IntBanPattern{Number: 11384}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(IntBanPattern{Number: 138}), Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
	as.Equal(checker.Scan(IntBanPattern{Number: 178}), Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
}