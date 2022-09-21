package userPorts

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(c *gin.Context)
}
