package debt

type Debt struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	DebtType      string `json:"type"`
	Total         string `json:"total"`
	MonthlyMin    string `json:"monthlyMin"`
	MonthlyActual string `json:"monthlyActual"`
	Interest      string `json:"interest"`
	DueDay        int    `json:"dueDay"`
}
