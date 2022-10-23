package dealershipPorts

import "github.com/brenos/qap/internal/core/domain"

type DealershipUseCase interface {
	Get(id string) *domain.Result
	List() *domain.Result
	ListByCountryAndOrState(country, state string) *domain.Result
	Create(dealershipRequest *domain.CreateDealershipRequest) *domain.Result
	Update(dealership *domain.Dealership) *domain.Result
	Delete(id string) *domain.Result
}
