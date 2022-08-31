package ports

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Get(c *gin.Context)
	List(c *gin.Context)
	Create(c *gin.Context)
}
