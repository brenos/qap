package healthservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// Liveness godoc
// @Summary Service alive
// @Schemes
// @Description Indicate that the service is alive
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} domain.Result
// @Router /liveness [get]
func (s *service) Liveness(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, s.usecase.Liveness())
}

// @BasePath /api

// Readiness godoc
// @Summary Service able
// @Schemes
// @Description Service able to receive requests
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router /readiness [get]
func (s *service) Readiness(c *gin.Context) {
	readiness := s.usecase.Readiness()
	c.IndentedJSON(readiness.Code, readiness)
}
