package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-server/config"
	"strings"
)

func main() {
	projectConfig := config.Config
	router := gin.Default()
	if strings.ToLower(projectConfig.Mode) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 查询字符串参数使用现有的底层 request 对象解析。
	// 请求响应匹配的 URL： /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guest")
		// 这个是 c.Request.URL.Query().Get("lastname") 的快捷方式。
		lastName := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})
	router.Run(":8005")
}