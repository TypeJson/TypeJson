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

type TypeJSON struct {
	defaultLabelList LabelList
}
type ParseInfo struct {
	Message string
}
type LabelList map[string]string
type Types map[string]TypesItem
type TypesItem struct {
	Default interface{}
	Check func(data TypeItemCheckData) (message string, pass bool)
	Label string
}
type TypeItemCheckData struct {
	valueNumber float64
	valueString string
	valueBool bool
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

func Create () (tjson TypeJSON , err error){
	return tjson, nil
}
func (self *TypeJSON) setDefaultLabel( labelList LabelList) {
	self.defaultLabelList = labelList
}
func (self *TypeJSON) parse(jsonstring string, data interface{}, types Types)(info ParseInfo, fail bool){
	json.Unmarshal([]byte(jsonstring), data)
	for key, schema := range types {
		attr := key
		attrList := strings.Split(attr, ".")
		lastAttr := attrList[len(attrList)-1:][0]
		targetResult := gjson.Get(jsonstring, key)
		isUndefinedValue := len(targetResult.Raw) == 0
		attrLastWord := string([]byte(attr)[len(attr)-1])
		required := true
		isEmptyValue := false
		attrHasNotRequiredToken := attrLastWord == "?"
		if attrHasNotRequiredToken {
			required = false
			removeNotRequiredTokenAttr := string([]byte(attr)[:len(attr)-1])
			attr = removeNotRequiredTokenAttr
		}
		shouldSetAttrNil := isUndefinedValue && !required
		shouldSetDefaultValue := !shouldSetAttrNil && schema.Default != nil
		shouldCheckValue := !shouldSetAttrNil && schema.Check != nil
		if  shouldSetAttrNil {
			attr = attr + "Nil"
			setValue(data, attr, true)
		}
		if shouldSetDefaultValue {
			setValue(data, attr, schema.Default)
		}
		if (isUndefinedValue) {
			isEmptyValue = true
		} else {
			switch targetResult.Type {
				case gjson.Number:
				case gjson.String:
					// targetResult.raw = `""`
					if len(targetResult.Raw) <=2 { isEmptyValue = true}
				case gjson.True:
				case gjson.False:
				case gjson.Null:
					isEmptyValue = true
				default:
						fmt.Print(targetResult.Type)
						log.Fatal("@TODO: 需要加上 object array 1")
			}
		}
		if required  && isEmptyValue && !shouldSetDefaultValue {
				var currentLabel string
				var currentDefaultLabel = self.defaultLabelList[lastAttr]
				if (len(schema.Label) != 0 ) {
					currentLabel = schema.Label
				} else if len(currentDefaultLabel) != 0 {
					currentLabel = currentDefaultLabel
				}
				fail = true
				info.Message = currentLabel + "必填"
				break
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
					log.Fatal("typejson: " + schema.Label + ":" + attr + " is nil!")
				default:
					fmt.Print(targetResult.Type)
					log.Fatal("@TODO: 需要加上 object array 2")
			}
			message, pass := schema.Check(data)
			if !pass {
				fail = true
				info.Message = message
				break
			}
		}
	}
	return
}
