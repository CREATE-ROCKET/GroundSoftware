package handler

import (
	"context"
	"log"

	"go.bug.st/serial"
)

var port serial.Port

type Serial struct{}

func (a *App) SerialStart() {
	a.ctx = context.WithValue(a.ctx, Serial{}, "start")
	if port == nil {
		// Retrieve the port list
		ports, err := serial.GetPortsList()
		// ports2, err := enumerator.GetDetailedPortsList()
		if err != nil {
			log.Println(err)
			return
		}
		if len(ports) == 0 {
			HUB.SendText("Serial:" + "No serial ports found!")
			return
		}
		// Open the first serial port detected at 115200bps N81
		mode := &serial.Mode{
			BaudRate: 115200,
			Parity:   serial.NoParity,
			DataBits: 8,
			StopBits: serial.OneStopBit,
		}
		// for _, port := range ports2 {
		// 	fmt.Printf("Found port: %s\n", port.Name)
		// 	if port.IsUSB {
		// 		log.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
		// 		log.Printf("   USB serial %s\n", port.SerialNumber)
		// 	}
		// }
		port, err = serial.Open(ports[0], mode)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Read and print the response
	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Println(err)
			return
		}
		// if n == 0 {
		// 	HUB.SendText("Serial:"+"\nEOF")
		// 	break
		// }
		HUB.SendText("Serial::" + string(buff[:n]))
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
