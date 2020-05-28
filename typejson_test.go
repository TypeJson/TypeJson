package tj_test

import (
	gtest "github.com/og/x/test"
	tj "github.com/typejson/go"
	"testing"
)

type RequiredOne struct {
	Name string
}
func (r RequiredOne) TJ(*tj.Rule){}
func Test_RequiredOne (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredOne{}), tj.Report{
		Fail:    true,
		Message: "Name必填",
	})
	as.Equal(c.Scan(RequiredOne{Name:"n"}), tj.Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredTwo struct {
	Name string
	Title string
}
func (r RequiredTwo) TJ(*tj.Rule){}
func Test_RequiredTwo (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredTwo{}), tj.Report{
		Fail:    true,
		Message: "Name必填",
	})
	as.Equal(c.Scan(RequiredTwo{Name:"n"}), tj.Report{
		Fail:    true,
		Message: "Title必填",
	})
	as.Equal(c.Scan(RequiredTwo{Name:"n",Title:"1"}), tj.Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredThree struct {
	Name string `tj:"nr"`
	Title string
}
func (r RequiredThree) TJ(*tj.Rule){}
func Test_RequiredThree (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredThree{}), tj.Report{
		Fail:    true,
		Message: "Title必填",
	})
	as.Equal(c.Scan(RequiredThree{Name:"n",Title:"1"}), tj.Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredFour struct {
	Name  string `sr:"姓名不能为空"`
	Title string `sr:"标题不能为空"`
}
func (r RequiredFour) TJ(*tj.Rule){}
func Test_RequiredFour (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredFour{}), tj.Report{
		Fail:    true,
		Message: "姓名不能为空",
	})
	as.Equal(c.Scan(RequiredFour{Name:"n",Title:""}), tj.Report{
		Fail:    true,
		Message: "标题不能为空",
	})
}
