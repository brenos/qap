package carservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// GetCars godoc
// @Summary Return cars from DB
// @Schemes
// @Description Return cars from DB, by ID, by dealership or by Brand and/or Model.
// @Tags cars
// @Accept json
// @Produce json
// @Param       token header string true "User token"
// @Param       id    query     string  false  "car search by id"
// @Param       idDealership    query     string  false  "car search by idDealership"
// @Param       brand    query     string  false  "car search by brand"
// @Param       model    query     string  false  "car search by model"
// @Success 200 {object} domain.Result
// @Failure 400 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router /car [get]
func (service service) GetProxy(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	idDealership := c.Request.URL.Query().Get("idDealership")
	brand := c.Request.URL.Query().Get("brand")
	model := c.Request.URL.Query().Get("model")

	if id != "" {
		service.get(c, id)
	} else if idDealership != "" {
		service.listByDealership(c, idDealership)
	} else if model != "" || brand != "" {
		service.listByBrandAndOrModel(c, brand, model)
	}
}

func (service service) get(c *gin.Context, id string) {
	if id == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "ID required",
		})
		return
	}

	message := service.usecase.Get(id)

	c.IndentedJSON(http.StatusOK, message)
}

func (service service) listByDealership(c *gin.Context, idDealership string) {
	if idDealership == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "idDealership id empty",
		})
	}

	message := service.usecase.ListByDealership(idDealership)

	c.IndentedJSON(http.StatusOK, message)
}

func (service service) listByBrandAndOrModel(c *gin.Context, brand, model string) {
	if brand == "" && model == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Brand or Model is empty",
		})
	}

	message := service.usecase.ListByBrandAndOrModel(brand, model)

	c.IndentedJSON(http.StatusOK, message)
}
