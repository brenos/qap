package di

import (
	"database/sql"

	"github.com/brenos/qap/internal/adapters/http/carservice"
	"github.com/brenos/qap/internal/adapters/postgres/carrepository"
	ports "github.com/brenos/qap/internal/core/ports/car"
	"github.com/brenos/qap/internal/core/usecases"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigCarDI(conn *sql.DB) (ports.CarService, ports.CarRepository) {
	carRepository := carrepository.NewCarPostgreRepo(conn)
	carUseCase := usecases.NewCarUseCase(carRepository)
	carService := carservice.New(carUseCase)

	return carService, carRepository
}
