package model

import (
	"log"

	"go.bug.st/serial"
)

var Port serial.Port

func SerialInit() {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	// ports2, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Println(err)
		HUB.SendError(err.Error())
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
	Port, err = serial.Open(ports[0], mode)
	if err != nil {
		log.Println(err)
		HUB.SendError(err.Error())
		return
	}
}
