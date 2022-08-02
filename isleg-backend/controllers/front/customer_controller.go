package controllers

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"strconv"

	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterCustomer(c *gin.Context) {

	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := models.ValidateCustomerData(customer.PhoneNumber, customer.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	hashPassword, err := models.HashPassword(customer.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	_, err = config.ConnDB().Exec("INSERT INTO customers (full_name,phone_number,password,email) VALUES ($1,$2,$3,$4)", customer.FullName, customer.PhoneNumber, hashPassword, customer.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	accessTokenString, err := auth.GenerateAccessToken(customer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	refreshTokenString, err := auth.GenerateRefreshToken(customer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"message":       "customer successfully added",
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})

}

// func RegisterCustomer(c *gin.Context) {

// 	fullName := c.PostForm("full_name")
// 	phoneNumber := c.PostForm("phone_number")
// 	password := c.PostForm("password")
// 	email := c.PostForm("email")

// 	// gender := c.PostForm("gender")
// 	// birthday := c.PostForm("birthday")
// 	// addresses := c.PostFormArray("addresses")

// 	err := models.ValidateCustomerData(fullName, phoneNumber, password, email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	hashPassword, err := models.HashPassword(password)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	_, err = config.ConnDB().Exec("INSERT INTO customers (full_name,phone_number,password,email) VALUES ($1,$2,$3,$4)", fullName, phoneNumber, hashPassword, email)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	accessTokenString, err := auth.GenerateAccessToken(phoneNumber)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	refreshTokenString, err := auth.GenerateRefreshToken(phoneNumber)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":        true,
// 		"message":       "customer successfully added",
// 		"access_token":  accessTokenString,
// 		"refresh_token": refreshTokenString,
// 	})

// }

func LoginCustomer(c *gin.Context) {

	phoneNumber := c.PostForm("phone_number")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "phone number is required",
		})
		return
	}

	_, err := strconv.Atoi(phoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if len(phoneNumber) != 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the length of the phone number must be 8",
		})
		return
	}

	password := c.PostForm("password")
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "password is required",
		})
		return
	}

	if len(password) < 5 || len(password) > 25 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "password length must be between 5 and 25",
		})
		return
	}

	// check if email exists and password is correct
	row, err := config.ConnDB().Query("SELECT password FROM customers WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	var oldPassword string

	for row.Next() {
		if err := row.Scan(&oldPassword); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
			return
		}
	}

	if oldPassword == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "this client does not exist"})
		return
	}

	credentialError := models.CheckPassword(password, oldPassword)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	accessTokenString, err := auth.GenerateAccessToken(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	refreshTokenString, err := auth.GenerateRefreshToken(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})

}
