package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaimForAdmin struct {
	PhoneNumber string `json:"phone_number"`
	AdminID     string `json:"admin_id"`
	Type        string `json:"type"`
	jwt.StandardClaims
}

func GenerateAccessTokenForAdmin(phoneNumber, adminID, adminType string) (accessTokenString string, err error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaimForAdmin{
		PhoneNumber: phoneNumber,
		AdminID:     adminID,
		Type:        adminType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err = token.SignedString(JwtKey)
	return

}

func GenerateRefreshTokenForAdmin(phoneNumber, adminID, adminType string) (refreshTokenString string, err error) {

	expirationTime := time.Now().Add(24 * time.Hour * 5)
	claims := &JWTClaimForAdmin{
		PhoneNumber: phoneNumber,
		AdminID:     adminID,
		Type:        adminType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err = token.SignedString(JwtKey)
	return

}
