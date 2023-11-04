package middlewares

import (
	"github.com/gin-gonic/gin"
	"main.go/auth"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "Invalid access token"})
			context.Abort()
			return
		}
		err, _ := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Err()})
			context.Abort()
			return
		}
		context.Next()
	}
}
