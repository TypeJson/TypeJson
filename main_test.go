package typejson

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValue(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
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
		validDemoA int
		validDemoB int
	}
	var nimo Author
	jsonstring := `
	{
		"name":"nimo",
		"likes": ["js", "go"],
		"validDemoA": 10,
		"validDemoB": 20
	}
	`
	Parse(jsonstring, &nimo, map[string]TypesItem{
		"age" : { Default: 27},
		"children.son" : { Default: "fifteen"},
		"children.daughter" : { Default: "wood"},
		"testEmptyNumber?" : {},
		"validDemoA": {
			check: func (value int) (failMessage string, pass bool) {
				if (value < 18) {
					return "未成年", false
				}
				return "", true
			},
		},
	})
	Convey("Name: ", func() {
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
	})
}