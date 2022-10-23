package dealershipPorts

import "github.com/brenos/qap/internal/core/domain"

type DealershipRepository interface {
	Get(id string) (*domain.Dealership, error)
	List() ([]domain.CleanDealership, error)
	ListByCountryAndState(country, state string) ([]domain.CleanDealership, error)
	Create(newDealership *domain.Dealership) (int64, error)
	Update(dealership *domain.Dealership) (int64, error)
	Delete(id string) (int64, error)
}
