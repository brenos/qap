package carservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) ListByDealership(c *gin.Context) {
	idDealership, founded := c.Params.Get("idDealership")
	if !founded {
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

func (service service) ListByBrandAndOrModel(c *gin.Context) {
	model, modelFounded := c.Params.Get("model")
	brand, brandFounded := c.Params.Get("brand")

	if !modelFounded && !brandFounded {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Brand and Model is empty",
		})
	}

	cars, err := service.usecase.ListByBrandAndOrModel(brand, model)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, cars)
}
