package main

import (
	docs "github.com/Promise-Devotion/GoTravelApplicationApi/docs"

	routes "github.com/Promise-Devotion/GoTravelApplicationApi/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@Title			travel项目api
//	@version		1.0
//	@description	GoTravelApplication前后端分离后端API.

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r = routes.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
