package main

import (
	"net/http"

	"example.com/gogo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// --- Using Gin package for http server with Logger and Recovery middleware attached
	server := gin.Default()

	// --- handlers for incoming requests ([endpoint], [action])
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// --- listen for incoming requests (localhost)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	// --- handle error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}