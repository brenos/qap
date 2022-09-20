package userservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) GetByEmail(c *gin.Context) {
	email := c.Param("email")

	if email == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Email required",
		})
		return
	}

	user, err := service.usecase.GetByEmail(email)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (service service) GetByToken(c *gin.Context) {
	token := c.Param("token")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Token required",
		})
		return
	}

	user, err := service.usecase.GetByToken(token)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
