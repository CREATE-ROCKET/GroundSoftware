package handler

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Luftalian/Computer_software/model"
)

// Start		2Byte	 						0F 5A (固定)
// Length		1Byte(<=254)					0E (nByte+13)
// MsgID		1Byte 							11
// MsgNo 		1Byte 							01
// DstID 		4Byte 	送信先デバイス 			XX XX XX XX
// SrcID 		4Byte 	送信元デバイス 			FF FF FF FF
// Parameter 	nByte(n=0～241) 				01

var DstId = "FFFFFFFF"
var DstIdFlag = false
var SrcId = "FFFFFFFF"
var SrcIdFlag = false
var MsgId uint64 = 0x55

type Configuration struct {
	DST_ID  string `json:"DST_ID"`
	SRC_ID  string `json:"SRC_ID"`
	POWER   string `json:"POWER"`
	CHANNEL string `json:"CHANNEL"`
	RF_BAND string `json:"RF_BAND"`
	CS_MODE string `json:"CS_MODE"`
}

func (a *App) ModuleStart(dstId string, srcId string) {
	DstId = dstId
	DstIdFlag = true
	SrcId = srcId
	SrcIdFlag = true
	model.HUB.SendText("Module: " + "DstId=" + DstId + " SrcId=" + SrcId)
}

func (a *App) ModuleSend(text string) {
	// log.Print("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if !DstIdFlag && !SrcIdFlag {
		model.HUB.SendError("Please set DstId and SrcId")
		return
	}
	// string to byte
	// log.Print(text)
	byteDate := []byte(text)
	// log.Print(byteDate)
	// a.ModuleStyleSerialSend(byteDate)
	// log.Println(byteDate)
	sendData := fmt.Sprintf("0F5A%02X21%02d%s%s", len(byteDate)+13, MsgId, DstId, SrcId)
	MsgId += 1
	// for _, b := range byteDate {
	// 	sendData += fmt.Sprintf("%02X ", b)
	// 	log.Print(b)
	// }

	// a.SerialTextSend(sendData)
	/// log.Println(sendData)
	s, err := hexStringToBytes(sendData)
	/// log.Println(s)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	s = append(s, byteDate...)
	a.SerialByteSend(s)
	model.HUB.SendText("Module: " + sendData)
	// return sendData
}

func (a *App) ModuleStyleSerialSend(byteDate []byte) {
	if !DstIdFlag && !SrcIdFlag {
		model.HUB.SendError("Please set DstId and SrcId")
		return
	}
	sendData := fmt.Sprintf("0F5A%02X21%02d%s%s", len(byteDate)+13, MsgId, DstId, SrcId)
	MsgId += 1
	// for _, b := range byteDate {
	// 	sendData += fmt.Sprintf("%02X", b)
	// 	log.Print(b)
	// }

	s, err := hexStringToBytes(sendData)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	s = append(s, byteDate...)
	a.SerialByteSend(s)
	model.HUB.SendText("Module: " + sendData)
	// return sendData
}

func (a *App) ModuleEnv() {

	filePath := "config.json"

	// JSONファイルの読み込み
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("ファイルの読み込みエラー: %v", err)
		model.HUB.SendError(err.Error())
	}

	// JSONデコード
	var config Configuration
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Printf("JSONデコードエラー: %v", err)
		model.HUB.SendError(err.Error())
	}

	// 読み込んだデータの表示
	fmt.Printf("DST_ID: %s\n", config.DST_ID)
	fmt.Printf("SRC_ID: %s\n", config.SRC_ID)
	fmt.Printf("POWER: %s\n", config.POWER)
	fmt.Printf("CHANNEL: %s\n", config.CHANNEL)
	fmt.Printf("RF_BAND: %s\n", config.RF_BAND)
	fmt.Printf("CS_MODE: %s\n", config.CS_MODE)

	if config.DST_ID == "" || config.SRC_ID == "" || config.POWER == "" || config.CHANNEL == "" || config.RF_BAND == "" || config.CS_MODE == "" {
		model.HUB.SendError("Please set all config")
		return
	}

	DstId = config.DST_ID
	DstIdFlag = true
	SrcId = config.SRC_ID
	SrcIdFlag = true

	s, err := hexStringToBytes(fmt.Sprintf("%s%s%s%s", config.POWER, config.CHANNEL, config.RF_BAND, config.CS_MODE))
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	a.ModuleStyleSerialSend(s)
}

func hexStringToBytes(hexString string) ([]byte, error) {
	// 文字列を[]byteに変換する
	byteString, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	return byteString, nil
}
