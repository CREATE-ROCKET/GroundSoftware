package handler

import (
	"context"
	"log"

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
		model.HUB.SendText("Serial::" + string(buff[:n]))
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

func (a *App) SerialSend(text string) {
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
