package dealershipPorts

import (
	"github.com/gin-gonic/gin"
)

type DealershipService interface {
	GetProxy(c *gin.Context)
	Create(c *gin.Context)
}
