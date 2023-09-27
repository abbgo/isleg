package middlewares

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"net/http"
	"strings"
	"time"

	ctx "context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth middleware gelen tokenin musdera degislimi ya-da dalmi sony barlayar
// eger gelen token dogry bolsa indi api gelen tokende musderinin id - sini alyp beryar
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
		defer db.Close()

		var customer_id string
		db.QueryRow(ctx.Background(), "SELECT id FROM customers WHERE id = $1 AND deleted_at IS NULL", claims.CustomerID).Scan(&customer_id)

		if customer_id == "" {
			context.AbortWithStatusJSON(404, gin.H{"message": "customer not found"})
			return
		}

		context.Set("customer_id", claims.CustomerID)
		context.Next()
	}
}

// IsSuperAdmin middleware dine super adminlere dostup beryar
// adminleri gecirmeyar
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
			context.AbortWithStatusJSON(400, gin.H{"message": "only super_admin can perform this task"})
			return
		}

		context.Next()
	}
}

// CheckAdmin middleware ahli adminlere dostup beryar
// gelen request - in admin tarapyndan gelip gelmedigini barlayar
// we admin bolmasa gecirmeyar
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
		defer db.Close()

		var admin_id string
		db.QueryRow(ctx.Background(), "SELECT id FROM admins WHERE id = $1 AND deleted_at IS NULL", claims.AdminID).Scan(&admin_id)

		if admin_id == "" {
			context.AbortWithStatusJSON(404, gin.H{"message": "admin not found"})
			return
		}

		context.Set("admin_id", claims.AdminID)
		context.Next()
	}
}
