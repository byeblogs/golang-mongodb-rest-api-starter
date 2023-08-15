package routes

import (
	"github.com/byeblogs/golang-mongodb-rest-api-starter/controllers"
	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.RouterGroup) {
	auth := router.Group("/ping")
	{
		auth.GET(
			"",
			controllers.Ping,
		)
	}
}
