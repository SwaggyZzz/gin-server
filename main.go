package main

import (
	"fmt"
	"gin-server/config"
	"gin-server/models"
	"gin-server/routes"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := routes.InitRouter()
	projectConfig := config.Config
	if strings.ToLower(projectConfig.Mode) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	defer models.CloseDB()

	err := router.Run(projectConfig.HttpConfig.Port)
	if err != nil {
		fmt.Println("router error = ", err)
	}
}
