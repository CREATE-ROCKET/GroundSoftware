package handler

import (
	"context"
	"os"
)

// App struct
type App struct {
	ctx          context.Context
	rawFileName  string
	quatFileName string
	lpsFileName  string
	openFileName string
	voltFileName string
	rawFile      *os.File
	quatFile     *os.File
	lpsFile      *os.File
	openFile     *os.File
	voltFile     *os.File
	timeFile string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}
