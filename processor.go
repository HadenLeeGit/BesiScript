package main

import (
	"math/rand"
	"strings"
)

func flowControl(funcName string, paramArray []string) string {
	var xml string
	// for logic block
	switch funcName {
	case "not", "and", "or", "nor", "nand", "xor", "xnor", "random", "srlatch", "dlatch", "counter", "edge":
		xml = logicProcess(funcName, paramArray)
	case "autosensor", "sensor":
		xml = sensorProcess(funcName, paramArray)
	case "autotimer", "htrtimer", "timer":
		xml = timerProcess(funcName, paramArray)
	case "autospeed", "speed":
		xml = speedProcess(funcName, paramArray)
	case "autoalti", "alti":
		xml = altiProcess(funcName, paramArray)
	case "autoangle", "angle":
		xml = angleProcess(funcName, paramArray)
	}

	return xml
}

// function to generate radon GIUD like afd7af7c-0bc7-489a-9303-7fd156203d70 format
func guidGenerate() string {
	var guid string
	var charArray = []rune("0123456789abcdef")
	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			guid += "-"
			continue
		}
		guid += string(charArray[rand.Intn(16)])
	}
	return guid
}

func funcProcess(s string) (string, []string) {

	//and(A+B, "AAA", "AAA", true)
	//get the function name
	var funcName string
	for i, v := range s {
		if v == '(' {
			funcName = s[:i]
			break
		}
	}

	//delete the space at the beginning and end of the line
	funcName = strings.TrimSpace(funcName)

	//get the parameters
	var params string
	for i, v := range s {
		if v == '(' {
			params = s[i+1:]
			break
		}
	}

	//get the parameters in string array
	var paramArray []string
	var param string
	var count int
	for _, v := range params {
		if v == '(' {
			count++
		}
		if v == ')' {
			count--
		}
		if v == ',' && count == 0 {
			paramArray = append(paramArray, param)
			param = ""
			continue
		}
		param += string(v)
	}

	//append the last parameter
	paramArray = append(paramArray, param)

	//delete ) char in the param
	for i, v := range paramArray {
		paramArray[i] = strings.Replace(v, ")", "", -1)
	}

	//delete the space at the beginning and end of each parameter
	for i, v := range paramArray {
		paramArray[i] = strings.TrimSpace(v)
	}

	return funcName, paramArray
}
