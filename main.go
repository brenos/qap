package main

import (
	"github.com/brenos/qap/di"
	"github.com/brenos/qap/internal/adapters/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := postgres.GetDbConnection()
	defer conn.Close()

	postgres.RunMigrations()

	//userService := di.ConfigUserDI(conn)
	carService, carRepository := di.ConfigCarDI(conn)
	dealershipService := di.ConfigDealershipDI(conn, carRepository)

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	//userGroup := api.Group("/user")
	//userGroup.GET("/:id", userService.Get)
	//userGroup.POST("/", userService.Create)

	carGroup := api.Group("/car")
	carGroup.GET("/", carService.GetProxy)
	carGroup.POST("/", carService.Create)
	carGroup.PUT("/", carService.Update)
	carGroup.DELETE("/:id", carService.Delete)

	dealershipGroup := api.Group("/dealership")
	dealershipGroup.GET("/", dealershipService.GetProxy)
	dealershipGroup.POST("/", dealershipService.Create)
	dealershipGroup.PUT("/", dealershipService.Update)
	dealershipGroup.DELETE("/:id", dealershipService.Delete)

	r.Run()
}
