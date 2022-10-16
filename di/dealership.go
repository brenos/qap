package di

import (
	"database/sql"

	"github.com/brenos/qap/internal/adapters/http/dealershipservice"
	"github.com/brenos/qap/internal/adapters/postgres/dealershiprepository"
	portsCar "github.com/brenos/qap/internal/core/ports/car"
	ports "github.com/brenos/qap/internal/core/ports/dealership"
	"github.com/brenos/qap/internal/core/usecases"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigDealershipDI(conn *sql.DB, carRepository *portsCar.CarRepository) ports.DealershipService {
	dealeshipRepository := dealershiprepository.NewDealershipPostgreRepo(conn, carRepository)
	dealeshipUseCase := usecases.NewDealershipUseCase(dealeshipRepository)
	dealeshipService := dealershipservice.New(dealeshipUseCase)

	return dealeshipService
}
