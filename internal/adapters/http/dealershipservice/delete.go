package dealershipservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// DeleteDealership godoc
// @Summary      Delete dealership from DB
// @Description  Delete dealership from DB by ID
// @Tags         dealerships
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Deslership ID"
// @Success      200  {object}  domain.Result
// @Failure      400  {object}  domain.Result
// @Failure      404  {object}  domain.Result
// @Failure      500  {object}  domain.Result
// @Router       /dealership/{id} [delete]
func (service service) Delete(c *gin.Context) {
	id := c.Param("id")

	message := service.usecase.Delete(id)

	c.IndentedJSON(http.StatusOK, message)
}
