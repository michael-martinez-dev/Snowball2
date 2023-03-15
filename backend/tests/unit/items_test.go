package unit

import (
	"DebtSnowball2/backend/debt"
	"testing"
)

// TestNewDebtItem tests the new debt item
func TestNewDebtItem(t *testing.T) {
	var id, due int = 1, 1
	var total, monthlyMin, monthlyActual, interest string = "1000.00", "100.00", "150.00", "0.00"
	var name, debtType string = "Test", "Test"
	var debtItem = debt.NewDebtItem(id, due, total, monthlyMin, monthlyActual, interest, name, debtType)
	if debtItem.ID != id {
		t.Errorf("Expected id to be %d, got %d", id, debtItem.ID)
	}
	if debtItem.DueDay != due {
		t.Errorf("Expected due to be %d, got %d", due, debtItem.DueDay)
	}
	if debtItem.Total != total {
		t.Errorf("Expected total to be %s, got %s", total, debtItem.Total)
	}
	if debtItem.MonthlyMin != monthlyMin {
		t.Errorf("Expected monthlyMin to be %s, got %s", monthlyMin, debtItem.MonthlyMin)
	}
	if debtItem.MonthlyActual != monthlyActual {
		t.Errorf("Expected monthlyActual to be %s, got %s", monthlyActual, debtItem.MonthlyActual)
	}
	if debtItem.Interest != interest {
		t.Errorf("Expected interest to be %s, got %s", interest, debtItem.Interest)
	}
	if debtItem.Name != name {
		t.Errorf("Expected name to be %s, got %s", name, debtItem.Name)
	}
	if debtItem.DebtType != debtType {
		t.Errorf("Expected debtType to be %s, got %s", debtType, debtItem.DebtType)
	}
}
