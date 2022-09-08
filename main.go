package main

import (
	"github.com/brenos/qap/di"
	"github.com/brenos/qap/internal/adapters/postgres"
	"github.com/brenos/qap/internal/adapters/postgres/carrepository"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := postgres.GetDbConnection()
	defer conn.Close()

	postgres.RunMigrations()

	userService := di.ConfigUserDI(conn)
	carRepository := carrepository.NewCarPostgreRepo(conn)
	carService := di.ConfigCarDI(conn, carRepository)
	dealershipService := di.ConfigDealershipDI(conn, carRepository)

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	userGroup := api.Group("/user")
	userGroup.GET("/:id", userService.Get)

	carGroup := api.Group("/car")
	carGroup.GET("/", carService.GetProxy)

	dealershipGroup := api.Group("/dealership")
	dealershipGroup.GET("/", dealershipService.GetProxy)

	r.Run()
}
