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

	user, err := service.usecase.Create(userRequest)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}
