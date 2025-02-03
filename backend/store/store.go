package store

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// DebtRecord is a representation of a debt
type DebtRecord struct {
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name"`
	DebtType      string `json:"type"`
	Total         string `json:"total"`
	MonthlyMin    string `json:"monthlyMin"`
	MonthlyActual string `json:"monthlyActual"`
	Interest      string `json:"interest"`
	DueDay        int    `json:"dueDay"`
}

// DebtStore is an interface for storing and retrieving debt
type DebtStore interface {
	Create(debt *DebtRecord) error
	Update(debt *DebtRecord) error
	Delete(id int) error
	Get(id int) (*DebtRecord, error)
	GetAll() ([]DebtRecord, error)
}

func BuildDebtStore(dbType, dbPath string) (DebtStore, error) {
	switch dbType {
	case "sqlite":
		return buildSqliteStore(dbPath)
	case "json":
		return NewJSONStore(dbPath), nil
	case "memory":
		return NewMemoryStore(), nil
	default:
		return nil, fmt.Errorf("unknown db type %s", dbType)
	}
}

func buildSqliteStore(dbPath string) (DebtStore, error) {
	home, _ := os.UserHomeDir()

	if !filepath.IsAbs(dbPath) {
		dbPath = filepath.Join(home, dbPath)
	}

	log.Infof("Openning sqlite db at %s", dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.AutoMigrate(&DebtRecord{}); err != nil {
		return nil, fmt.Errorf("failed to migrate db: %w", err)
	}

	log.Infof("Using SQLite store at %s", dbPath)

	return NewSQLiteStore(db), nil
}
