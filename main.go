package typejson

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"regexp"
)


type TokenInfo struct {
	Type string
	Value string
}
func Parse(jsonString string) (data interface{}, err error){

	jsonByte := []byte(jsonString)

	firstWord := string(jsonByte[0:])
	isFirstWordCorrect, err := regexp.MatchString(`(\[|\{)`, firstWord)
	if err != nil { log.Fatal(err) }
	if !isFirstWordCorrect {
		err = errors.New(`typejson: jsonString first word must be "{" or "[", your error json is:` + "\r\n" + jsonString + "\r\n")
		return
	}
	lastWord := string(jsonByte[len(jsonByte)-1])
	isLastWordCorrect, err := regexp.MatchString(`(\]|\})`, lastWord)
	if err != nil { log.Fatal(err) }
	if !isLastWordCorrect {
		err = errors.New(`typejson: jsonString last word must be "}" or "]", your error json is:` + "\r\n" + jsonString + "\r\n")
		return
	}
	//var dataJSONType string
	//if firstWord == "[" {
	//	dataJSONType = "array"
	//} else {
	//	dataJSONType = "object"
	//}
	// {"name":"nimo"}
	handleJSONstring := jsonByte
	var tokenInfoList = []TokenInfo{}
	beginObject, _ := regexp.Compile(`{`)
	indexRange := beginObject.FindIndex(handleJSONstring)
	if (len(indexRange) != 0) {
		newToken := TokenInfo{
			Type: "BEGIN_OBJECT",
			Value: "",
		}
		handleJSONstring = handleJSONstring[indexRange[1]:]
		tokenInfoList = append(tokenInfoList, newToken)
	}
	fmt.Print(string(handleJSONstring))
	return data, nil

}
