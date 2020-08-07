package tj

import (
	gtest "github.com/og/x/test"
	"testing"
)

type ArrayMinLen struct {
	Skills []string
}
func (v ArrayMinLen) TJ(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MinLen: Int(2),
	})
}
func TestArrayMinLen(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(ArrayMinLen{Skills: []string{}}), Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	as.Equal(checker.Scan(ArrayMinLen{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "skills长度不能小于2",
	})
	as.Equal(checker.Scan(ArrayMinLen{Skills: []string{"c","d"}}), Report{
		Fail:    false,
		Message: "",
	})
}

type ArrayMinLenMessage struct {
	Skills []string
}
func (v ArrayMinLenMessage) TJ(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MinLen: Int(2),
		MinLenMessage: "skills 长度必须 < {{MinLen}}",
	})
}
func TestArrayMinLenMessage(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(ArrayMinLenMessage{Skills: []string{}}), Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	as.Equal(checker.Scan(ArrayMinLenMessage{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "skills 长度必须 < 2",
	})
	as.Equal(checker.Scan(ArrayMinLenMessage{Skills: []string{"c","d"}}), Report{
		Fail:    false,
		Message: "",
	})
}


type ArrayMaxLen struct {
	Skills []string
}
func (v ArrayMaxLen) TJ(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MaxLen: Int(2),
	})
}
func TestArrayMaxLen(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(ArrayMaxLen{Skills: []string{}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMaxLen{Skills: []string{"c"}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMaxLen{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMaxLen{Skills: []string{"c","d","e"}}), Report{
		Fail:    true,
		Message: "skills长度不能大于2",
	})
}

type ArrayMaxLenMessage struct {
	Skills []string
}
func (v ArrayMaxLenMessage) TJ(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "skills",
		MaxLen: Int(2),
		MaxLenMessage: "skills 长度必须 > {{MaxLen}}",
	})
}
func TestArrayMaxLenMessage(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(ArrayMaxLenMessage{Skills: []string{}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMaxLenMessage{Skills: []string{"c"}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMaxLenMessage{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMaxLenMessage{Skills: []string{"c","d","e"}}), Report{
		Fail:    true,
		Message: "skills 长度必须 > 2",
	})
}


type ArrayMinMax struct {
	Skills []string
}
func (v ArrayMinMax) TJ(r *Rule) {
	r.Array(len(v.Skills), ArraySpec{
		Name: "技能",
		MinLen: Int(2),
		MaxLen: Int(4),
	})
}
func TestArrayMinMax(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(ArrayMinMax{Skills: []string{}}), Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	as.Equal(checker.Scan(ArrayMinMax{Skills: []string{"c"}}), Report{
		Fail:    true,
		Message: "技能长度不能小于2",
	})
	as.Equal(checker.Scan(ArrayMinMax{Skills: []string{"c", "d"}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMinMax{Skills: []string{"c","d","e"}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMinMax{Skills: []string{"c","d","e","f"}}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(ArrayMinMax{Skills: []string{"c","d","e","f", "g"}}), Report{
		Fail:    true,
		Message: "技能长度不能大于4",
	})
}