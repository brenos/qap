package di

import (
	"database/sql"

	"github.com/brenos/qap/internal/adapters/http/carservice"
	ports "github.com/brenos/qap/internal/core/ports/car"
	"github.com/brenos/qap/internal/core/usecases"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigCarDI(conn *sql.DB, carRepository ports.CarRepository) ports.CarService {
	carUseCase := usecases.NewCarUseCase(carRepository)
	carService := carservice.New(carUseCase)

	return carService
}
