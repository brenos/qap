package dealershipservice

import (
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (service service) Update(c *gin.Context) {
	dealershipRequest, err := domain.FromJSONDealershipRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = service.usecase.Update(dealershipRequest)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := domain.NewMessageResponse("Dealership updated")

	c.IndentedJSON(http.StatusOK, response)
}