package tj_test

import (
	gtest "github.com/og/x/test"
	tj "github.com/typejson/go"
	"testing"
)

type SpecStringMinLen struct {
	Name string `tj:"nr"`
}
func (s SpecStringMinLen) TJ(r *tj.Rule) {
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		MinRuneLen:        4,
	})
};
type SpecStringMinLenCustomMessage struct {
	Name string
}
func (s SpecStringMinLenCustomMessage) TJ(r *tj.Rule) {
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		MinRuneLen:        4,
		MinRuneLenMessage: "姓名长度不能小于{{MinRuneLen}}位,你输入的是{{Value}}",
	})
}
func Test_SpecString_MinLen(t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(SpecStringMinLen{Name:"ni"}), tj.Report{
		Fail:    true,
		Message: "姓名长度不能小于4",
	})
	as.Equal(c.Scan(SpecStringMinLen{Name:"nimo"}), tj.Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(c.Scan(SpecStringMinLenCustomMessage{Name:"ni"}), tj.Report{
		Fail:    true,
		Message: "姓名长度不能小于4位,你输入的是ni",
	})
}
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
	Name string `sr:"姓名不能为空"`
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
