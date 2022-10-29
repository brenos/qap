package carservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// DeleteCar godoc
// @Summary      Delete car from DB
// @Description  Delete car from DB by ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Car ID"
// @Success      200  {object}  domain.Result
// @Failure      400  {object}  domain.Result
// @Failure      404  {object}  domain.Result
// @Failure      500  {object}  domain.Result
// @Router       /car/{id} [delete]
func (service service) Delete(c *gin.Context) {
	id := c.Param("id")

	message := service.usecase.Delete(id)

	c.IndentedJSON(http.StatusOK, message)
}
