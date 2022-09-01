package dealershipPorts

import "github.com/brenos/qap/internal/core/domain"

type DealershipRepository interface {
	Get(id string) (*domain.Dealership, error)
	List() ([]domain.Dealership, error)
	ListByState(state string) ([]domain.Dealership, error)
	ListByCountry(country string) ([]domain.Dealership, error)
	Create(newDealership *domain.Dealership) (*domain.Dealership, error)
}
