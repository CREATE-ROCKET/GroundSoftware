package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Luftalian/Computer_software/model"
	"go.bug.st/serial"
)

type Serial struct{}

func (a *App) SerialStart() {
	a.ctx = context.WithValue(a.ctx, Serial{}, "start")
	// var port string
	if model.Port == nil {
		// port = model.SerialInit()
		_, err := model.SerialInit("")
		if err != nil {
			log.Println(err)
			// model.HUB.SendError(err.Error())
			return
		}
	}

	// Read and print the response
	buff := make([]byte, 100)
	for {
		n, err := model.Port.Read(buff)
		if err != nil {
			log.Println(err)
			model.HUB.SendError(err.Error())
			return
		}
		// if n == 0 {
		// 	HUB.SendText("Serial:"+"\nEOF")
		// 	break
		// }
		// log.Println("oooooooooooooooooooooooooooo")
		// log.Println(buff)
		jjj := findData(buff[:n])
		parseData(jjj)
		model.HUB.SendText("Serial::" + byteArrayToString(jjj))
		log.Print(buff[:n])
		// If we receive a newline stop reading
		// if strings.Contains(string(buff[:n]), "\n") {
		// 	break
		// }
		if a.ctx.Value(Serial{}) == "stop" {
			break
		}
	}
}

func (a *App) SerialStop() {
	a.ctx = context.WithValue(a.ctx, Serial{}, "stop")
}

func (a *App) SerialTextSend(text string) {
	if model.Port == nil {
		_, err := model.SerialInit("")
		if err != nil {
			log.Println(err)
			// model.HUB.SendError(err.Error())
			return
		}
	}
	// Send string to the serial port
	n, err := model.Port.Write([]byte(text))
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		// if err ==  {

		// }
	}
	log.Printf("Sent %v bytes\n", n)
}

func (a *App) SerialByteSend(byteDate []byte) {
	if model.Port == nil {
		_, err := model.SerialInit("")
		if err != nil {
			log.Println(err)
			// model.HUB.SendError(err.Error())
			return
		}
	}
	// Send string to the serial port
	n, err := model.Port.Write(byteDate)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		// if err ==  {

		// }
	}
	log.Printf("Sent %v bytes\n", n)
}

func (a *App) PortList() []string {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return nil
	}
	if len(ports) == 0 {
		model.HUB.SendText("Serial:" + "No serial ports found!")
		return nil
	}
	log.Println("Found ports:", len(ports))
	return ports
}

func (a *App) SelectedPort(port string) {
	model.SerialInit(port)
}

func byteArrayToString(bytes []byte) string {
	result := "["
	for i, b := range bytes {
		result += strconv.Itoa(int(b))
		if i < len(bytes)-1 {
			result += " "
		}
	}
	result += "]"
	return result
}

dataLength := a

func findData(input []byte) []byte {
	index := findStartIndex(input)
	if index == -1 {
		log.Println("Not found in input: 0x0f, 0x5a")
		model.HUB.SendError("Not found in input: 0x0f, 0x5a")
		return input
	}

	// 0xbbの値を取得してデータを読み取る
	dataLength := int(input[index+2])
	if dataLength > len(input) {
		log.Println("Data length is too long")
		model.HUB.SendError("Data length is too long")
		return input
	}
	data := input[index+3 : index+3+dataLength]
	return data
}

// 0x0f, 0x5aを探す関数
func findStartIndex(input []byte) int {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == 0x0f && input[i+1] == 0x5a {
			return i
		}
	}
	return -1
}

// データを解析する関数
func parseData(data []byte) {
	header := data[0]
	fmt.Printf("Header: 0x%02x\n", header)

	offset := 1
	for i := 0; i < 8; i++ {
		time := data[offset : offset+4]
		offset += 4

		accel := data[offset : offset+4]
		offset += 4

		fmt.Printf("Data %d: %d, %d\n", i+1, time, accel)

	}

	lpsTime := data[offset : offset+4]
	offset += 4
	fmt.Printf("lpsTime: %d\n", lpsTime)

	pressure := data[offset : offset+3]
	offset += 3
	fmt.Printf("pressure: %d\n", pressure)

	openRate := data[offset : offset+2]
	offset += 2
	fmt.Printf("openRate: %d\n", openRate)

	undefined := data[offset : offset+4]
	fmt.Printf("undefined: %d\n", undefined)
}
