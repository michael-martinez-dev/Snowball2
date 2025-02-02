package main

import (
	"DebtSnowball2/backend/config"
	"DebtSnowball2/backend/debt"
	"DebtSnowball2/backend/logger"
	"DebtSnowball2/backend/notes"
	"DebtSnowball2/backend/store"
	"fmt"
	"path/filepath"

	"embed"

	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
	cfgService  *config.ConfigService
	debtService *debt.DebtService
}

func NewApp(cfgService *config.ConfigService, debtService *debt.DebtService) *App {
	return &App{
		cfgService:  cfgService,
		debtService: debtService,
	}
}

func (app *App) SwitchDebtStore(
	migrateData bool,
	newDBType, newDBPath string,
) error {
	newStore, err := store.BuildDebtStore(newDBType, newDBPath)
	if err != nil {
		return fmt.Errorf("failed to build new store: %w", err)
	}

	if migrateData {
		if err := app.debtService.MigrateTo(newStore); err != nil {
			return fmt.Errorf("failed to migrate data: %w", err)
		}
	}

	app.debtService.SetStore(newStore)

	current := app.cfgService.GetConfig()
	current.DBType = newDBType
	current.DBPath = newDBPath
	app.cfgService.UpdateConfig(current)

	log.Infof("Switched to %s store at %s", newDBType, newDBPath)
	return nil
}

func main() {
	logger.InitLogger("Snowball2.log")

	initialConfig := &config.Config{
		DBType:    "sqlite",
		DBPath:    filepath.Join("G:\\", "My Drive", "Personal", "Debts", "Debts.sqlite"),
		NotesPath: filepath.Join("G:\\", "My Drive", "Personal", "Debts"),
	}

	cfgService := config.NewConfigService(initialConfig)

	debtStore, err := store.BuildDebtStore(initialConfig.DBType, initialConfig.DBPath)
	if err != nil {
		log.Fatalf("Failed to build debt store: %v", err)
	}

	debtService := debt.NewDebtService(debtStore)
	notesService := notes.NewNotesService(initialConfig.NotesPath)

	app := NewApp(cfgService, debtService)

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
