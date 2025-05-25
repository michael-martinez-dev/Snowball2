package store

import (
	"fmt"
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
		return BuildSqliteStore(dbPath)
	case "json":
		return NewJSONStore(dbPath), nil
	case "memory":
		return NewMemoryStore(), nil
	case "pocketbase":
		return NewPocketbaseStore(dbPath), nil
	default:
		return nil, fmt.Errorf("unknown db type %s", dbType)
	}
}
