package middlewares

import (
	"net/http"

	"example.com/gogo/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
		return
	}

	// --- add data to the Context
	context.Set("userId", userId)
	// --- ensure the next request handler in line will execute
	context.Next()
}