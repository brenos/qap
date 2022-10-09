package main

import (
	"github.com/brenos/qap/di"
	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/adapters/postgres"
	sendgrid "github.com/brenos/qap/internal/adapters/sendGrid"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := postgres.GetDbConnection()
	defer conn.Close()

	postgres.RunMigrations()

	tokenUseCase := di.ConfigTokenDi()
	emailAdapter := sendgrid.NewEmailAdapter()
	userService, userUserCase := di.ConfigUserDI(conn, tokenUseCase, emailAdapter)
	carService, carRepository := di.ConfigCarDI(conn)
	dealershipService := di.ConfigDealershipDI(conn, carRepository)

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	userGroup := api.Group("/user")
	userGroup.POST("/", userService.Create)

	middlewareGroup := api.Group("/")
	middlewareGroup.Use(helpers.ValidateTokenMiddleware(*userUserCase, *&tokenUseCase))
	middlewareGroup.Use(helpers.ValidateUserAndIncrementRequestCountMiddleware(*userUserCase, *&tokenUseCase))

	carGroup := middlewareGroup.Group("/car")
	carGroup.GET("/", carService.GetProxy)
	carGroup.POST("/", carService.Create)
	carGroup.PUT("/", carService.Update)
	carGroup.DELETE("/:id", carService.Delete)

	dealershipGroup := middlewareGroup.Group("/dealership")
	dealershipGroup.GET("/", dealershipService.GetProxy)
	dealershipGroup.POST("/", dealershipService.Create)
	dealershipGroup.PUT("/", dealershipService.Update)
	dealershipGroup.DELETE("/:id", dealershipService.Delete)

	r.Run()
}

func NewEmailAdapter() {
	panic("unimplemented")
}
