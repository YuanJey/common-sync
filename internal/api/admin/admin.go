package admin

import (
	"common-sync/pkg/config"
	"github.com/YuanJey/goutils2/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	//检查请求头key的值
	if c.GetHeader("key") != "admin" {
		c.JSON(401, "unauthorized")
		c.Abort()
	}
}
func SaveConfig(c *gin.Context) {
	operationID := utils.OperationIDGenerator()
	fileName := c.Query("fileName")
	if fileName == "" {
		c.JSON(400, gin.H{
			"operationID": operationID,
			"message":     "fileName is required",
		})
		return
	}
	err := config.ServerConfig.Save(operationID, fileName)
	if err != nil {
		c.JSON(500, gin.H{
			"operationID": operationID,
			"message":     err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"operationID": operationID,
		"message":     "save config file success",
	})
}
func Dept(c *gin.Context) {
	c.HTML(200, "dept.html", nil)
}
func User(c *gin.Context) {
	c.HTML(200, "user.html", nil)
}
