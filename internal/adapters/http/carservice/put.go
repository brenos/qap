package carservice

import (
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (service service) Update(c *gin.Context) {
	carRequest, err := domain.FromJSONCarRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = service.usecase.Update(carRequest)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := domain.NewResultMessageAndContext("Car updated!", carRequest)

	c.IndentedJSON(http.StatusOK, message)
}
