package unit

import (
	"DebtSnowball2/backend/debt"
	"os"
	"testing"
)

// TestStore tests the store
func TestStore(t *testing.T) {
	var store = debt.NewDebtStore("test", "json")
	if store.Name != "test.json" {
		t.Errorf("Expected store name to be test.json, got %s", store.Name)
	}
	if store.Type != "json" {
		t.Errorf("Expected store type to be json, got %s", store.Type)
	}
	store = debt.NewDebtStore("test", "sqlite")
	if store.Name != "test.db" {
		t.Errorf("Expected store name to be test.db, got %s", store.Name)
	}
	if store.Type != "sqlite" {
		t.Errorf("Expected store type to be sqlite, got %s", store.Type)
	}
	store = debt.NewDebtStore("test", "memory")
	if store.Name != "test" {
		t.Errorf("Expected store name to be test, got %s", store.Name)
	}
	if store.Type != "memory" {
		t.Errorf("Expected store type to be memory, got %s", store.Type)
	}
}

// TestLoadGivenJsonType tests the load function given a json type
func TestLoadGivenJsonType(t *testing.T) {
	var store = debt.NewDebtStore("test", "json")
	var debts = debt.Load(store)
	if len(debts) != 0 {
		t.Errorf("Expected debts to be empty, got %d", len(debts))
	}

	// clean up
	home, _ := os.UserHomeDir()
	err := os.Remove(home + "/test.json")
	if err != nil {
		t.Errorf("Error removing test.json")
	}
}

// TestLoadGivenSqliteType tests the load function given a sqlite type
func TestLoadGivenSqliteType(t *testing.T) {
	var store = debt.NewDebtStore("test", "sqlite")
	var debts = debt.Load(store)
	if len(debts) != 0 {
		t.Errorf("Expected debts to be empty, got %d", len(debts))
	}

	// clean up
	home, _ := os.UserHomeDir()
	err := os.Remove(home + "/test.db")
	if err != nil {
		t.Errorf("Error removing test.db")
	}
}

// TestLoadGivenMemoryType tests the load function given a memory type
func TestLoadGivenMemoryType(t *testing.T) {
	var store = debt.NewDebtStore("test", "memory")
	var debts = debt.Load(store)
	if len(debts) != 0 {
		t.Errorf("Expected debts to be empty, got %d", len(debts))
	}
}

// TestSaveGivenJsonType tests the save function given a json type
func TestSaveGivenJsonType(t *testing.T) {
	var store = debt.NewDebtStore("test", "json")
	var debts = debt.Load(store)
	debts = append(
		debts,
		*debt.NewDebtItem(
			1,
			1,
			"10000",
			"100",
			"150",
			"1.0",
			"test",
			"Test"))
	debt.Save(store, debts)
	debts = debt.Load(store)
	if len(debts) != 1 {
		t.Errorf("Expected debts to have 1 item, got %d", len(debts))
	}

	// clean up
	home, _ := os.UserHomeDir()
	err := os.Remove(home + "/test.json")
	if err != nil {
		t.Errorf("Error removing test.json")
	}
}

// TestSaveGivenSqliteType tests the save function given a sqlite type
func TestSaveGivenSqliteType(t *testing.T) {
	var store = debt.NewDebtStore("test", "sqlite")
	var debts = debt.Load(store)
	debts = append(
		debts,
		*debt.NewDebtItem(
			1,
			1,
			"10000",
			"100",
			"150",
			"1.0",
			"test",
			"Test"))
	debt.Save(store, debts)
	debts = debt.Load(store)
	if len(debts) != 1 {
		t.Errorf("Expected debts to have 1 item, got %d", len(debts))
	}

	// clean up
	home, _ := os.UserHomeDir()
	err := os.Remove(home + "/test.db")
	if err != nil {
		t.Errorf("Error removing test.db")
	}
}

// TestSaveGivenMemoryType tests the save function given a memory type
func TestSaveGivenMemoryType(t *testing.T) {
	var store = debt.NewDebtStore("test", "memory")
	var debts = debt.Load(store)
	debts = append(
		debts,
		*debt.NewDebtItem(
			1,
			1,
			"10000",
			"100",
			"150",
			"1.0",
			"test",
			"Test"))
	debt.Save(store, debts)
	debts = debt.Load(store)
	if len(debts) != 1 {
		t.Errorf("Expected debts to have 1 item, got %d", len(debts))
	}
}
