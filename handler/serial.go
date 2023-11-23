package handler

import (
	"context"
	"log"
	"os"

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

	var err error
	a.rawFile, err = os.OpenFile(a.rawFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	defer a.rawFile.Close()

	a.quatFile, err = os.OpenFile(a.quatFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	defer a.quatFile.Close()

	a.lpsFile, err = os.OpenFile(a.lpsFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	defer a.lpsFile.Close()

	a.openFile, err = os.OpenFile(a.openFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	defer a.openFile.Close()

	a.voltFile, err = os.OpenFile(a.voltFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		return
	}
	defer a.voltFile.Close()

	receivedData := make(chan []byte)
	// バックグラウンドでデータを受信し解析するゴルーチン
	go a.ReceiveData(receivedData)

	// Read and print the response
	buff := make([]byte, 512)
	for {
		n, err := model.Port.Read(buff)
		if err != nil {
			log.Println(err)
			model.HUB.SendError(err.Error())
			// return
		}
		// log.Printf("buffer length: %d\n", n)

		// model.HUB.SendText("Serial::" + byteArrayToString(buff[:n]))

		// if n == 0 {
		// 	HUB.SendText("Serial:"+"\nEOF")
		// 	break
		// }
		// log.Println("oooooooooooooooooooooooooooo")
		// log.Println(buff)

		// jjj := findData(buff[:n])
		// parseData(jjj)
		// model.HUB.SendText("Serial::" + byteArrayToString(jjj))
		// log.Print(buff[:n])

		// If we receive a newline stop reading
		// if strings.Contains(string(buff[:n]), "\n") {
		// 	break
		// }

		receivedData <- buff[:n]
		// 生データの保存
		// err = model.AppendToFile(buff[:n], a.rawFileName)
		err = model.AppendToFileDirect(buff[:n], a.rawFile)
		if err != nil {
			log.Println(err)
			model.HUB.SendError(err.Error())
			// return
		}

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
			model.HUB.SendError(err.Error())
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
			model.HUB.SendError(err.Error())
			return
		}
	}
	// Send string to the serial port
	_, err := model.Port.Write(byteDate)
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
		// if err ==  {

		// }
	}
	/// log.Printf("Sent %v bytes\n", n)
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
