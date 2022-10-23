package dealershipservice

import (
	"net/http"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	dealershipRequest := domain.CreateDealershipRequest{}
	resultErr := helpers.ValidateOrCreateBodyRequest(c, &dealershipRequest)

	if resultErr != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resultErr)
		return
	}

	message := service.usecase.Create(&dealershipRequest)

	c.IndentedJSON(http.StatusCreated, message)
}
