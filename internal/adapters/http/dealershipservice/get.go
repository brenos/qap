package dealershipservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) GetProxy(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	country := c.Request.URL.Query().Get("country")
	state := c.Request.URL.Query().Get("state")

	if id != "" {
		service.get(c, id)
	} else if country != "" || state != "" {
		service.listByCountryAndOrState(c, country, state)
	} else {
		service.listAll(c)
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

func (service service) listAll(c *gin.Context) {
	dealerships, err := service.usecase.List()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, dealerships)
}

func (service service) listByCountryAndOrState(c *gin.Context, country, state string) {
	if country == "" && state == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Country or State is empty",
		})
	}

	dealerships, err := service.usecase.ListByCountryAndOrState(country, state)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, dealerships)
}
