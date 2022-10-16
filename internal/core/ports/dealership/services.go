package dealershipPorts

import (
	"github.com/gin-gonic/gin"
)

type DealershipService interface {
	GetProxy(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
