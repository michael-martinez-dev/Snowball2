package store

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type jsonDebtStore struct {
	filePath string
	mu       sync.Mutex
}

// NewJSONDebtStore creates a new JSONDebtStore
func NewJSONStore(filePath string) DebtStore {
	return &jsonDebtStore{
		filePath: filePath,
	}
}

func (s *jsonDebtStore) Create(d *DebtRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	debts, err := s.readAll()
	if err != nil {
		return err
	}
	// Generate an ID if none provided (simple approach)
	maxID := 0
	for _, debt := range debts {
		if debt.ID > maxID {
			maxID = debt.ID
		}
	}
	d.ID = maxID + 1
	debts = append(debts, *d)
	return s.writeAll(debts)
}

func (s *jsonDebtStore) Update(d *DebtRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	debts, err := s.readAll()
	if err != nil {
		return err
	}
	for i, bill := range debts {
		if bill.ID == d.ID {
			debts[i] = *d
			return s.writeAll(debts)
		}
	}
	return fmt.Errorf("debt with ID %d not found", d.ID)
}

func (s *jsonDebtStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	debts, err := s.readAll()
	if err != nil {
		return err
	}
	newList := make([]DebtRecord, 0, len(debts))
	found := false
	for _, bill := range debts {
		if bill.ID == id {
			found = true
			continue
		}
		newList = append(newList, bill)
	}
	if !found {
		return fmt.Errorf("debt with ID %d not found", id)
	}
	return s.writeAll(newList)
}

func (s *jsonDebtStore) GetAll() ([]DebtRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.readAll()
}

func (s *jsonDebtStore) Get(id int) (*DebtRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	debts, err := s.readAll()
	if err != nil {
		return nil, err
	}
	for _, bill := range debts {
		if bill.ID == id {
			return &bill, nil
		}
	}
	return nil, fmt.Errorf("debt with ID %d not found", id)
}

// Helper: readAll from file
func (s *jsonDebtStore) readAll() ([]DebtRecord, error) {
	// Ensure the file exists
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []DebtRecord{}, nil
	}

	bytes, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}
	var debts []DebtRecord
	if len(bytes) == 0 {
		return []DebtRecord{}, nil
	}
	err = json.Unmarshal(bytes, &debts)
	return debts, err
}

// Helper: writeAll to file
func (s *jsonDebtStore) writeAll(debts []DebtRecord) error {
	data, err := json.MarshalIndent(debts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0600)
}
