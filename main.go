package main

import (
	"DebtSnowball2/backend"
	"DebtSnowball2/backend/debt"
	"DebtSnowball2/backend/notes"
	"embed"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	println()

	initLogger()
	backend.InitProperties()

	log.Infoln("Setting up Backend...")
	snowball := debt.NewDebt("debt")
	snowballNotes := notes.NewNote("debt-notes", "")

	log.Infoln("Setting up Frontend...")
	err := wails.Run(&options.App{
		Title:  "Debt Snowball 2",
		Width:  1050,
		Height: 550,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			snowball,
			snowballNotes,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func initLogger() {
	logFile, err := os.Create("Snowball2.log")
	if err != nil {
		log.Warnln("Could not create log file")
		return
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)
}
