package userservice

import (
	"net/http"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	userRequest := domain.CreateUserRequest{}
	resultErr := helpers.ValidateOrCreateBodyRequest(c, &userRequest)

	if resultErr != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resultErr)
		return
	}

	message := service.usecase.Create(&userRequest)

	c.IndentedJSON(http.StatusCreated, message)
}
