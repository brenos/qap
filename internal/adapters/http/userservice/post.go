package userservice

import (
	"net/http"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	userRequest, err := domain.FromJSONCreateUserRequest(c.Request.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message := service.usecase.Create(userRequest)

	c.IndentedJSON(http.StatusCreated, message)
}
