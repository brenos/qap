package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/liveness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run()
}
