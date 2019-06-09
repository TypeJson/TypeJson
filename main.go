package typejson

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"reflect"
	"strings"
	"unicode"
)

type TokenInfo struct {
	Type string
	Value string
}
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
type TypesItem struct { Default interface{} }

func Parse(jsonstring string, data interface{}, types map[string]TypesItem){
	json.Unmarshal([]byte(jsonstring), data)
	for key, schema := range types {
		attr := key
		isEmptyValue := len(gjson.Get(jsonstring, key).Raw) == 0
		if isEmptyValue {
			attrLastWord := string([]byte(attr)[len(attr)-1])
			isNotRequired := false
			if attrLastWord == "?" {
				isNotRequired = true
				attr = string([]byte(attr)[:len(attr)-1])
			}
			if isNotRequired {
				attr = attr + "Nil"
				setValue(data, attr, true)
			} else {
				if schema.Default != nil {
					setValue(data, attr, schema.Default)
				}
			}
		}
	}
}
