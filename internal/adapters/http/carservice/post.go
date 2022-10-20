package carservice

import (
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	carRequest, err := domain.FromJSONCreateCarRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := service.usecase.Create(carRequest)

	c.IndentedJSON(http.StatusCreated, message)
}
