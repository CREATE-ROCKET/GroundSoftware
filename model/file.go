package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var filename = ""

func CreateFileWithTimestamp() error {
	timestamp := time.Now().Format("2006-01-02-15-04-05") // フォーマット例: 2023-10-26-14-30-00
	filename = filepath.Join("log", "file_"+timestamp+".txt")

	// ディレクトリの存在を確認
	if _, err := os.Stat("log"); err == nil {
	} else if os.IsNotExist(err) {
		err = os.Mkdir("log", 0755)
		if err != nil {
			return err
		}
	} else {
		HUB.SendError(err.Error())
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		HUB.SendError(err.Error())
		return err
	}
	defer file.Close()

	return nil
}

func AppendDataWithTimeToFile(data string) error {
	if filename == "" {
		err := CreateFileWithTimestamp()
		if err != nil {
			HUB.SendError(err.Error())
			return err
		}
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		HUB.SendError(err.Error())
		return err
	}
	defer file.Close()

	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	if _, err := file.WriteString(timeStamp + " - " + data + "\n"); err != nil {
		return err
	}

	return nil
}

// func main() {
// 	dataFromFrontEnd := "新しいデータ"
// 	filename := "追記したいファイルのパス"

// 	err := AppendDataWithTimeToFile(dataFromFrontEnd, filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func AppendToFile(data []byte, fileLabel string) error {
	file, err := os.OpenFile(fileLabel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// if _, err := file.WriteString(bytesToHexSpaceSeparated(data)); err != nil {
	if _, err := file.Write(data); err != nil {
		return err
	}
	return nil
}

func AppendStringToFile(data string, fileLabel string) error {
	file, err := os.OpenFile(fileLabel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return err
	}
	return nil
}

func AppendStringToFileDirect(data string, file *os.File) error {
	// if file == nil {
	// 	os.Open(filename)
	// }
	if _, err := file.WriteString(data); err != nil {
		return err
	}
	return nil
}

func bytesToHexSpaceSeparated(data []byte) string {
	// 16進数文字列のスライスを作成
	hexStrings := make([]string, len(data))
	for i, b := range data {
		hexStrings[i] = fmt.Sprintf("%02X ", b)
	}

	// 文字列に変換
	result := strings.Join(hexStrings, "")

	return result
}
