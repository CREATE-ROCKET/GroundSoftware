package handler

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	log.Print("-------------------------------------------sssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
	a.ModuleEnv()

	currentTime := time.Now()
	timestamp := currentTime.Format("2006-01-02-15-04-05") // フォーマット例: 2023-10-26-14-30-00

	if _, err := os.Stat("data"); err == nil {
	} else if os.IsNotExist(err) {
		err = os.Mkdir("data", 0755)
		if err != nil {
			return
		}
	} else {
		return
	}

	a.rawFileName = "data/" + timestamp + "/raw_" + timestamp + ".txt"
	err := makeDirAndFile("data/"+timestamp, "data/"+timestamp+"/raw_"+timestamp+".txt")
	if err != nil {
		log.Println(err)
	}

	a.quatFileName = "data/" + timestamp + "/quat_" + timestamp + ".txt"
	err = makeDirAndFile("data/"+timestamp, "data/"+timestamp+"/quat_"+timestamp+".txt")
	if err != nil {
		log.Println(err)
	}

	a.lpsFileName = "data/" + timestamp + "/lps_" + timestamp + ".txt"
	err = makeDirAndFile("data/"+timestamp, "data/"+timestamp+"/lps_"+timestamp+".txt")
	if err != nil {
		log.Println(err)
	}

	a.openFileName = "data/" + timestamp + "/open_" + timestamp + ".txt"
	err = makeDirAndFile("data/"+timestamp, "data/"+timestamp+"/open_"+timestamp+".txt")
	if err != nil {
		log.Println(err)
	}
	a.voltFileName = "data/" + timestamp + "/volt_" + timestamp + ".txt"
	err = makeDirAndFile("data/"+timestamp, "data/"+timestamp+"/volt_"+timestamp+".txt")
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

func makeDirAndFile(dir string, file string) error {
	if _, err := os.Stat(dir); err == nil {
	} else if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	rawFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer rawFile.Close()
	return nil
}
