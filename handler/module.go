package handler

import (
	"fmt"

	"github.com/Luftalian/Computer_software/model"
)

// Start		2Byte	 						0F 5A (固定)
// Length		1Byte(<=254)					0E (nByte+13)
// MsgID		1Byte 							11
// MsgNo 		1Byte 							01
// DstID 		4Byte 	送信先デバイス 			XX XX XX XX
// SrcID 		4Byte 	送信元デバイス 			FF FF FF FF
// Parameter 	nByte(n=0～241) 				01

var DstId = "00 00 00 00"
var DstIdFlag = false
var SrcId = "FF FF FF FF"
var SrcIdFlag = false
var MsgId uint64 = 0x00

func (a *App) ModuleStart(dstId string, srcId string) {
	DstId = dstId
	DstIdFlag = true
	SrcId = srcId
	SrcIdFlag = true
	model.HUB.SendText("Module: " + "DstId=" + DstId + " SrcId=" + SrcId)
}

func (a *App) ModuleSend(text string) {
	// string to byte
	byteDate := []byte(text)
	// log.Println(byteDate)
	if !DstIdFlag && !SrcIdFlag {
		model.HUB.SendError("Please set DstId and SrcId")
		return
	}

	sendData := fmt.Sprintf("0F 5A %02X 11 01 %s %s ", len(byteDate)+13, DstId, SrcId)
	for _, b := range byteDate {
		sendData += fmt.Sprintf("%02X ", b)
	}

	a.SerialSend(sendData)
	model.HUB.SendText("Module: " + sendData)
	// return sendData
}
