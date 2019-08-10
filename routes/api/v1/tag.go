package v1

import (
	"gin-server/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		maps["state"] = state
	}

	//data["list"] = models.GetTags()
}

func AddTags(c *gin.Context) {

}

func EditTags(c *gin.Context) {

}

func DeleteTags(c *gin.Context) {

}
