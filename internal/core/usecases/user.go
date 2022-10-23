package usecases

import (
	"fmt"
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/brenos/qap/internal/core/domain/result"
	emailPorts "github.com/brenos/qap/internal/core/ports/email"
	tokenPorts "github.com/brenos/qap/internal/core/ports/token"
	ports "github.com/brenos/qap/internal/core/ports/user"
)

type userUseCase struct {
	userRepo     ports.UserRepository
	tokenUseCase tokenPorts.TokenUseCase
	emailAdapter emailPorts.EmailAdapter
}

func NewUserUseCase(userRepo ports.UserRepository, tokenUseCase tokenPorts.TokenUseCase, emailAdapter emailPorts.EmailAdapter) ports.UserUseCase {
	return &userUseCase{
		userRepo:     userRepo,
		tokenUseCase: tokenUseCase,
		emailAdapter: emailAdapter,
	}
}

func (u *userUseCase) Create(userRequest *domain.CreateUserRequest) *domain.Result {
	errValidateEmail := helpers.ValidateEmail(userRequest.Email)
	if errValidateEmail != nil {
		log.Println(errValidateEmail.Error())
		return domain.NewResultMessageAndCode(errValidateEmail.Error(), result.CodeInternalError)
	}

	var userId = helpers.RandomUUIDAsString()
	newUser := domain.NewUser(userId, userRequest.Email, false, 0)

	_, err := u.userRepo.Create(newUser)
	if err != nil {
		errTxt := fmt.Sprintf("Error creating user")
		log.Panicf("%s from repo - %s", errTxt, err)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	token, errToken := u.tokenUseCase.GenerateToken(newUser)
	if errToken != nil {
		u.Delete(userId)
		errTxt := fmt.Sprintf("Error creating user")
		log.Panicf("Error creating token - %s", errToken)
		return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
	}

	if len(token) > 0 {
		errSendEmail := u.emailAdapter.SendEmail(userRequest.Email, token)
		if errSendEmail != nil {
			u.Delete(userId)
			errTxt := fmt.Sprintf("Error creating user")
			return domain.NewResultMessageAndCode(errTxt, result.CodeInternalError)
		}
	}

	return domain.NewResultMessageContextCode("User created!", newUser, result.CodeCreated)
}

func (u *userUseCase) GetById(id string) (*domain.User, error) {
	userGetted, err := u.userRepo.GetById(id)
	if err != nil {
		log.Panicf("Error getting from repo - %s", err)
		return nil, err
	}
	return userGetted, nil
}

func (u *userUseCase) GetByEmail(email string) (*domain.User, error) {
	userGetted, err := u.userRepo.GetByEmail(email)
	if err != nil {
		log.Panicf("Error getting from repo - %s", err)
		return nil, err
	}
	return userGetted, nil
}

func (u *userUseCase) UpdateRequestCount(id string) error {
	err := u.userRepo.UpdateRequestCount(id)
	if err != nil {
		log.Panicf("Error on update request count - %s", err)
	}
	return err
}

func (u *userUseCase) Delete(id string) error {
	_, err := u.userRepo.Delete(id)

	if err != nil {
		log.Panicf("Error deleting user from repo - %s", err)
	}

	return err
}
