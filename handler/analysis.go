package handler

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"

	"github.com/Luftalian/Computer_software/model"
)

var quatIndex int = 0
var lpsIndex int = 0
var openIndex int = 0
var voltIndex int = 0

func (a *App) QuatAndTimeToFile(timeData []byte, quatData []byte) {
	time := binary.LittleEndian.Uint32(timeData)

	floatQuat1 := math.Float32frombits(binary.LittleEndian.Uint32(quatData[0:4]))
	floatQuat2 := math.Float32frombits(binary.LittleEndian.Uint32(quatData[4:8]))
	floatQuat3 := math.Float32frombits(binary.LittleEndian.Uint32(quatData[8:12]))
	floatQuat4 := math.Float32frombits(binary.LittleEndian.Uint32(quatData[12:16]))

	// err := model.AppendStringToFile(fmt.Sprintf("%d,%g,%g,%g,%g,\n", time, floatQuat1, floatQuat2, floatQuat3, floatQuat4), a.quatFileName)
	err := model.AppendStringToFileDirect(fmt.Sprintf("%d,%g,%g,%g,%g,\n", time, floatQuat1, floatQuat2, floatQuat3, floatQuat4), a.quatFile)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	if quatIndex%8 == 0 {
		model.HUB.SendText(fmt.Sprintf("Quat::%d,%g,%g,%g,%g,\n", time, floatQuat1, floatQuat2, floatQuat3, floatQuat4))
		// model.AppendStringToFile(fmt.Sprintf("%s,%s\n", byteArrayToString(timeData), byteArrayToString(quatData[:16])), a.quatFileName)
		quatIndex = 0
	}
	quatIndex++
}

func (a *App) LpsAndTimeToFile(timeData []byte, lpsData []byte) {
	time := binary.LittleEndian.Uint32(timeData)
	lps := binary.LittleEndian.Uint16(lpsData[0:3])

	// err := model.AppendStringToFile(fmt.Sprintf("%d,%d,\n", time, lps), a.lpsFileName)
	err := model.AppendStringToFileDirect(fmt.Sprintf("%d,%d,\n", time, lps), a.lpsFile)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	if lpsIndex%100 == 0 {
		model.HUB.SendText(fmt.Sprintf("Lps::%d,%d,\n", time, lps))
		// model.AppendStringToFile(fmt.Sprintf("%s,%s,\n", byteArrayToString(timeData), byteArrayToString(lpsData[:3])), a.lpsFileName)
		lpsIndex = 0
	}
	lpsIndex++
}

func (a *App) OpenAndTimeToFile(timeData []byte, openData []byte) {
	time := binary.LittleEndian.Uint32(timeData)

	var open int16

	// バイト列から signed int に変換
	err := binary.Read(bytes.NewReader(openData[:2]), binary.LittleEndian, &open)
	if err != nil {
		log.Println("Error converting quat1:", err)
		return
	}

	// err = model.AppendStringToFile(fmt.Sprintf("%d,%d,\n", time, open), a.openFileName)
	err = model.AppendStringToFileDirect(fmt.Sprintf("%d,%d,\n", time, open), a.openFile)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	if openIndex%100 == 0 {
		model.HUB.SendText(fmt.Sprintf("OpenRate::%d,%d,\n", time, open))
		// model.AppendStringToFile(fmt.Sprintf("%s,%s,\n", byteArrayToString(timeData), byteArrayToString(openData[:2])), a.openFileName)
		openIndex = 0
	}
	openIndex++
}

func (a *App) VoltageToFile(voltageData []byte) {
	voltage1 := binary.LittleEndian.Uint16(voltageData[0:3])
	voltage2 := binary.LittleEndian.Uint16(voltageData[3:6])
	voltage3 := binary.LittleEndian.Uint16(voltageData[6:9])

	// err := model.AppendStringToFile(fmt.Sprintf("%d,%d,%d,\n", voltage1, voltage2, voltage3), a.voltFileName)
	err := model.AppendStringToFileDirect(fmt.Sprintf("%d,%d,%d,\n", voltage1, voltage2, voltage3), a.voltFile)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	if voltIndex%100 == 0 {
		model.HUB.SendText(fmt.Sprintf("Voltage::%d,%d,%d,\n", voltage1, voltage2, voltage3))
		voltIndex = 0
	}
	voltIndex++
}
