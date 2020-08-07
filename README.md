# typejson

> 类型安全的结构体校验器，远离不类型安全的 struct tag 验证。

```go
package main
import (
	tj "github.com/typejson/go"
)
type RequestCreateUser struct {
	Name string
	Nickname string
	Age int
	Skills []string
	Address RequestCreateUserAddress
}
func (v RequestCreateUser) TJ(r *tj.Rule) {
	r.String(v.Name, tj.StringSpec{
		Name:              "姓名",
		MinRuneLen:        2,
		MaxRuneLen:        10,
	})
	r.String(v.Nickname, tj.StringSpec{
		Name:              "昵称",
		AllowEmpty:        true, // 昵称非必填
		MinRuneLen:        2,
		MaxRuneLen:        10,
	})
	r.Int(v.Age, tj.IntSpec{
		Name:           "年龄",
		Min:            tj.Int(18),
		MinMessage:     "只允许成年人注册",
	})
	r.Array(len(v.Skills), tj.ArraySpec{
		Name:          "技能",
		MaxLen:        tj.Int(10),
		MaxLenMessage: "最多填写{{MaxLen}}项",
	})
	for _, skill := range v.Skills {
		r.String(skill, tj.StringSpec{
			Name:              "技能项",
		})
	}
	// Address由 RequestCreateUserAddress{}.TJ() 实现
}
type RequestCreateUserAddress struct {
	Province string
	Detail string
}
func (v RequestCreateUserAddress) TJ(r *tj.Rule) {
	r.String(v.Province, tj.StringSpec{
		Name:              "省",
	})
	r.String(v.Detail, tj.StringSpec{
		Name: "详细地址",
	})
}

func main() {
	checker := tj.NewCN()
	createUser := RequestCreateUser{
		Name: "张三",
		Nickname: "三儿",
		Age: 20,
		Skills: []string{"clang", "go"},
		Address: RequestCreateUserAddress{
			Province: "上海",
			Detail:   "人民广场一号",
		},
	}
	report := checker.Scan(createUser)
	if report.Fail {
		log.Print(report.Message)
	} else {
		log.Panic("验证通过")
	}
}
```