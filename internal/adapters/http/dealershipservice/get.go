package dealershipservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// GetDealership(s) godoc
// @Summary Return dealership(s) from DB
// @Schemes
// @Description Return dealership(s) from DB, by ID, by Country and/or Sate or All Dealerships.
// @Tags dealerships
// @Accept json
// @Produce json
// @Param       token header string true "User token"
// @Param       id    query     string  false  "Dealership search by id"
// @Param       country    query     string  false  "Dealership search by country"
// @Param       state    query     string  false  "Dealership search by state"
// @Success 200 {object} domain.Result
// @Failure 400 {object} domain.Result
// @Failure 500 {object} domain.Result
// @Router /dealership [get]
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
