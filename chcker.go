package tj

import (
	"fmt"
	"log"
	"reflect"
)
type Checker struct {
	Format Formatter
}

type Data interface {
	TJ(r *Rule)
}
type Report struct {
	Fail bool
	Message string
}
func (checker Checker) Scan(data Data) (report Report) {
	rValue := reflect.ValueOf(data)
	rType := rValue.Type()
	return checker.reflectScan(rValue, rType)
}
func (checker Checker) reflectScan(rValue reflect.Value, rType reflect.Type) (report Report) {
	checkMethod := rValue.MethodByName("TJ")
	if checkMethod.IsValid() {
		rule := Rule{
			Format: checker.Format,
		}
		var rValueList []reflect.Value
		rValueList = append(rValueList, reflect.ValueOf(&rule))
		checkMethod.Call(rValueList)
		if rule.Fail {
			report.Fail = true
			report.Message = rule.Message
			return
		}
	}
	for i:=0; i<rType.NumField();i++ {
		rValueItem := rValue.Field(i)
		structField := rType.Field(i)
		switch structField.Type.Kind() {
		case reflect.String:
			tjValue,hasTJ := structField.Tag.Lookup("tj")
			if hasTJ && tjValue == "" {
				log.Print("typejson/go: " + structField.Name + ` struct tag tj can not be tj:"" (empty string)\r\n` + fmt.Sprintf("%#v", rValue.Interface()))
			}
			if tjValue == "nr" { continue }
			if rValueItem.IsZero() {
				stringRequredMessage := structField.Tag.Get("sr")
				message := ""
				if stringRequredMessage != "" {
					message = stringRequredMessage
				} else {
					message = checker.Format.StringRequiredFailMessage(structField.Name)
				}
				report.Fail = true
				report.Message = message
				return
			}
		case reflect.Struct:
			report = checker.reflectScan(rValueItem, structField.Type)
			if report.Fail {
				return
			}
		}
	}
	return
}

