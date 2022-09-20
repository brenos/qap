package userPorts

import "github.com/brenos/qap/internal/core/domain"

type UserUseCase interface {
	Create(userRequest *domain.CreateUserRequest) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetByToken(token string) (*domain.User, error)
	UpdateRequestCount(token string) error
}
