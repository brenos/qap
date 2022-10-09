package di

import (
	"database/sql"

	"github.com/brenos/qap/internal/adapters/http/userservice"
	"github.com/brenos/qap/internal/adapters/postgres/userrepository"
	emailPorts "github.com/brenos/qap/internal/core/ports/email"
	tokenPorts "github.com/brenos/qap/internal/core/ports/token"
	ports "github.com/brenos/qap/internal/core/ports/user"
	"github.com/brenos/qap/internal/core/usecases"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigUserDI(conn *sql.DB, tokenUseCase tokenPorts.TokenUseCase, emailAdapter emailPorts.EmailAdapter) (ports.UserService, *ports.UserUseCase) {
	userRepository := userrepository.NewUserPostgreRepo(conn)
	userUseCase := usecases.NewUserUseCase(userRepository, tokenUseCase, emailAdapter)
	userService := userservice.New(userUseCase)

	return userService, &userUseCase
}
