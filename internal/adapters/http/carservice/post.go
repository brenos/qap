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

	err = service.usecase.Create(carRequest)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := domain.NewResultMessage("Car created!")

	c.IndentedJSON(http.StatusCreated, message)
}
