package typejson

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"reflect"
	"strings"
	"testing"
	"unicode"

	//. "github.com/smartystreets/goconvey/convey"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func stringFirstWordToUpper(str string) (output string) {
	wordList := []rune{}
	for index, word := range str {
		if index == 0 {
			word = unicode.ToUpper(word)
		}
		wordList = append(wordList, word)
	}
	output = string(wordList)
	return
}
func setValue(target *Author, attr string, value interface{}) {
	attrList := strings.Split(attr, ".")
	for _, targetAttr := range attrList {
		targetAttr = stringFirstWordToUpper(targetAttr)
		targetRef := reflect.ValueOf(target)
		tempPointer := targetRef.Elem().FieldByName(targetAttr)
		tempPointer.Set(reflect.ValueOf(value))
	}
}

type Author struct {
		Name  string   `json:"name"`
		Age   int      `json:"age",omitempty`
		Likes []string `json:"likes"`
		Hot bool `json:"hot"`
	}
func TestValue(t *testing.T) {
	//Convey("validator json firstWord", t, func() {
	//	_, err := Parse("a")
	//	So(err, ShouldBeError)
	//	So(err.Error(), ShouldEqual, `typejson: jsonString first word must be "{" or "[", your error json is:`+"\r\na\r\n")
	//})
	//Convey("validator json lastWord", t, func() {
	//	_, err := Parse("[a")
	//	So(err, ShouldBeError)
	//	So(err.Error(), ShouldEqual, `typejson: jsonString last word must be "}" or "]", your error json is:`+"\r\n[a\r\n")
	//})
	// jsonstring := `{"name":"nimo", "age": 27, "likes": ["js", "go"], "hot": true}`
	jsonstring := `{"name":"nimo", "likes": ["js", "go"]}`
	var nimo Author
	json.Unmarshal([]byte(jsonstring), &nimo)
	type TjsonSchame struct { Default interface{} }
	types := map[string]TjsonSchame{
		"age": { Default: 18 },
	}
	for key, schema := range types {
		attr := key
		isEmptyValue := len(gjson.Get(jsonstring, key).Raw) == 0
		if isEmptyValue && schema.Default != nil{
			setValue(&nimo, attr, schema.Default)
		}
		fmt.Print(nimo)
	}




	//age := gjson.Get(jsonstring, "age")
	//fmt.Print(age.Float())
	// fmt.Print(nimo.Age) // 0
	//Convey("sompile json", t, func() {
	//	value, err := Parse(`{"name": "nimo"}`)
	//	So(err, ShouldBeNil)
	//	fmt.Print(value)
	//})
}
