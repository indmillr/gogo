package routes

import (
	"example.com/gogo/middlewares"
	"github.com/gin-gonic/gin"
)

// --- handlers (routes) for incoming requests ([endpoint], [action])
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// --- define one 'authenticated' Group that will Use middleware Authenticate
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// it is possible add items from middleware (auth) in line for the route one at a time, but Group is more efficient
	// server.POST("/events", middlewares.Authenticate, createEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}