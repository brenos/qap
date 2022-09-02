package carPorts

import "github.com/brenos/qap/internal/core/domain"

type CarUseCase interface {
	Get(id string) (*domain.Car, error)
	ListByDealership(idDealership string) ([]domain.Car, error)
	ListByBrand(brand string) ([]domain.Car, error)
	ListByModel(model string) ([]domain.Car, error)
	ListByBrandAndOrModel(brand, model string) ([]domain.Car, error)
	Create(newCar *domain.CreateCarRequest) (*domain.Car, error)
}
