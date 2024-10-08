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

type JWTClaimForAdmin struct {
	PhoneNumber string `json:"phone_number"`
	AdminID     string `json:"admin_id"`
	Type        string `json:"type"`
	jwt.StandardClaims
}

func GenerateAccessTokenForAdmin(phoneNumber, adminID, adminType string) (string, string, error) {

	accessTokenTimeOut, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIMEOUT"))
	if err != nil {
		return "", "", err
	}
	expirationTimeAccessToken := time.Now().Add(time.Duration(accessTokenTimeOut) * time.Second)

	claimsAccessToken := &JWTClaimForAdmin{
		PhoneNumber: phoneNumber,
		AdminID:     adminID,
		Type:        adminType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeAccessToken.Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccessToken)
	accessTokenString, err := accessToken.SignedString(JwtKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenTimeOut, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIMEOUT"))
	if err != nil {
		return "", "", err
	}
	expirationTimeRefreshToken := time.Now().Add(time.Duration(refreshTokenTimeOut) * time.Second)
	claimsRefreshToken := &JWTClaimForAdmin{
		PhoneNumber: phoneNumber,
		AdminID:     adminID,
		Type:        adminType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeRefreshToken.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
	refreshTokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}

func RefreshTokenForAdmin(c *gin.Context) {

	tokenStr := c.GetHeader("RefreshToken")
	tokenString := strings.Split(tokenStr, " ")[1]

	if tokenString == "" {
		helpers.HandleError(c, 401, "request does not contain an refresh token")
		return
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaimForAdmin{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
	)

	if err != nil {
		helpers.HandleError(c, 403, err.Error())
		return
	}

	claims, ok := token.Claims.(*JWTClaimForAdmin)

	if !ok {
		helpers.HandleError(c, 400, "couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		helpers.HandleError(c, 403, "token expired")
		return
	}

	accessTokenString, refreshTokenString, err := GenerateAccessTokenForAdmin(claims.PhoneNumber, claims.AdminID, claims.Type)
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
