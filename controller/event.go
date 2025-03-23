package controller

import (
	"eventapp/db"
	"eventapp/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {
	client, err := db.ConnectDB()
	if err != nil {
		log.Println("Database did not connected...", err)
	}
	ctx := c.Request.Context()
	events, err := models.GetEvent(client, ctx)
	if err != nil {
		log.Println("Get Event giving error")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to fetch event app"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event data fetched successfully",
		"event":   events,
	})

}

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := db.ConnectDB()
	if err != nil {
		log.Println("Database did not connected...", err)
	}
	ctx := c.Request.Context()

	event.Created_At = time.Now()
	err = models.InsertEvent(client, ctx, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert Event"})

	}
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully", "event": event})
}

func DeleteEvent(c *gin.Context) {

}
