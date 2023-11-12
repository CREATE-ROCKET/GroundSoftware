package handler

import (
	"context"
)

// App struct
type App struct {
	ctx          context.Context
	rawFileName  string
	quatFileName string
	lpsFileName  string
	openFileName string
	voltFileName string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}
