package main

import "strings"

func logicProcess(funcName string, paramArray []string) string {

	var xml string = logicBlock

	//logic
	//and(A+B, "AAA", "AAA", true)
	//or(A+B, "AAA", "AAA", true)
	//nor(A+B, "AAA", "AAA", true)
	//nand(A+B, "AAA", "AAA", true)
	//xor(A+B, "AAA", "AAA", true)
	//xnor(A+B, "AAA", "AAA", true)
	//random(A+B, "AAA", true)
	//srlatch(A+B, "AAA", "AAA")
	//dlatch(A+B, "AAA", "AAA")
	//counter(A+B, "AAA", "AAA")
	//edge(A+B, "AAA", true)
	//not(A+B, "AAA", true)
	//"not", "and", "or", "nor", "nand", "xor", "xnor", "random", "srlatch", "dlatch", "counter", "edge"
	switch funcName {
	case "not":
		//get the value
		var input string = stringProcess(paramArray[0])
		var output string = stringProcess(paramArray[1])

		//replace the %INPUT0% with the input
		xml = strings.Replace(xml, "%INPUT0%", input, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "0", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[2], -1)

	case "and":

		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "1", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[3], -1)

	case "or":

		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "2", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[3], -1)

	case "nor":
		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "3", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[3], -1)

	case "nand":
		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "4", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[3], -1)

	case "xor":
		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "5", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[3], -1)

	case "xnor":
		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "6", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[3], -1)

	case "random":
		//get the value
		var input string = stringProcess(paramArray[0])
		var output string = stringProcess(paramArray[1])

		//replace the %INPUT0% with the input
		xml = strings.Replace(xml, "%INPUT0%", input, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "7", -1)
		//replace the %TOGGLE% with the toggle mode
		xml = strings.Replace(xml, "%TOGGLE%", paramArray[2], -1)

	case "srlatch":
		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "8", -1)

	case "dlatch":
		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "9", -1)

	case "counter":
		//get the value
		var input0 string = stringProcess(paramArray[0])
		var input1 string = stringProcess(paramArray[1])
		var output string = stringProcess(paramArray[2])

		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %INPUT1% with the input1
		xml = strings.Replace(xml, "%INPUT1%", input1, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "10", -1)

	case "edge":
		//get the value
		var input string = stringProcess(paramArray[0])
		var output string = stringProcess(paramArray[1])

		//replace the %INPUT0% with the input
		xml = strings.Replace(xml, "%INPUT0%", input, -1)
		//replace the %OUTPUT% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %GATE% with the gate number
		xml = strings.Replace(xml, "%GATE%", "11", -1)
		//replace the %INVERTED% with the INVERTED mode
		xml = strings.Replace(xml, "%INVERTED%", paramArray[2], -1)
	}

	//replace the %INVERTED% with the init mode
	xml = strings.Replace(xml, "%INVERTED%", "False", -1)

	//replace the %TOGGLE% with the init mode
	xml = strings.Replace(xml, "%TOGGLE%", "True", -1)

	//%TRANSFORM%
	xml = strings.Replace(xml, "%TRANSFORM%", transformProcess(3), -1)

	//replace the %GUID% with the guid
	xml = strings.Replace(xml, "%GUID%", guidGenerate(), -1)

	//delete all unplaced %%
	xml = strings.Replace(xml, "%", "", -1)

	return xml

}
