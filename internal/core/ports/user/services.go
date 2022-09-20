package userPorts

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(c *gin.Context)
	GetByEmail(c *gin.Context)
	GetByToken(c *gin.Context)
}
