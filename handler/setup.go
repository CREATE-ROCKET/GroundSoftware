package handler

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	log.Print("-------------------------------------------sssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
	a.ModuleEnv()

	timestamp := time.Now().Format("2006-01-02-15-04-05") // フォーマット例: 2023-10-26-14-30-00

	_, err := os.Stat("data")
	if err == nil {
	} else if os.IsNotExist(err) {
		err = os.Mkdir("data", 0755)
		if err != nil {
			return
		}
	} else {
		return
	}

	a.rawFileName = filepath.Join("data", timestamp, "raw_"+timestamp+".txt")
	a.rawFile, err = makeDirAndFile(filepath.Join("data", timestamp), a.rawFileName)
	if err != nil {
		log.Println(err)
	}

	a.quatFileName = filepath.Join("data", timestamp, "quat_"+timestamp+".txt")
	a.quatFile, err = makeDirAndFile(filepath.Join("data", timestamp), a.quatFileName)
	if err != nil {
		log.Println(err)
	}

	a.lpsFileName = filepath.Join("data", timestamp, "lps_"+timestamp+".txt")
	a.lpsFile, err = makeDirAndFile(filepath.Join("data", timestamp), a.lpsFileName)
	if err != nil {
		log.Println(err)
	}

	a.openFileName = filepath.Join("data", timestamp, "open_"+timestamp+".txt")
	a.openFile, err = makeDirAndFile(filepath.Join("data", timestamp), a.openFileName)
	if err != nil {
		log.Println(err)
	}
	a.voltFileName = filepath.Join("data", timestamp, "volt_"+timestamp+".txt")
	a.voltFile, err = makeDirAndFile(filepath.Join("data", timestamp), a.voltFileName)
	if err != nil {
		log.Println(err)
	}
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

func makeDirAndFile(dir string, file string) (*os.File, error) {
	if _, err := os.Stat(dir); err == nil {
	} else if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	rawFile, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	defer rawFile.Close()
	return rawFile, nil
}
