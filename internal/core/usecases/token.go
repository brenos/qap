package usecases

import (
	"errors"
	"log"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	ports "github.com/brenos/qap/internal/core/ports/token"
	"github.com/golang-jwt/jwt"
)

type tokenUseCase struct{}

func NewTokenUseCase() ports.TokenUseCase {
	return &tokenUseCase{}
}

func (t *tokenUseCase) GenerateToken(user *domain.User) (string, error) {
	tokenKey := []byte(helpers.TOKEN_KEY())
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	return token.SignedString(tokenKey)

}

func (t *tokenUseCase) VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.TOKEN_KEY()), nil
	})
	if err != nil {
		log.Panicf("Error validate token %s - Err:  %s", tokenString, err)
	}
	return token.Valid
}

func (t *tokenUseCase) GetUserIdByToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.TOKEN_KEY()), nil
	})
	if err == nil {
		claims := token.Claims.(jwt.MapClaims)
		return claims["id"].(string), nil
	}
	return "", errors.New("Error on get user id by token.")
}
