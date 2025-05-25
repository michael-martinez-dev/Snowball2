package store

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDebtStore struct {
	db *gorm.DB
}

func NewSQLiteStore(db *gorm.DB) DebtStore {
	return &sqliteDebtStore{
		db: db,
	}
}

func BuildSqliteStore(dbPath string) (DebtStore, error) {
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
func (s *sqliteDebtStore) Create(d *DebtRecord) error {
	return s.db.Create(d).Error
}

func (s *sqliteDebtStore) Update(d *DebtRecord) error {
	return s.db.Save(d).Error
}

func (s *sqliteDebtStore) Delete(id int) error {
	res := s.db.Delete(&DebtRecord{}, id)
	return res.Error
}

func (s *sqliteDebtStore) GetAll() ([]DebtRecord, error) {
	var debts []DebtRecord
	err := s.db.Find(&debts).Error
	return debts, err
}

func (s *sqliteDebtStore) Get(id int) (*DebtRecord, error) {
	var bill DebtRecord
	err := s.db.First(&bill, id).Error
	if err != nil {
		return nil, err
	}
	return &bill, nil
}
