package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

// @Tags ping Demo
// @Sumarry demo示例
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} string "操作成功"
// @Router /ping/demo [get]
func (con PingController) PingDemo(c *gin.Context) {
	userid := c.Request.FormValue("id")
	fmt.Println(userid + "------------------")
	c.JSON(http.StatusOK, gin.H{
		"message": "操作成功",
	})
}
