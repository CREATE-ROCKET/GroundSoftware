package handler

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/menu"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

var Addr = flag.String("addr", ":3007", "http service address")

func ServeHome(w http.ResponseWriter, r *http.Request) {
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

func OpenFile(*menu.CallbackData) {
}
