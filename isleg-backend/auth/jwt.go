package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateAccessToken(phoneNumber string) (accessTokenString string, err error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &JWTClaim{
		Email: phoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err = token.SignedString(jwtKey)
	return
}

func GenerateRefreshToken(phoneNumber string) (refreshTokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: phoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err = token.SignedString(jwtKey)
	return
}
