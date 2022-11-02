package middlewares

import (
	"github/abbgo/isleg/isleg-backend/auth"
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
			context.AbortWithStatusJSON(401, gin.H{"message": "request does not contain an access token"})
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
			context.AbortWithStatusJSON(403, gin.H{"message": err.Error()})
			return
		}
		claims, ok := token.Claims.(*auth.JWTClaim)
		if !ok {
			context.AbortWithStatusJSON(400, gin.H{"message": "couldn't parse claims"})
			return
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			context.AbortWithStatusJSON(403, gin.H{"message": "token expired"})
			return
		}
		context.Set("customer_id", claims.CustomerID)
		context.Next()
	}
}

func IsSuperAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader("AuthorizationAdmin")
		tokenString := strings.Split(tokenStr, " ")[1]

		if tokenString == "" {
			context.AbortWithStatusJSON(401, gin.H{"message": "request does not contain an access token"})
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&auth.JWTClaimForAdmin{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.JwtKey), nil
			},
		)
		if err != nil {
			context.AbortWithStatusJSON(403, gin.H{"message": err.Error()})
			return
		}
		claims, ok := token.Claims.(*auth.JWTClaimForAdmin)
		if !ok {
			context.AbortWithStatusJSON(400, gin.H{"message": "couldn't parse claims"})
			return
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			context.AbortWithStatusJSON(403, gin.H{"message": "token expired"})
			return
		}
		// context.Set("admin_id", claims.AdminID)

		if claims.Type != "super_admin" {
			context.AbortWithStatusJSON(400, gin.H{"message": "only super_admin can add admin and super_admin"})
			return
		}

		context.Next()
	}
}
