package main

import "strings"

func sensorProcess(funcName string, paramArray []string) string {

	//autosensor(5, 0.5, "AAA", true, false)
	//sensor(5, 0.5, A+B, true, false, A+B, false)

	var xml string = sensorBlock

	switch funcName {
	case "autosensor":
		//get the value
		var distance string = paramArray[0]
		var radius string = paramArray[1]
		var output string = stringProcess(paramArray[2])
		var reverse string = paramArray[3]
		var static string = paramArray[4]

		//replace the %DISTANCE% with the distance
		xml = strings.Replace(xml, "%DISTANCE%", distance, -1)
		//replace the %RADIUS% with the radius
		xml = strings.Replace(xml, "%RADIUS%", radius, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %STATIC% with the static mode
		xml = strings.Replace(xml, "%STATIC%", static, -1)
		//replace the %AUTO% with the false
		xml = strings.Replace(xml, "%AUTO%", "False", -1)
	case "sensor":
		//get the value
		var distance string = paramArray[0]
		var radius string = paramArray[1]
		var output string = stringProcess(paramArray[2])
		var reverse string = paramArray[3]
		var static string = paramArray[4]
		var input0 string = stringProcess(paramArray[5])
		var hold string = paramArray[6]

		//replace the %DISTANCE% with the distance
		xml = strings.Replace(xml, "%DISTANCE%", distance, -1)
		//replace the %RADIUS% with the radius
		xml = strings.Replace(xml, "%RADIUS%", radius, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %STATIC% with the static mode
		xml = strings.Replace(xml, "%STATIC%", static, -1)
		//replace the %AUTO% with the true
		xml = strings.Replace(xml, "%AUTO%", "True", -1)
		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %HOLD% with the hold mode
		xml = strings.Replace(xml, "%HOLD%", hold, -1)

	}

	//%HOLD% with init mode
	xml = strings.Replace(xml, "%HOLD%", "False", -1)

	//%TRANSFORM%
	xml = strings.Replace(xml, "%TRANSFORM%", transformProcess(2), -1)

	//replace the %GUID% with the guid
	xml = strings.Replace(xml, "%GUID%", guidGenerate(), -1)

	//delete all unplaced %%
	xml = strings.Replace(xml, "%", "", -1)

	return xml
}

func timerProcess(funcName string, paramArray []string) string {
	// autotimer(5, 5, "AAA", true)
	// htrtimer(5, 5, "AAA", true, A+B)
	// timer(5, 5, A+B, true, B, true)

	var xml string = timerBlock

	switch funcName {
	case "autotimer":
		//get the value
		var wait string = paramArray[0]
		var time string = paramArray[1]
		var output string = stringProcess(paramArray[2])
		var loop string = paramArray[3]

		//replace the %WAIT% with the wait
		xml = strings.Replace(xml, "%WAIT%", wait, -1)
		//replace the %TIME% with the time
		xml = strings.Replace(xml, "%TIME%", time, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %LOOP% with the loop mode
		xml = strings.Replace(xml, "%LOOP%", loop, -1)
		//replace the %AUTO% with the true
		xml = strings.Replace(xml, "%AUTO%", "True", -1)
	case "htrtimer":
		//get the value
		var wait string = paramArray[0]
		var time string = paramArray[1]
		var output string = stringProcess(paramArray[2])
		var loop string = paramArray[3]
		var input0 string = stringProcess(paramArray[4])

		//replace the %WAIT% with the wait
		xml = strings.Replace(xml, "%WAIT%", wait, -1)
		//replace the %TIME% with the time
		xml = strings.Replace(xml, "%TIME%", time, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %LOOP% with the loop mode
		xml = strings.Replace(xml, "%LOOP%", loop, -1)
		//replace the %AUTO% with the false
		xml = strings.Replace(xml, "%AUTO%", "False", -1)
		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %HOLD% with true
		xml = strings.Replace(xml, "%HOLD%", "True", -1)
	case "timer":
		//get the value
		var wait string = paramArray[0]
		var time string = paramArray[1]
		var output string = stringProcess(paramArray[2])
		var loop string = paramArray[3]
		var input0 string = stringProcess(paramArray[4])
		var stop string = paramArray[5]

		//replace the %WAIT% with the wait
		xml = strings.Replace(xml, "%WAIT%", wait, -1)
		//replace the %TIME% with the time
		xml = strings.Replace(xml, "%TIME%", time, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %LOOP% with the loop mode
		xml = strings.Replace(xml, "%LOOP%", loop, -1)
		//replace the %AUTO% with the false
		xml = strings.Replace(xml, "%AUTO%", "False", -1)
		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %HOLD% with False
		xml = strings.Replace(xml, "%HOLD%", "False", -1)
		//replace the %STOP% with the stop mode
		xml = strings.Replace(xml, "%STOP%", stop, -1)
	}

	//%STOP% with init mode
	xml = strings.Replace(xml, "%STOP%", "False", -1)

	//%HOLD% with init mode
	xml = strings.Replace(xml, "%HOLD%", "False", -1)

	//%TRANSFORM%
	xml = strings.Replace(xml, "%TRANSFORM%", transformProcess(2), -1)

	//replace the %GUID% with the guid
	xml = strings.Replace(xml, "%GUID%", guidGenerate(), -1)

	//delete all unplaced %%
	xml = strings.Replace(xml, "%", "", -1)

	return xml
}

func altiProcess(funcName string, paramArray []string) string {

	//autoalti(5, "AAA", false)
	//alti(5, "AAA", false, B, false)

	var xml string = altimeterBlock

	switch funcName {
	case "autoalti":
		//get the value
		var height string = paramArray[0]
		var output string = stringProcess(paramArray[1])
		var reverse string = paramArray[2]

		//replace the %HEIGHT% with the height
		xml = strings.Replace(xml, "%HEIGHT%", height, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %AUTO% with the false
		xml = strings.Replace(xml, "%AUTO%", "False", -1)
	case "alti":
		//get the value
		var height string = paramArray[0]
		var output string = stringProcess(paramArray[1])
		var reverse string = paramArray[2]
		var input0 string = stringProcess(paramArray[3])
		var hold string = paramArray[4]

		//replace the %HEIGHT% with the height
		xml = strings.Replace(xml, "%HEIGHT%", height, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %AUTO% with the true
		xml = strings.Replace(xml, "%AUTO%", "True", -1)
		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %HOLD% with the hold mode
		xml = strings.Replace(xml, "%HOLD%", hold, -1)

	}

	//%HOLD% with init mode
	xml = strings.Replace(xml, "%HOLD%", "False", -1)

	//%TRANSFORM%
	xml = strings.Replace(xml, "%TRANSFORM%", transformProcess(2), -1)

	//replace the %GUID% with the guid
	xml = strings.Replace(xml, "%GUID%", guidGenerate(), -1)

	//delete all unplaced %%
	xml = strings.Replace(xml, "%", "", -1)

	return xml

}

func speedProcess(funcName string, paramArray []string) string {
	//autospeed(5, "AAA", false)
	//speed(5, "AAA", false, B, false)

	var xml string = speedmeterBlock

	switch funcName {
	case "autospeed":
		//get the value
		var speed string = paramArray[0]
		var output string = stringProcess(paramArray[1])
		var reverse string = paramArray[2]

		//replace the %SPEED% with the height
		xml = strings.Replace(xml, "%SPEED%", speed, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %AUTO% with the false
		xml = strings.Replace(xml, "%AUTO%", "False", -1)
	case "speed":
		//get the value
		var speed string = paramArray[0]
		var output string = stringProcess(paramArray[1])
		var reverse string = paramArray[2]
		var input0 string = stringProcess(paramArray[3])
		var hold string = paramArray[4]

		//replace the %SPEED% with the height
		xml = strings.Replace(xml, "%SPEED%", speed, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %AUTO% with the true
		xml = strings.Replace(xml, "%AUTO%", "True", -1)
		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %HOLD% with the hold mode
		xml = strings.Replace(xml, "%HOLD%", hold, -1)

	}

	//%HOLD% with init mode
	xml = strings.Replace(xml, "%HOLD%", "False", -1)

	//%TRANSFORM%
	xml = strings.Replace(xml, "%TRANSFORM%", transformProcess(2), -1)

	//replace the %GUID% with the guid
	xml = strings.Replace(xml, "%GUID%", guidGenerate(), -1)

	//delete all unplaced %%
	xml = strings.Replace(xml, "%", "", -1)

	return xml

}

func angleProcess(funcName string, paramArray []string) string {

	//autoangle(45, -45, "AAA", false)
	//angle(45, -45, "AAA", false, B, false)
	var xml string = anglometerBlock

	switch funcName {
	case "autoangle":
		//get the value
		var start string = paramArray[0]
		var end string = paramArray[1]
		var output string = stringProcess(paramArray[2])
		var reverse string = paramArray[3]

		//replace the %START% with the height
		xml = strings.Replace(xml, "%START%", start, -1)
		//replace the %END% with the end
		xml = strings.Replace(xml, "%END%", end, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %AUTO% with the false
		xml = strings.Replace(xml, "%AUTO%", "False", -1)
	case "angle":
		//get the value
		var start string = paramArray[0]
		var end string = paramArray[1]
		var output string = stringProcess(paramArray[2])
		var reverse string = paramArray[3]
		var input0 string = stringProcess(paramArray[4])
		var hold string = paramArray[5]

		//replace the %START% with the height
		xml = strings.Replace(xml, "%START%", start, -1)
		//replace the %END% with the end
		xml = strings.Replace(xml, "%END%", end, -1)
		//replace the %output% with the output
		xml = strings.Replace(xml, "%OUTPUT%", output, -1)
		//replace the %REVERSE% with the reverse mode
		xml = strings.Replace(xml, "%REVERSE%", reverse, -1)
		//replace the %AUTO% with the true
		xml = strings.Replace(xml, "%AUTO%", "True", -1)
		//replace the %INPUT0% with the input0
		xml = strings.Replace(xml, "%INPUT0%", input0, -1)
		//replace the %HOLD% with the hold mode
		xml = strings.Replace(xml, "%HOLD%", hold, -1)

	}

	//%HOLD% with init mode
	xml = strings.Replace(xml, "%HOLD%", "False", -1)

	//%TRANSFORM%
	xml = strings.Replace(xml, "%TRANSFORM%", transformProcess(2), -1)

	//replace the %GUID% with the guid
	xml = strings.Replace(xml, "%GUID%", guidGenerate(), -1)

	//delete all unplaced %%
	xml = strings.Replace(xml, "%", "", -1)

	return xml
}
