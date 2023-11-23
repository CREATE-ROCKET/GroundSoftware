package main

import (
	"embed"
	"flag"
	"net/http"

	"github.com/Luftalian/Computer_software/handler"
	"github.com/Luftalian/Computer_software/model"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := handler.NewApp()

	// websocket server
	flag.Parse()
	hub := model.NewHub()
	model.HUB = hub
	go handler.Run(hub)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(hub, w, r)
	})
	go http.ListenAndServe(*model.Addr, nil)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "CREATE Serial Monitor",
		Width:  10240,
		Height: 7680,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:             app.ApplicationMenu(),
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		OnDomReady:       app.Domready,
		OnShutdown:       app.Shutdown,
		OnBeforeClose:    app.BeforeClose,
		Bind: []interface{}{
			app,
			&model.Hub{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
