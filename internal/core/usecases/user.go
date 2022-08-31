package usecases

import (
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/user"
)

type userUserCase struct {
	userRepo ports.UserRepository
}

func NewUserUseCase(userRepo ports.UserRepository) ports.UserUseCase {
	return &userUserCase{
		userRepo: userRepo,
	}
}

func (u *userUserCase) Get(id string) (*domain.User, error) {
	userGetted, err := u.userRepo.Get(id)
	if err != nil {
		log.Panicf("Error getting from repo - %s", err)
		return nil, err
	}
	return userGetted, nil
}

func (u *userUserCase) List() ([]domain.User, error) {
	users, err := u.userRepo.List()
	if err != nil {
		log.Panicf("Error listing from repo - %s", err)
		return nil, err
	}

	return users, nil
}

func (u *userUserCase) Create(userRequest *domain.CreateUserRequest) (*domain.User, error) {
	var userId = helpers.RandomUUIDAsString()
	var token = userId
	newUser := domain.NewUser(userId, userRequest.Email, token, userRequest.IsPaidUser, 0)

	_, err := u.userRepo.Create(newUser)
	if err != nil {
		log.Panicf("Error creating from repo - %s", err)
		return nil, err
	}

	return newUser, nil
}
