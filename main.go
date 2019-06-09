package typejson

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
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
type TypesItem struct {
	Default interface{}
	check func(data TypeItemCheckData) (message string, pass bool)
}
type TypeItemCheckData struct {
	valueNumber float64
	valueString string
	valueBool bool
}
type ParseInfo struct {
	Message string
}
func Parse(jsonstring string, data interface{}, types map[string]TypesItem)(info ParseInfo, fail bool){
	json.Unmarshal([]byte(jsonstring), data)
	for key, schema := range types {
		attr := key
		targetResult := gjson.Get(jsonstring, key)
		targetValue := targetResult.Value()
		_ = targetValue
		isEmptyValue := len(gjson.Get(jsonstring, key).Raw) == 0
		attrLastWord := string([]byte(attr)[len(attr)-1])
		isNotRequired := false
		attrHasNotRequiredToken := attrLastWord == "?"
		if attrHasNotRequiredToken {
			isNotRequired = true
			removeNotRequiredTokenAttr := string([]byte(attr)[:len(attr)-1])
			attr = removeNotRequiredTokenAttr
		}
		shouldSetAttrNil := isEmptyValue && isNotRequired
		shouldSetDefaultValue := !shouldSetAttrNil && schema.Default != nil
		shouldCheckValue := !shouldSetAttrNil && schema.check != nil
		if  shouldSetAttrNil {
			attr = attr + "Nil"
			setValue(data, attr, true)
		}
		if shouldSetDefaultValue {
			setValue(data, attr, schema.Default)
		}
		if (shouldCheckValue) {
			var data TypeItemCheckData
			switch targetResult.Type {
				case gjson.Number:
					data.valueNumber = targetResult.Float()
				case gjson.String:
					data.valueString = targetResult.String()
				case gjson.True:
				case gjson.False:
					data.valueBool = targetResult.Bool()
				case gjson.Null:
					log.Fatal("typejson: " + attr + " is nil!")
				default:
					fmt.Print(targetResult.Type)
					log.Fatal("@TODO: 需要加上 object array")
			}
			message, pass := schema.check(data)
			if !pass {
				fail = true
				info.Message = message
				break
			}
		}
	}
	return
}
