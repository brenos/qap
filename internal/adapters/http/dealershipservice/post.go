package dealershipservice

import (
	"net/http"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// CreateDealership godoc
// @Summary      Create a dealership in DB
// @Description  Create a dealership in DB
// @Tags dealerships
// @Accept       json
// @Produce      json
// @Param        token header string true "User token"
// @Param        dealership body domain.CreateDealershipRequest true "Dealership"
// @Success 200 {object} domain.Result
// @Failure 400 {object} domain.Result
// @Failure 422 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router       /dealership [post]
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
