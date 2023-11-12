package handler

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/Luftalian/Computer_software/model"
)

func (a *App) QuatAndTimeToFile(timeData []byte, quatData []byte) {
	time := binary.LittleEndian.Uint32(timeData)
	quat1 := binary.LittleEndian.Uint32(quatData[0:4])
	quat2 := binary.LittleEndian.Uint32(quatData[4:8])
	quat3 := binary.LittleEndian.Uint32(quatData[8:12])
	quat4 := binary.LittleEndian.Uint32(quatData[12:16])

	floatQuat1 := math.Float32frombits(quat1)
	floatQuat2 := math.Float32frombits(quat2)
	floatQuat3 := math.Float32frombits(quat3)
	floatQuat4 := math.Float32frombits(quat4)

	err := model.AppendStringToFile(fmt.Sprintf("%d,%g,%g,%g,%g,\n", time, floatQuat1, floatQuat2, floatQuat3, floatQuat4), a.quatFileName)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	model.HUB.SendText(fmt.Sprintf("Quat:: %d,%s\n", time, byteArrayToString(quatData[:16])))
	// model.AppendStringToFile(fmt.Sprintf("%s,%s\n", byteArrayToString(timeData), byteArrayToString(quatData[:16])), a.quatFileName)
}

func (a *App) LpsAndTimeToFile(timeData []byte, lpsData []byte) {
	time := binary.LittleEndian.Uint32(timeData)
	lps := binary.LittleEndian.Uint16(lpsData[0:3])

	err := model.AppendStringToFile(fmt.Sprintf("%d,%d,\n", time, lps), a.lpsFileName)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	model.HUB.SendText(fmt.Sprintf("Lps:: %s,%s,\n", byteArrayToString(timeData), byteArrayToString(lpsData[:3])))
	// model.AppendStringToFile(fmt.Sprintf("%s,%s,\n", byteArrayToString(timeData), byteArrayToString(lpsData[:3])), a.lpsFileName)
}

func (a *App) OpenAndTimeToFile(timeData []byte, openData []byte) {
	time := binary.LittleEndian.Uint32(timeData)

	var open int16

	// バイト列から signed int に変換
	err := binary.Read(bytes.NewReader(openData[:2]), binary.LittleEndian, &open)
	if err != nil {
		fmt.Println("Error converting quat1:", err)
		return
	}

	err = model.AppendStringToFile(fmt.Sprintf("%d,%d,\n", time, open), a.openFileName)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	model.HUB.SendText(fmt.Sprintf("Open:: %d,%s,\n", time, byteArrayToString(openData[:2])))
	// model.AppendStringToFile(fmt.Sprintf("%s,%s,\n", byteArrayToString(timeData), byteArrayToString(openData[:2])), a.openFileName)
}

func (a *App) VoltageToFile(voltageData []byte) {
	voltage1 := binary.LittleEndian.Uint16(voltageData[0:3])
	voltage2 := binary.LittleEndian.Uint16(voltageData[3:6])
	voltage3 := binary.LittleEndian.Uint16(voltageData[6:9])

	err := model.AppendStringToFile(fmt.Sprintf("%d,%d,%d,\n", voltage1, voltage2, voltage3), a.voltFileName)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	model.HUB.SendText(fmt.Sprintf("Voltage:: %d,%d,%d,\n", voltage1, voltage2, voltage3))
}
