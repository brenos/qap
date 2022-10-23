package carPorts

import "github.com/brenos/qap/internal/core/domain"

type CarUseCase interface {
	Get(id string) *domain.Result
	ListByDealership(idDealership string) *domain.Result
	ListByBrandAndOrModel(brand, model string) *domain.Result
	Create(newCar *domain.CreateCarRequest) *domain.Result
	Update(car *domain.Car) *domain.Result
	Delete(id string) *domain.Result
}
