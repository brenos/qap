package dealershipservice

import (
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Update(c *gin.Context) {
	dealershipRequest, err := domain.FromJSONDealershipRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := service.usecase.Update(dealershipRequest)

	c.IndentedJSON(http.StatusOK, message)
}
