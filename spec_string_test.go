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
	as.Equal(c.Scan(SpecStringMinLen{Name:"nim"}), tj.Report{
		Fail:    true,
		Message: "姓名长度不能小于4",
	})
	as.Equal(c.Scan(SpecStringMinLen{Name:"nimo"}), tj.Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(c.Scan(SpecStringMinLen{Name:"nimoc"}), tj.Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(c.Scan(SpecStringMinLenCustomMessage{Name:"ni"}), tj.Report{
		Fail:    true,
		Message: "姓名长度不能小于4位,你输入的是ni",
	})
}

type SpecStringMaxLen struct {
	Name string `tj:"nr"`
}
func (s SpecStringMaxLen) TJ(r *tj.Rule) {
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		MaxRuneLen:        4,
	})
};
type SpecStringMaxLenCustomMessage struct {
	Name string
}
func (s SpecStringMaxLenCustomMessage) TJ(r *tj.Rule) {
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		MaxRuneLen:        4,
		MaxRuneLenMessage: "姓名长度不能大于{{MaxRuneLen}}位,你输入的是{{Value}}",
	})
}
func Test_SpecString_MaxLen(t *testing.T) {
	c := tj.NewCN()
	as := gtest.NewAS(t)
	as.Equal(c.Scan(SpecStringMaxLen{Name:"nimoc"}), tj.Report{
		Fail:    true,
		Message: "姓名长度不能大于4",
	})
	as.Equal(c.Scan(SpecStringMaxLen{Name:"nimo"}), tj.Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(c.Scan(SpecStringMaxLen{Name:"nim"}), tj.Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(c.Scan(SpecStringMaxLenCustomMessage{Name:"nimoc"}), tj.Report{
		Fail:    true,
		Message: "姓名长度不能大于4位,你输入的是nimoc",
	})
}
type SpecStringPattern struct {
	Name string
}
func (s SpecStringPattern) TJ (r *tj.Rule){
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		Pattern:		   "^nimo",
		PatternMessage:    "",
	})
}
func TestSpecStringPattern(t *testing.T) {
	as := gtest.NewAS(t)
	c := tj.NewCN()
	{
		as.Equal(c.Scan(SpecStringPattern{
			Name: "nimo",
		}), tj.Report{
			Fail:    false,
			Message: "",
		})
	}
	{
		as.Equal(c.Scan(SpecStringPattern{
			Name: "xnimo",
		}), tj.Report{
			Fail:    true,
			Message: "姓名格式错误",
		})
	}
}