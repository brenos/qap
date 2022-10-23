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

	message := service.usecase.Get(id)

	c.IndentedJSON(http.StatusOK, message)
}

func (service service) listAll(c *gin.Context) {
	message := service.usecase.List()

	c.IndentedJSON(http.StatusOK, message)
}

func (service service) listByCountryAndOrState(c *gin.Context, country, state string) {
	if country == "" && state == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Country or State is empty",
		})
	}

	message := service.usecase.ListByCountryAndOrState(country, state)

	c.IndentedJSON(http.StatusOK, message)
}
