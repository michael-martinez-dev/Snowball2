package store

import "sync"

type memoryStore struct {
	mu     sync.Mutex
	debts  []DebtRecord
	nextID int
}

// NewMemoryStore creates a new MemoryStore
func NewMemoryStore() DebtStore {
	return &memoryStore{
		debts:  []DebtRecord{},
		nextID: 1,
	}
}

func (m *memoryStore) Create(d *DebtRecord) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Generate an ID in a naive way
	maxID := 0
	for _, debt := range m.debts {
		if debt.ID > maxID {
			maxID = debt.ID
		}
	}
	d.ID = maxID + 1
	m.debts = append(m.debts, *d)
	return nil
}

func (m *memoryStore) Update(d *DebtRecord) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, bill := range m.debts {
		if bill.ID == d.ID {
			m.debts[i] = *d
			return nil
		}
	}
	return nil
}

func (m *memoryStore) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, bill := range m.debts {
		if bill.ID == id {
			m.debts = append(m.debts[:i], m.debts[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *memoryStore) GetAll() ([]DebtRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Return a copy
	copyItems := make([]DebtRecord, len(m.debts))
	copy(copyItems, m.debts)
	return copyItems, nil
}

func (m *memoryStore) Get(id int) (*DebtRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, bill := range m.debts {
		if bill.ID == id {
			b := bill
			return &b, nil
		}
	}
	return nil, nil
}
