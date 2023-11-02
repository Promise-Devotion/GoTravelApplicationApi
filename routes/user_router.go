package routes

import (
	"github.com/Promise-Devotion/GoTravelApplicationApi/controllers/admin"
	"github.com/gin-gonic/gin"
)

func UserRuoterRegister(e *gin.Engine) {

	UserRouters := e.Group("/user")
	{
		UserRouters.GET("/getUserInfo", admin.UserController{}.GetUserInfo)
		UserRouters.GET("/getUserList", admin.UserController{}.GetUserList)
	}
}
