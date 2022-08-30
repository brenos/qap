package main

import (
	"github.com/brenos/qap/di"
	"github.com/brenos/qap/internal/adapters/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := postgres.GetDbConnection()
	defer conn.Close()

	//postgres.RunMigrations()

	userService := di.ConfigUserDI(conn)

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	userGroup := api.Group("/user")
	userGroup.GET("/:id", userService.Get)

	r.Run()
}
