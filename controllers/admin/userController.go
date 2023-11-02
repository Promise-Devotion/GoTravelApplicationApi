package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// @Tags 用户中心
// @Summary Get user by ID
// @Description Get user by ID
// @ID get-user-by-id
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {string} string "操作成功"
// @Router /user/getUserInfo [get]
func (con UserController) GetUserInfo(c *gin.Context) {
	userid := c.Request.FormValue("id")
	fmt.Println(userid + "------------------")
	c.JSON(http.StatusOK, gin.H{
		"message": "操作成功",
	})
}

func (con UserController) GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "操作成功",
	})
}
