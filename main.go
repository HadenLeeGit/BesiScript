package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nameFlag bool = false

func main() {

	//field
	var fileName string = "input.txt"

	//if there is a parameter, it should be the file name
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	// read the input file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//Scanner
	scanner := bufio.NewScanner(file)

	//variables to store all lines
	var fieldLines []string
	var funcLines []string

	//read by line
	for scanner.Scan() {
		line := scanner.Text()

		//delete the space at the beginning and end of the line
		line = strings.TrimSpace(line)

		//if the line is empty, it should be ignored
		if len(line) == 0 {
			continue
		}

		//if the line starts with // , it is a comment and should be ignored
		if strings.HasPrefix(line, "//") {
			continue
		}

		//if the line contains //, the part after // should be ignored
		if strings.Contains(line, "//") {
			line = strings.Split(line, "//")[0]
		}

		//ignore multiline comments
		if strings.HasPrefix(line, "/*") {
			for scanner.Scan() {
				line = scanner.Text()
				if strings.Contains(line, "*/") {
					break
				}
			}
			continue
		}

		//if the line start with $, it is a field
		if strings.HasPrefix(line, "$") {
			fieldLines = append(fieldLines, line)
		} else {
			funcLines = append(funcLines, line)
		}
	}

	// input fieldLines in splitField to get the key and value and seartch for the key in the funcLines to replace the value
	for _, fieldLine := range fieldLines {
		key, value := splitField(fieldLine)
		for i, funcLine := range funcLines {
			if strings.Contains(funcLine, key) {
				funcLines[i] = strings.Replace(funcLine, key, value, -1)
			}
		}
	}

	//replace all true to True, false to False
	for i, funcLine := range funcLines {
		funcLines[i] = strings.Replace(funcLine, "true", "True", -1)
		funcLines[i] = strings.Replace(funcLine, "false", "False", -1)
	}

	//call the initFunc to process the init function
	funcLines = initFunc(funcLines)

	//call the mainProcess to process the function
	finalXml := mainProcess(funcLines)

	//print the finalXml
	//fmt.Println(finalXml)

	//if the nameFlag is true, it should be output to the file with the name
	if nameFlag {
		fileName = initBlockInfo.name + ".bsg"
	} else {
		fileName = "output.bsg"
	}

	//output the finalXml to the file
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer outputFile.Close()

	outputFile.WriteString(finalXml)

}

// mainProcess to process the function
func mainProcess(funcLines []string) string {
	var finalXml string = bsgHead

	//replace the %NAME% with the name in init function
	finalXml = strings.Replace(finalXml, "%NAME%", initBlockInfo.name, -1)

	//baseblock
	if initBlockInfo.baseBlock {
		var xml = baseBlock

		//replace the %GUID% with the guidGenerate function
		xml = strings.Replace(xml, "%GUID%", guidGenerate(), -1)

		//replace the %TRANSFORM% with the transform function
		xml = strings.Replace(xml, "%TRANSFORM%", transformProcess(1), -1)

		finalXml += xml
	}

	//loop the funcLines to flowControl function to add to finalXml
	for _, funcLine := range funcLines {
		finalXml += flowControl(funcProcess(funcLine))
	}

	finalXml += bsgTail

	return finalXml

}

// init
func initFunc(funcLines []string) []string {
	//find these functions init(0,5,0) baseblock(true) instrumentinit(0.1,0.1,0.1) logicinit(0.1,0.1,0.1) name(test) and input the value to struct
	var initValue = [3]float64{0, 5, 0}
	var baseBlock bool = false
	var instrumentinitValue = [3]float64{1, 1, 1}
	var logicinitValue = [3]float64{0.1, 0.1, 0.1}
	var nameValue = "output"
	for _, funcLine := range funcLines {
		if strings.Contains(funcLine, "init") {
			_, temp := funcProcess(funcLine)

			//delete value in initValue
			initValue = [3]float64{}

			//convert the string value to float64
			for i, v := range temp {
				initValue[i] = stringToFloat64(v)
			}

		}
		if strings.Contains(funcLine, "baseblock") {
			_, temp := funcProcess(funcLine)
			baseBlock, _ = strconv.ParseBool(temp[0])
		}
		if strings.Contains(funcLine, "instrumentinit") {

			_, temp := funcProcess(funcLine)

			//delete value in instrumentinitValue
			instrumentinitValue = [3]float64{}

			//convert the string value to float64
			for i, v := range temp {
				instrumentinitValue[i] = stringToFloat64(v)
			}

		}
		if strings.Contains(funcLine, "logicinit") {
			_, temp := funcProcess(funcLine)

			//delete value in logicinitValue
			logicinitValue = [3]float64{}

			//convert the string value to float64
			for i, v := range temp {
				logicinitValue[i] = stringToFloat64(v)
			}
		}
		if strings.Contains(funcLine, "name") {
			_, temp := funcProcess(funcLine)
			nameValue = temp[0]

			//set the nameFlag to true
			nameFlag = true
		}

	}

	//store the value in struct
	initBlockInfo.name = nameValue
	initBlockInfo.baseBlock = baseBlock
	initBlockInfo.px = initValue[0]
	initBlockInfo.py = initValue[1]
	initBlockInfo.pz = initValue[2]

	instrumentScaleInfo.sx = instrumentinitValue[0]
	instrumentScaleInfo.sy = instrumentinitValue[1]
	instrumentScaleInfo.sz = instrumentinitValue[2]

	instrumentPositionInfo.px = initValue[0]
	instrumentPositionInfo.py = initValue[1]
	instrumentPositionInfo.pz = initValue[2] + 2

	logicScaleInfo.sx = logicinitValue[0]
	logicScaleInfo.sy = logicinitValue[1]
	logicScaleInfo.sz = logicinitValue[2]

	logicPositionInfo.px = initValue[0] + 2
	logicPositionInfo.py = initValue[1]
	logicPositionInfo.pz = initValue[2]

	//delete these functions init(0,5,0) baseblock(true) instrumentinit(0.1,0.1,0.1) logicinit(0.1,0.1,0.1) name(test) in funcLines
	for i, funcLine := range funcLines {
		if strings.Contains(funcLine, "init") {
			funcLines[i] = ""
		}
		if strings.Contains(funcLine, "baseblock") {
			funcLines[i] = ""
		}
		if strings.Contains(funcLine, "instrumentinit") {
			funcLines[i] = ""
		}
		if strings.Contains(funcLine, "logicinit") {
			funcLines[i] = ""
		}
		if strings.Contains(funcLine, "name") {
			funcLines[i] = ""
		}
	}

	//delete the empty line in funcLines
	funcLines = deleteEmpty(funcLines)

	return funcLines
}

func deleteEmpty(lines []string) []string {
	var ret []string
	for _, line := range lines {
		if line != "" {
			ret = append(ret, line)
		}
	}
	return ret
}
