package dealershipservice

import (
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (service service) Delete(c *gin.Context) {
	id := c.Param("id")
	err := service.usecase.Delete(id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := domain.NewMessageResponse("Dealership deleted")

	c.IndentedJSON(http.StatusOK, response)
}
