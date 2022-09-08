package dealershipPorts

import (
	"github.com/gin-gonic/gin"
)

type DealershipService interface {
	Get(c *gin.Context)
	List(c *gin.Context)
	ListByCountryAndOrState(c *gin.Context)
	Create(c *gin.Context)
}
