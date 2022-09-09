package usecases

import (
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/car"
)

type carUseCase struct {
	carRepo ports.CarRepository
}

func NewCarUseCase(carRepo ports.CarRepository) ports.CarUseCase {
	return &carUseCase{
		carRepo: carRepo,
	}
}

func (c *carUseCase) Get(id string) (*domain.Car, error) {
	car, err := c.carRepo.Get(id)
	if err != nil {
		log.Panicf("Error getting car by id %s from repo - %s", id, err)
		return nil, err
	}
	return car, nil
}

func (c *carUseCase) ListByDealership(idDealership string) ([]domain.CleanCar, error) {
	cars, err := c.carRepo.ListByDealership(idDealership)
	if err != nil {
		log.Panicf("Error getting cars by dealershipId %s from repo - %s", idDealership, err)
		return nil, err
	}
	return cars, nil
}

func (c *carUseCase) ListByBrandAndOrModel(brand, model string) ([]domain.Car, error) {
	cars, err := c.carRepo.ListByBrandAndOrModel(brand, model)
	if err != nil {
		log.Panicf("Error getting cars by brand %s and model %s from repo - %s", brand, model, err)
		return nil, err
	}
	return cars, nil
}

func (c *carUseCase) Create(carDto *domain.CarRequest) error {
	var carId = helpers.RandomUUIDAsString()
	newCar := domain.NewCar(carId, carDto.Brand, carDto.Model, carDto.FuelType, carDto.IdDealerShip, carDto.Year, carDto.Price)

	err := c.carRepo.Create(newCar)

	if err != nil {
		log.Panicf("Error on creating car from repo - %s", err)
	}

	return err
}

func (c *carUseCase) Update(car *domain.CarRequest) error {
	carToUpdate := domain.NewCar(car.ID, car.Brand, car.Model, car.FuelType, car.IdDealerShip, car.Year, car.Price)

	err := c.carRepo.Update(carToUpdate)

	if err != nil {
		log.Panicf("Error on updating car from repo - %s", err)
	}

	return err
}

func (c *carUseCase) Delete(id string) error {
	err := c.carRepo.Delete(id)

	if err != nil {
		log.Panicf("Error on deleting car from repo - %s", err)
	}

	return err
}
