package app

import (
	"DebtSnowball2/backend/config"
	"DebtSnowball2/backend/debt"
	"DebtSnowball2/backend/store"

	"fmt"

	log "github.com/sirupsen/logrus"
)

type AppManager struct {
	cfgService  *config.ConfigService
	debtService *debt.DebtService
}

func NewAppManager(cfgService *config.ConfigService, debtService *debt.DebtService) *AppManager {
	return &AppManager{
		cfgService:  cfgService,
		debtService: debtService,
	}
}

func (app *AppManager) SwitchDebtStore(
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
