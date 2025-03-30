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
		EventGroup.PUT("/updateevent/:id", controller.UpdateEvent)
		EventGroup.DELETE("/deleteevent/:id", controller.DeleteEvent)
	}
}
