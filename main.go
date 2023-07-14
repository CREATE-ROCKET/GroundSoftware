package main

import (
	"embed"
	"flag"
	"net/http"

	"github.com/Luftalian/Computer_software/handler"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := handler.NewApp()

	flag.Parse()
	hub := handler.NewHub()
	handler.HUB = hub
	go hub.Run()
	http.HandleFunc("/", handler.ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(hub, w, r)
	})

	go http.ListenAndServe(*handler.Addr, nil)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		Bind: []interface{}{
			app,
			&handler.Hub{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
