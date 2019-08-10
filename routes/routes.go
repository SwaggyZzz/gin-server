package routes

import (
	"gin-server/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTags)
		apiV1.PUT("/tags/:id", v1.EditTags)
		apiV1.DELETE("/tags/:id", v1.DeleteTags)
	}

	return r
}
