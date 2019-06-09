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
func setValue(target interface{}, attr string, value interface{}) {
	attrList := strings.Split(attr, ".")
	var tempPointer reflect.Value
	for deepLevel, targetAttr := range attrList {
		isFirstAttr := deepLevel == 0
		isLastAttr := deepLevel == len(attrList) - 1
		targetAttr = stringFirstWordToUpper(targetAttr)
		if isFirstAttr {
			tempPointer = reflect.ValueOf(target).Elem()
		}
		tempPointer = tempPointer.FieldByName(targetAttr)
		if isLastAttr {
			tempPointer.Set(reflect.ValueOf(value))
		}
	}
}

type Author struct {
		Name  string   `json:"name"`
		A struct {
			B int `json:"b"`
		} `json:"a"`
		Age   int      `json:"age",omitempty`
		Likes []string `json:"likes"`
		Hot bool `json:"hot"`
	}
func TestValue(t *testing.T) {
	jsonstring := `{"name":"nimo", "likes": ["js", "go"]}`
	var nimo Author
	json.Unmarshal([]byte(jsonstring), &nimo)
	type TjsonSchame struct { Default interface{} }
	types := map[string]TjsonSchame{
		//"age" : { Default: 18 },
		"a.b" : { Default: 10},
	}
	for key, schema := range types {
		attr := key
		isEmptyValue := len(gjson.Get(jsonstring, key).Raw) == 0
		if isEmptyValue && schema.Default != nil {
			setValue(&nimo, attr, schema.Default)
		}
		fmt.Print("\r\n")
		fmt.Print(nimo)
		fmt.Print("\r\n")
	}
}