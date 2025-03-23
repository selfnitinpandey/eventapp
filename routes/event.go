package routes

import (
	"eventapp/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine) {
	{
		EventGroup := router.Group("/events")
		EventGroup.GET("/getevent", controller.GetEvent)
		EventGroup.POST("/createevent", controller.CreateEvent)
	}
}
