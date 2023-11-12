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
	quat1 := binary.LittleEndian.Uint16(quatData[0:4])
	quat2 := binary.LittleEndian.Uint16(quatData[4:8])
	quat3 := binary.LittleEndian.Uint16(quatData[8:12])
	quat4 := binary.LittleEndian.Uint16(quatData[12:16])

	floatQuat1 := math.Float32frombits(uint32(quat1))
	floatQuat2 := math.Float32frombits(uint32(quat2))
	floatQuat3 := math.Float32frombits(uint32(quat3))
	floatQuat4 := math.Float32frombits(uint32(quat4))

	err := model.AppendStringToFile(fmt.Sprintf("%d,%g,%g,%g,%g,\n", time, floatQuat1, floatQuat2, floatQuat3, floatQuat4), a.quatFileName)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	model.HUB.SendText(fmt.Sprintf("Quat:: %d,%g,%g,%g,%g,\n", time, floatQuat1, floatQuat2, floatQuat3, floatQuat4))
}

func (a *App) LpsAndTimeToFile(timeData []byte, lpsData []byte) {
	time := binary.LittleEndian.Uint32(timeData)
	lps := binary.LittleEndian.Uint16(lpsData[0:2])

	err := model.AppendStringToFile(fmt.Sprintf("%d,%d,\n", time, lps), a.lpsFileName)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	model.HUB.SendText(fmt.Sprintf("Lps:: %d,%d,\n", time, lps))
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

	model.HUB.SendText(fmt.Sprintf("Open:: %d,%d,\n", time, open))
}

func (a *App) VoltageToFile(voltageData []byte) {
	voltage := binary.LittleEndian.Uint16(voltageData[0:2])

	err := model.AppendStringToFile(fmt.Sprintf("%d,\n", voltage), a.voltFileName)
	if err != nil {
		model.HUB.SendError(err.Error())
	}

	model.HUB.SendText(fmt.Sprintf("Voltage:: %d,\n", voltage))
}
