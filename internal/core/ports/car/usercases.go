package carPorts

import "github.com/brenos/qap/internal/core/domain"

type CarUseCase interface {
	Get(id string) (*domain.Car, error)
	ListByDealership(idDealership string) ([]domain.CleanCar, error)
	ListByBrandAndOrModel(brand, model string) ([]domain.Car, error)
	Create(newCar *domain.CarRequest) error
	Update(car *domain.CarRequest) error
	Delete(id string) error
}
