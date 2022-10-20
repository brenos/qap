package userPorts

import "github.com/brenos/qap/internal/core/domain"

type UserUseCase interface {
	Create(userRequest *domain.CreateUserRequest) *domain.Result
	GetById(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	UpdateRequestCount(id string) error
	Delete(id string) error
}
