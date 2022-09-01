package carPorts

import "github.com/gin-gonic/gin"

type CarService interface {
	Get(c *gin.Context)
	ListByDealership(c *gin.Context)
	ListByBrand(c *gin.Context)
	ListByModel(c *gin.Context)
	ListByBrandAndOrModel(c *gin.Context)
	Create(c *gin.Context)
}
