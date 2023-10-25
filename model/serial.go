package model

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

var Port serial.Port

func SerialInit(selectedPort string) (string, error) {
	// Retrieve the port list
	if Port != nil {
		Port.Close()
	}
	ports, err := serial.GetPortsList()
	// ports2, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Println(err)
		HUB.SendError(err.Error())
		return "", err
	}
	if len(ports) == 0 {
		HUB.SendText("Serial:" + "No serial ports found!")
		return "", fmt.Errorf("No serial ports found!")
	}
	log.Println("Found ports:", len(ports))
	HUB.SendText("Serial:" + "Found ports:" + string(len(ports)))
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
	if selectedPort == "" {
		for _, port := range ports {
			Port, err = serial.Open(port, mode)
			if err != nil {
				log.Println(err)
				HUB.SendError(err.Error())
				HUB.SendText(port)
				if portErr, ok := err.(serial.PortError); ok {
					if portErr.Code() == serial.PortBusy {
						log.Printf("Port %s busy\n", port)
						HUB.SendText("Serial:" + "Port " + port + " busy")
						continue
					}
				}
			} else {
				log.Printf("Opened port %s\n", port)
				HUB.SendText("Serial:" + "Opened port " + port)
				return port, nil
			}
		}
	} else {
		Port, err = serial.Open(selectedPort, mode)
		if err != nil {
			log.Println(err)
			HUB.SendError(err.Error())
			HUB.SendText(selectedPort)
			// if portErr, ok := err.(*serial.PortError); ok {
			// 	if portErr.Code() == serial.PortBusy {
			// 		log.Printf("Port %s busy\n", selectedPort)
			// 		HUB.SendText("Serial:" + "Port " + selectedPort + " busy")
			// 	}
			// }
		} else {
			log.Printf("Opened port %s\n", selectedPort)
			HUB.SendText("Serial:" + "Opened port " + selectedPort)
			return selectedPort, nil
		}
	}
	return "", fmt.Errorf("No serial ports found!")
}
