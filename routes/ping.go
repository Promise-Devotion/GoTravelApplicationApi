package routes

import (
	"github.com/Promise-Devotion/GoTravelApplicationApi/controllers/admin"

	"github.com/gin-gonic/gin"
)

func PingRegister(e *gin.Engine) {
	PingRouters := e.Group("/ping")
	{
		PingRouters.GET("/demo", admin.PingController{}.PingDemo)
		// ……
	}
}
