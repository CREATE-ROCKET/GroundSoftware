package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"go.bug.st/serial"
)

// App struct
type App struct {
	ctx context.Context
}

var Conn *websocket.Conn

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

var port serial.Port

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
}

var addr = flag.String("addr", ":3007", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// http.ServeFile(w, r, "home.html")
}

func (a *App) SerialMonitor() {
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

type Serial struct{}

func (a *App) SerialStop() {
	a.ctx = context.WithValue(a.ctx, Serial{}, "stop")
}
