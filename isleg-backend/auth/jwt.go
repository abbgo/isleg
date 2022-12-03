package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaim struct {
	PhoneNumber string `json:"phone_number"`
	CustomerID  string `json:"customer_id"`
	jwt.StandardClaims
}

func GenerateTokenForCustomer(phoneNumber, customerID string) (string, string, error) {

	expirationTimeAccessToken := time.Now().Add(30 * time.Minute)

	claimsAccessToken := &JWTClaim{
		PhoneNumber: phoneNumber,
		CustomerID:  customerID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeAccessToken.Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccessToken)
	accessTokenString, err := accessToken.SignedString(JwtKey)
	if err != nil {
		return "", "", nil
	}

	expirationTimeRefreshToken := time.Now().Add(12 * time.Hour)
	claimsRefreshToken := &JWTClaim{
		PhoneNumber: phoneNumber,
		CustomerID:  customerID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeRefreshToken.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
	refreshTokenString, err := refreshToken.SignedString(JwtKey)
	if err != nil {
		return "", "", nil
	}

	return accessTokenString, refreshTokenString, nil

}

func Refresh(c *gin.Context) {

	tokenStr := c.GetHeader("RefreshToken")
	tokenString := strings.Split(tokenStr, " ")[1]

	if tokenString == "" {
		c.JSON(401, gin.H{
			"message": "request does not contain an refresh token",
		})
		// c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
	)

	if err != nil {
		c.JSON(403, gin.H{
			"message": err.Error(),
		})
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("couldn't parse claims")})
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		c.JSON(403, gin.H{
			"message": errors.New("token expired"),
		})
		return
	}

	accessTokenString, refreshTokenString, err := GenerateTokenForCustomer(claims.PhoneNumber, claims.CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})

}

// func ValidateToken(signedToken string) (err error) {

// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&JWTClaim{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(JwtKey), nil
// 		},
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	claims, ok := token.Claims.(*JWTClaim)
// 	if !ok {
// 		err = errors.New("couldn't parse claims")
// 		return
// 	}
// 	if claims.ExpiresAt < time.Now().Local().Unix() {
// 		err = errors.New("token expired")
// 		return
// 	}
// 	return nil

// }

// func ValidateRefreshToken(signedToken string) (err error) {

// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&JWTClaim{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(jwtKey), nil
// 		},
// 	)
// 	if err != nil {
// 		return
// 	}
// 	claims, ok := token.Claims.(*JWTClaim)
// 	if !ok {
// 		err = errors.New("couldn't parse claims")
// 		return
// 	}
// 	if claims.ExpiresAt < time.Now().Local().Unix() {
// 		err = errors.New("token expired")
// 		return
// 	}
// 	return

// }
