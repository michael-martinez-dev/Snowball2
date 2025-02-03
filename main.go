package main

import (
	"DebtSnowball2/backend/app"
	"DebtSnowball2/backend/config"
	"DebtSnowball2/backend/debt"
	"DebtSnowball2/backend/logger"
	"DebtSnowball2/backend/notes"
	"DebtSnowball2/backend/store"
	"path/filepath"

	"embed"

	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger.InitLogger("Snowball2.log")

	initialConfig := &config.Config{
		DBType:    "sqlite",
		DBPath:    filepath.Join("G:\\", "My Drive", "Personal", "Debts"),
		DBFile:    "debts.db",
		NotesPath: filepath.Join("G:\\", "My Drive", "Personal", "Debts"),
	}

	cfgService := config.NewConfigService(initialConfig)

	debtStore, err := store.BuildDebtStore(
		initialConfig.DBType,
		filepath.Join(initialConfig.DBPath, initialConfig.DBFile),
	)
	if err != nil {
		log.Fatalf("Failed to build debt store: %v", err)
	}

	debtService := debt.NewDebtService(debtStore)
	notesService := notes.NewNotesService(initialConfig.NotesPath)

	app := app.NewAppManager(cfgService, debtService)

	log.Infoln("Setting up Frontend...")
	err = wails.Run(&options.App{
		Title:  "Debt Snowball 2",
		Width:  1050,
		Height: 550,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			debtService,
			notesService,
			cfgService,
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
