package dealershipservice

import (
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// UpdateDealership godoc
// @Summary      Update a dealership in DB
// @Description  Update a dealership in DB
// @Tags dealerships
// @Accept       json
// @Produce      json
// @Param        token header string true "User token"
// @Param        dealership body domain.Dealership true "Dealership"
// @Success 200 {object} domain.Result
// @Failure 400 {object} domain.Result
// @Failure 422 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router       /dealership [put]
func (service service) Update(c *gin.Context) {
	dealershipRequest, err := domain.FromJSONDealershipRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := service.usecase.Update(dealershipRequest)

	c.IndentedJSON(http.StatusOK, message)
}
