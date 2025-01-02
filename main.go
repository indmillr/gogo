package main

import (
	"example.com/gogo/db"
	"example.com/gogo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// --- establish DB connection
	db.InitDB()

	// --- Using Gin package for http server with Logger and Recovery middleware attached
	server := gin.Default()

	routes.RegisterRoutes(server)

	// --- listen for incoming requests (localhost)
	server.Run(":8080")
}

