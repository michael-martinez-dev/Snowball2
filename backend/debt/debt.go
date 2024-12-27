package debt

import (
	"DebtSnowball2/backend"
	"DebtSnowball2/backend/debt/models"
	"DebtSnowball2/backend/debt/store"
)

// Debt struct
type Debt struct {
	debtItems []models.NewBill
	dataStore *store.Store
}

// NewDebt creates a new Debt struct
func NewDebt(storeFile string) *Debt {
	database := getDatabaseToUse()
	storage := store.NewDebtStore(storeFile, database)
	return &Debt{
		debtItems: store.Get(storage),
		dataStore: storage,
	}
}

func getDatabaseToUse() string {
	return backend.DEFAULT_DB
}

// CreateDebtItem creates a new debt item
func (d *Debt) CreateDebtItem(id, due int, total, monthlyMin, monthlyActual, interest, name, debtType string) {
	d.debtItems = append(d.debtItems,
		*models.NewDebtItem(
			id,
			due,
			total,
			monthlyMin,
			monthlyActual,
			interest,
			name,
			debtType,
		),
	)
	store.Save(d.dataStore, d.debtItems)
}

// RetrieveDebts retrieves all debt items
func (d *Debt) RetrieveDebts() []models.NewBill {
	d.debtItems = store.Get(d.dataStore)
	return d.debtItems
}

// UpdateDebtItem updates a debt item by id
func (d *Debt) UpdateDebtItem(id, due int, total, monthlyMin, monthlyActual, interest, name, debtType string) {
	for i, debtItem := range d.debtItems {
		if debtItem.ID == id {
			d.debtItems[i] = *models.NewDebtItem(
				id,
				due,
				total,
				monthlyMin,
				monthlyActual,
				interest,
				name,
				debtType,
			)
		}
	}
	store.Save(d.dataStore, d.debtItems)
}

// DeleteDebtItem deletes a debt item by id
func (d *Debt) DeleteDebtItem(id int) {
	for i, debtItem := range d.debtItems {
		if debtItem.ID == id {
			d.debtItems = append(d.debtItems[:i], d.debtItems[i+1:]...)
		}
	}
	store.Delete(d.dataStore, id)
}

// findDebtItem finds a debt item by name and returns true if found
//func (d *Debt) findDebtItem(name string) bool {
//	found := false
//	for _, debtItem := range d.debtItems {
//		if debtItem.Name == name {
//			found = true
//		}
//	}
//	return found
//}
