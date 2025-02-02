package debt

import (
	"DebtSnowball2/backend/store"

	log "github.com/sirupsen/logrus"
)

// DebtService implements business logic for debt
type DebtService struct {
	repo store.DebtStore
}

// NewDebtService creates a new DebtService
func NewDebtService(repo store.DebtStore) *DebtService {
	return &DebtService{
		repo: repo,
	}
}

func (s *DebtService) SetStore(newRepo store.DebtStore) {
	s.repo = newRepo
	log.Infof("DebtStore switched at runtime.")
}

func (s *DebtService) MigrateTo(newStore store.DebtStore) error {
	oldDebts, err := s.GetAllDebts()
	if err != nil {
		return err
	}

	for _, d := range oldDebts {
		gb := toGoBill(d)
		if err := newStore.Create(&gb); err != nil {
			return err
		}
	}
	log.Infof("Migrated %d debts to new store.", len(oldDebts))
	return nil
}

func (s *DebtService) CreateDebt(d Debt) error {
	log.Debugf("Creating new debt: %+v", d)
	gb := toGoBill(d)
	return s.repo.Create(&gb)
}

func (s *DebtService) GetAllDebts() ([]Debt, error) {
	bills, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return toDomainList(bills), nil
}

func (s *DebtService) GetDebtByID(id int) (*Debt, error) {
	bill, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return toDomain(*bill), nil
}

func (s *DebtService) UpdateDebt(d Debt) error {
	gb := toGoBill(d)
	return s.repo.Update(&gb)
}

func (s *DebtService) DeleteDebt(id int) error {
	return s.repo.Delete(id)
}

func toGoBill(d Debt) store.GoBill {
	return store.GoBill{
		ID:            d.ID,
		Name:          d.Name,
		DebtType:      d.DebtType,
		Total:         d.Total,
		MonthlyMin:    d.MonthlyMin,
		MonthlyActual: d.MonthlyActual,
		Interest:      d.Interest,
		DueDay:        d.DueDay,
	}
}

func toDomain(gb store.GoBill) *Debt {
	return &Debt{
		ID:            gb.ID,
		Name:          gb.Name,
		DebtType:      gb.DebtType,
		Total:         gb.Total,
		MonthlyMin:    gb.MonthlyMin,
		MonthlyActual: gb.MonthlyActual,
		Interest:      gb.Interest,
		DueDay:        gb.DueDay,
	}
}

func toDomainList(gbs []store.GoBill) []Debt {
	var debts []Debt
	for _, gb := range gbs {
		debts = append(debts, *toDomain(gb))
	}
	return debts
}
