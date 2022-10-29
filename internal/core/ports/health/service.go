package healthPorts

import "github.com/gin-gonic/gin"

type HealthService interface {
	Liveness(c *gin.Context)
	Readiness(c *gin.Context)
}
