package helpers

import (
	"log"
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	userPorts "github.com/brenos/qap/internal/core/ports/user"
	"github.com/gin-gonic/gin"
)

func ValidateUserMiddleware(userUseCase userPorts.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) <= 0 {
			log.Printf("Token is not sended!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewResultMessage("You are not authorized!"))
			return
		}
		user, err := userUseCase.GetByToken(token)
		if err != nil || user == nil {
			log.Printf("ERROR | %s", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewResultMessage("You are not authorized!"))
			return
		}
		c.Next()
	}
}

func IncrementRequestCountMiddleware(userUseCase userPorts.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) <= 0 {
			log.Printf("Token is not sended!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewResultMessage("You are not authorized!"))
			return
		}
		err := userUseCase.UpdateRequestCount(token)
		if err != nil {
			log.Printf("ERROR | %s", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.NewResultMessage("We are a error, please try again!"))
			return
		}
		c.Next()
	}
}
