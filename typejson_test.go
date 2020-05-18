package tj_test

import (
	gtest "github.com/og/x/test"
	tj "github.com/typejson/go"
	"testing"
)

type SpecStringMinLen struct {
	Name string
}
func (s SpecStringMinLen) TJ(r *tj.Rule) {
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		MinRuneLen:        4,
	})
}
type SpecStringMinLenCustomMessage struct {
	Name string
}
func (s SpecStringMinLenCustomMessage) TJ(r *tj.Rule) {
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		MinRuneLen:        4,
		MinRuneLenMessage: "姓名长度不能小于四位",
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
		Message: "姓名长度不能小于四位",
	})
}