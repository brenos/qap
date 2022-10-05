package userPorts

import "github.com/brenos/qap/internal/core/domain"

type UserRepository interface {
	Create(newUser *domain.User) (*domain.User, error)
	GetById(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	UpdateRequestCount(id string) error
}
