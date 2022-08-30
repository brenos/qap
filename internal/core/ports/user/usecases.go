package ports

import "github.com/brenos/qap/internal/core/domain"

type UserUseCase interface {
	Get(id string) (*domain.User, error)
	List() ([]domain.User, error)
	Create(userRequest *domain.CreateUserRequest) (*domain.User, error)
}
