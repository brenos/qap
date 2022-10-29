package userservice

import (
	"net/http"

	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// CreateUser godoc
// @Summary      Create a user in DB
// @Description  Create a user in DB
// @Tags users
// @Accept       json
// @Produce      json
// @Param        user body domain.CreateUserRequest true "User"
// @Success 200 {object} domain.Result
// @Failure 400 {object} domain.Result
// @Failure 422 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router       /user [post]
func (service service) Create(c *gin.Context) {
	userRequest := domain.CreateUserRequest{}
	resultErr := helpers.ValidateOrCreateBodyRequest(c, &userRequest)

	if resultErr != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resultErr)
		return
	}

	message := service.usecase.Create(&userRequest)

	c.IndentedJSON(http.StatusCreated, message)
}
