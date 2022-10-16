package carPorts

import "github.com/gin-gonic/gin"

type CarService interface {
	GetProxy(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
