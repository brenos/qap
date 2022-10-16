package carservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	car, err := service.usecase.Get(id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, car)
}

func (service service) listByDealership(c *gin.Context, idDealership string) {
	if idDealership == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "idDealership id empty",
		})
	}

	cars, err := service.usecase.ListByDealership(idDealership)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, cars)
}

func (service service) listByBrandAndOrModel(c *gin.Context, brand, model string) {
	if brand == "" && model == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Brand or Model is empty",
		})
	}

	cars, err := service.usecase.ListByBrandAndOrModel(brand, model)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, cars)
}
