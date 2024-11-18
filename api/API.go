package api

import (
	"goShop/api/routes"
	docs "goShop/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router = gin.Default()

func APIRegister() {
	println("API Registering")

	docs.SwaggerInfo.BasePath = "api/v1"
	v1Api := Router.Group("/api/v1")

	routes.RegisterUserRoutes(v1Api.Group("/users"))
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
