package dealershipservice

import (
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	dealershipRequest, err := domain.FromJSONCreateDealershipRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = service.usecase.Create(dealershipRequest)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := domain.NewResultMessage("Dealership created!")

	c.IndentedJSON(http.StatusCreated, message)
}
