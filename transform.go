package main

import "strings"

type initBlock struct {
	name      string
	baseBlock bool
	px        float64
	py        float64
	pz        float64
}

type logicPosition struct {
	px float64
	py float64
	pz float64
}

type logicScale struct {
	sx float64
	sy float64
	sz float64
}

type instrumentPosition struct {
	px float64
	py float64
	pz float64
}

type instrumentScale struct {
	sx float64
	sy float64
	sz float64
}

// for transform info
var initBlockInfo initBlock
var instrumentScaleInfo instrumentScale
var instrumentPositionInfo instrumentPosition
var logicScaleInfo logicScale
var logicPositionInfo logicPosition

// transform field
var logicCount int

// function for transform
func transformProcess(blockType int) string {
	var transform string = transform

	//if the blockType is 1, it should be a base block
	if blockType == 1 {
		//convert px to string and replace the %PX% with the px
		transform = strings.Replace(transform, "%PX%", float64ToString(initBlockInfo.px), -1)
		//convert py to string and replace the %PY% with the py
		transform = strings.Replace(transform, "%PY%", float64ToString(initBlockInfo.py), -1)
		//convert pz to string and replace the %PZ% with the pz
		transform = strings.Replace(transform, "%PZ%", float64ToString(initBlockInfo.pz), -1)
		//convert sx to string and replace the %SX% with the sx
		transform = strings.Replace(transform, "%SX%", "1", -1)
		//convert sy to string and replace the %SY% with the sy
		transform = strings.Replace(transform, "%SY%", "1", -1)
		//convert sz to string and replace the %SZ% with the sz
		transform = strings.Replace(transform, "%SZ%", "1", -1)
	}

	//if the blockType is 2, it should be a instrument block
	if blockType == 2 {

		instrumentPositionInfo.pz = instrumentPositionInfo.pz + 2*instrumentScaleInfo.sz

		//convert px to string and replace the %PX% with the px
		transform = strings.Replace(transform, "%PX%", float64ToString(instrumentPositionInfo.px), -1)
		//convert py to string and replace the %PY% with the py
		transform = strings.Replace(transform, "%PY%", float64ToString(instrumentPositionInfo.py), -1)
		//convert pz to string and replace the %PZ% with the pz
		transform = strings.Replace(transform, "%PZ%", float64ToString(instrumentPositionInfo.pz), -1)
		//convert sx to string and replace the %SX% with the sx
		transform = strings.Replace(transform, "%SX%", float64ToString(instrumentScaleInfo.sx), -1)
		//convert sy to string and replace the %SY% with the sy
		transform = strings.Replace(transform, "%SY%", float64ToString(instrumentScaleInfo.sy), -1)
		//convert sz to string and replace the %SZ% with the sz
		transform = strings.Replace(transform, "%SZ%", float64ToString(instrumentScaleInfo.sz), -1)
	}

	//if the blockType is 3, it should be a logic block
	if blockType == 3 {

		if logicCount%50 == 0 {
			logicPositionInfo.pz = initBlockInfo.pz
			logicPositionInfo.px = logicPositionInfo.px + 2*logicScaleInfo.sx
		} else {
			logicPositionInfo.pz = logicPositionInfo.pz + 2*logicScaleInfo.sx
		}

		//convert px to string and replace the %PX% with the px
		transform = strings.Replace(transform, "%PX%", float64ToString(logicPositionInfo.px), -1)
		//convert py to string and replace the %PY% with the py
		transform = strings.Replace(transform, "%PY%", float64ToString(logicPositionInfo.py), -1)
		//convert pz to string and replace the %PZ% with the pz
		transform = strings.Replace(transform, "%PZ%", float64ToString(logicPositionInfo.pz), -1)
		//convert sx to string and replace the %SX% with the sx
		transform = strings.Replace(transform, "%SX%", float64ToString(logicScaleInfo.sx), -1)
		//convert sy to string and replace the %SY% with the sy
		transform = strings.Replace(transform, "%SY%", float64ToString(logicScaleInfo.sy), -1)
		//convert sz to string and replace the %SZ% with the sz
		transform = strings.Replace(transform, "%SZ%", float64ToString(logicScaleInfo.sz), -1)

		logicCount++
	}

	return transform
}
