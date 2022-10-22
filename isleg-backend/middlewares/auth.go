package middlewares

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader("Authorization")
		tokenString := strings.Split(tokenStr, " ")[1]

		if tokenString == "" {
			context.JSON(401, gin.H{"message": "request does not contain an access token"})
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&auth.JWTClaim{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.JwtKey), nil
			},
		)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		claims, ok := token.Claims.(*auth.JWTClaim)
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse claims"})
			return
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			context.JSON(http.StatusBadRequest, gin.H{"message": "token expired"})
			return
		}
		context.Set("customer_id", claims.CustomerID)
		context.Next()
	}
}
