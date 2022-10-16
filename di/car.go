package di

import (
	"database/sql"

	"github.com/brenos/qap/internal/adapters/postgres/carrepository"

	"github.com/brenos/qap/internal/adapters/http/carservice"
	ports "github.com/brenos/qap/internal/core/ports/car"
	"github.com/brenos/qap/internal/core/usecases"
)

func ConfigCarDI(conn *sql.DB) (ports.CarService, *ports.CarRepository) {
	carRepository := carrepository.NewCarPostgreRepo(conn)
	carUseCase := usecases.NewCarUseCase(carRepository)
	carService := carservice.New(carUseCase)

	return carService, &carRepository
}
