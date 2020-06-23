package tj_test

import (
	gtest "github.com/og/x/test"
	tj "github.com/typejson/go"
	"testing"
)

type RequiredOne struct {
	Name string
}
func (v RequiredOne) TJ(r *tj.Rule){
	r.String(v.Name, tj.StringSpec{
		Name: "姓名",
		Path: "name",
	})
}
func Test_RequiredOne (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredOne{}), tj.Report{
		Fail:    true,
		Message: "姓名必填",
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
func (v RequiredTwo) TJ(r *tj.Rule){
	r.String(v.Name, tj.StringSpec{
		Name: "姓名",
		Path: "name",
	})
	r.String(v.Title, tj.StringSpec{
		Name: "标题",
		Path: "title",
	})
}
func Test_RequiredTwo (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredTwo{}), tj.Report{
		Fail:    true,
		Message: "姓名必填",
	})
	as.Equal(c.Scan(RequiredTwo{Name:"n"}), tj.Report{
		Fail:    true,
		Message: "标题必填",
	})
	as.Equal(c.Scan(RequiredTwo{Name:"n",Title:"1"}), tj.Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredThree struct {
	Name string
	Title string
}
func (v RequiredThree) TJ(r *tj.Rule){
	r.String(v.Name, tj.StringSpec{
		Name: "姓名",
		Path: "name",
		AllowEmpty: true,
	})
	r.String(v.Title, tj.StringSpec{
		Name: "标题",
		Path: "title",
	})
}
func Test_RequiredThree (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredThree{}), tj.Report{
		Fail:    true,
		Message: "标题必填",
	})
	as.Equal(c.Scan(RequiredThree{Name:"n",Title:"1"}), tj.Report{
		Fail:    false,
		Message: "",
	})
}
type RequiredFour struct {
	Name  string
	Title string
}
func (v RequiredFour) TJ(r *tj.Rule){
	r.String(v.Name, tj.StringSpec{
		Name: "姓名",
		Path: "name",
	})
	r.String(v.Title, tj.StringSpec{
		Name: "标题",
		Path: "title",
	})
}
func Test_RequiredFour (t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(RequiredFour{}), tj.Report{
		Fail:    true,
		Message: "姓名必填",
	})
	as.Equal(c.Scan(RequiredFour{Name:"n",Title:""}), tj.Report{
		Fail:    true,
		Message: "标题必填",
	})
}
