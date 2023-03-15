package debt

// GoBill struct for debt items
type GoBill struct {
	ID            int    `json:"id" gorm:"primary_key;auto_increment;unique;index;"`
	DueDay        int    `json:"dueDay"`
	Name          string `json:"name"`
	DebtType      string `json:"type"`
	Total         string `json:"total"`
	MonthlyMin    string `json:"monthlyMin"`
	MonthlyActual string `json:"monthlyActual"`
	Interest      string `json:"interest"`
}

type OldDebtItem struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Link     string `json:"link"`
	Total    string `json:"total"`
	Monthly  string `json:"monthly"`
	Due      string `json:"due"`
	Interest string `json:"interest"`
}

// NewDebtItem creates a new DebtItem
func NewDebtItem(
	id, due int,
	total, monthlyMin, monthlyActual, interest,
	name, debtType string) *GoBill {
	return &GoBill{
		ID:            id,
		Name:          name,
		DebtType:      debtType,
		Total:         total,
		MonthlyMin:    monthlyMin,
		MonthlyActual: monthlyActual,
		DueDay:        due,
		Interest:      interest,
	}
}
