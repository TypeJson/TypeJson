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
	Name string 
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
	Title string
	More string 
}
func (s SpecStringPattern) TJ (r *tj.Rule){
	r.String(s.Name, tj.StringSpec{
		Name:              "姓名",
		Path:              "name",
		Pattern:		   []string{"^nimo"},
	})
	r.String(s.Title, tj.StringSpec{
		Name: "标题",
		Path: "title",
		Pattern: []string{`abc$`},
		PatternMessage: "{{Name}}必须以abc为结尾",
	})
	r.String(s.More, tj.StringSpec{
		AllowEmpty: true,
		Name: "更多",
		Pattern:[]string{`^a`, `a$`},
		PatternMessage: "{{Name}}开始结尾必须是a",
	})
}
func TestSpecStringPattern(t *testing.T) {
	as := gtest.NewAS(t)
	c := tj.NewCN()
	{
		as.Equal(c.Scan(SpecStringPattern{
			Name: "nimo",
			Title: "abc",
		}), tj.Report{
			Fail:    true,
			Message: "更多开始结尾必须是a",
		})
	}
	{
		as.Equal(c.Scan(SpecStringPattern{
			Name: "xnimo",
			Title: "abc",
		}), tj.Report{
			Fail:    true,
			Message: "姓名格式错误",
		})
	}
	{
		as.Equal(c.Scan(SpecStringPattern{
			Name: "nimo",
			Title: "abcd",
		}), tj.Report{
			Fail:    true,
			Message: "标题必须以abc为结尾",
		})
	}
	{
		as.Equal(c.Scan(SpecStringPattern{
			Name: "nimo",
			Title: "abcd",
			More: "c",
		}), tj.Report{
			Fail:    true,
			Message: "标题必须以abc为结尾",
		})
	}
}