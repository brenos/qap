package carPorts

import "github.com/brenos/qap/internal/core/domain"

type CarRepository interface {
	Get(id string) (*domain.Car, error)
	ListByDealership(idDealership string) ([]domain.CleanCar, error)
	ListByBrandAndOrModel(brand, model string) ([]domain.Car, error)
	Create(newCar *domain.Car) (int64, error)
	Update(car *domain.Car) (int64, error)
	Delete(id string) (int64, error)
}
