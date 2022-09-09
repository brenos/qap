package dealershipservice

import (
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	dealershipRequest, err := domain.FromJSONDealershipRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = service.usecase.Create(dealershipRequest)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := domain.NewMessageResponse("Dealership created")

	c.IndentedJSON(http.StatusCreated, response)
}
