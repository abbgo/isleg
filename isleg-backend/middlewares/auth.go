package middlewares

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenStr := context.GetHeader("Authorization")
		if tokenStr == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Token is required")
			return
		}
		var tokenString string

		splitToken := strings.Split(tokenStr, "Bearer ")
		if len(splitToken) > 1 {
			tokenString = splitToken[1]
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
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

		db, err := config.ConnDB()
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}
		defer func() {
			if err := db.Close(); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
				return
			}
		}()

		rowCustomer, err := db.Query("SELECT id FROM customers WHERE phone_number = $1 AND deleted_at IS NULL", claims.PhoneNumber)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}
		defer func() {
			if err := rowCustomer.Close(); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
				return
			}
		}()

		var customer_id string

		for rowCustomer.Next() {
			if err := rowCustomer.Scan(&customer_id); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
				return
			}
		}

		if customer_id == "" {
			context.AbortWithStatusJSON(400, gin.H{"message": "customer not found"})
			return
		}

		context.Set("customer_id", claims.CustomerID)
		context.Next()

	}

}

func IsSuperAdmin() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenStr := context.GetHeader("Authorization")
		if tokenStr == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Token is required")
			return
		}
		var tokenString string

		splitToken := strings.Split(tokenStr, "Bearer ")
		if len(splitToken) > 1 {
			tokenString = splitToken[1]
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
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

func CheckAdmin() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenStr := context.GetHeader("Authorization")
		if tokenStr == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Token is required")
			return
		}
		var tokenString string

		splitToken := strings.Split(tokenStr, "Bearer ")
		if len(splitToken) > 1 {
			tokenString = splitToken[1]
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
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

		db, err := config.ConnDB()
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}
		defer func() {
			if err := db.Close(); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
				return
			}
		}()

		rowAdmin, err := db.Query("SELECT id FROM admins WHERE id = $1 AND deleted_at IS NULL", claims.AdminID)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}
		defer func() {
			if err := rowAdmin.Close(); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
				return
			}
		}()

		var admin_id string

		for rowAdmin.Next() {
			if err := rowAdmin.Scan(&admin_id); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
				return
			}
		}

		if admin_id == "" {
			context.AbortWithStatusJSON(400, gin.H{"message": "admin not found"})
			return
		}

		// context.Set("admin_id", claims.AdminID)
		context.Next()

	}

}
