package main

import (
	"github.com/brenos/qap/di"
	_ "github.com/brenos/qap/docs"
	"github.com/brenos/qap/helpers"
	"github.com/brenos/qap/internal/adapters/postgres"
	sendgrid "github.com/brenos/qap/internal/adapters/sendGrid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          Quality Assurance Platform
// @version        1.0
// @description    Free API's that was developed to help IT roles to learn, test and/or use API's to make a tests for new positions
// @contact.url    http://brenos.github.io
// @BasePath /api
func main() {
	conn := postgres.GetDbConnection()
	defer conn.Close()

	postgres.RunMigrations()

	tokenUseCase := di.ConfigTokenDi()
	emailAdapter := sendgrid.NewEmailAdapter()
	userService, userUserCase := di.ConfigUserDI(conn, tokenUseCase, emailAdapter)
	carService, carRepository := di.ConfigCarDI(conn)
	dealershipService := di.ConfigDealershipDI(conn, carRepository)
	healthService := di.ConfigHealthDI(conn)

	r := gin.Default()

	api := r.Group("/api")
	api.GET("/liveness", healthService.Liveness)
	api.GET("/readiness", healthService.Readiness)

	userGroup := api.Group("/user")
	userGroup.POST("/", userService.Create)

	middlewareGroup := api.Group("/")
	middlewareGroup.Use(helpers.ValidateTokenMiddleware(*userUserCase, tokenUseCase))
	middlewareGroup.Use(helpers.ValidateUserAndIncrementRequestCountMiddleware(*userUserCase, tokenUseCase))

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func NewEmailAdapter() {
	panic("unimplemented")
}
