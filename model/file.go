package model

import (
	"os"
	"time"
)

var filename = ""

func CreateFileWithTimestamp() error {
	currentTime := time.Now()
	timestamp := currentTime.Format("2006-01-02-15-04-05") // フォーマット例: 2023-10-26-14-30-00
	filename = "log/file_" + timestamp + ".txt"

	dirPath := "log"

	// ディレクトリの存在を確認
	if _, err := os.Stat(dirPath); err == nil {
	} else if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func AppendDataWithTimeToFile(data string) error {
	if filename == "" {
		err := CreateFileWithTimestamp()
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	currentTime := time.Now()
	timeStamp := currentTime.Format("2006-01-02 15:04:05")
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
