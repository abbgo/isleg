package auth

import (
	"github/abbgo/isleg/isleg-backend/helpers"
	"net/http"
	"os"
	"strconv"
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

	accessTokenTimeOut, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIMEOUT"))
	if err != nil {
		return "", "", err
	}
	expirationTimeAccessToken := time.Now().Add(time.Duration(accessTokenTimeOut) * time.Second)

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

	refreshTokenTimeOut, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIMEOUT"))
	if err != nil {
		return "", "", err
	}
	expirationTimeRefreshToken := time.Now().Add(time.Duration(refreshTokenTimeOut) * time.Second)
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
		helpers.HandleError(c, 401, "request does not contain an refresh token")
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
		helpers.HandleError(c, 403, err.Error())
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		helpers.HandleError(c, 400, "couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		helpers.HandleError(c, 403, "token expired")
		return
	}

	accessTokenString, refreshTokenString, err := GenerateTokenForCustomer(claims.PhoneNumber, claims.CustomerID)
	if err != nil {
		helpers.HandleError(c, 500, err.Error())
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
