package carservice

import (
	"net/http"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// CreateCar godoc
// @Summary      Create a car in DB
// @Description  Create a car in DB
// @Tags cars
// @Accept       json
// @Produce      json
// @Param        token header string true "User token"
// @Param        car body domain.CreateCarRequest true "Car"
// @Success 200 {object} domain.Result
// @Failure 400 {object} domain.Result
// @Failure 422 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router       /car [post]
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
