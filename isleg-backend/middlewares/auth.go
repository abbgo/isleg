package middlewares

import (
	"fmt"
	"github/abbgo/isleg/isleg-backend/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader("Authorization")

		tokenString := strings.Split(tokenStr," ")[1]

		fmt.Println("gelen token: ",tokenString)
		if tokenString == "" {
			context.JSON(401, gin.H{"message": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"message": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
