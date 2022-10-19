package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaim struct {
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}

func GenerateAccessToken(phoneNumber string) (accessTokenString string, err error) {

	// expirationTime := time.Now().Add(1 * time.Hour)
	expirationTime := time.Now().Add(10 * time.Second)

	claims := &JWTClaim{
		PhoneNumber: phoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err = token.SignedString(jwtKey)
	return

}

// func GenerateRefreshToken(phoneNumber string) (refreshTokenString string, err error) {

// 	// expirationTime := time.Now().Add(5 * time.Hour)
// 	expirationTime := time.Now().Add(10 * time.Hour)

// 	claims := &JWTClaim{
// 		PhoneNumber: phoneNumber,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	refreshTokenString, err = token.SignedString(jwtKey)
// 	return

// }

// func Refresh(c *gin.Context) {

// 	tokenStr := c.GetHeader("RefreshToken")
// 	fmt.Println("resfresh token: ", tokenStr)
// 	// tokenString := strings.Split(tokenStr, " ")[1]
// 	if tokenStr == "" {
// 		c.JSON(401, gin.H{
// 			"message": "request does not contain an refresh token",
// 		})
// 		// c.Abort()
// 		return
// 	}

// 	token, err := jwt.ParseWithClaims(
// 		tokenStr,
// 		&JWTClaim{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(jwtKey), nil
// 		},
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 		})
// 		// c.Abort()
// 		return
// 	}

// 	claims, ok := token.Claims.(*JWTClaim)

// 	if !ok {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": errors.New("couldn't parse claims")})
// 		// c.Abort()
// 		return
// 	}

// 	if claims.ExpiresAt < time.Now().Local().Unix() {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": errors.New("token expired"),
// 		})
// 		// c.Abort()
// 		return
// 	}

// 	accessTokenString, err := GenerateAccessToken(claims.PhoneNumber)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		// c.Abort()
// 		return
// 	}

// 	refreshTokenString, err := GenerateRefreshToken(claims.PhoneNumber)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		// c.Abort()
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"access_token":  accessTokenString,
// 		"refresh_token": refreshTokenString,
// 	})

// }

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return errors.New("yalnyslyk")
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return nil
}

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
