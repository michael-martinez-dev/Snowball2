package debt

// Debt struct
type Debt struct {
	debtItems []GoBill
	store     *store
}

// NewDebt creates a new Debt struct
func NewDebt(storeFile string) *Debt {
	storage := NewDebtStore(storeFile, "sqlite")
	return &Debt{
		debtItems: Load(storage),
		store:     storage,
	}
}

// CreateDebtItem creates a new debt item
func (d *Debt) CreateDebtItem(id, due int, total, monthlyMin, monthlyActual, interest, name, debtType string) {
	d.debtItems = append(d.debtItems,
		*NewDebtItem(
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
	Save(d.store, d.debtItems)
}

// RetrieveDebts retrieves all debt items
func (d *Debt) RetrieveDebts() []GoBill {
	d.debtItems = Load(d.store)
	return d.debtItems
}

// UpdateDebtItem updates a debt item by id
func (d *Debt) UpdateDebtItem(id, due int, total, monthlyMin, monthlyActual, interest, name, debtType string) {
	for i, debtItem := range d.debtItems {
		if debtItem.ID == id {
			d.debtItems[i] = *NewDebtItem(
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
	Save(d.store, d.debtItems)
}

// DeleteDebtItem deletes a debt item by id
func (d *Debt) DeleteDebtItem(id int) {
	for i, debtItem := range d.debtItems {
		if debtItem.ID == id {
			d.debtItems = append(d.debtItems[:i], d.debtItems[i+1:]...)
		}
	}
	Delete(d.store, id)
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
