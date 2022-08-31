package di

import (
	"database/sql"

	"github.com/brenos/qap/internal/adapters/http/userservice"
	"github.com/brenos/qap/internal/adapters/postgres/userrepository"
	ports "github.com/brenos/qap/internal/core/ports/user"
	"github.com/brenos/qap/internal/core/usecases"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigUserDI(conn *sql.DB) ports.UserService {
	userRepository := userrepository.NewUserPostgreRepo(conn)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userService := userservice.New(userUseCase)

	return userService
}
