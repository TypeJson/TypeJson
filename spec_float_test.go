package tj

import (
	gtest "github.com/og/x/test"
	"testing"
)


type FloatMin struct {
	Age float64
}
func (v FloatMin) TJ(r *Rule) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Min: Float(18.2),
	})
}
func TestFloatMin(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(FloatMin{Age:17}), Report{
		Fail:    true,
		Message: "年龄不能小于18.2",
	})
	as.Equal(checker.Scan(FloatMin{Age:18.1}), Report{
		Fail:    true,
		Message: "年龄不能小于18.2",
	})
	as.Equal(checker.Scan(FloatMin{Age:18.2}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMin{Age:18.3}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMin{Age:19}), Report{
		Fail:    false,
		Message: "",
	})
}

type FloatMinMessage struct {
	Age float64
}

func (v FloatMinMessage) TJ(r *Rule) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Min: Float(18.2),
		MinMessage:"年龄不可以小于{{Min}}",
	})
}
func TestFloatMinMessage(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(FloatMinMessage{Age:17}), Report{
		Fail:    true,
		Message: "年龄不可以小于18.2",
	})
	as.Equal(checker.Scan(FloatMinMessage{Age:18}), Report{
		Fail:    true,
		Message: "年龄不可以小于18.2",
	})
	as.Equal(checker.Scan(FloatMinMessage{Age:19}), Report{
		Fail:    false,
		Message: "",
	})
}


type FloatMax struct {
	Age float64
}
func (v FloatMax) TJ(r *Rule) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Max: Float(18.2),
	})
}
func TestFloatMax(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(FloatMax{Age:17}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMax{Age:18.2}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMax{Age:18.3}), Report{
		Fail:    true,
		Message: "年龄不能大于18.2",
	})
	as.Equal(checker.Scan(FloatMax{Age:18.4}), Report{
		Fail:    true,
		Message: "年龄不能大于18.2",
	})
	as.Equal(checker.Scan(FloatMax{Age:19}), Report{
		Fail:    true,
		Message: "年龄不能大于18.2",
	})
}

type FloatMaxMessage struct {
	Age float64
}
func (v FloatMaxMessage) TJ(r *Rule) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Max: Float(18),
		MaxMessage:"年龄不可以大于{{Max}}",
	})
}
func TestFloatMaxMessage(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(FloatMaxMessage{Age:17}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMaxMessage{Age:18}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMaxMessage{Age:19}), Report{
		Fail:    true,
		Message: "年龄不可以大于18",
	})
}
type FloatMinMax struct {
	Age float64
}
func (v FloatMinMax) TJ (r *Rule) {
	r.Float(v.Age, FloatSpec{
		Name: "年龄",
		Min: Float(2),
		Max: Float(4),
	})
}
func TestFloatMinMax(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(FloatMinMax{Age: 0}), Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	as.Equal(checker.Scan(FloatMinMax{Age: 1}), Report{
		Fail:    true,
		Message: "年龄不能小于2",
	})
	as.Equal(checker.Scan(FloatMinMax{Age: 2}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMinMax{Age: 3}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMinMax{Age: 4}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatMinMax{Age: 5}), Report{
		Fail:    true,
		Message: "年龄不能大于4",
	})

}
type FloatPattern struct {
	Number float64
}
func (v FloatPattern) TJ (r *Rule) {
	r.Float(v.Number, FloatSpec{
		Name: "号码",
		Pattern: []string{`^138`},
		PatternMessage: "{{Name}}必须以138开头",
	})
}
func TestFloatPattern(t *testing.T) {
	as := gtest.NewAS(t)
	checker := NewCN()
	as.Equal(checker.Scan(FloatPattern{Number: 11384}), Report{
		Fail:    true,
		Message: "号码必须以138开头",
	})
	as.Equal(checker.Scan(FloatPattern{Number: 138}), Report{
		Fail:    false,
		Message: "",
	})
}

type FloatBanPattern struct {
	Number float64
}
func (v FloatBanPattern) TJ (r *Rule) {
	r.Float(v.Number, FloatSpec{
		Name: "号码",
		BanPattern: []string{`^138`, `^178`},
		PatternMessage: "{{Name}}不允许以138和178开头",
	})
}
func TestFloatBanPattern(t *testing.T) {
	as := gtest.NewAS(t)
	_=as
	checker := NewCN()
	as.Equal(checker.Scan(FloatBanPattern{Number: 11384}), Report{
		Fail:    false,
		Message: "",
	})
	as.Equal(checker.Scan(FloatBanPattern{Number: 138}), Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
	as.Equal(checker.Scan(FloatBanPattern{Number: 178}), Report{
		Fail:    true,
		Message: "号码不允许以138和178开头",
	})
}