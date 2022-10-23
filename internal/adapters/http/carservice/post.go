package carservice

import (
	"net/http"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	carRequest := domain.CreateCarRequest{}
	resultErr := helpers.ValidateOrCreateBodyRequest(c, &carRequest)

	if resultErr != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resultErr)
		return
	}

	message := service.usecase.Create(&carRequest)

	c.IndentedJSON(http.StatusCreated, message)
}
