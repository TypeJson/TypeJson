package typejson

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParse(t *testing.T) {
	Convey("default NotRequired",t,  func() {
		type Author struct {
			Name  string   `json:"name"`
			Age   int      `json:"age",omitempty`
			IsFriendly bool `json:"is_firendly"`
			Likes []string `json:"likes"`
			Children struct {
				Son string `json:"son"`
				Daughter string `json:"daughter"`
			} `json:"a"`
			TestEmptyNumberNil bool
			TestEmptyNumber int
		}
		var nimo Author
		jsonstring := `
		{
			"name":"nimo",
			"likes": ["js", "go"]
		}
		`
		Parse(jsonstring, &nimo, map[string]TypesItem{
			"age" : { Default: 27},
			"children.son" : { Default: "fifteen"},
			"children.daughter" : { Default: "wood"},
			"testEmptyNumber?" : {},
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
	Convey("check",t,  func() {
		checkTypes := map[string]TypesItem{
			"age": {
				check: func (data TypeItemCheckData) (message string, pass bool) {
					if data.valueNumber < 18 {
						message = "未成年"; return
					}
					pass = true; return
				},
			},
		}
		type Query struct {
			age int
		}
		var wrongQuery Query
		queryInfo, queryPassFail := Parse(` { "age": 10 }`, &wrongQuery, checkTypes)
		So(queryPassFail, ShouldEqual, true)
		So(queryInfo.Message, ShouldEqual, "未成年")

		var correctQuery Query
		correctQueryInfo, correctQueryPassFail := Parse(` { "age": 20 }`, &correctQuery, checkTypes)
		So(correctQueryPassFail, ShouldEqual, false)
		So(correctQueryInfo.Message, ShouldEqual, "")
	})
}