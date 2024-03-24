package main

import (
	"strconv"
	"strings"
)

//$msg = "Hello World"

// funtion to split the string by =
func splitField(s string) (string, string) {
	var key, value string
	for i, v := range s {
		if v == '=' {
			key = s[:i]
			value = s[i+1:]
			break
		}
	}
	//delete the space at the beginning and end of the line
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	return key, value
}

// funtion to convert the value like "AAA", [A, B], A to xml
func stringProcess(s string) string {

	var valueArray []string
	var typeID int

	//call the valueProcess function to process the value
	valueArray, typeID = valueProcess(s)

	//if the typeID is 1, it should be a string
	if typeID == 1 {

		var msgVarRet string
		//get msgVar from processor.go and replace %VAR% with the value
		msgVarRet = strings.Replace(msgVar, "%VAR%", valueArray[0], -1)
		return msgVarRet
	}

	//if the typeID is 2, it should be a string array
	if typeID == 2 {
		var btnVar = "<String>%VAR%</String>"
		var btnVarRet string
		//get btnVar from processor.go and replace %VAR% with the value
		for _, v := range valueArray {
			btnVarRet += strings.Replace(btnVar, "%VAR%", v, -1) + "\n"
		}
		return btnVarRet
	}

	//if the typeID is 3
	return valueArray[0]

}

// funtion to process the value like "AAA", [A, B], A
func valueProcess(value string) ([]string, int) {
	//a variable to store the value in string array
	var valueArray []string
	var typeID int

	//if the value starts with " and ends with ", it should be removed
	if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
		value = value[1 : len(value)-1]
		valueArray = append(valueArray, value)
		typeID = 1

		// if the value has +, it should be split by + and remove the space at the beginning and end of each part
	} else if strings.Contains(value, "+") {

		parts := strings.Split(value, "+")
		for i, v := range parts {
			parts[i] = strings.TrimSpace(v)
		}

		valueArray = parts
		typeID = 2
	} else {
		valueArray = append(valueArray, value)
		typeID = 3
	}

	return valueArray, typeID
}

// function to convert float64 to string
func float64ToString(f float64) string {
	//use strconv to convert float64 to string
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// function to convert string to float64
func stringToFloat64(s string) float64 {
	//use strconv to convert string to float64
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
