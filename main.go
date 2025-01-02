package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// --- Using Gin package for http server with Logger and Recovery middleware attached
	server := gin.Default()

	// --- handler for incoming req
	server.GET("/events", getEvents)

	// --- listen for incoming requests (localhost)
	server.Run(":8080")
}

// ----- For GET Reqs...
func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}