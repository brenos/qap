package carservice

import (
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// UpdateCar godoc
// @Summary      Update a car in DB
// @Description  Update a car in DB
// @Tags cars
// @Accept       json
// @Produce      json
// @Param        token header string true "User token"
// @Param        car body domain.Car true "Car"
// @Success 200 {object} domain.Result
// @Failure 400 {object} domain.Result
// @Failure 422 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router       /car [put]
func (service service) Update(c *gin.Context) {
	carRequest, err := domain.FromJSONCarRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := service.usecase.Update(carRequest)

	c.IndentedJSON(http.StatusOK, message)
}
