package store

import (
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

func (s *sqliteDebtStore) Create(d *GoBill) error {
	return s.db.Create(d).Error
}

func (s *sqliteDebtStore) Update(d *GoBill) error {
	return s.db.Save(d).Error
}

func (s *sqliteDebtStore) Delete(id int) error {
	res := s.db.Delete(&GoBill{}, id)
	return res.Error
}

func (s *sqliteDebtStore) GetAll() ([]GoBill, error) {
	var debts []GoBill
	err := s.db.Find(&debts).Error
	return debts, err
}

func (s *sqliteDebtStore) Get(id int) (*GoBill, error) {
	var bill GoBill
	err := s.db.First(&bill, id).Error
	if err != nil {
		return nil, err
	}
	return &bill, nil
}
