package helpers

import (
	"log"
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	tokenPorts "github.com/brenos/qap/internal/core/ports/token"
	userPorts "github.com/brenos/qap/internal/core/ports/user"
	"github.com/gin-gonic/gin"
)

func ValidateTokenMiddleware(userUseCase userPorts.UserUseCase, tokenUseCase tokenPorts.TokenUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) <= 0 {
			log.Panicln("ERROR | Token is not sended!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewResultMessage("You are not authorized!"))
			return
		}
		isValid := tokenUseCase.VerifyToken(token)
		if !isValid {
			log.Panicln("ERROR | Token is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewResultMessage("You are not authorized!"))
			return
		}
		c.Next()
	}
}

func ValidateUserAndIncrementRequestCountMiddleware(userUseCase userPorts.UserUseCase, tokenUseCase tokenPorts.TokenUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) <= 0 {
			log.Panicln("Token is not sended!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewResultMessage("You are not authorized!"))
			return
		}
		userId, errToken := tokenUseCase.GetUserIdByToken(token)
		if errToken != nil {
			log.Panicf("ERROR | %s", errToken.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewResultMessage("Token error!"))
			return
		}
		user, errUser := userUseCase.GetById(userId)
		if errUser != nil || user == nil {
			log.Panicf("ERROR | %s", errUser.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewResultMessage("You are not authorized!"))
			return
		}
		err := userUseCase.UpdateRequestCount(userId)
		if err != nil {
			log.Panicf("ERROR | %s", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewResultMessage("We are a error, please try again!"))
			return
		}
		c.Next()
	}
}
