package typejson

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParse(t *testing.T) {
	Convey("default NotRequired", t, func() {
		type Author struct {
			Name       string   `json:"name"`
			Age        int      `json:"age",omitempty`
			IsFriendly bool     `json:"is_firendly"`
			Likes      []string `json:"likes"`
			Children   struct {
				Son      string `json:"son"`
				Daughter string `json:"daughter"`
			} `json:"a"`
			TestEmptyNumberNil bool
			TestEmptyNumber    int
		}
		var nimo Author
		jsonstring := `
		{
			"name":"nimo",
			"likes": ["js", "go"]
		}
		`
		tjson := Create()
		tjson.parse(jsonstring, &nimo, Types{
			"age":               {Default: 27},
			"children.son":      {Default: "fifteen"},
			"children.daughter": {Default: "wood"},
			"testEmptyNumber?":  {},
		})
		So(nimo.Name, ShouldEqual, "nimo")
		So(nimo.Age, ShouldEqual, 27)
		So(nimo.IsFriendly, ShouldEqual, false)
		So(nimo.Likes[0], ShouldEqual, "js")
		So(nimo.Likes[1], ShouldEqual, "go")
		So(nimo.Children.Son, ShouldEqual, "fifteen")
		So(nimo.Children.Daughter, ShouldEqual, "wood")
		So(nimo.TestEmptyNumberNil, ShouldEqual, true)
		So(nimo.TestEmptyNumber, ShouldEqual, 0)
	})
	Convey("Check", t, func() {
		CheckTypes := Types{
			"age": {
				Check: func(data TypeItemCheckData) (message string, pass bool) {
					if data.valueNumber < 18 {
						message = "未成年"
						return
					}
					pass = true
					return
				},
			},
		}
		type Query struct {
			age int
		}
		tjson := Create()
		var wrongQuery Query
		queryInfo, queryPassFail := tjson.parse(` { "age": 10 }`, &wrongQuery, CheckTypes)
		So(queryPassFail, ShouldEqual, true)
		So(queryInfo.Message, ShouldEqual, "未成年")

		var correctQuery Query
		correctQueryInfo, correctQueryPassFail := tjson.parse(` { "age": 20 }`, &correctQuery, CheckTypes)
		So(correctQueryPassFail, ShouldEqual, false)
		So(correctQueryInfo.Message, ShouldEqual, "")
	})
	Convey("requied", t, func() {
		type Query struct {
			Page string `json:"page"`
		}
		var query Query
		queryTypes := Types{
			"page": {
				Label: "页码",
			},
		}
		tjson := Create()
		queryInfo, queryFail := tjson.parse(`{}`, &query, queryTypes)
		So(queryFail, ShouldEqual, true)
		So(queryInfo.Message, ShouldEqual, "页码必填")
	})
}
func TestParseArray(t *testing.T) {
	Convey("array", t, func() {
		type Query struct {
			PersonList []struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
				Vip bool `json:"vip"`
			} `json:"personList"`
		}
		jsonstr := `{
			"personList": [
			{
				"name": "nimo",
				"age": 27
			},
			{
				"name": "nico",
				"age": 18
			}
			]
		}`
		tjson := Create()
		var query Query
		types := Types{
			"personList.*.vip": {
				Default: true,
			},
		}
		queryInfo, queryFail := tjson.parse(jsonstr, &query, types)
		So(queryFail, ShouldEqual, false)
		So(queryInfo.Message, ShouldEqual, "")
		So(len(query.PersonList), ShouldEqual, 2)
		resultJSON, err := json.Marshal(query)
		if err != nil {
			log.Print(err)
		}
		So(string(resultJSON), ShouldEqual, `{"personList":[{"name":"nimo","age":27,"vip":true},{"name":"nico","age":18,"vip":true}]}`)
	})
}
func TestParselabel(t *testing.T) {
	Convey("default NotRequired", t, func() {
		DefaultLabelList := LabelList{
			"name": "名字",
		}
		tjson := Create()
		tjson.setDefaultLabel(DefaultLabelList)
		type People struct {
			Name string `json:"name"`
		}
		var some People
		types := Types{
			"name": {},
		}
		parseInfo, fail := tjson.parse(`{"name":""}`, &some, types)
		So(fail, ShouldEqual, true)
		So(parseInfo.Message, ShouldEqual, "名字必填")
		fmt.Print("@TODO 安全的取空 [] ")
	})
}
