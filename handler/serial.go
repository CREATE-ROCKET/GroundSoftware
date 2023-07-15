package handler

import (
	"context"
	"log"

	"github.com/Luftalian/Computer_software/model"
)

type Serial struct{}

func (a *App) SerialStart() {
	a.ctx = context.WithValue(a.ctx, Serial{}, "start")
	if model.Port == nil {
		model.SerialInit()
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
		model.SerialInit()
	}
	// Send string to the serial port
	n, err := model.Port.Write([]byte(text))
	if err != nil {
		log.Println(err)
		model.HUB.SendError(err.Error())
	}
	log.Printf("Sent %v bytes\n", n)
}
