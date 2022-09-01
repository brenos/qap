package dealershipPorts

import "github.com/brenos/qap/internal/core/domain"

type DealershipUseCaseinterface interface {
	Get(id string) (*domain.Dealership, error)
	List() ([]domain.Dealership, error)
	ListByState(state string) ([]domain.Dealership, error)
	ListByCountry(country string) ([]domain.Dealership, error)
	Create(dealershipRequest *domain.CreateDealershipRequest) (*domain.Dealership, error)
}
