package userservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) List(c *gin.Context) {
	users, err := service.usecase.List()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
