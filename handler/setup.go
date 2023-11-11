package handler

import (
	"context"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	log.Print("-------------------------------------------sssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
	a.ModuleEnv()
}

func (a *App) Shutdown(ctx context.Context) {
}

func (a *App) Domready(ctx context.Context) {
}

func (a *App) BeforeClose(ctx context.Context) bool {
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Quit?",
		Message: "Are you sure you want to quit?",
	})

	if err != nil {
		return false
	}
	return dialog != "Yes"
}
