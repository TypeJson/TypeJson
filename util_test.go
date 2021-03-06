package tj

import (
	gtest "github.com/og/x/test"
	"testing"
)

type TestUtilEnumValues struct {
	Type string
}
func (TestUtilEnumValues) Dict() (dict struct{
	Type struct {
		Normal string
		Danger string
	}
}) {
	dict.Type.Normal = "normal"
	dict.Type.Danger = "danger"
	return
}
func (v TestUtilEnumValues) TJ(r *Rule) {
	r.String(v.Type, StringSpec{
		Name:              "类型",
		Enum:              EnumValues(v.Dict().Type),
	})
}
func Test_UtilEnumValues(t *testing.T) {
	data := TestUtilEnumValues{}
	as := gtest.NewAS(t)
	as.Equal(EnumValues(data.Dict().Type), []string{"normal", "danger"})
}
