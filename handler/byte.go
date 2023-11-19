package handler

import (
	"errors"
)

// func byteArrayToString(bytes []byte) string {
// 	result := "["
// 	for i, b := range bytes {
// 		result += strconv.Itoa(int(b))
// 		if i < len(bytes)-1 {
// 			result += " "
// 		}
// 	}
// 	result += "]"
// 	return result
// }

// var dataLength int

// 送られてきた配列が最初のコマンド2つから始まって、データ列の指定された大きさと送られてきた配列の大きさが同じとき
// 送られてきた配列が最初のコマンド2つから始まって、データ列の指定された大きさと送られてきた配列の大きさが小さいとき
// 送られてきた配列が最初のコマンド2つから始まって、データ列の指定された大きさと送られてきた配列の大きさが大きいとき
// 送られてきた配列が最初のコマンドを含まないとき
// 送られてきた配列が最初のコマンドを途中に1つ含むとき
// 送られてきた配列が最初のコマンドを途中に複数含むとき

// 最初のコマンドをずっと探して、そのデータの長さを取得して、その長さの間はコマンドを探さないようにする
// そのデータの長さを超えたらまたコマンドを探すようにする
// データの長さの間での解析をする。長さが187 bytesのときのみ解析する

// 最初のコマンド含むけど、長さコマンド含まないとき

// {

// 	// テスト用のデータを送信するチャネル
// 	receivedData := make(chan []byte)

// 	// バックグラウンドでデータを受信し解析するゴルーチン
// 	go ReceiveData(receivedData)
//  receivedData <- testData
// }

// Custom errors
var (
	ErrStartCommandNotFound       = errors.New("start command not found")
	ErrInvalidDataLength          = errors.New("invalid data length")
	ErrInsufficientDataLength     = errors.New("insufficient data length")
	ErrLengthCommandNotFound      = errors.New("length command not found")
	ErrStartCommandFoundOnlyFirst = errors.New("start command found only first")
)

// ReceiveData はデータを受け取り、解析します。
func (a *App) ReceiveData(receivedData chan []byte) {
	var buffer []byte
	for data := range receivedData {
		// 受け取ったデータをバッファに追加
		buffer = append(buffer, data...)
		for len(buffer) > 0 {
			start, end, err := StartCommand(buffer)
			// if err != nil {
			// 	// スタートコマンドが見つからないか、データの長さが不正な場合
			// 	break
			// }
			if err == ErrStartCommandNotFound {
				buffer = nil
				break
			} else if err == ErrStartCommandFoundOnlyFirst {
				buffer = buffer[len(buffer)-1:]
				break
			} else if err == ErrLengthCommandNotFound {
				buffer = buffer[start-1:]
				break
			} else if err == ErrInsufficientDataLength {
				buffer = buffer[start-2:]
				break
			} else if err != nil {
				// Undefined error
				break
			} else if err == nil {
				// データを取り出して解析
				parsedData := buffer[start:end]
				/// log.Println(parsedData)
				go a.ParseData(parsedData) // go文に

				// 解析が終わったデータをバッファから削除
				buffer = buffer[end:]
			}
		}
	}
}

// StartCommand は指定されたバイト列からスタートコマンドを探し、見つかればそのデータの長さと次に読み取るべき位置を返します。
func StartCommand(data []byte) (int, int, error) {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x0f && data[i+1] == 0x5a {
			if i+2 < len(data) {
				// スタートコマンドの次にデータの長さがあるか確認
				length := int(data[i+2]) - 2 // 最初のコマンドを除いた長さ
				if i+2+length < len(data) {
					return i + 2, i + 2 + length, nil
				} else {
					/// log.Println("aaaaaaaaaasfffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", i+2+length, len(data))
					// データの長さが足りない
					// 最初のコマンドと長さコマンド含むけど、データが途中までしかないとき
					return i + 2, i + 2 + length, ErrInsufficientDataLength
				}
			} else {
				// データの長さが足りない
				// 最初のコマンド含むけど、長さコマンド含まないとき
				return i + 2, 0, ErrLengthCommandNotFound
			}
		}
	}
	if data[len(data)-1] == 0x0f {
		return 0, 0, ErrStartCommandFoundOnlyFirst
	}
	// スタートコマンドが見つからない
	return 0, 0, ErrStartCommandNotFound
}

// データを解析する関数
// 来るデータは長さコマンドから始まる
func (a *App) ParseData(data []byte) {
	if len(data) < 11 {
		return
	}
	header := data[11]
	/// log.Printf("Header: 0x%02x\n", header)
	if header != 0x41 || data[0] != 0xb7 {
		if header == 0x50 && data[0] == 0x17 {
			// 電圧データなし
			// 電圧用のコード
			// model.HUB.SendText("Voltage: " + byteArrayToString(data[offset:offset+9])) // offset=12
			// log.Println(data[12:21])
			// a.VoltageToFile(data[12:21])
			return
		} else if header == 0x51 && data[0] == 0x17 {
			// 電圧データあり
			// model.HUB.SendText("Voltage: " + byteArrayToString(data[offset:offset+9])) // offset=12
			// log.Println(data[12:21])
			a.VoltageToFile(data[12:21])
			return
		}
		return
	}

	offset := 12
	for i := 0; i < 8; i++ {
		time := data[offset : offset+4]
		offset += 4

		quat := data[offset : offset+4*4]
		offset += 4 * 4

		// log.Printf("Data %d: %d, %d\n", i+1, time, quat)
		a.QuatAndTimeToFile(time, quat)
	}

	lpsTime := data[offset : offset+4]
	offset += 4
	// log.Printf("lpsTime: %d\n", lpsTime)

	pressure := data[offset : offset+3]
	offset += 3
	// log.Printf("pressure: %d\n", pressure)
	a.LpsAndTimeToFile(lpsTime, pressure)

	openRate := data[offset : offset+2]
	// offset += 2
	// log.Printf("openRate: %d\n", openRate)
	a.OpenAndTimeToFile(lpsTime, openRate)
}
