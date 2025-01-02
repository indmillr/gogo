package routes

import "github.com/gin-gonic/gin"

// --- handlers (routes) for incoming requests ([endpoint], [action])
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

	server.PUT("/events/:id", updateEvent)
}