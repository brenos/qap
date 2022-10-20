package carservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) Delete(c *gin.Context) {
	id := c.Param("id")

	message := service.usecase.Delete(id)

	c.IndentedJSON(http.StatusOK, message)
}
