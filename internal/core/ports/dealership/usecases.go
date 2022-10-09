package dealershipPorts

import "github.com/brenos/qap/internal/core/domain"

type DealershipUseCase interface {
	Get(id string) (*domain.Dealership, error)
	List() ([]domain.CleanDealership, error)
	ListByCountryAndOrState(country, state string) ([]domain.CleanDealership, error)
	Create(dealershipRequest *domain.CreateDealershipRequest) error
	Update(dealership *domain.Dealership) error
	Delete(id string) error
}
